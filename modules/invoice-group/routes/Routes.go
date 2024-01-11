package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/invoice-group/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	InvoiceGroupRoute(conn)
}

// InvoiceGroupRoute is used for
func InvoiceGroupRoute(conn *connections.Connections) {
	invoiceGroupRoute := mux.NewRouter()
	invoiceGroupRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListInvoiceGroupEndpoint(conn),
		transport.InvoiceGroupDecodeRequest,
		transport.InvoiceGroupListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceGroupRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateInvoiceGroupEndpoint(conn),
		transport.InvoiceGroupDecodeRequest,
		transport.InvoiceGroupSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceGroupRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateInvoiceGroupEndpoint(conn),
		transport.InvoiceGroupDecodeRequest,
		transport.InvoiceGroupSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceGroupRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteInvoiceGroupEndpoint(conn),
		transport.InvoiceGroupDecodeRequest,
		transport.InvoiceGroupSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/invoice-group", invoiceGroupRoute)
}
