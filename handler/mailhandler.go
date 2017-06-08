package handler

import (
	"encoding/json"
	"fmt"
	"github.com/TechChallengeLyke/EmailService/action"
	"github.com/TechChallengeLyke/EmailService/data"
	"goji.io/pat"
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
	//TODO
}
