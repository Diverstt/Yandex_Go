package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCalc(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedResult float64
		expectedError  string
	}{
		{
			name:           "Valid expression",
			requestBody:    `{"expression": "3+4"}`,
			expectedStatus: http.StatusOK,
			expectedResult: 7,
		},
		{
			name:           "Invalid request body",
			requestBody:    `{"expression": "3+4`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid request body",
		},
		{
			name:           "Invalid expression",
			requestBody:    `{"expression": "3-4+"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedError:  "Expression is not valid",
		},
		{
			name:           "Unknown operator",
			requestBody:    `{"expression": "3^4"}`,
			expectedStatus: http.StatusUnprocessableEntity,
			expectedError:  "Expression is not valid",
		},
		{
			name:           "Division by zero",
			requestBody:    `{"expression": "12/0"}`,
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "Internal server error",
		},
	}

	service := &CalcService{}

	handler := NewCalcHandler(service)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/calculate", bytes.NewBufferString(tt.requestBody))
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			rec := httptest.NewRecorder()

			handler.HandlerCalc(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d; got %d", tt.expectedStatus, rec.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var response map[string]float64
				if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
					t.Fatalf("could not decode response: %v", err)
				}

				if response["result"] != tt.expectedResult {
					t.Errorf("expected result %f; got %f", tt.expectedResult, response["result"])
				}
			}

			if tt.expectedStatus != http.StatusOK {
				if rec.Body.String() != tt.expectedError+"\n" {
					t.Errorf("expected error message %q; got %q", tt.expectedError, rec.Body.String())
				}
			}
		})
	}
}
