package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headers    http.Header
		wantAPIKey string
		wantErr    error
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey abc123"},
			},
			wantAPIKey: "abc123",
			wantErr:    nil,
		},
		{
			name:       "missing authorization header",
			headers:    http.Header{},
			wantAPIKey: "",
			wantErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed header - wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer abc123"},
			},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
		{
			name: "malformed header - no space",
			headers: http.Header{
				"Authorization": []string{"ApiKeyabc123"},
			},
			wantAPIKey: "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, gotErr := GetAPIKey(tt.headers)

			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() apiKey = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}

			if (gotErr == nil && tt.wantErr != nil) || (gotErr != nil && tt.wantErr == nil) || (gotErr != nil && gotErr.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
