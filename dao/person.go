package dao

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/rootemanuel/api-fiber-crud/entity"
)

type PersonDao struct {
	GenericDao
}

const (
	COLL_CRUD = "crud"
)

func (m *PersonDao) CreatePerson(personEntity entity.PersonEntity) error {
	session := GetSession()
	defer session.Close()

	err := session.DB(DB).C(COLL_CRUD).Insert(personEntity)
	return err
}

func (m *PersonDao) GetPerson(cpf string) (*entity.PersonEntity, error) {
	session := GetSession()
	defer session.Close()

	var result entity.PersonEntity
	queryAccount := make(bson.M, 0)

	if cpf != "" {
		queryAccount["cpf"] = cpf
	}

	err := session.DB(DB).C(COLL_CRUD).Find(queryAccount).One(&result)

	return &result, err
}

func (m *PersonDao) GetPersons() ([]entity.PersonEntity, error) {
	session := GetSession()
	defer session.Close()

	result := make([]entity.PersonEntity, 0)
	err := session.DB(DB).C(COLL_CRUD).Find(nil).All(&result)

	return result, err
}

func (m *PersonDao) DeletePerson(cpf string) error {
	session := GetSession()
	defer session.Close()

	selector := bson.M{
		"cpf": cpf,
	}

	return session.DB(DB).C(COLL_CRUD).Remove(selector)
}

func (m *PersonDao) UpdatePerson(personEntity entity.PersonEntity) error {
	session := GetSession()
	defer session.Close()

	selector := bson.M{
		"cpf": personEntity.Cpf,
	}

	updateFields := bson.M{}

	updateFields["nome"] = personEntity.Nome

	update := bson.M{
		"$set": updateFields,
	}

	return session.DB(DB).C(COLL_CRUD).Update(selector, update)
}
