package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/company-summary/datastruct"
	"backendbillingdashboard/modules/company-summary/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// CompanySummaryServices provides operations for endpoint

// ListCompanySummary is use for
func ListCompanySummary(ctx context.Context, req dt.CompanySummaryRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("CompanySummaryService.ListCompanySummary Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listCompanySummary, err := processors.GetListCompanySummary(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listCompanySummary {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}
