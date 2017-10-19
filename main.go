package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"./gpio"
	"os/exec"
	"fmt"
	"os"
)

const (
	PinParam     string = "/{pin:\\d{1,2}}"
	DurParam     string = "/{dur:s|ms}"
	PatternParam string = "/{pat:[\\d,]+}"
)

var store = make(map[string]gpio.Led)

func main() {

	if _, err := exec.LookPath("gpio"); err != nil{
		fmt.Println("Wiringpi not installed")
		os.Exit(1)
	}else {
		fmt.Println("GPIO Ready!")
	}

	r := mux.NewRouter()
	r.HandleFunc("/set"+PinParam+DurParam+PatternParam, setLed)
	r.HandleFunc("/stop"+PinParam, stopLed)

	http.ListenAndServe(":8000", r)

}

func stopLed(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	if l, ok := store[v["pin"]]; ok {
		l.Stop()
		delete(store, v["pin"])
	}else{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("notfound"))
	}
}

func setLed(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)


	if l, ok := store[v["pin"]]; ok {
		l.Update(v["dur"], v["pat"])
	} else {
		l := gpio.Make(v["pin"], v["dur"], v["pat"])
		l.Blink()
		store[v["pin"]] = l
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))

}

