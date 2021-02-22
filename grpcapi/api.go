package grpcapi

import (
	"context"
	"errors"
	"fmt"
	"userservice/app"
	pb "userservice/pb"
)

type API struct {
	App *app.App
	pb.UnimplementedUserServiceServer
}

func Init(s *app.Server) *API {
	a := app.New(app.ServerConnector(s))
	api := &API{App: a}
	return api
}

func (api *API) SessionHasPermissionTo(ctx context.Context, data *pb.Session) (*pb.AuthorizationResult, error) {
	err := api.App.MakeAuthenticationError()

	if data.Token == "" || data.UserId == "" {
		err.DetailedError = "Empty token or userId"
		return nil, errors.New(err.ToJson())
	}

	permission, err1 := api.App.Srv().Store.Permission().GetByName(data.PermissionName)
	if err1 != nil {
		err.DetailedError = "Get permission failed, " + err.DetailedError
		return nil, errors.New(err.ToJson())
	}

	if !api.App.UserHasPermissionTo(data.UserId, permission) {
		return nil, errors.New(api.App.MakePermissionError(permission).ToJson())
	}

	return &pb.AuthorizationResult{
		StatusCode: pb.StatusCode_OK,
		Message:    "Can Access",
	}, nil
}

func (api *API) SessionHasPermissionToCamera(ctx context.Context, sessionCamera *pb.SessionCamera) (*pb.AuthorizationResult, error) {
	return nil, nil
}

func (api *API) GetCamerasByUserId(ctx context.Context, data *pb.UserId) (*pb.CameraList, error) {
	fmt.Printf("User %v data\n", data.UserId)
	cameras, err := api.App.GetCamerasByUserId(data.UserId)
	if err != nil {
		return nil, errors.New(err.ToJson())
	}
	return &pb.CameraList{Cameras: cameras}, nil
}

func (api *API) DeleteCameraFromUser(ctx context.Context, data *pb.CameraId) (*pb.Empty, error) {
	err := api.App.DeleteCameraFromUser(data.CameraId)
	if err != nil {
		return nil, errors.New(err.ToJson())
	}

	return &pb.Empty{}, nil
}