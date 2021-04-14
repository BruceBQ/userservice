package app

import (
	"net"
	"net/http"
	"sync"
	"userservice/clog"
	"userservice/config"
	"userservice/jobs"
	"userservice/services/grpcservice"
	"userservice/store"
	"userservice/store/mongostore"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type Server struct {
	mongoStore         *mongostore.MongoSupplier
	Store              store.Store
	AppInitializedOnce sync.Once

	newStore func() store.Store

	SessionCache *redis.Client

	RootRouter *mux.Router
	Router     *mux.Router

	HTTPServer *http.Server
	ListenAdd  *net.TCPAddr

	configStore config.Store

	GRPCServer  *grpc.Server
	GRPCService *grpcservice.GRPCService
	Log         *clog.Logger

	Jobs *jobs.JobServer
}

func NewServer(options ...Option) (*Server, error) {
	rootRouter := mux.NewRouter()

	s := &Server{
		RootRouter: rootRouter,
		GRPCServer: grpc.NewServer(),
	}

	for _, option := range options {
		if err := option(s); err != nil {
			return nil, errors.Wrap(err, "failed to apply options")
		}
	}

	if s.configStore == nil {
		configStore, err := config.NewFileStore("config.json")
		if err != nil {
			return nil, errors.Wrap(err, "failed to load config")
		}

		s.configStore = configStore
	}

	if err := s.initLogging(); err != nil {
		clog.Error(err.Error())
	}

	clog.Info("Server is initializing...")

	// fakeApp := New(ServerConnector(s))
	// fakeApp.Srv()

	s.GRPCService = grpcservice.Init(s.Config().GRPCSettings)

	if s.newStore == nil {
		s.newStore = func() store.Store {
			s.mongoStore = mongostore.NewMongoSupplier(s.Config().MongoSettings)
			return store.New(s.mongoStore)
		}
	}

	s.Store = s.newStore()

	// s.GRPCService = grpcservice.Init(s.Config().GRPCSettings)
	s.Router = s.RootRouter.PathPrefix("/").Subrouter()

	if s.SessionCache == nil {
		client := redis.NewClient(&redis.Options{
			Addr:     *s.Config().RedisSettings.Address,
			Password: *s.Config().RedisSettings.Password,
			DB:       *s.Config().RedisSettings.DB,
		})
		s.SessionCache = client
	}

	s.InitJobs()

	return s, nil
}

func (s *Server) Start() error {
	clog.Info("Starting Server...")

	var handler http.Handler = s.RootRouter

	corsWrapper := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut},
		AllowedHeaders: []string{"*"},
	})

	handler = corsWrapper.Handler(handler)

	s.HTTPServer = &http.Server{
		Handler: handler,
	}

	httpAddr := *s.Config().ServiceSettings.HTTPAddress
	listener, _ := net.Listen("tcp", httpAddr)

	go func() {
		clog.Info("Serve http server...")
		s.HTTPServer.Serve(listener)
	}()

	grpcAddr := *s.Config().ServiceSettings.GRPCAddress
	gprcListner, _ := net.Listen("tcp", grpcAddr)

	s.GRPCServer.Serve(gprcListner)

	return nil
}

func (s *Server) Shutdown() error {
	return nil
}

func (s *Server) AppOptions() []AppOption {
	return []AppOption{
		ServerConnector(s),
	}
}

func (s *Server) initLogging() error {
	if s.Log == nil {
		s.Log = clog.NewLogger()
	}

	clog.InitGlobalLogger(s.Log)
	return nil
}

// StopHTTPServer Stop HTTPServer
func (s *Server) StopHTTPServer() {
	if s.HTTPServer != nil {
		s.HTTPServer.Close()
		s.HTTPServer = nil
	}
}

// StopGRPCServer Stop GRPC Server
func (s *Server) StopGRPCServer() {
	if s.GRPCServer != nil {
		s.GRPCServer.Stop()
		s.GRPCServer = nil
	}
}

func (s *Server) InitJobs() {
	s.Jobs = jobs.NewJobServer(s.Store)
}

func (s *Server) RunJobs() {
	if err := s.Jobs.StartWorkers(); err != nil {
		clog.Error("Failed to start job server workers" + err.Error())
	}

	if err := s.Jobs.StartSchedulers(); err != nil {
		clog.Error("Failed to start job server schedulers" + err.Error())
	}
}
