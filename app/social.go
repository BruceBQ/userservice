package app

import (
	"errors"
	"net/http"
	"userservice/model"
	"userservice/store"
)

func (a *App) SocialLogin(user *model.SocialUser) (*model.SocialUserWithSession, *model.AppError) {
	var ruser *model.SocialUser

	ruser, err := a.Srv().Store.SocialUser().GetByEmail(user.Email)
	if err != nil {
		var nfErr *store.ErrNotFound
		switch {
		case errors.As(err, &nfErr):
			user.PreSave()
			id, err := a.Srv().Store.SocialUser().Create(user)
			if err != nil {
				return nil, model.NewAppError("SocialLogin", "app.social.create.app_err", err.Error(), nil, "", http.StatusInternalServerError)
			}

			ruser, err = a.Srv().Store.SocialUser().GetById(id)
			if err != nil {
				return nil, model.NewAppError("SocialLogin", "app.social.get_by_id.app_err", err.Error(), nil, "", http.StatusInternalServerError)
			}
		default:
			return nil, model.NewAppError("SocialLogin", "app.social.get_by_email.internal_err", err.Error(), nil, "", http.StatusInternalServerError)
		}
	}

	session := &model.Session{
		UserID: ruser.Id,
		IsAuth: true,
	}

	a.SetSessionExpireInDays(session, *a.Config().ServiceSettings.SessionLengthMobileInDays)

	_, err1 := a.CreateSession(session)
	if err1 != nil {
		return nil, err1
	}

	// ruser.Sanitize()
	rSocicalUser := &model.SocialUserWithSession{
		User:  ruser,
		Token: session.Token,
	}

	return rSocicalUser, nil
}

func (a *App) GetPlateNumbers(id string) ([]*string, *model.AppError) {
	plateNumbers, err := a.Srv().Store.SocialUser().GetPlateNumbers(id)
	if err != nil {
		return nil, model.NewAppError("GetPlateNumbers", "app.social.get_plate_numbers.app_error", "Lấy danh sách biển số thất bại", nil, "", http.StatusInternalServerError)
	}

	return plateNumbers, nil
}

func (a *App) AddPlateNumber(id string, plateNumber string) *model.AppError {
	plateNumbers, err := a.Srv().Store.SocialUser().GetPlateNumbers(id)
	if err != nil {
		return model.NewAppError("AddPlateNumber", "app.social.add_plate_number.app_error", "Thêm biển số thất bại.", nil, "", http.StatusInternalServerError)
	}
	if plateNumberIsExist(plateNumber, plateNumbers) {
		return model.NewAppError("AddPlateNumber", "app.social.add_plate_number.app_error", "Biến số đã được thêm trong danh sách của bạn.", map[string]interface{}{"plateNumber": "Biển số đã tồn tại."}, "", http.StatusBadRequest)
	}

	err = a.Srv().Store.SocialUser().AddPlateNumber(id, plateNumber)

	if err != nil {
		return model.NewAppError("AddPlateNumber", "app.social.add_plate_number.app_error", "Thêm biển số thất bại.", nil, "", http.StatusInternalServerError)
	}

	return nil
}

func (a *App) DeletePlateNumber(id string, plateNumber string) *model.AppError {
	plateNumbers, err := a.Srv().Store.SocialUser().GetPlateNumbers(id)
	if err != nil {
		return model.NewAppError("DeletePlateNumber", "app.social.delete_plate_number.app_error", "Xóa biển số thất bại.", nil, "", http.StatusInternalServerError)
	}

	if !plateNumberIsExist(plateNumber, plateNumbers) {
		return model.NewAppError("DeletePlateNumber", "app.social.delete_plate_number.app_error", "Biến số không tồn tại trong danh sách của bạn.", map[string]interface{}{"plateNumber": "Biển số không tồn tại."}, "", http.StatusBadRequest)
	}

	err = a.Srv().Store.SocialUser().DeletePlateNumber(id, plateNumber)
	if err != nil {
		return model.NewAppError("DeletePlateNumber", "app.social.delete_plate_number.app_error", "Xóa biển số thất bại.", nil, "", http.StatusInternalServerError)
	}

	return nil
}

func plateNumberIsExist(plateNumber string, plateNumbers []*string) bool {
	for _, plate := range plateNumbers {
		if *plate == plateNumber {
			return true
		}
	}

	return false
}
