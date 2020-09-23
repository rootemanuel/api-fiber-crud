package service

import (
	"net/http"
	"strings"

	"github.com/rootemanuel/api-fiber-crud/dao"
	"github.com/rootemanuel/api-fiber-crud/dto"
	"github.com/rootemanuel/api-fiber-crud/entity"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	"gopkg.in/mgo.v2"
)

var validate = validator.New()

type PersonService struct{}

func (m *PersonService) GetPersons(c *fiber.Ctx) error {

	dao := dao.PersonDao{}

	personsResult, errFindPerson := dao.GetPersons()
	if errFindPerson == mgo.ErrNotFound {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.Status(http.StatusOK).JSON(personsResult)
}

func (m *PersonService) GetPerson(c *fiber.Ctx) error {

	dao := dao.PersonDao{}
	cpf := c.Params("cpf")

	if cpf == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	personResult, errFindPerson := dao.GetPerson(cpf)
	if errFindPerson == mgo.ErrNotFound {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.Status(http.StatusOK).JSON(personResult)
}

func (m *PersonService) DeletePerson(c *fiber.Ctx) error {

	dao := dao.PersonDao{}
	cpf := c.Params("cpf")

	if cpf == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	errDeletePerson := dao.DeletePerson(cpf)
	if errDeletePerson != nil {

		if errDeletePerson == mgo.ErrNotFound {
			return c.SendStatus(http.StatusNotFound)
		}

		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (m *PersonService) UpdatePerson(c *fiber.Ctx) error {

	dao := dao.PersonDao{}
	cpf := c.Params("cpf")
	req := dto.PersonUpdateReq{}

	if cpf == "" {
		return c.SendStatus(http.StatusBadRequest)
	}

	if err := c.BodyParser(&req); err != nil {
		errors := strings.Split(err.Error(), ";")

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	personEntity := entity.PersonEntity{
		Cpf:  cpf,
		Nome: req.Nome,
	}

	errDeletePerson := dao.UpdatePerson(personEntity)
	if errDeletePerson != nil {

		if errDeletePerson == mgo.ErrNotFound {
			return c.SendStatus(http.StatusNotFound)
		}

		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusNoContent)
}

func (m *PersonService) CreatePerson(c *fiber.Ctx) error {

	req := dto.PersonCreateReq{}
	dao := dao.PersonDao{}

	if err := c.BodyParser(&req); err != nil {
		errors := strings.Split(err.Error(), ";")

		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	personEntity := entity.PersonEntity{
		Cpf:  req.Cpf,
		Nome: req.Nome,
	}

	errCreatePerson := dao.CreatePerson(personEntity)
	if errCreatePerson != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusCreated)
}

func (m *PersonService) Ping(c *fiber.Ctx) error {
	return c.SendString("0WN3D 👋!")
}
