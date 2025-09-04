package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joaoeman/First-CRUD/models"
)

type TaskHandler struct {
	DB *sql.DB
}

func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
}

//preenchendo o campo da taskhandler, basicamente, está servindo como um construtor

func (TaskHandler *TaskHandler) ReadTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	rows, err := TaskHandler.DB.Query("SELECT * FROM tasks") //vai preencher com todas as linhas
	if err != nil {                                          //se nao houver, aciona o erro
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for rows.Next() { //enquanto houver linhas
		var taks models.Task
		err := rows.Scan(&taks.ID, &taks.Title, &taks.Description, &taks.Status) //vai verificar se todos os campos foram preenchidos, se sim, preenche eles
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, taks)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (taskHandler *TaskHandler) CreateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validação básica
	if task.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// Query correta para PostgreSQL
	query := `
        INSERT INTO tasks (title, description, status) 
        VALUES ($1, $2, $3) 
        RETURNING id
    `

	// Executar a query e obter o ID retornado
	err = taskHandler.DB.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (TaskHandler *TaskHandler) RemoveTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if id == 0 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	result, err := TaskHandler.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (TaskHandler *TaskHandler) UpdateTasks(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	task.ID = id

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if task.ID == 0 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	query := `
		UPDATE tasks 
		SET title = $1, description = $2, status = $3 
		WHERE id = $4
		RETURNING id, title, description, status
	`

	row := TaskHandler.DB.QueryRow(query, task.Title, task.Description, task.Status, task.ID)
	err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err == sql.ErrNoRows {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
