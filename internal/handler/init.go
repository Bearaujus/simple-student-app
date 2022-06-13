package handler

import (
	"net/http"

	util "github.com/Bearaujus/simple-student-app/util"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := handler(w, r); err != nil {
		resp, _ := util.ParseResponseToJSON(false, []string{err.Error()}, nil)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resp)
		return
	}

	w.WriteHeader(http.StatusOK)
}
