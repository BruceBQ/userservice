package grpcapi

import (
	"context"
	"errors"
	"userservice/app"
	"userservice/model"
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

func (api *API) SessionHasPermissionToCamera(ctx context.Context, data *pb.SessionCamera) (*pb.AuthorizationResult, error) {
	err := api.App.MakeAuthenticationError()

	permission, err1 := api.App.Srv().Store.Permission().GetByName(data.PermissionName)
	if err1 != nil {
		err.DetailedError = "Get permission failed, " + err.DetailedError
		return nil, errors.New(err.ToJson())
	}

	if !api.App.UserHasPermissionToCamera(data.UserId, data.CameraId, permission) {
		return nil, errors.New(api.App.MakePermissionError(permission).ToJson())
	}

	return &pb.AuthorizationResult{
		StatusCode: pb.StatusCode_OK,
		Message:    "Can Access",
	}, nil
}

func (api *API) GetCamerasByUserId(ctx context.Context, data *pb.UserId) (*pb.CameraList, error) {
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

func (api *API) LogAudit(ctx context.Context, data *pb.AuditData) (*pb.String, error) {
	audit := model.Audit{
		UserID:         data.UserId,
		PermissionName: data.PermissionName,
		Data:           data.Data,
	}

	err := api.App.SaveAudit(&audit)
	if err != nil {
		return nil, errors.New(err.ToJson())
	}

	return &pb.String{Response: "OK"}, nil
}
