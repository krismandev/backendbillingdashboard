package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/reconciliation/datastruct"
	"backendbillingdashboard/modules/reconciliation/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// ReconServices provides operations for endpoint

// ListRecon is use for
func ListRecon(ctx context.Context, req dt.ReconRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("ReconService.ListRecon Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listRecon, err := processors.GetListRecon(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listRecon {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateRecon is use for
func CreateRecon(ctx context.Context, req dt.ReconRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ReconService.CreateRecon Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	// if len(req.ReconID) == 0 || len(req.ReconName) == 0 {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
	// 	return response
	// }

	// // block request if data is already exists
	// err = models.CheckReconDuplicate("", conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
	// 	return response
	// }

	// process input
	response.Data, err = processors.InsertRecon(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateRecon is use for
func UpdateRecon(ctx context.Context, req dt.ReconRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ReconService.UpdateRecon Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	// if len(req.ReconID) == 0 || len(req.ReconName) == 0 {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
	// 	return response
	// }

	// // block request if old data is not exists
	// err = models.CheckReconExists(req.ReconID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// // block request if data is already exists
	// err = models.CheckReconDuplicate(req.ReconID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process input
	response.Data, err = processors.UpdateRecon(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteRecon is use for
func DeleteRecon(ctx context.Context, req dt.ReconRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("ReconService.DeleteRecon Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	// if len(req.ReconID) == 0 {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
	// 	return response
	// }

	// run
	err = processors.DeleteRecon(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
