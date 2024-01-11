package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/company-summary/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	CompanySummaryRoute(conn)
}

// CompanySummaryRoute is used for
func CompanySummaryRoute(conn *connections.Connections) {
	stubRoute := mux.NewRouter()
	stubRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListCompanySummaryEndpoint(conn),
		transport.CompanySummaryDecodeRequest,
		transport.CompanySummaryListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/company-summary", stubRoute)
}
