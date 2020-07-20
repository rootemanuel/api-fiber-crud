package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type PersonEntity struct {
	ObjectID bson.ObjectId `json:"-"bson:"_id,omitempty"`
	Cpf      string        `bson:"cpf" json:"cpf"`
	Nome     string        `bson:"nome" json:"nome"`
}
