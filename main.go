package main

//função chamada na aplicaçção inicial
//tem que iniciar o db e mapear as rotas dos endpoints
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaoeman/First-CRUD/config"
	"github.com/joaoeman/First-CRUD/handlers"
	"github.com/joaoeman/First-CRUD/models"
)

func main() {
	dbConnection := config.SetupDataBase() // assim q importa uma função
	//ta iniciando o dbbC
	defer dbConnection.Close()
	_, err := dbConnection.Exec(models.CreateTableSQL) //verifica se a tabela foi criada

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter() //gerenciador de rotas/endpointes da api

	// router.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte("Hello world"))
	// }).Methods("GET") //Ela ta ouvindo
	taskHandler := handlers.NewTaskHandler(dbConnection)
	router.HandleFunc("/tasks", taskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTasks).Methods("POST")
	//path parameter
	// PUT /tasks/26589
	// DELETE /tasks/26589
	router.HandleFunc("/tasks/{id}", taskHandler.RemoveTasks).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTasks).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
