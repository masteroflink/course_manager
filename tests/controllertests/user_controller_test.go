package controllertests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"main/api/models"
	"main/helpers"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/assert.v1"
)

type Name struct {
	first  string
	middle string
	last   string
}

func TestCreateUser(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	samples := []struct {
		inputJSON    string
		statusCode   int
		name         Name
		email        string
		errorMessage string
	}{
		{
			inputJSON:  `{"name": {"first": "Mickey", "middle": "", "last": "Mouse"}, "email": "mickey@example.com", "password": "password"}`,
			statusCode: 201,
			name: Name{
				first:  "Mickey",
				middle: "",
				last:   "Mouse",
			},
			email:        "mickey@example.com",
			errorMessage: "",
		},
		{
			inputJSON:  `{"name": {"first": "Frank", "middle": "", "last": "Burt"}, "email": "mickey@example.com", "password": "password"}`,
			statusCode: 500,
			name: Name{
				first:  "Frank",
				middle: "",
				last:   "Burt",
			},
			email:        "mickey@example.com",
			errorMessage: "Email Already Taken",
		},
		{
			inputJSON:  `{"name": {"first": "Mickey", "middle": "", "last": "Mouse"}, "email": "example.com", "password": "password"}`,
			statusCode: 422,
			name: Name{
				first:  "Mickey",
				middle: "",
				last:   "Mouse",
			},
			email:        "example.com",
			errorMessage: "Invalid Email",
		},
		{
			inputJSON:  `{"name": {"first": "Mickey", "middle": "", "last": "Mouse"}, "email": "", "password": "password"}`,
			statusCode: 422,
			name: Name{
				first:  "Mickey",
				middle: "",
				last:   "Mouse",
			},
			email:        "",
			errorMessage: "Required Email",
		},
		{
			inputJSON:  `{"name": {"first": "Mickey", "middle": "", "last": "Mouse"}, "email": "mickey@example.com", "password": "}`,
			statusCode: 422,
			name: Name{
				first:  "Mickey",
				middle: "",
				last:   "Mouse",
			},
			email:        "mickey@example.com",
			errorMessage: "Required password",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/users", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("this is the error: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.CreateUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}
		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 201 {
			assert.Equal(t, responseMap["email"], v.email)
		}
		if v.statusCode == 422 || v.statusCode == 500 && v.errorMessage != "" {
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}

func TestGetUsers(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}
	err = helpers.SeedUsers(server.DB)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.GetUsers)
	handler.ServeHTTP(rr, req)

	var users []models.User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		log.Fatalf("Cannot convert to json: %v", err)
	}
	assert.Equal(t, rr.Code, http.StatusOK)
	assert.Equal(t, len(users), 2)
}

func TestGetUserByID(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}
	user, err := helpers.SeedOneUser(server.DB)
	if err != nil {
		log.Fatal(err)
	}
	userSample := []struct {
		id           string
		statusCode   int
		name         Name
		email        string
		errorMessage string
	}{
		{
			id:         strconv.Itoa(int(user.ID)),
			statusCode: 200,
			name: Name{
				first:  "Mickey",
				middle: "",
				last:   "Mouse",
			},
			email: "mickey@example.com",
		},
		{
			id:         "unknown",
			statusCode: 422,
		},
	}

	for _, v := range userSample {
		req, err := http.NewRequest("GET", "/users", nil)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		req = mux.SetURLVars(req, map[string]string{"id": v.id})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.GetUser)
		handler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			log.Fatalf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, v.statusCode)

		if v.statusCode == 200 {
			assert.Equal(t, user.Name.First, responseMap["name"].(map[string]interface{})["first"])
			assert.Equal(t, user.Name.Last, responseMap["name"].(map[string]interface{})["last"])
			assert.Equal(t, user.Email, responseMap["email"])
		}
	}
}
