package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TechChallengeLyke/EmailService/action"
	"github.com/TechChallengeLyke/EmailService/data"
	"goji.io/pat"
	"html/template"
	"net/http"
	"strconv"
)

func SendMail(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	email := &data.Email{}
	err := decoder.Decode(email)
	if err != nil {
		http.Error(w, "error : invalid json", http.StatusBadRequest)
		return
	}

	err = action.SendMail(email)
	if err != nil {
		http.Error(w, "error : "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "OK\n")
	return
}

func GetMails(w http.ResponseWriter, r *http.Request) {
	getMails(w, r, 0)
}

func GetMailsWithStartingPoint(w http.ResponseWriter, r *http.Request) {
	from := pat.Param(r, "from")
	fromNumber, err := strconv.Atoi(from)
	if err != nil {
		http.Error(w, "error : "+err.Error(), http.StatusBadRequest)
		return
	}
	getMails(w, r, fromNumber)
}

func getMails(w http.ResponseWriter, r *http.Request, from int) {
	numberStr := pat.Param(r, "number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		http.Error(w, "error : "+err.Error(), http.StatusBadRequest)
		return
	}
	mails := data.GetMails(from, number)

	json, err := json.Marshal(mails)
	if err != nil {
		http.Error(w, "error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func GetMetrics(w http.ResponseWriter, r *http.Request) {
	in_progress, failure, success := data.GetMailMetrics()
	metrics := struct {
		InProgressMails int
		Failures int
		Success int
	}{InProgressMails:in_progress, Failures:failure, Success:success}
	json, err := json.Marshal(metrics)
	if err != nil {
		http.Error(w, "error : "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(json)
}

func ShowGetMails(w http.ResponseWriter, r *http.Request) {
	mails := data.GetMails(0, 100)
	model := struct {
		Mails []data.Email
	}{Mails: *mails}

	renderTemplate(w, "templates/list.html", "list.html", &model)
}


func renderTemplate(w http.ResponseWriter, tmpl string, name string, p interface{}) {
	t := template.New(name) // Create a template.
	t, _ = t.ParseFiles(tmpl)      // Parse template file.
	err := t.Execute(w, p)
	if err != nil {
		fmt.Println("error : ", err)
		http.Error(w, "error : "+err.Error(), http.StatusInternalServerError)
	}
}
