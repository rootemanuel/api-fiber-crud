package service

import (
	"net/http"
	"strings"

	"github.com/rootemanuel/api-fiber-crud/dao"
	"github.com/rootemanuel/api-fiber-crud/dto"
	"github.com/rootemanuel/api-fiber-crud/entity"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber"

	"gopkg.in/mgo.v2"
)

var validate = validator.New()

type PersonService struct{}

func (m *PersonService) GetPersons(c *fiber.Ctx) {

	dao := dao.PersonDao{}

	personsResult, errFindPerson := dao.GetPersons()
	if errFindPerson == mgo.ErrNotFound {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK).JSON(personsResult)
}

func (m *PersonService) GetPerson(c *fiber.Ctx) {

	dao := dao.PersonDao{}
	cpf := c.Params("cpf")

	if cpf == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	personResult, errFindPerson := dao.GetPerson(cpf)
	if errFindPerson == mgo.ErrNotFound {
		c.Status(http.StatusNotFound)
		return
	}

	c.Status(http.StatusOK).JSON(personResult)
}

func (m *PersonService) DeletePerson(c *fiber.Ctx) {

	dao := dao.PersonDao{}
	cpf := c.Params("cpf")

	if cpf == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	errDeletePerson := dao.DeletePerson(cpf)
	if errDeletePerson != nil {

		if errDeletePerson == mgo.ErrNotFound {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (m *PersonService) UpdatePerson(c *fiber.Ctx) {

	dao := dao.PersonDao{}
	cpf := c.Params("cpf")
	req := dto.PersonUpdateReq{}

	if cpf == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	if err := c.BodyParser(&req); err != nil {
		errors := strings.Split(err.Error(), ";")

		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})

		return
	}

	personEntity := entity.PersonEntity{
		Cpf:  cpf,
		Nome: req.Nome,
	}

	errDeletePerson := dao.UpdatePerson(personEntity)
	if errDeletePerson != nil {

		if errDeletePerson == mgo.ErrNotFound {
			c.Status(http.StatusNotFound)
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (m *PersonService) CreatePerson(c *fiber.Ctx) {

	req := dto.PersonCreateReq{}
	dao := dao.PersonDao{}

	if err := c.BodyParser(&req); err != nil {
		errors := strings.Split(err.Error(), ";")

		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})

		return
	}

	personEntity := entity.PersonEntity{
		Cpf:  req.Cpf,
		Nome: req.Nome,
	}

	errCreatePerson := dao.CreatePerson(personEntity)
	if errCreatePerson != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
