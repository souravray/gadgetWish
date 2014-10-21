/*
* @Author: souravray
* @Date:   2014-10-20 16:34:01
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 03:40:27
 */

package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Products []Product

type Product struct {
	//BaseModel
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Img         string        `json:"img"  bson:"img"`
	Description string        `json:"description" bson:"description"`
	Price       string        `json:"price"  bson:"price"`
	PriceUnit   string        `json:"price_unit"  bson:"price_unit"`
}
