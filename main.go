package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/XenZi/ARS-2022-23/data"
	"github.com/XenZi/ARS-2022-23/model"
	routerConfig "github.com/XenZi/ARS-2022-23/router"
	"github.com/XenZi/ARS-2022-23/utils"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	router := routerConfig.HandleRequests()
	cf1 := model.Config{Id: utils.CreateId(), Entries: map[string]string{
		"test":  "test",
		"test2": "test2",
	}}
	cf2 := model.Config{Id: utils.CreateId(), Entries: map[string]string{
		"test3": "test3",
		"test4": "test4",
	}}
	cf1p := &cf1
	cf2p := &cf2
	cslice := []*model.Config{cf1p, cf2p}
	db := data.NewDataInstance()
	db.Service = model.Service{}
	db.ConfigGroups = []*model.ConfigGroup{
		{
			Id:    utils.CreateId(),
			Group: []*model.Config{cf1p},
		},
	}
	db.Service.Data = map[string][]*model.Config{
		"1": cslice,
	}
	// start server
	srv := &http.Server{Addr: "0.0.0.0:8000", Handler: router}
	go func() {
		log.Println("server starting")
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-quit

	log.Println("service shutting down ...")

	// gracefully stop server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server stopped")

}
