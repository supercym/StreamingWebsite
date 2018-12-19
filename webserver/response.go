package main

import (
	"StreamingWebsite/webserver/defs"
	"encoding/json"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse)  {
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalReaponse(w http.ResponseWriter, resp string, sc int)  {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}

