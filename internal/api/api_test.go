package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_healthHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Health Handler", args{
			w: httptest.NewRecorder(),
			r: httptest.NewRequest("GET", "/api/v1/health", nil),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			healthHandler(tt.args.w, tt.args.r)
			assert.Equal(t, http.StatusOK, tt.args.w.(*httptest.ResponseRecorder).Code)
			assert.Equal(t, "OK", tt.args.w.(*httptest.ResponseRecorder).Body.String())
		})
	}
}

func Test_appState_getUsersHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		as   appState
		args args
	}{
		// TODO: Add test cases.
		{
			name: "Test Get Users Handler",
			as: appState{
				userService: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.as.getUsersHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_appState_getUserHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		as   appState
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.as.getUserHandler(tt.args.w, tt.args.r)
		})
	}
}
