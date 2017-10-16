package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/abu-sayed/go-rest/services"
	"github.com/abu-sayed/go-rest/models"
	"strconv"
)

func GetTodo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		resErrors := []models.Error{models.Error{Message: "Invalid todo ID: "+params.ByName("id"), Code: "INVALID_INPUT"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	todo, err := services.GetTodo(id)

	if err != nil {
		resErrors := []models.Error{models.Error{Message: err.Error(), Code: "NOT_FOUND"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	jsonResponse(res, todo, nil)
}

func GetTodos(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

}

func CreateTodo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var todo models.Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		resErrors := []models.Error{models.Error{Message: err.Error(), Code: "INVALID_JSON_REQUEST_BODY"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	if todo.Title == "" {
		resErrors := []models.Error{models.Error{Message: "Todo title field is required", Code: "REQUIRED_FIELD_MISSING"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	todo, err = services.CreateTodo(todo)

	if err != nil {
		resErrors := []models.Error{models.Error{Message: err.Error(), Code: "FAILED"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	jsonResponse(res, "Created Successfully", nil)
}

func UpdateTodo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	var todo models.Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		resErrors := []models.Error{models.Error{Message: err.Error(), Code: "INVALID_JSON_REQUEST_BODY"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	todo.ID, err = strconv.Atoi(params.ByName("id"))

	if todo.Title == "" || todo.ID == 0 {
		resErrors := []models.Error{models.Error{Message: "Todo id, title field are required", Code: "REQUIRED_FIELD_MISSING"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	_, err = services.UpdateTodo(todo)

	if err != nil {
		resErrors := []models.Error{models.Error{Message: err.Error(), Code: "FAILED"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	jsonResponse(res, "Updated Successfully", nil)
}

func DeleteTodo(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		resErrors := []models.Error{models.Error{Message: "Invalid todo ID: "+params.ByName("id"), Code: "INVALID_INPUT"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	_, err = services.DeleteTodo(id)

	if err != nil {
		resErrors := []models.Error{models.Error{Message: err.Error(), Code: "FAILED"}}
		jsonResponse(res, nil, resErrors)
		return
	}

	jsonResponse(res, "Deleted Successfully", nil)
}
