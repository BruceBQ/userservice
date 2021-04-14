package api

import (
	"net/http"
	"userservice/model"
)

func (api *API) InitAudit() {
	api.BaseRoutes.Audit.Handle("", api.ApiHandler(getAudits)).Methods("POST")
	api.BaseRoutes.Audit.Handle("/filter_data", api.ApiHandler(getAuditFilterData)).Methods("GET")
}

func getAudits(c *Context, w http.ResponseWriter, r *http.Request) {
	if !c.App.UserHasPermissionTo(r.Header.Get("user_id"), model.PERMISSION_GET_AUDIT_LOGS) {
		c.SetPermissionError(model.PERMISSION_GET_AUDIT_LOGS)
		return
	}

	props := model.AuditRequestFromJson(r.Body)

	before := model.GetMillisFromTime(props.Before)

	audits, err := c.App.GetAudits(before, props.Page, props.UserId, props.Permission)
	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(audits.ToJson()))
}

func getAuditFilterData(c *Context, w http.ResponseWriter, r *http.Request) {
	filter, err := c.App.GetAuditFilterData()
	if err != nil {
		c.Err = err
		return
	}

	w.Write([]byte(filter.ToJson()))
}
