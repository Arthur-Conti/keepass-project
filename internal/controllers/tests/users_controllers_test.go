	package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Arthur-Conti/keepass-project/internal/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
	settingUpEnvs()
	recorder := httptest.NewRecorder()
	context := getTestGinContext(recorder)
	params := gin.Params{}
	url := url.Values{}
	data := `{
		"first_name": "arthur",
		"last_name": "teste2",
		"email": "arthur_teste2@teste.com",
		"password": "12345"
	}`
	makePostContext(context, params, url, data)
	controllers.NewUserController().CreateUserHandler(context)
	expectedStatusCode := http.StatusCreated
	assert.EqualValues(t, expectedStatusCode, recorder.Code, "error checking status code: " + recorder.Body.String())
	var response controllers.BaseSuccessResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err, "error unmarshaling json: " + recorder.Body.String())
	expectedBody := controllers.BaseSuccessResponse{
		Data: "Created",
		Success: true,
	}
	assert.EqualValues(t, expectedBody, response, "error checking response body: " + recorder.Body.String())
}

func TestListAllUsersHandler(t *testing.T) {
	settingUpEnvs()
	recorder := httptest.NewRecorder()
	context := getTestGinContext(recorder)
	params := gin.Params{}
	url := url.Values{}
	makeGetContext(context, params, url)
	controllers.NewUserController().ListAllUsersHandler(context)
	expectedStatusCode := http.StatusOK
	assert.EqualValues(t, expectedStatusCode, recorder.Code, "error checking status code: " + recorder.Body.String())
	var response controllers.BaseSuccessResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err, "error unmarshaling json: " + recorder.Body.String())
}

func TestGetUserByIDHandler(t *testing.T) {
	settingUpEnvs()
	recorder := httptest.NewRecorder()
	context := getTestGinContext(recorder)
	params := gin.Params{
		{
			Key: "id",
			Value: "ebad42a262a74789cd063009e58f668b",
		},
	}
	url := url.Values{}
	makeGetContext(context, params, url)
	controllers.NewUserController().GetUserByIDHandler(context)
	expectedStatusCode := http.StatusOK
	assert.EqualValues(t, expectedStatusCode, recorder.Code, "error checking status code: " + recorder.Body.String())
	var response controllers.BaseSuccessResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err, "error unmarshaling json: " + recorder.Body.String())
}
