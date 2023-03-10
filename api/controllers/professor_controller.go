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

func (server *Server) CreateProfessor(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professor := models.Professor{}
	err = json.Unmarshal(body, &professor)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professor.Prepare()
	err = professor.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professorCreated, err := professor.SaveProfessor(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, professorCreated.ID))
	responses.JSON(w, http.StatusCreated, professorCreated)
}

func (server *Server) GetProfessors(w http.ResponseWriter, r *http.Request) {
	professor := models.Professor{}
	professors, err := professor.GetAllProfessors(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, professors)
}

func (server *Server) GetProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professor := models.Professor{}
	professorFound, err := professor.GetProfessor(server.DB, uint32(pid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, professorFound)
}

func (server *Server) UpdateProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professor := models.Professor{}
	err = json.Unmarshal(body, &professor)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	professor.Prepare()

	updatedProfessor, err := professor.UpdateProfessor(server.DB, uint32(pid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedProfessor)
}

func (server *Server) DeleteProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	professor := models.Professor{}
	rowsAffected, err := professor.DeleteProfessor(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, rowsAffected)
}
