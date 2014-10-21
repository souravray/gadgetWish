/*
* @Author: souravray
* @Date:   2014-10-20 16:36:50
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 11:26:24
 */

package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/souravray/gadgetWish/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"strings"
)

// public method
func SignIn(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "wishlist-session")

	s, err := mgo.Dial(conf.DbURI)
	if err != nil {
		dispatchError(w, "database not responding")
		return
	}
	s.SetMode(mgo.Monotonic, true)
	c := s.DB(conf.DbName).C("user")

	var user models.User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	email := strings.TrimSpace(user.Email)

	c.Find(bson.M{"email": email}).One(&user)

	if len(user.Img) == 0 {
		user.Email = email
		user.Img, err = userImage(email)
		if err != nil {
			dispatchError(w, "user gravatar image cannot be created")
			return
		}

		err = user.Validator()
		if err != nil {
			dispatchError(w, err.Error())
			return
		}

		c.Upsert(bson.M{"email": user.Email}, user)
	}
	s.Close()

	session.Values["user"] = user
	session.Save(r, w)
	dispatchJSON(w, user)
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "wishlist-session")
	session.Options.MaxAge = -1
	session.Save(r, w)
	w.Write([]byte(""))
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	session, _ := store.Get(r, "wishlist-session")
	if session.Values["user"] == nil {
		dispatchError(w, "user session doesn't exists")
		return
	}
	user = session.Values["user"].(*models.User)
	dispatchJSON(w, user)
}

// private methods
func userImage(email string) (imageurl string, err error) {
	baseUrl := "http://www.gravatar.com/avatar/"
	h := md5.New()
	_, err = io.WriteString(h, email)

	if err == nil {
		imageurl = fmt.Sprintf("%s%x", baseUrl, h.Sum(nil))
	}
	return
}
