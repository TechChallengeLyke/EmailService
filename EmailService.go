package main

import (
	"flag"
	"fmt"
	"github.com/TechChallengeLyke/EmailService/action"
	"github.com/TechChallengeLyke/EmailService/handler"
	"github.com/TechChallengeLyke/EmailService/provider"
	"goji.io"
	"goji.io/pat"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

//initialize number of workers with a sensible default, if it is not set
var numberOfWorkers = flag.Int("workers", 4, "The number of workers to start")

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

//Keep all resources that are needed during the live time of the application
//in a config object
//this simplifies cleanup and makes mocking for unit tests easier
type GlobalConfig struct {
	ProviderList map[string]provider.EmailProvider
	Log          log.Logger
	//db connection
	//cache connection
	//etc
}

var globalConfig = GlobalConfig{}

func main() {

	flag.Parse()
	fmt.Printf("Starting with %v Workers\n", *numberOfWorkers)

	//globalConfig.ProviderList
	var err error
	globalConfig.ProviderList, err = action.InitializeImplementations()
	if err != nil {
		fmt.Printf("Email Provider Initialization failed : ", err.Error())
		return
	}

	action.StartWorkers(globalConfig.ProviderList, *numberOfWorkers)
	setupCleanupProcess()

	//email := data.Email{Subject: "test subject"}
	//action.SendMail(email)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Post("/sendmail"), handler.SendMail)
	mux.HandleFunc(pat.Get("/getmails/:number"), handler.GetMails)
	mux.HandleFunc(pat.Get("/getmails/:number/:from"), handler.GetMailsWithStartingPoint)
	mux.Handle(pat.Get("/*"), http.FileServer(http.Dir(filepath.Join(".", "public"))))
	http.ListenAndServe("localhost:8000", mux)

}

//catch os.Interrupt to make sure the current work is done before closing the service
func setupCleanupProcess() {

	c := make(chan os.Signal, 1)
	signal.Reset()
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Interrupt received ... shutting down")
		action.StopWorkers()
		os.Exit(1)
	}()

}
