package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"main/api/formaterror"
	"main/api/models"
	"main/api/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateStudent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	student := models.Student{}
	err = json.Unmarshal(body, &student)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	student.Prepare()
	err = student.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	studentCreated, err := student.SaveStudent(server.DB)

	if err != nil {
		// err.Error() describes itself as a string
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, studentCreated.ID))
	responses.JSON(w, http.StatusCreated, studentCreated)
}

func (server *Server) GetStudents(w http.ResponseWriter, r *http.Request) {
	student := models.Student{}
	students, err := student.GetAllStudents(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, students)
}

func (server *Server) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	student := models.Student{}
	studentFound, err := student.GetStudent(server.DB, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, studentFound)
}

func (server *Server) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	student := models.Student{}
	err = json.Unmarshal(body, &student)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	student.Prepare()

	updatedStudent, err := student.UpdateStudent(server.DB, uint32(sid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedStudent)
}

func (server *Server) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	student := models.Student{}
	rowsAffected, err := student.DeleteStudent(server.DB, uint32(sid))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, rowsAffected)
}
