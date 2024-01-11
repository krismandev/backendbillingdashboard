package services

import (
	"backendbillingdashboard/config"
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	"backendbillingdashboard/modules/invoice/datastruct"
	dt "backendbillingdashboard/modules/invoice/datastruct"
	"backendbillingdashboard/modules/proforma-invoice/models"
	"backendbillingdashboard/modules/proforma-invoice/processors"
	"context"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// InvoiceServices provides operations for endpoint

// ListInvoice is use for
func ListInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("InvoiceService.ListInvoice Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	listInvoice, err := processors.GetListInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range listInvoice {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

// CreateInvoice is use for
func CreateInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.CreateInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CompanyID) == 0 || len(req.InvoiceTypeID) == 0 || len(req.InvoiceDate) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// process input
	err = processors.InsertInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}

func CreateCustomInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.CreateCustomInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.CompanyID) == 0 || len(req.InvoiceTypeID) == 0 || len(req.InvoiceDate) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// process input
	err = processors.InsertCustomInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}

// UpdateInvoice is use for
func UpdateInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.UpdateInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.InvoiceDate) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	_, err = models.ValidateRevisionCounter(req.InvoiceID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNotAuthorized, err.Error(), err)
		return response
	}
	// block request if old data is not exists
	err = models.CheckInvoiceExists(req.InvoiceID, conn)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
		return response
	}

	// process input
	err = processors.UpdateInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}

// DeleteInvoice is use for
func DeleteInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.DeleteInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// run
	err = processors.DeleteInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

func CancelInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.CancelInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.CancelDesc) < 10 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, "Incomplete Request. Cancel description min 10 character", err)
		return response
	}

	// run
	err = processors.CancelInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
		return response
	}

	return response
}

func PrintInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.PrintInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	list, err := processors.GetListInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	var invoice datastruct.InvoiceDataStruct
	for _, each := range list {
		if each.InvoiceID == req.InvoiceID {
			invoice = each
		}
	}
	printCounter, _ := strconv.Atoi(invoice.PrintCounter)
	if printCounter >= config.Param.MaxPrintInvoice {
		core.ErrorGlobalSingleResponse(&response, core.ErrNotAuthorized, "Cannot print invoice more than "+strconv.Itoa(config.Param.MaxPrintInvoice), err)
		return response
	}

	// block request if old data is not exists
	// err = models.CheckInvoiceExists(req.InvoiceID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckInvoiceDuplicate(req.InvoiceID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process
	err = processors.PrintInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

func ReceivedDate(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.UpdateReceivedDate Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.ReceivedDate) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// process
	err = processors.UpdateReceivedDate(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

func GenerateInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.GenerateInvoice Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.MonthUse) == 0 || len(req.PaymentMethod) == 0 || len(req.DueDate) == 0 || len(req.InvoiceDate) == 0 || len(req.CurrencyCode) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// block request if old data is not exists
	// err = models.CheckInvoiceExists(req.InvoiceID, conn)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrNoData, core.DescNoData, err)
	// 	return response
	// }

	// block request if data is already exists
	// err = models.CheckInvoiceDuplicate(req.InvoiceID, conn, req)
	// if err != nil {
	// 	core.ErrorGlobalSingleResponse(&response, core.ErrDataExists, err.Error(), err)
	// 	return response
	// }

	// process
	err = processors.GenerateInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, core.DescServer, err)
	}

	return response
}

func SummaryInvoice(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("InvoiceService.SummaryInvoice Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	summaryInvoice, err := processors.SummaryInvoice(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range summaryInvoice {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

func InquiryPayment(ctx context.Context, req dt.InquiryPaymentRequest, conn *connections.Connections) core.GlobalListResponse {
	log.Infof("InvoiceService.InquiryPayment Request : %+v", req)
	var response = core.DefaultGlobalListResponse(ctx)
	var err error

	if len(req.MonthUse) == 0 {
		core.ErrorGlobalListResponse(&response, core.ErrServer, core.DescServer, err)
		return response
	}

	summaryInvoice, err := processors.InquiryPayment(conn, req)
	if err != nil {
		core.ErrorGlobalListResponse(&response, core.ErrServer, err.Error(), err)
		return response
	} else {
		response.Data.Page = req.Param.Page
		response.Data.PerPage = req.Param.PerPage
	}

	// append list data as []interface{}
	for _, ls := range summaryInvoice {
		response.Data.List = append(response.Data.List, ls)
	}

	return response
}

func AdjustmentConfirmation(ctx context.Context, req dt.InvoiceRequest, conn *connections.Connections) core.GlobalSingleResponse {
	log.Infof("InvoiceService.AdjustmentConfirmation Request : %+v", req)
	var response = core.DefaultGlobalSingleResponse(ctx)
	var err error

	// validate input
	if len(req.InvoiceID) == 0 || len(req.LastUpdateUsername) == 0 {
		core.ErrorGlobalSingleResponse(&response, core.ErrIncompleteRequest, core.DescIncompleteRequest, err)
		return response
	}

	// process input
	err = processors.AdjustmentConfirmation(conn, req)
	if err != nil {
		core.ErrorGlobalSingleResponse(&response, core.ErrServer, err.Error(), err)
	}

	return response
}
