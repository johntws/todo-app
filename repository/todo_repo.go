package repository

import (
	"database/sql"
	"todo-app/dto"
	"todo-app/model"
)

var DB *sql.DB

func FindAllTodos(limit int, offset int) ([]model.Todo, error) {
	var todos []model.Todo

	rows, err := DB.Query("SELECT * FROM todo limit ?, ?", limit, offset)

	defer rows.Close()

	if err != nil {
		return todos, err
	}

	for rows.Next() {
		var todo model.Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func FindTodoById(id int64) (model.Todo, error) {
	var todo model.Todo

	rows, err := DB.Query("SELECT * FROM todo where id = ?", id)

	defer rows.Close()

	if err != nil {
		return todo, err
	}

	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return todo, err
		}
	}

	return todo, nil
}

func CreateTodo(todo dto.Todo) (int64, error) {
	result, err := DB.Exec("INSERT INTO todo (title, description) VALUES (?, ?)", todo.Title, todo.Description)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateTodo(tx *sql.Tx, todo model.Todo) (int64, error) {
	result, err := tx.Exec("UPDATE todo set title = ?, description = ? where id = ?", todo.Title, todo.Description, todo.ID)

	if err != nil {
		return 0, err
	}

	id, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteTodo(tx *sql.Tx, id int64) (int64, error) {
	result, err := tx.Exec("DELETE FROM todo where id = ?", id)

	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rows, nil
}
