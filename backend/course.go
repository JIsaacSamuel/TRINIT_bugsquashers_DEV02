package main

import (
	"database/sql"
	"encoding/json"
	"internal/database"
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handleCourseCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		TutorID    uuid.UUID `json:"ID"`
		Language   string    `json:"lang"`
		Price      int32     `json:"price"`
		CourseName string    `json:"coursename"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// tutor, err := cfg.DB.GetTutorCred(r.Context(), params.TutorID)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// if tutor.ID != params.TutorID {
	// 	respondWithError(w, http.StatusBadRequest, "Invalid tutor")
	// 	return
	// }

	course, err := cfg.DB.CreateCourse(r.Context(), database.CreateCourseParams{
		ID:         uuid.New(),
		Langtaught: params.Language,
		TutorID:    params.TutorID,
		Price:      params.Price,
		Takenby:    0,
		Coursename: sql.NullString{String: params.CourseName, Valid: true},
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Course creation failed")
		return
	}

	respondWithJSON(w, http.StatusOK, course)
}

func (cfg *apiConfig) handleSearchCourses(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		// StudentID uuid.UUID `json:"studentid"`
		Language string `json:"lang"`
		MaxPrice int32  `json:"price"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	reqCourses, err := cfg.DB.AllCourseLang(r.Context(), database.AllCourseLangParams{
		Langtaught: params.Language,
		Price:      params.MaxPrice,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "No results")
		return
	}

	respondWithJSON(w, http.StatusOK, reqCourses)
}

func (cfg *apiConfig) handleCourseSubscribing(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		StudentID uuid.UUID `json:"studentid"`
		CourseID  uuid.UUID `json:"courseid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	res, err := cfg.DB.CreateCourseSub(r.Context(), database.CreateCourseSubParams{
		ID:        uuid.New(),
		StudentID: params.StudentID,
		CourseID:  params.CourseID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	respondWithJSON(w, http.StatusOK, res)
}
