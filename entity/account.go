package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type AccountEntity struct {
	ObjectID   bson.ObjectId     `json:"_id,omitempty"bson:"_id,omitempty"`
	Document   string            `bson:"document" json:"document"`
	Name       string            `bson:"name" json:"name"`
	Users      []UserEntity      `bson:"users" json:"users"`
	WareHouses []WarehouseEntity `bson:"warehouses" json:"warehouses"`
	Lastupdate time.Time         `bson:"lastupdate" json:"lastupdate"`
	Register   time.Time         `bson:"register" json:"register"`
	Active     bool              `bson:"active" json:"active"`
}

type UserEntity struct {
	Name       string        `bson:"name" json:"name"`
	Email      string        `bson:"email" json:"email"`
	Password   string        `bson:"password" json:"password"`
	Groups     []GroupEntity `bson:"groups" json:"groups"`
	Lastupdate time.Time     `bson:"lastupdate" json:"lastupdate"`
	Register   time.Time     `bson:"register" json:"register"`
	Active     bool          `bson:"active" json:"active"`
}

type GroupEntity struct {
	Code        string    `bson:"code" json:"code"`
	Description string    `bson:"description" json:"description"`
	Lastupdate  time.Time `bson:"lastupdate" json:"lastupdate"`
	Register    time.Time `bson:"register" json:"register"`
}

type WarehouseEntity struct {
	Code        string    `bson:"code" json:"code"`
	Description string    `bson:"description" json:"description"`
	Lastupdate  time.Time `bson:"lastupdate" json:"lastupdate"`
	Register    time.Time `bson:"register" json:"register"`
}
