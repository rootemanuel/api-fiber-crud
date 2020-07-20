package dto

type PersonCreateReq struct {
	Cpf  string `json:"cpf" validate:"required"`
	Nome string `json:"nome" validate:"required"`
}

type PersonUpdateReq struct {
	Nome string `json:"nome" validate:"required"`
}
