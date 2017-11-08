package gerencianet

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockClient struct {
	authorized bool
	DoFunc     func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, errors.New("auth error")
}

func Test_newAuth(t *testing.T) {
	actualResult := newAuth("cidTest", "csTest", true, 10)
	if actualResult.clientID != "cidTest" ||
		actualResult.clientSecret != "csTest" ||
		actualResult.sandbox != true ||
		actualResult.timeout != 10 {
		t.Error("Error on constructor")
	}
}

func Test_getAccessTokenToReturnTokenData(t *testing.T) {

	httpClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("test"))),
			}, nil
		},
	}

	auth := auth{"cidTest", "csTest", true, 10, httpClient}
	_, err := auth.getAccessToken()
	if err != nil {
		t.Error("Expected err to be nil")
	}
}

func Test_getAccessTokenToReturnUnauthorizedError(t *testing.T) {

	httpClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 401,
				Status:     "Unauthorized",
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("test"))),
			}, nil
		},
	}

	auth := auth{"cidTest", "csTest", false, 10, httpClient}
	expected := "Status: Unauthorized Could not authenticate. \nPlease make sure you are using correct credentials and if you are using then in the correct environment."
	_, err := auth.getAccessToken()
	if err.Error() != expected {
		t.Errorf("Error actual = %v, and Expected = %v.", err, expected)
	}
}

func Test_getAccessTokenToReturnHttpRequestError(t *testing.T) {

	httpClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("Test")
		},
	}

	auth := auth{"cidTest", "csTest", true, 10, httpClient}
	_, err := auth.getAccessToken()
	if err == nil {
		t.Errorf("Expected error but got no one")
	}
}
