package transport

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/invoice/datastruct"
	"backendbillingdashboard/modules/proforma-invoice/services"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

// ListInvoiceEndpoint is as request middleware
func ListInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.ListInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// CreateInvoiceEndpoint is as request middleware
func CreateInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.CreateInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil

	}
}

func CreateCustomInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.CreateCustomInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil

	}
}

// UpdateInvoiceEndpoint is as request middleware
func UpdateInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.UpdateInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// DeleteInvoiceEndpoint is as request middleware
func DeleteInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.DeleteInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func CancelInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.CancelInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func PrintInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.PrintInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func ReceivedDateEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.ReceivedDate(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func GenerateInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.GenerateInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func SummaryInvoiceEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.SummaryInvoice(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func InquiryPaymentEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InquiryPaymentRequest); ok {
			return services.InquiryPayment(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

func AdjustmentConfirmationEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceRequest); ok {
			return services.AdjustmentConfirmation(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}
