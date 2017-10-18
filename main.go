package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"./gpio"
	"strconv"
)

const (
	PinParam     string = "/{pin:\\d{1,2}}"
	DurParam     string = "/{dur:s|ms}"
	PatternParam string = "/{pat:[\\d,]+}"
)

var store = make(map[int]gpio.Led)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/set"+PinParam+DurParam+PatternParam, setLed)
	r.HandleFunc("/stop"+PinParam, stopLed)

	http.ListenAndServe(":8000", r)

}

func stopLed(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	p, _ := strconv.Atoi(v["pin"])

	if l, ok := store[p]; ok {
		l.Stop()
		delete(store, p)
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("notfound"))
	}
}

func setLed(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	p, _ := strconv.Atoi(v["pin"])

	if l, ok := store[p]; ok {
		l.Update(v["dur"], v["pat"])
	} else {
		l := gpio.Make(p, v["dur"], v["pat"])
		l.Blink()
		store[p] = l
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

}

