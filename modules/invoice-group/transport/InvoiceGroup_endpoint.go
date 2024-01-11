package transport

import (
	connections "backendbillingdashboard/connections"
	"backendbillingdashboard/core"
	dt "backendbillingdashboard/modules/invoice-group/datastruct"
	"backendbillingdashboard/modules/invoice-group/services"
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	log "github.com/sirupsen/logrus"
)

// ListInvoiceGroupEndpoint is as request middleware
func ListInvoiceGroupEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalListResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalListResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceGroupRequest); ok {
			return services.ListInvoiceGroup(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalListResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// CreateInvoiceGroupEndpoint is as request middleware
func CreateInvoiceGroupEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceGroupRequest); ok {
			return services.CreateInvoiceGroup(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil

	}
}

// UpdateInvoiceGroupEndpoint is as request middleware
func UpdateInvoiceGroupEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceGroupRequest); ok {
			return services.UpdateInvoiceGroup(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}

// DeleteInvoiceGroupEndpoint is as request middleware
func DeleteInvoiceGroupEndpoint(conn *connections.Connections) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var errResp core.GlobalSingleResponse
		errNoJWT, errJWT := core.HandleJWT(ctx)
		if errJWT != nil {
			core.ErrorGlobalSingleResponse(&errResp, errNoJWT, errJWT.Error(), errJWT)
			return errResp, nil
		}

		if req, ok := request.(dt.InvoiceGroupRequest); ok {
			return services.DeleteInvoiceGroup(ctx, req, conn), nil
		}

		log.Error("Unhandled error occured: request is in unknown format")
		core.ErrorGlobalSingleResponse(&errResp, core.ErrOthers, core.DescOthers, errors.New("Request is in unknown format"))
		return errResp, nil
	}
}
