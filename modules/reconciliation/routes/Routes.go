package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/reconciliation/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ReconRoute(conn)
}

// ReconRoute is used for
func ReconRoute(conn *connections.Connections) {
	stubRoute := mux.NewRouter()
	stubRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListReconEndpoint(conn),
		transport.ReconDecodeRequest,
		transport.ReconListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateReconEndpoint(conn),
		transport.ReconDecodeRequest,
		transport.ReconSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateReconEndpoint(conn),
		transport.ReconDecodeRequest,
		transport.ReconSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	stubRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteReconEndpoint(conn),
		transport.ReconDecodeRequest,
		transport.ReconSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/reconciliation", stubRoute)
}
