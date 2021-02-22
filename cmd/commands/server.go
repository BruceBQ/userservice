package commands

import (
	"os"
	"os/signal"
	"syscall"
	"userservice/api"
	"userservice/app"
	"userservice/clog"
	"userservice/config"
	"userservice/grpcapi"
	pb "userservice/pb"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func serverCmdF(command *cobra.Command, args []string) error {
	configDNS := viper.GetString("config")

	interruptChan := make(chan os.Signal, 1)

	configStore, err := config.NewStore(configDNS)

	if err != nil {
		return errors.Wrap(err, "Failed to load configuration")
	}

	return runServer(configStore, interruptChan)
}

func runServer(configStore config.Store, interruptChan chan os.Signal) error {
	options := []app.Option{app.ConfigStore(configStore)}

	server, err := app.NewServer(options...)
	if err != nil {
		clog.Critical(err.Error())
		return err
	}
	defer server.Shutdown()

	api.Init(server.AppOptions, server.RootRouter)
	// web.New(server.AppOptions, server.RootRouter)
	rpcapi := grpcapi.Init(server)
	pb.RegisterUserServiceServer(server.GRPCServer, rpcapi)

	serverErr := server.Start()
	if serverErr != nil {
		return serverErr
	}

	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan
	return nil
}
