/*
* @Author: souravray
* @Date:   2014-10-20 16:34:23
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 03:39:16
 */

package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Wishes []Wish

type Wish struct {
	//BaseModel
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Product     bson.ObjectId `json:"product_id" bson:"product_id"`
	UserEmail   string        `json:"-" bson:"user_email"`
	Name        string        `json:"name" bson:"name"`
	Img         string        `json:"img"  bson:"img"`
	Description string        `json:"description" bson:"description"`
	Price       string        `json:"price"  bson:"price"`
	PriceUnit   string        `json:"price_unit"  bson:"price_unit"`
	Timestamp   time.Time     `json:"creation_time"  bson:"creation_time,omitempty"`
}

var UserIndex = mgo.Index{
	Key:        []string{"user_email"},
	Unique:     false,
	DropDups:   false,
	Background: true,
	Sparse:     true,
}

func (wish *Wish) Validator() error {
	if len(wish.UserEmail) == 0 {
		return ErrNotFilled
	}
	if !EmailRegexp.MatchString(wish.UserEmail) {
		return ErrInvalidEmail
	}
	if len(wish.Name) == 0 {
		return ErrorProductNameNotFilled
	}
	if len(wish.Img) == 0 {
		return ErrorProductImgNotFilled
	}
	if len(wish.Description) == 0 {
		return ErrorProductDescriptionNotFilled
	}
	if len(wish.Price) == 0 {
		return ErrorProductPriceNotFilled
	}
	if len(wish.PriceUnit) == 0 {
		return ErrorProductPriceUnitNotFilled
	}
	return nil
}
