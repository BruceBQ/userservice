package app

import (
	"math"
	"net/http"
	"userservice/model"
)

func (a *App) GetAudits(before int64, page int, userId string, permission string) (*model.AuditSearch, *model.AppError) {
	audits, err := a.Srv().Store.Audit().Get(before, page, userId, permission)
	if err != nil {
		return nil, model.NewAppError("GetAudits", "app.audit.get_audits", "Lấy danh sách nhật ký người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	count, err := a.Srv().Store.Audit().Count(before, userId, permission)
	if err != nil {
		return nil, model.NewAppError("GetAudits", "app.audit.get_audits", "Lấy danh sách nhật ký người dùng thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	auditSearch := &model.AuditSearch{
		AuditLogEntries: audits,
		Page:            &page,
		PerPage:         model.AuditPerPage,
		TotalPage:       int64(math.Ceil(float64(count / model.AuditPerPage))),
	}

	return auditSearch, nil
}

func (a *App) SaveAudit(audit *model.Audit) *model.AppError {
	err := a.Srv().Store.Audit().Save(audit)
	if err != nil {
		return model.NewAppError("SaveAudit", "app.audit.save_audit", "Lưu nhật ký thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (a *App) GetAuditFilterData() (*model.AuditFilter, *model.AppError) {
	users, err := a.Srv().Store.User().GetFilterAudit()
	if err != nil {
		return nil, model.NewAppError("GetAuditFilterData", "app.audit.get_audit_filter_data", "Lấy dữ liệu để lọc nhật ký chính sửa thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	permissions, err := a.Srv().Store.Permission().GetFilterAudit()
	if err != nil {
		return nil, model.NewAppError("GetAuditFilterData", "app.audit.get_audit_filter_data", "Lấy dữ liệu để lọc nhật ký chính sửa thất bại.", nil, err.Error(), http.StatusInternalServerError)
	}

	filter := &model.AuditFilter{
		Users:       users,
		Permissions: permissions,
	}

	return filter, nil
}
