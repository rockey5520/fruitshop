package controllers

import (
	"fruitshop/server/api/models"
	"net/http"

	"github.com/gorilla/mux"
)

// FindFruit returns fruit details if fruit exists in the store
// @Summary Show details of all fruits
// @Description Get details of all fruits
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Fruit
// @Router /server/api/v1/fruits [get]
func (server *Server) FindFruit(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name, err := vars["name"]
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	fruit := models.Fruit{}
	fruitFetched, err := fruit.FindFruitByName(server.DB, name)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, fruitGotten)
}

// FindFruits will retuen all fruits exists within the fruitshop
// @Summary Show details of a fruit item
// @Description Get details of a fruit item
// @Accept  json
// @Produce  json
// @Param name path string true "Fruit identifier"
// @Success 200 {object} models.Fruit
// @Failure 400 {string} string "Bad input"
// @Router /server/api/v1/fruits/{name} [get]
func (server *Server) FindFruits(w http.ResponseWriter, r *http.Request) {

	fruit := models.Fruit{}

	fruits, err := fruit.FindAllFruits(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, fruits)
}
