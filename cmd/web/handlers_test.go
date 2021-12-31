package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestDisplayCoffee(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name string
		urlPath string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/coffee/1", http.StatusOK, []byte("Coffee")},
		{"Non-existent ID", "/coffee/2", http.StatusNotFound, nil},
        {"Negative ID", "/coffee/-1", http.StatusNotFound, nil},
        {"Decimal ID", "/coffee/1.23", http.StatusNotFound, nil},
        {"String ID", "/coffee/foo", http.StatusNotFound, nil},
        {"Empty ID", "/coffee/", http.StatusNotFound, nil},
        {"Trailing slash", "/coffee/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contains %q", tt.wantBody)
			}
		})
	}
}