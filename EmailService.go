package main

import (
	"flag"
	"fmt"
	"github.com/TechChallengeLyke/EmailService/action"
	"github.com/TechChallengeLyke/EmailService/data"
	"goji.io"
	"goji.io/pat"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
)

//initialize number of workers with a sensible default, if it is not set
var numberOfWorkers = flag.Int("workers", 4, "The number of workers to start")

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func sendMail(w http.ResponseWriter, r *http.Request) {
	name := "test"
	fmt.Fprintf(w, "Sending mail to %v!", name)
}

func main() {

	flag.Parse()
	fmt.Printf("Starting with %v Workers\n", *numberOfWorkers)

	setupCleanupProcess()
	action.StartWorkers(*numberOfWorkers)

	email := data.Email{Subject: "test subject"}
	action.SendMail(email)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Post("/sendmail"), sendMail)
	mux.Handle(pat.Get("/*"), http.FileServer(http.Dir(filepath.Join(".", "public"))))
	http.ListenAndServe("localhost:8000", mux)

}

//catch os.Interrupt to make sure the current work is done before closing the service
func setupCleanupProcess() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		action.StopWorkers()
		os.Exit(1)
	}()

}
