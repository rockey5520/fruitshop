package controllers

import (
	"net/http"

	"fruitshop/api/models"
	"fruitshop/api/responses"
)

// GetFruits fetches all fruits from the fruit table which is meta table where data loaded during application start-up
func (server *Server) GetFruits(w http.ResponseWriter, r *http.Request) {

	fruit := models.Fruit{}

	fruits, err := fruit.FindAllFruits(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fruits)
}
