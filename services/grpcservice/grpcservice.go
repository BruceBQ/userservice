package grpcservice

import (
	"userservice/model"
	pb "userservice/pb"
)

type GRPCService struct {
	settings      model.GRPCSettings
	CameraService *CameraService
}

func Init(settings model.GRPCSettings) *GRPCService {
	g := &GRPCService{
		settings: settings,
	}

	g.CameraService = newCameraServiceConn(settings)

	return g
}

func (g *GRPCService) Camera() pb.CameraClient {
	return g.CameraService.Client
}
