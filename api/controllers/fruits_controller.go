package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
)

// GetFruits is
func (server *Server) GetFruits(w http.ResponseWriter, r *http.Request) {

	fruit := models.Fruit{}

	fruits, err := fruit.FindAllFruits(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fruits)
}
