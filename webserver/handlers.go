package main

import (
	"StreamingWebsite/webserver/dbops"
	"StreamingWebsite/webserver/defs"
	"StreamingWebsite/webserver/session"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	sid := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true, SessionId:sid}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		sendNormalReaponse(w, string(resp), 201)
	}

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	// router在parse url的时候会把：后面的东西解析为ByName
	uname := p.ByName("user_name")
	io.WriteString(w, "Hello, welcome, ALOHA!!!. \n this is the page of login.")
	io.WriteString(w, uname)
}


