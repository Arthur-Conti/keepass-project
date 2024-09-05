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

func TestCreateVaultHandler(t *testing.T) {
	settingUpEnvs()
	recorder := httptest.NewRecorder()
	context := getTestGinContext(recorder)
	params := gin.Params{}
	url := url.Values{}
	body := `{
		"user_id": "ebad42a262a74789cd063009e58f668b",
		"name": "vault_test",
		"description": "A vault to test",
		"password": "12345" 
	}`
	makePostContext(context, params, url, body)
	controllers.NewVaultsController().CreateVaultHandler(context)
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

func TestListAllVaultsByUserID(t *testing.T) {
	settingUpEnvs()
	recorder := httptest.NewRecorder()
	context := getTestGinContext(recorder)
	params := gin.Params{
		{
			Key: "user_id",
			Value: "ebad42a262a74789cd063009e58f668b",
		},
	}
	url := url.Values{}
	makeGetContext(context, params, url)
	controllers.NewVaultsController().ListAllVaultsByUserIDHandler(context)
	expectedStatusCode := http.StatusOK
	assert.EqualValues(t, expectedStatusCode, recorder.Code, "error checking status code: " + recorder.Body.String())
	var response controllers.BaseSuccessResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err, "error unmarshaling json: " + recorder.Body.String())
}

func TestGetVaultByIDHandler(t *testing.T) {
	settingUpEnvs()
	recorder := httptest.NewRecorder()
	context := getTestGinContext(recorder)
	params := gin.Params{
		{
			Key: "id",
			Value: "a5cd8ac17c36382eaab5fb30e395a7e2",
		},
	}
	url := url.Values{}
	makeGetContext(context, params, url)
	controllers.NewVaultsController().GetVaultByIDHandler(context)
	expectedStatusCode := http.StatusOK
	assert.EqualValues(t, expectedStatusCode, recorder.Code, "error checking status code: " + recorder.Body.String())
	var response controllers.BaseSuccessResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err, "error unmarshaling json: " + recorder.Body.String())
}