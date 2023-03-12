package modeltests

import (
	"errors"
	"log"
	"main/api/models"
	"main/helpers"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetAllUsers(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.SeedUsers(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.GetAllUsers(server.DB)

	if err != nil {
		t.Errorf("Error getting all users: %v\n", err)
		return
	}

	assert.Equal(t, len(*users), 3)
}

func TestGetUser(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.SeedUsers(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	user, err := helpers.SeedOneUser(server.DB)

	if err != nil {
		log.Fatal(err)
	}

	foundUser, err := userInstance.GetUser(server.DB, user.ID)

	if err != nil {
		t.Errorf("Error getting user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
}

func TestSaveUser(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		Name: models.Name{
			First: "Mickey",
			Last:  "Mouse",
		},
		Email:    "mickey@example.com",
		Password: "password",
	}

	savedUser, err := newUser.SaveUser(server.DB)

	if err != nil {
		t.Errorf("Error saving user: %v", err)
		return
	}

	assert.Equal(t, savedUser.Email, newUser.Email)
	assert.Equal(t, savedUser.Name.First, newUser.Name.First)
	assert.Equal(t, savedUser.Name.Last, newUser.Name.Last)
}

func TestDeleteUser(t *testing.T) {
	err := helpers.RefreshUserTable(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.SeedUsers(server.DB)
	if err != nil {
		log.Fatal(err)
	}

	user, err := helpers.SeedOneUser(server.DB)

	if err != nil {
		log.Fatal(err)
	}

	numRowsDeleted, err := user.DeleteUser(server.DB, user.ID)

	if err != nil {
		t.Errorf("Error deleting user %v: %v", user.ID, err)
		return
	}

	assert.Equal(t, numRowsDeleted, int64(1))

	_, err = userInstance.GetUser(server.DB, user.ID)

	assert.Equal(t, err, errors.New("User Not Found"))
}
