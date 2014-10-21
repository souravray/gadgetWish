/*
* @Author: souravray
* @Date:   2014-10-20 16:19:40
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-22 01:09:19
 */

package main

import (
	"encoding/gob"
	"github.com/gorilla/mux"
	"github.com/souravray/gadgetWish/models"
	"github.com/souravray/gadgetWish/router"
	"log"
	"net/http"
	"os"
)

func main() {
	// seting up router
	rtr := mux.NewRouter()
	router.Routes(rtr)
	http.Handle("/", rtr)

	//registr User object to be serialized in seesion object
	gob.Register(&models.User{})

	// start server here
	log.Println("Listening...")

	//http.ListenAndServe(":8080", nil) os.Getenv("PORT")
	http.ListenAndServe(GetPort(), nil)
}

func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
