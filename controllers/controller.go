package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/abu-sayed/go-rest/models"
)

func jsonResponse(res http.ResponseWriter, data interface{}, errors []models.Error)  {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(models.Response{Data:data, Errors: errors});
}