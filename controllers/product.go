/*
* @Author: souravray
* @Date:   2014-10-20 16:37:04
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 09:46:48
 */

package controllers

import (
	"github.com/souravray/gadgetWish/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func ProductList(w http.ResponseWriter, r *http.Request) {
	var products models.Products
	products = []models.Product{}
	session, _ := store.Get(r, "wishlist-session")
	if session.Values["user"] != nil {
		s, err := mgo.Dial(conf.DbURI)
		if err != nil {
			dispatchError(w, "database not responding")
			return
		}

		s.SetMode(mgo.Monotonic, true)
		c := s.DB(conf.DbName).C("products")

		err = c.Find(bson.M{}).Sort("-price").All(&products)
		s.Close()
	}
	dispatchJSON(w, products)
}
