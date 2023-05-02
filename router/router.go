package router

import (
	"github.com/XenZi/ARS-2022-23/handlers"
	"github.com/gorilla/mux"
)

/*
Ova funkcija nam sluzi nesto nalik kontroleru gde cemo da izvlacimo funkcije iz handlers foldera, tj., kada implementirate nesto novo vezano za nacin handleovanja, vi pozovete iz handlers metodu (obicno ce to biti service koji ce da radi sa DAO slojem)
*/
func HandleRequests() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/config", handlers.AddConfig).Methods("POST")
	router.HandleFunc("/api/config", handlers.GetAllConfigs).Methods("GET")
	router.HandleFunc("/api/config/{id}", handlers.GetOneConfig).Methods("GET")
	router.HandleFunc("/api/group-config", handlers.AddConfigGroup).Methods("POST")
	router.HandleFunc("/api/group-config", handlers.GetAllGroupConfigs).Methods("GET")
	router.HandleFunc("/api/group-config/{id}", handlers.GetOneConfigGroup).Methods("GET")
	router.HandleFunc("/api/group-config/{id}", handlers.RemoveConfigGroup).Methods("DELETE")
	router.HandleFunc("/api/group-config/{id}", handlers.AddConfigIntoGroup).Methods("PUT")
	return router
}
