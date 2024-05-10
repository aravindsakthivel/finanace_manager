package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	envPkg "github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	envPkg.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := chi.NewRouter()

	var srv *http.Server = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Printf("Server started on port %s", port)
	var err = srv.ListenAndServe()

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
