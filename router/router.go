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
	"github.com/XenZi/ARS-2022-23/metrics"
	"github.com/XenZi/ARS-2022-23/repository"
	"github.com/XenZi/ARS-2022-23/tracing"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
)

/*
Ova funkcija nam sluzi nesto nalik kontroleru gde cemo da izvlacimo funkcije iz handlers foldera, tj., kada implementirate nesto novo vezano za nacin handleovanja, vi pozovete iz handlers metodu (obicno ce to biti service koji ce da radi sa DAO slojem)
*/
func HandleRequests() *mux.Router {
	router := mux.NewRouter()
	createdRepository, err := repository.New()
	if err != nil {
		panic(err.Error())
	}
	trace, closer := tracing.Init("configHandler")
	opentracing.SetGlobalTracer(trace)
	configHandler := handlers.ConfigHandler{
		Repo:   createdRepository,
		Tracer: trace,
		Closer: closer,
	}

	configGroupHandler := handlers.ConfigGroupHandler{
		Repo:   createdRepository,
		Tracer: trace,
		Closer: closer,
	}
	router.HandleFunc("/api/config", metrics.Count("api/config", configHandler.AddConfig)).Methods("POST")
	router.HandleFunc("/api/config", metrics.Count("api/config", configHandler.GetAll)).Methods("GET")
	router.HandleFunc("/api/config/{id}/{version}", metrics.Count("api/config/{id}/{version}", configHandler.GetOneConfig)).Methods("GET")
	router.HandleFunc("/api/config/{id}/{version}", metrics.Count("api/config/{id}/{version}", configHandler.DeleteOneConfig)).Methods("DELETE")
	router.HandleFunc("/api/group-config", metrics.Count("/api/group-config", configGroupHandler.AddConfigGroup)).Methods("POST")
	router.HandleFunc("/api/group-config", metrics.Count("/api/group-config", configGroupHandler.GetAllGroupConfigs)).Methods("GET")
	router.HandleFunc("/api/group-config/{id}/{version}/", metrics.Count("/api/group-config/{id}/{version}/", configGroupHandler.GetOneConfigGroup)).Methods("GET")
	//router.HandleFunc("/api/group-config/{id}/{version}", handlers.RemoveConfigGroup).Methods("DELETE")
	router.HandleFunc("/api/group-config/{id}/{version}/{label}", metrics.Count("/api/group-config/{id}/{version}/{label}", configGroupHandler.GetAllConfigsInGroupByLabel)).Methods("GET")
	router.HandleFunc("/swagger.yaml", handlers.SwaggerHandler).Methods("GET")
	router.Path("/metrics").Handler(metrics.MetricsHandler())

	// SwaggerUI
	optionsDevelopers := middleware.SwaggerUIOpts{SpecURL: "swagger.yaml"}
	developerDocumentationHandler := middleware.SwaggerUI(optionsDevelopers, nil)
	router.Handle("/docs", developerDocumentationHandler)
	return router
}
