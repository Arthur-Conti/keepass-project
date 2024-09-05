package tests

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func settingUpEnvs() {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}
}

func getTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)
	context, _ := gin.CreateTestContext(recorder)
	context.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{},
	}
	return context
}

func makeGetContext(ginContext *gin.Context, params gin.Params, URL url.Values) {
	ginContext.Request.Method = "GET"
	ginContext.Request.Header.Set("Content-Type", "application/json")
	ginContext.Params = params
	ginContext.Request.URL.RawQuery = URL.Encode()
}

func makePostContext(ginContext *gin.Context, params gin.Params, URL url.Values, data string) {
	ginContext.Request.Method = "GET"
	ginContext.Request.Header.Set("Content-Type", "application/json")
	ginContext.Params = params
	ginContext.Request.URL.RawQuery = URL.Encode()
	body := bytes.NewBuffer([]byte(data))
	ginContext.Request.Body = io.NopCloser(body)
}