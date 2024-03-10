package main

import (
	"encoding/json"
	"internal/database"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleFlashCardCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		CourseID uuid.UUID `json:"courseid"`
		Word     string    `json:"word"`
		Meaning  string    `json:"meaning"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	flashc, err := cfg.DB.CreateFlashCard(r.Context(), database.CreateFlashCardParams{
		ID:       uuid.New(),
		CourseID: params.CourseID,
		Word:     params.Word,
		Meaning:  params.Meaning,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Unable to create flashcard")
		return
	}

	respondWithJSON(w, http.StatusOK, flashc)
}

func (cfg *apiConfig) handleFlashCardDisplay(w http.ResponseWriter, r *http.Request) {
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

	flashes, err := cfg.DB.ViewFlashCard(r.Context(), params.CourseID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't fetch flash cards now")
		return
	}

	respondWithJSON(w, http.StatusOK, flashes)
}
