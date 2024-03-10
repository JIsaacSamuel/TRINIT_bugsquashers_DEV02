package main

import (
	"encoding/json"
	"internal/database"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerTutorCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"username"`
		EmailID  string `json:"emailid"`
		Passcode string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	tuts, err := cfg.DB.CreateTutor(r.Context(), database.CreateTutorParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Emailid:   params.EmailID,
		CreatedAt: time.Now(),
		Passcode:  params.Passcode,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, tuts)
}

func (cfg *apiConfig) handlerTutorAuth(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		EmailID  string `json:"username"`
		Passcode string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	pass, err := cfg.DB.GetTutorPass(r.Context(), params.EmailID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "User not registered")
		return
	}

	if pass.Passcode != params.Passcode {
		respondWithError(w, http.StatusBadRequest, "Incorrect password")
		return
	}

	respondWithJSON(w, http.StatusOK, pass)
}

func (cfg *apiConfig) handleListCourses(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		TutorID uuid.UUID `json:"tutorid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	courses, err := cfg.DB.AllcourseTutor(r.Context(), params.TutorID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't find courses")
		return
	}

	respondWithJSON(w, http.StatusOK, courses)
}

func (cfg *apiConfig) handleListStudents(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CourseID uuid.UUID `json:"courseid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	students, err := cfg.DB.ListStudents(r.Context(), params.CourseID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch students")
		return
	}

	respondWithJSON(w, http.StatusOK, students)
}
