package app

import (
	"context"
	"net/http"
	"userservice/model"
	pb "userservice/pb"
)

func (a *App) checkCamera(cameras []string) *model.AppError {
	if len(cameras) > 0 && cameras[0] != "*" {
		camMap := make(map[string]string)
		for _, camID := range cameras {
			camMap[camID] = camID
		}

		response, err := a.srv.GRPCService.Camera().Get(context.Background(), &pb.GetCameraRequest{Ids: camMap})
		if err != nil {
			return model.NewAppError("checkCamera", "app.camera.check_camera.app_error", "Kiểm tra danh sách camera thất bại", nil, err.Error(), http.StatusInternalServerError)
		}

		for _, camID := range cameras {
			cams := response.GetCameras()
			if cams[camID].Id == "" {
				return model.NewAppError("checkCamera", "api.camera.check_camera.invalid_camera", "Danh sách camrea không hợp lệ", map[string]interface{}{"cameras": "Danh sách camera không hợp lệ"}, "", http.StatusBadRequest)
			}
		}

	}

	return nil
}

func (a *App) GetCamerasByUserId(userID string) ([]string, *model.AppError) {
	user, err := a.Srv().Store.User().Get(userID)
	if err != nil {
		return nil, model.NewAppError("GetCamerasByUserId", "", "Get camera by userID failed", nil, err.Error(), http.StatusInternalServerError)
	}

	return user.Cameras, nil
}

func (a *App) DeleteCameraFromUser(cameraId string) *model.AppError {
	err := a.Srv().Store.User().DeleteCameraFromUser(cameraId)
	if err != nil {
		return model.NewAppError("DeleteCameraFromUser", "app.camera.store.app_error", "Xóa camera khỏi danh sách người dùng thất bại", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}
