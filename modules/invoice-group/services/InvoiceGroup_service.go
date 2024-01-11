package services

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/invoice-group/datastruct"
	"backendbillingdashboard/modules/invoice-group/models"
	"backendbillingdashboard/modules/invoice-group/processors"
	"context"

	log "github.com/sirupsen/logrus"
)

// InvoiceGroupServices provides operations for endpoint

// ListInvoiceGroup is use for
func ListInvoiceGroup(ctx context.Context, req dt.InvoiceGroupRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("InvoiceGroupService.ListInvoiceGroup Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listInvoiceGroup, err := processors.GetListInvoiceGroup(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listInvoiceGroup {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateInvoiceGroup is use for
func CreateInvoiceGroup(ctx context.Context, req dt.InvoiceGroupRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceGroupService.CreateInvoiceGroup Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.GroupName) == 0 || len(req.InvoiceTypeID) == 0 || len(req.CompanyID) == 0 || len(req.InvoiceGroupDetail) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if data is already exists
	err = models.CheckInvoiceGroupDuplicate("", conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, core.DescDataExists, err)
		return response
	}

	// process input
	response.Data, err = processors.InsertInvoiceGroup(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// UpdateInvoiceGroup is use for
func UpdateInvoiceGroup(ctx context.Context, req dt.InvoiceGroupRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceGroupService.UpdateInvoiceGroup Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.GroupID) == 0 || len(req.GroupName) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	err = models.CheckInvoiceGroupExists(req.GroupID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// block request if data is already exists
	// err = models.CheckInvoiceGroupDuplicate(req.GroupID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process input
	response.Data, err = processors.UpdateInvoiceGroup(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

// DeleteInvoiceGroup is use for
func DeleteInvoiceGroup(ctx context.Context, req dt.InvoiceGroupRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceGroupService.DeleteInvoiceGroup Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.GroupID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteInvoiceGroup(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}
