/*
* @Author: souravray
* @Date:   2014-10-20 16:36:25
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-22 00:18:55
 */

package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/sessions"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

type Config struct {
	DbURI  string `json:"db_uri"`
	DbName string `json:"db_name"`
}

var conf *Config

var store = sessions.NewCookieStore([]byte("secret-wish-sting"))

var templates = template.Must(template.ParseFiles(
	"static/index.html",
))

func render(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dispatchError(w http.ResponseWriter, message string) {
	err := errors.New(message)
	http.Error(w, err.Error(), 500)
}

func dispatchJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func init() {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// change for heroku deployment
	conf.DbURI = os.Getenv("MONGO_URL")
	conf.DbName = os.Getenv("MDB_NAME")

	// err = json.Unmarshal(content, &conf)
	// if err != nil {
	// 	panic(err)
	// }
}

//
func Landing(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html")
	return
}
