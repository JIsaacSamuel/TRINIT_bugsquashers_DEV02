package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"internal/database"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	PORT := "5000"
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("CONN")
	if dbURL == "" {
		log.Fatal("Unable to load dbURL")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	subr1 := chi.NewRouter()
	// signaling
	subr1.Post("/studentsignup", apiCfg.handlerStudentCreate)
	subr1.Post("/tutorsignup", apiCfg.handlerTutorCreate)
	subr1.Post("/auth/stulogin", apiCfg.handlerStudentAuth)
	subr1.Post("/auth/tutorlogin", apiCfg.handlerTutorAuth)
	subr1.Post("/tutor/createcourse", apiCfg.handleCourseCreate)
	subr1.Post("/tutor/listcourse", apiCfg.handleListCourses)
	subr1.Post("/tutor/getstudents", apiCfg.handleListStudents)
	subr1.Post("/tutor/createflashcard", apiCfg.handleFlashCardCreate)
	subr1.Post("/student/listcourses", apiCfg.handleCoursesSubscribed)
	subr1.Post("/student/requiredcourse", apiCfg.handleSearchCourses)
	subr1.Post("/student/coursesubsribing", apiCfg.handleCourseSubscribing)
	subr1.Post("/student/viewflashcards", apiCfg.handleFlashCardDisplay)
	router.Mount("/api", subr1)

	srv := http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(srv.ListenAndServe())
}

// helper functions
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
