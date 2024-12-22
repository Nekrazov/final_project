package application_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/Nekrazov/final_project/app"
	"github.com/Nekrazov/final_project/calc"
)

type RequestBody struct {
	Expression string `json:"expression"`
}
func TestCalcHandler_Success(t *testing.T) {

	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	requestBody := RequestBody{
		Expression: "1+1",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var response application.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	expectedResult := "2.000000"
	if response.Result != expectedResult {
		t.Fatalf("Expected result %s, got %s", expectedResult, response.Result)
	}
}

func TestCalcHandler_InvalidExpression(t *testing.T) {

	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	requestBody := RequestBody{
		Expression: "1+/",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}

	var response application.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calc.ErrInvalidExpression.Error() {
		t.Fatalf("Expected error %v, got %v", calc.ErrInvalidExpression, response.Error)
	}
}

func TestCalcHandler_DivisionByZero(t *testing.T) {
	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	requestBody := RequestBody{
		Expression: "1/0",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}

	var response application.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calc.ErrDivisionByZero.Error() {
		t.Fatalf("Expected error %v, got %v", calc.ErrDivisionByZero, response.Error)
	}
}

func TestCalcHandler_EmptyExpression(t *testing.T) {
	handler := http.HandlerFunc(application.CalcHandler)
	server := httptest.NewServer(handler)
	defer server.Close()

	requestBody := RequestBody{
		Expression: "",
	}
	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Error marshalling request body: %v", err)
	}

	req, err := http.NewRequest("POST", server.URL+"/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Fatalf("Expected status 422, got %d", resp.StatusCode)
	}

	var response application.Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}
	if response.Error != calc.ErrEmptyExpression.Error() {
		t.Fatalf("Expected error %v, got %v", calc.ErrEmptyExpression, response.Error)
	}
}
