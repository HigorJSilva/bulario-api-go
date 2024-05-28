package types

type SearchMedicinePayload struct {
	NomeProduto string `json:"nomeProduto"`
	NumeroRegistro string `json:"numeroRegistro"`
}

type MedicinesAPIResponse struct {
	Content          []Content   `json:"content"`
	TotalPages       int64       `json:"totalPages"`
	TotalElements    int64       `json:"totalElements"`
	Last             bool        `json:"last"`
	NumberOfElements int64       `json:"numberOfElements"`
	First            bool        `json:"first"`
	Sort             interface{} `json:"sort"`
	Size             int64       `json:"size"`
	Number           int64       `json:"number"`
}

type Content struct {
	IDProduto                   int64  `json:"idProduto"`
	NumeroRegistro              string `json:"numeroRegistro"`
	NomeProduto                 string `json:"nomeProduto"`
	Expediente                  string `json:"expediente"`
	RazaoSocial                 string `json:"razaoSocial"`
	Cnpj                        string `json:"cnpj"`
	NumeroTransacao             string `json:"numeroTransacao"`
	Data                        string `json:"data"`
	NumProcesso                 string `json:"numProcesso"`
	IDBulaPacienteProtegido     string `json:"idBulaPacienteProtegido"`
	IDBulaProfissionalProtegido string `json:"idBulaProfissionalProtegido"`
	DataAtualizacao             string `json:"dataAtualizacao"`
}

type GetMedicinesResponse struct {
	IDProduto                   int64  `json:"idProduto"`
	NumeroRegistro              string `json:"numeroRegistro"`
	NomeProduto                 string `json:"nomeProduto"`
	RazaoSocial                 string `json:"razaoSocial"`
	Data                        string `json:"data"`
}
