package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/slothmakeout/movies-project/data/dbs"
	"github.com/slothmakeout/movies-project/handlers"
	moviesService "github.com/slothmakeout/movies-project/pkg/movies"
	reviewsService "github.com/slothmakeout/movies-project/pkg/reviews"
)

func main() {
	// Инициализация базы данных
	err := dbs.InitializeDatabaseLayer()
	if err != nil {
		panic("Failed to initialize database")
	}

	l := log.New(os.Stdout, "movies-project-microservice-api ", log.LstdFlags)
	db := dbs.GetDB()

	// Инициализация сервисов
	moviesService := moviesService.GetService(db)
	reviewsService := reviewsService.GetReviewsService(db)

	// Инициализация обработчиков
	moviesHandler := handlers.NewMovies(l, moviesService)
	reviewsHandler := handlers.NewReviews(l, reviewsService)

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/movies", moviesHandler.GetMovies)
	getRouter.HandleFunc("/movies/{id:[0-9]+}", moviesHandler.GetMovieById)
	getRouter.HandleFunc("/reviews", reviewsHandler.GetReviews)
	getRouter.HandleFunc("/reviews/{id:[0-9]+}", reviewsHandler.GetReviewById)
	getRouter.HandleFunc("/movies/{id:[0-9]+}/reviews", reviewsHandler.GetReviewsByMovieId)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/reviews", reviewsHandler.AddReview)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/reviews/{id:[0-9]+}", reviewsHandler.UpdateReview)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/reviews/{id:[0-9]+}", reviewsHandler.DeleteReview)

	// CORS
	corsHandler := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"Content-Type"}),
	)

	s := &http.Server{
		Addr:         "localhost:9090",
		Handler:      corsHandler(sm), // set the default handler
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Создаём канал чтобы принимать сигналы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Запускает сервер в отдельной горутине
	go func() {
		l.Printf("Server listening on %s\n", s.Addr)

		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			l.Fatalf("listen: %s\n", err)
		}
	}()

	// Ждём когда сигнал завершит работу сервера
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// Создаём контекст с таймаутом
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// завершаем работу сервера
	s.Shutdown(tc)
}
