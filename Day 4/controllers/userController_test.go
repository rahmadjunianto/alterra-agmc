package controllers

import (
	"day4/config"
	"day4/lib/database/seeder"
	"day4/mocks"
	"day4/models"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

var (
	echoMock = mocks.EchoMock{E: echo.New()}
)

func setupTest(t *testing.T) {
	//load env
	if err := godotenv.Load("../.env"); err != nil {
		t.Error("Error3 loading .env file")
	}

	//setup database
	config.InitDB()

	// clear database
	s := seeder.NewUserSeeder()
	s.Delete()
	s.Seed()

}

func TestLoginUsersControllersSuccess(t *testing.T) {
	setupTest(t)

	//create json body
	body := models.Users{
		Email:    "user1@mail.com",
		Password: "user1",
	}

	//setup request
	b, _ := json.Marshal(body)
	c, rec := echoMock.RequestMock(http.MethodGet, "/", strings.NewReader(string(b)))

	asserts := assert.New(t)
	if asserts.NoError(Login(c)) {
		asserts.Equal(200, rec.Code)
		responseBody := rec.Body.String()
		result := map[string]interface{}{}
		assert.NoError(t, json.Unmarshal([]byte(responseBody), &result))
		data := result["data"].(map[string]interface{})
		assert.Equal(t, true, result["success"])
		assert.NotNil(t, data["name"])
		assert.NotEmpty(t, data["name"])
		assert.Equal(t, body.Email, data["email"])
		assert.NotNil(t, data["token"])
		assert.NotEmpty(t, data["token"])
	}
}
func TestLoginUsersControllersInvalid(t *testing.T) {
	setupTest(t)
	//create json body
	body := models.Users{
		Email:    "test1@mail.com",
		Password: "test2",
	}

	//setup request
	b, _ := json.Marshal(body)

	c, _ := echoMock.RequestMock(http.MethodGet, "/", strings.NewReader(string(b)))
	//test
	err := Login(c)
	response, ok := err.(*echo.HTTPError)
	assert.True(t, ok)
	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "these credentials do not match our records", response.Message)
}
func TestGetUser(t *testing.T) {
	setupTest(t)
	c, rec := echoMock.RequestMock(http.MethodGet, "/v1/users", nil)
	//test
	assert.NoError(t, GetUser(c))
	fmt.Println(rec.Code)
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetUserById(t *testing.T) {
	setupTest(t)
	c, rec := echoMock.RequestMock(http.MethodGet, "/", nil)
	//set params
	c.SetParamNames("id")
	c.SetParamValues("1")
	assert.NoError(t, GetUserById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestGetUserByIdNotFound(t *testing.T) {
	setupTest(t)
	c, _ := echoMock.RequestMock(http.MethodGet, "/", nil)
	//set params
	c.SetParamNames("id")
	c.SetParamValues("3")
	assert.Error(t, GetUserById(c))
}
func TestCreateUserSuccess(t *testing.T) {

	setupTest(t)
	//create json body
	body := models.Users{
		Email:    "test3@mail.com",
		Password: "test3",
		Name:     "test3",
	}
	//setup request
	b, _ := json.Marshal(body)
	c, rec := echoMock.RequestMock(http.MethodPost, "/", strings.NewReader(string(b)))
	assert.NoError(t, CreateUser(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestCreateUserInvalid(t *testing.T) {

	setupTest(t)
	//create json body
	body := models.Users{
		Email:    "test3@mail.com",
		Password: "test3",
	}
	//setup request
	b, _ := json.Marshal(body)
	c, _ := echoMock.RequestMock(http.MethodPost, "/", strings.NewReader(string(b)))
	assert.Error(t, CreateUser(c))
}
func TestUpdateUserById(t *testing.T) {
	setupTest(t)

	//create json body
	body := models.Users{
		Email:    "test3@mail.com",
		Password: "test3",
		Name:     "test3",
	}
	//setup request

	//create json body
	bodyLogin := models.Users{
		Email:    "user1@mail.com",
		Password: "user1",
	}
	//login
	bl, _ := json.Marshal(bodyLogin)
	c, rec := echoMock.RequestMock(http.MethodGet, "/", strings.NewReader(string(bl)))

	asserts := assert.New(t)
	var token string
	if asserts.NoError(Login(c)) {
		asserts.Equal(200, rec.Code)
		responseBody := rec.Body.String()
		result := map[string]interface{}{}
		assert.NoError(t, json.Unmarshal([]byte(responseBody), &result))
		data := result["data"].(map[string]interface{})
		assert.Equal(t, true, result["success"])
		assert.NotNil(t, data["name"])
		assert.NotEmpty(t, data["name"])
		assert.Equal(t, bodyLogin.Email, data["email"])
		assert.NotNil(t, data["token"])
		assert.NotEmpty(t, data["token"])
		token = fmt.Sprintf("%v", data["token"])
	}

	//setup request
	b, _ := json.Marshal(body)
	c, rec = echoMock.RequestMock(http.MethodPost, "/", strings.NewReader(string(b)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.SetParamNames("id")
	c.SetParamValues("1")
	assert.NoError(t, UpdateUserById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestUpdateUserByIdUnauthorizedUser(t *testing.T) {
	setupTest(t)

	//create json body
	body := models.Users{
		Email:    "test3@mail.com",
		Password: "test3",
		Name:     "test3",
	}
	//setup request

	//create json body
	bodyLogin := models.Users{
		Email:    "user1@mail.com",
		Password: "user1",
	}
	//login
	bl, _ := json.Marshal(bodyLogin)
	c, rec := echoMock.RequestMock(http.MethodGet, "/", strings.NewReader(string(bl)))

	asserts := assert.New(t)
	var token string
	if asserts.NoError(Login(c)) {
		asserts.Equal(200, rec.Code)
		responseBody := rec.Body.String()
		result := map[string]interface{}{}
		assert.NoError(t, json.Unmarshal([]byte(responseBody), &result))
		data := result["data"].(map[string]interface{})
		assert.Equal(t, true, result["success"])
		assert.NotNil(t, data["name"])
		assert.NotEmpty(t, data["name"])
		assert.Equal(t, bodyLogin.Email, data["email"])
		assert.NotNil(t, data["token"])
		assert.NotEmpty(t, data["token"])
		token = fmt.Sprintf("%v", data["token"])
	}

	//setup request
	b, _ := json.Marshal(body)
	c, rec = echoMock.RequestMock(http.MethodPost, "/", strings.NewReader(string(b)))
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.SetParamNames("id")
	c.SetParamValues("2")
	assert.Error(t, UpdateUserById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
func TestDeleteUserById(t *testing.T) {
	setupTest(t)
	//setup request
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")
	assert.NoError(t, DeleteUserById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteUserByIdInvalid(t *testing.T) {
	setupTest(t)
	//setup request
	c, rec := echoMock.RequestMock(http.MethodDelete, "/", nil)
	c.SetParamNames("id")
	c.SetParamValues("3")
	assert.Error(t, DeleteUserById(c))
	assert.Equal(t, http.StatusOK, rec.Code)
}
