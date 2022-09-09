package main

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go-movies-crud/handler"
	"go-movies-crud/models"
	"go-movies-crud/repository"

	"net/http"
)

var movies []models.Movie

var PORT int // global config vars

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error occurred reading config: %s", err.Error())
	}

	movies = append(movies,
		models.Movie{
			Id:       "1",
			Isbn:     "543511",
			Title:    "Titapic",
			Director: &models.Director{FirstName: "Aytal", LastName: "Tuprin"},
		},
		models.Movie{
			Id:       "2",
			Isbn:     "111999",
			Title:    "Vinnie da Pooh",
			Director: &models.Director{FirstName: "Timur", LastName: "Petrov"},
		},
	)

	PORT = viper.Get("PORT").(int)

}

func main() {
	r := mux.NewRouter()

	repo := repository.NewRepository()
	handlers := handler.NewHandler(repo)

	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovie).Methods("GET")
	r.HandleFunc("/movie", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf(":%d", PORT),
			r,
		),
	)
}
