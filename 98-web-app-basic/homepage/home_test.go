package homepage

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers_HomeHandler(t *testing.T) {
	// type fields struct {
	// 	logger *log.Logger
	// }
	// type args struct {
	// 	w http.ResponseWriter
	// 	r *http.Request
	// }
	tests := []struct {
		name           string
		in             *http.Request
		out            *httptest.ResponseRecorder
		expectedStatus int
		expectedMsg    string
		// fields         fields
		// args           args
	}{
		{
			name:           "test home page",
			in:             httptest.NewRequest("GET", "/", nil),
			out:            httptest.NewRecorder(),
			expectedStatus: http.StatusOK,
			expectedMsg:    "Welcome! to Go World",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandlers(nil)
			h.HomeHandler(tt.out, tt.in)
			if tt.out.Code != tt.expectedStatus {
				t.Errorf("expected status code %v but got %v", tt.expectedStatus, tt.out.Code)
				t.Fail()
			}
			output := tt.out.Body.String()
			if output != tt.expectedMsg {
				t.Errorf("expected output %v but got %v", tt.expectedMsg, output)
				t.Fail()
			}
		})
	}
}
