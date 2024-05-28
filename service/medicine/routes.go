package medicine

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/higorjsilva/goapi/service/medicineapi"
	"github.com/higorjsilva/goapi/types"
	"github.com/higorjsilva/goapi/utils"
)

type Handler struct{

}

func NewHandler() *Handler{
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/search", h.handleSearch).Methods("POST")
	router.HandleFunc("/side-effects", h.handleSideEffects).Methods("POST")
}


func (h *Handler) handleSearch(w http.ResponseWriter, r *http.Request){
	api := medicineapi.NewApi()

	var payload types.SearchMedicinePayload

	if err := utils.ParseJSON(r, &payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := api.GetMedicines(payload)

	if err!= nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	responses := setMedicinesResponse(response.Content)

	utils.WriteJSON(w, http.StatusAccepted, responses)
	return
}

func (h *Handler) handleSideEffects(w http.ResponseWriter, r *http.Request){
	api := medicineapi.NewApi()

	var payload types.SearchMedicinePayload

	if err := utils.ParseJSON(r, &payload); err != nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := api.GetSideEffects(payload)

	if err!= nil{
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, response)
}

func setMedicinesResponse(contents []types.Content) []types.GetMedicinesResponse {
	var responses []types.GetMedicinesResponse
	for _, content := range contents {
		response := types.GetMedicinesResponse{
			IDProduto:      content.IDProduto,
			NumeroRegistro: content.NumeroRegistro,
			NomeProduto:    content.NomeProduto,
			RazaoSocial:    content.RazaoSocial,
			Data:           content.Data,
		}
		responses = append(responses, response)
	}
	return responses
}
