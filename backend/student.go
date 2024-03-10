package main

import (
	"encoding/json"
	"internal/database"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerStudentCreate(w http.ResponseWriter, r *http.Request) {
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

	stud, err := cfg.DB.CreateStudent(r.Context(), database.CreateStudentParams{
		ID:       uuid.New(),
		Name:     params.Name,
		Emailid:  params.EmailID,
		Passcode: params.Passcode,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create student")
		return
	}

	respondWithJSON(w, http.StatusOK, stud)
}

func (cfg *apiConfig) handlerStudentAuth(w http.ResponseWriter, r *http.Request) {
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

	pass, err := cfg.DB.GetStudentPass(r.Context(), params.EmailID)
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

func (cfg *apiConfig) handleCoursesSubscribed(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		StudentID uuid.UUID `json:"studentid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	courses, err := cfg.DB.ListCoursesSubs(r.Context(), params.StudentID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch courses")
		return
	}

	respondWithJSON(w, http.StatusOK, courses)
}
