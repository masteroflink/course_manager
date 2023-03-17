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

func (server *Server) CreateCourse(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course := models.Course{}
	err = json.Unmarshal(body, &course)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course.Prepare()
	err = course.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	courseCreated, err := course.SaveCourse(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, courseCreated.ID))
	responses.JSON(w, http.StatusCreated, courseCreated)
}

func (server *Server) GetCourses(w http.ResponseWriter, r *http.Request) {
	course := models.Course{}
	courses, err := course.GetAllCourses(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, courses)
}

func (server *Server) GetCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course := models.Course{}
	courseFound, err := course.GetCourse(server.DB, uint32(cid))

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, courseFound)
}

func (server *Server) UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course := models.Course{}
	err = json.Unmarshal(body, &course)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	course.Prepare()

	updatedProfessor, err := course.UpdateCourse(server.DB, uint32(cid))

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	responses.JSON(w, http.StatusOK, updatedProfessor)
}

func (server *Server) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}
	rowsAffected, err := course.DeleteCourse(server.DB, uint32(cid))
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	responses.JSON(w, http.StatusOK, rowsAffected)
}

func (server *Server) EnrollStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["cid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sid, err := strconv.ParseUint(vars["sid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}
	enrolledCourse, err := course.EnrollStudent(server.DB, uint32(cid), uint32(sid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, enrolledCourse)
}

func (server *Server) RemoveStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["cid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	sid, err := strconv.ParseUint(vars["sid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}
	enrolledCourse, err := course.RemoveStudent(server.DB, uint32(cid), uint32(sid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, enrolledCourse)
}

func (server *Server) AssignProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["cid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	pid, err := strconv.ParseUint(vars["pid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}
	enrolledCourse, err := course.AssignProfessor(server.DB, uint32(cid), uint32(pid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, enrolledCourse)
}

func (server *Server) RemoveProfessor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid, err := strconv.ParseUint(vars["cid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	pid, err := strconv.ParseUint(vars["pid"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}
	enrolledCourse, err := course.RemoveProfessor(server.DB, uint32(cid), uint32(pid))

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, enrolledCourse)
}
