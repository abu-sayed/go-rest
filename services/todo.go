package services

import (
	"github.com/abu-sayed/go-rest/models"
	"github.com/abu-sayed/go-rest/services/db"
	"database/sql"
	"errors"
	"strconv"
)

func GetTodo(id int) (models.Todo, error){
	var todo models.Todo
	err := db.GetDb().QueryRow("SELECT * FROM todo where todo_id = ?", id).Scan(&todo.ID, &todo.Title, &todo.Status, &todo.DeadlineDate)

	if err != nil && err == sql.ErrNoRows{
		return todo, errors.New("Todo not found by id:"+ strconv.Itoa(id))
	}

	return todo, nil
}

func CreateTodo(todo models.Todo) (models.Todo, error){
	stm, err := db.GetDb().Prepare("INSERT INTO todo(title, status, deadline_date) VALUES(?, ?, ?)")
	_, err = stm.Exec(todo.Title, todo.Status, todo.DeadlineDate)

	if err != nil {
		return todo, errors.New("Failed to create todo:"+ err.Error())
	}

	return todo, nil
}

func UpdateTodo(todo models.Todo) (bool, error){
	stm, err := db.GetDb().Prepare("UPDATE todo SET title=?, status=?, deadline_date=? WHERE todo_id=?")
	_, err = stm.Exec(todo.Title, todo.Status, todo.DeadlineDate, todo.ID)

	if err != nil {
		return false, errors.New("Failed to update todo:"+ err.Error())
	}

	return true, nil
}

func DeleteTodo(id int) (bool, error){
	stm, err := db.GetDb().Prepare("DELETE FROM todo WHERE todo_id=?")
	_, err = stm.Exec(id)

	if err != nil {
		return false, errors.New("Failed to delete todo:"+ err.Error())
	}

	return true, nil
}