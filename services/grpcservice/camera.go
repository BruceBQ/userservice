package grpcservice

import (
	"fmt"
	"time"
	"userservice/model"
	pb "userservice/pb"

	"google.golang.org/grpc"
)

type CameraService struct {
	settings model.GRPCSettings
	Conn     *grpc.ClientConn
	Client   pb.CameraClient
}

func newCameraServiceConn(setttings model.GRPCSettings) *CameraService {
	c := &CameraService{
		settings: setttings,
	}

	go c.initConnection()
	time.Sleep(2 * time.Second)
	return c
}

func (c *CameraService) initConnection() error {
	conn, err := grpc.Dial(*c.settings.CameraServiceAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Printf("Connect to grpc camera server failed: %s", err.Error())
		return err
	}

	c.Conn = conn
	c.Client = pb.NewCameraClient(conn)
	return nil
}
