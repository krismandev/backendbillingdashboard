package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/proforma-invoice/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	InvoiceRoute(conn)
	// CancelInvoiceRoute(conn)
	PrintInvoiceRoute(conn)
	// GenerateInvoiceRoute(conn)
	// SummaryInvoiceRoute(conn)
	// InquiryPaymentRoute(conn)
	// CustomInvoiceRoute(conn)
	// AdjustmentConfirmationRoute(conn)
	// ReceivedDateRoute(conn)
}

// InvoiceRoute is used for
func InvoiceRoute(conn *connections.Connections) {
	invoiceRoute := mux.NewRouter()
	invoiceRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	invoiceRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/proforma-invoice", invoiceRoute)
}

// func CustomInvoiceRoute(conn *connections.Connections) {
// 	invoiceRoute := mux.NewRouter()
// 	invoiceRoute.Methods("POST").Handler(httptransport.NewServer(
// 		transport.CreateCustomInvoiceEndpoint(conn),
// 		transport.InvoiceDecodeRequest,
// 		transport.InvoiceSingleEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/custom-invoice", invoiceRoute)
// }

// func CancelInvoiceRoute(conn *connections.Connections) {
// 	cancelInvoiceRoute := mux.NewRouter()
// 	cancelInvoiceRoute.Methods("GET").Handler(httptransport.NewServer(
// 		transport.CancelInvoiceEndpoint(conn),
// 		transport.CancelInvoiceDecodeRequest,
// 		transport.InvoiceSingleEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/invoice/cancel", cancelInvoiceRoute)
// }

func PrintInvoiceRoute(conn *connections.Connections) {
	printInvoiceRoute := mux.NewRouter()
	printInvoiceRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.PrintInvoiceEndpoint(conn),
		transport.InvoiceDecodeRequest,
		transport.InvoiceSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))

	http.Handle("/proforma-invoice/print", printInvoiceRoute)
}

// func ReceivedDateRoute(conn *connections.Connections) {
// 	printInvoiceRoute := mux.NewRouter()
// 	printInvoiceRoute.Methods("POST").Handler(httptransport.NewServer(
// 		transport.ReceivedDateEndpoint(conn),
// 		transport.InvoiceDecodeRequest,
// 		transport.InvoiceSingleEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/invoice/received-date", printInvoiceRoute)
// }

// func GenerateInvoiceRoute(conn *connections.Connections) {
// 	printInvoiceRoute := mux.NewRouter()
// 	printInvoiceRoute.Methods("POST").Handler(httptransport.NewServer(
// 		transport.GenerateInvoiceEndpoint(conn),
// 		transport.InvoiceDecodeRequest,
// 		transport.InvoiceSingleEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/invoice/generate", printInvoiceRoute)
// }

// func SummaryInvoiceRoute(conn *connections.Connections) {
// 	summaryInvoiceRoute := mux.NewRouter()
// 	summaryInvoiceRoute.Methods("GET").Handler(httptransport.NewServer(
// 		transport.SummaryInvoiceEndpoint(conn),
// 		transport.InvoiceDecodeRequest,
// 		transport.InvoiceListEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/invoice/summary", summaryInvoiceRoute)
// }

// func InquiryPaymentRoute(conn *connections.Connections) {
// 	summaryInvoiceRoute := mux.NewRouter()
// 	summaryInvoiceRoute.Methods("GET").Handler(httptransport.NewServer(
// 		transport.InquiryPaymentEndpoint(conn),
// 		transport.InquiryPaymentDecodeRequest,
// 		transport.InvoiceListEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/inquiry-payment", summaryInvoiceRoute)
// }

// func AdjustmentConfirmationRoute(conn *connections.Connections) {
// 	confirmationRoute := mux.NewRouter()
// 	confirmationRoute.Methods("POST").Handler(httptransport.NewServer(
// 		transport.AdjustmentConfirmationEndpoint(conn),
// 		transport.InvoiceDecodeRequest,
// 		transport.InvoiceSingleEncodeResponse,
// 		httptransport.ServerBefore(httptransport.PopulateRequestContext),
// 		httptransport.ServerBefore(core.GetRequestInformation),
// 	))

// 	http.Handle("/invoice/adjustment-confirmation", confirmationRoute)
// }
