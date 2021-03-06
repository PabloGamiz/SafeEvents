package api

import (
	"net/http"

	"github.com/PabloGamiz/SafeEvents-Backend/api/client"
	"github.com/PabloGamiz/SafeEvents-Backend/api/event"
	"github.com/PabloGamiz/SafeEvents-Backend/api/event/feedback"
	"github.com/PabloGamiz/SafeEvents-Backend/api/radar"
	"github.com/PabloGamiz/SafeEvents-Backend/api/ticket"
	"github.com/gorilla/mux"
)

// NewServer builds a brand new server
func NewServer() Server {
	api := &api{}

	router := mux.NewRouter()
	//router.Headers("Access-Control-Allow-Origin", "*")
	/*
		cors := handlers.CORS(
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Access-Control-Allow-Origin"}),
			handlers.AllowedOrigins([]string{"*"}))

		router.Use(cors)
	*/
	// Client router Handlers
	router.HandleFunc(client.APISigninPath, client.HandleSigninRequest).Methods(http.MethodPost)
	router.HandleFunc(client.APILogoutPath, client.HandleLogoutRequest).Methods(http.MethodPut)
	router.HandleFunc(client.APIClientInfoPath, client.HandleClientInfoRequest).Methods(http.MethodPut)
	router.HandleFunc(client.APIAddFavPath, client.HandleClientAddFavRequest).Methods(http.MethodPost)
	router.HandleFunc(client.APIDelFavPath, client.HandleClientDelFavRequest).Methods(http.MethodPost)
	router.HandleFunc(client.APIStatusPath, client.HandleStatusRequest).Methods(http.MethodPost)

	// Events router Handlers
	router.HandleFunc(event.APIListEvents, event.HandleListEventsRequest).Methods(http.MethodGet)
	router.HandleFunc(event.APIListEventsByType, event.HandleListEventsByTypeRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APIPubliEvent, event.HandlePublicaEventRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APIListFavorites, event.HandleListFavoritesRequest).Methods(http.MethodPut)
	router.HandleFunc(event.APIModificaEvent, event.HandleModificaEventRequest).Methods(http.MethodPut)
	router.HandleFunc(event.APIGetEvent, event.HandleGetEventRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APIGetEventAn, event.HandleGetEventAnonimRequest).Methods(http.MethodPost)
	router.HandleFunc(event.APIRecomanaEvents, event.HandleRecomanaEventsRequest).Methods(http.MethodPut)

	// Ticket router Handlers
	router.HandleFunc(ticket.APIPurchasePath, ticket.HandlePurchaseRequest).Methods(http.MethodPost)
	router.HandleFunc(ticket.APIActivatePath, ticket.HandleActivateRequest).Methods(http.MethodPut)
	router.HandleFunc(ticket.APIGetTicketsPath, ticket.HandleGetTicketsRequest).Methods(http.MethodPut)
	router.HandleFunc(ticket.APICheckPath, ticket.HandleCheckRequest).Methods(http.MethodPut)
	router.HandleFunc(ticket.APIPaypalPath, ticket.HandlePaypalPurchase).Methods(http.MethodPut)
	router.HandleFunc(ticket.APITicketPrice, ticket.HandleTicketPriceRequest).Methods(http.MethodPut)

	//Feedbacks router Handlers
	router.HandleFunc(feedback.APIPOSTFeedback, feedback.HandlePOSTFeedbackRequest).Methods(http.MethodPost)
	router.HandleFunc(feedback.APIPUTFeedback, feedback.HandlePUTFeedbackRequest).Methods(http.MethodPut)
	router.HandleFunc(feedback.APIDELETEFeedback, feedback.HandleDELETEFeedbackRequest).Methods(http.MethodDelete)
	router.HandleFunc(feedback.APIGETFeedbacks, feedback.HandleGETFeedbacksRequest).Methods(http.MethodPost)

	// Radar
	router.HandleFunc(radar.APIActivatePath, radar.HandleActivateRequest).Methods(http.MethodPut)
	router.HandleFunc(radar.APIActivatePath, radar.HandleDeactivateRequest).Methods(http.MethodDelete)
	router.HandleFunc(radar.APIInteractionPath, radar.HandleInteractionRequest).Methods(http.MethodPost)

	api.router = router
	return api
}
