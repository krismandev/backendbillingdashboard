package routes

import (
	"backendbillingdashboard/connections"
	"backendbillingdashboard/modules/server-data/transport"
	"net/http"

	"backendbillingdashboard/core"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func InitRoutes(conn *connections.Connections) {
	ServerDataRoute(conn)
	LoadServerDataRoute(conn)
	SenderRoute(conn)
	AccountRoute(conn)
	UserRoute(conn)
	ServerDataInquiryPaymentRoute(conn)
}

// ServerDataRoute is used for
func ServerDataRoute(conn *connections.Connections) {
	serverDataRoute := mux.NewRouter()
	serverDataRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverDataRoute.Methods("POST").Handler(httptransport.NewServer(
		transport.CreateServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverDataRoute.Methods("PATCH").Handler(httptransport.NewServer(
		transport.UpdateServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	serverDataRoute.Methods("DELETE").Handler(httptransport.NewServer(
		transport.DeleteServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataSingleEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data", serverDataRoute)
}

func LoadServerDataRoute(conn *connections.Connections) {
	loadServerDataRoute := mux.NewRouter()
	loadServerDataRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.LoadServerDataEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data/load", loadServerDataRoute)
}

func SenderRoute(conn *connections.Connections) {
	senderRoute := mux.NewRouter()
	senderRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListSenderEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data/sender", senderRoute)
}

func AccountRoute(conn *connections.Connections) {
	senderRoute := mux.NewRouter()
	senderRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListAccountEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data/account", senderRoute)
}

func UserRoute(conn *connections.Connections) {
	userRoute := mux.NewRouter()
	userRoute.Methods("GET").Handler(httptransport.NewServer(
		transport.ListUserEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data/user", userRoute)
}

func ServerDataInquiryPaymentRoute(conn *connections.Connections) {
	route := mux.NewRouter()
	route.Methods("GET").Handler(httptransport.NewServer(
		transport.ListInquiryPaymentEndpoint(conn),
		transport.ServerDataDecodeRequest,
		transport.ServerDataListEncodeResponse,
		httptransport.ServerBefore(httptransport.PopulateRequestContext),
		httptransport.ServerBefore(core.GetRequestInformation),
	))
	http.Handle("/server-data/inquiry-payment", route)
}
