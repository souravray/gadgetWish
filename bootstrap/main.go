/*
* @Author: souravray
* @Date:   2014-10-20 19:50:13
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-20 20:55:13
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbURI  = "mongodb://localhost/wishdb"
	dbName = "wishdb"
)

type Product struct {
	//BaseModel
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Img         string        `json:"img"  bson:"img"`
	Description string        `json:"description" bson:"description"`
	Price       int64         `json:"price"  bson:"price"`
	PriceUnit   string        `json:"price_unit"  bson:"price_unit"`
}

func main() {
	content, err := ioutil.ReadFile("products.json")
	if err != nil {
		fmt.Print("Error:", err)
	}
	var products []Product
	err = json.Unmarshal(content, &products)
	if err != nil {
		fmt.Print("Error:", err)
	}
	s, err := mgo.Dial(dbURI)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	c := s.DB(dbName).C("products")
	for _, product := range products {
		c.Insert(product)
	}
	s.Close()
}
