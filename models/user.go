/*
* @Author: souravray
* @Date:   2014-10-20 16:33:21
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 03:38:38
 */

package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Email string        `json:"email" bson:"email"`
	Img   string        `json:"img"  bson:"img,omitempty"`
}

func (user *User) Validator() error {
	if len(user.Email) == 0 {
		return ErrNotFilled
	}
	if !EmailRegexp.MatchString(user.Email) {
		return ErrInvalidEmail
	}
	return nil
}
