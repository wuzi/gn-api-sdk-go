package gerencianet

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type MockAuth struct {
	authorized bool
}

func (a *MockAuth) getAccessToken() (authResponseBody, error) {
	var response authResponseBody
	if a.authorized == false {
		return response, errors.New("401 error")
	}
	response.AccessToken = "token_test"
	return response, nil
}

func Test_newRequestToReturnSandboxRequester(t *testing.T) {
	requester := newRequester("cidTest", "csTest", true, 10)
	if requester.url != UrlSandbox ||
		requester.timeout != 10 {
		t.Error("Error on constructor")
	}
}

func Test_newRequestToReturnProductionRequester(t *testing.T) {
	requester := newRequester("cidTest", "csTest", false, 10)
	if requester.url != UrlProduction ||
		requester.timeout != 10 {
		t.Error("Error on constructor")
	}
}

func Test_requestToReturnString(t *testing.T) {

	httpClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(req.URL.String()))),
			}, nil
		},
	}

	auth := auth{"cidTest", "csTest", true, 10, httpClient}
	requester := requester{auth, "sandbox.com", 10, "token", time.Time{}, httpClient}
	params := map[string]string{"id": "1", "token": "2", "data": "test"}
	requestBody := map[string]interface{} {"test": "test"}

	res, err := requester.request("/test/:id", "POST", params, requestBody)
	if err != nil {
		t.Error("Expected err to be nil")
	}
	if res != "sandbox.com/test/1?token=2&data=test" {
		t.Error("Expected parse params on route")
	}
}

func Test_requestToReturnAuthError(t *testing.T) {

	httpClient := &MockClient{}
	auth := auth{"cidTest", "csTest", true, 10, httpClient}
	requester := requester{auth, "sandbox.com", 10, "token", time.Time{}, httpClient}
	_, err := requester.request("/test", "GET", nil, nil)
	if err.Error() != "auth error" {
		t.Error("Expected err to be an auth error")
	}
}


func Test_requestToReturnServerError(t *testing.T) {
	authClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("auth test"))),
			}, nil
		},
	}
	requestClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("erro 500 test"))),
			}, nil
		},
	}
	auth := auth{"cidTest", "csTest", true, 10, authClient}
	requester := requester{auth, "sandbox.com", 10, "token", time.Time{}, requestClient}
	_, err := requester.request("test", "PUT", nil, nil)
	if err.Error() != "erro 500 test" {
		t.Error("Expected err to be an 500 http error")
	}
}

func Test_requestToReturnHttpDoError(t *testing.T) {
	authClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte("auth test"))),
			}, nil
		},
	}
	requestClient := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("http Do error")
		},
	}
	auth := auth{"cidTest", "csTest", true, 10, authClient}
	requester := requester{auth, "sandbox.com", 10, "token", time.Time{}, requestClient}
	_, err := requester.request("test", "PUT", nil, nil)
	if err.Error() != "http Do error" {
		t.Error("Expected err to be an http/net error")
	}
}
