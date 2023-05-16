// Config Service API
//
//	   Title: Config Service API
//
//	   Schemes: http
//		  Version: 0.0.1
//		  BasePath: /
//
//		  Produces:
//			- application/json
//
// swagger:meta
package router

import (
	"github.com/XenZi/ARS-2022-23/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

/*
Ova funkcija nam sluzi nesto nalik kontroleru gde cemo da izvlacimo funkcije iz handlers foldera, tj., kada implementirate nesto novo vezano za nacin handleovanja, vi pozovete iz handlers metodu (obicno ce to biti service koji ce da radi sa DAO slojem)
*/
func HandleRequests() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/config", handlers.AddConfig).Methods("POST")
	router.HandleFunc("/api/config", handlers.GetAllConfigs).Methods("GET")
	router.HandleFunc("/api/config/{id}/{version}", handlers.GetOneConfig).Methods("GET")
	router.HandleFunc("/api/config/{id}/{version}", handlers.DeleteOneConfig).Methods("DELETE")
	router.HandleFunc("/api/group-config", handlers.AddConfigGroup).Methods("POST")
	router.HandleFunc("/api/group-config", handlers.GetAllGroupConfigs).Methods("GET")
	router.HandleFunc("/api/group-config/{id}/{version}", handlers.GetOneConfigGroup).Methods("GET")
	router.HandleFunc("/api/group-config/{id}/{version}", handlers.RemoveConfigGroup).Methods("DELETE")
	router.HandleFunc("/api/group-config/{id}/{version}/{label}", handlers.GetAllConfigsInGroupByLabel).Methods("GET")
	router.HandleFunc("/swagger.yaml", handlers.SwaggerHandler).Methods("GET")

	// SwaggerUI
	optionsDevelopers := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	developerDocumentationHandler := middleware.SwaggerUI(optionsDevelopers, nil)
	router.Handle("/docs", developerDocumentationHandler)
	return router
}
