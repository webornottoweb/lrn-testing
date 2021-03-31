package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	tt := []struct {
		name   string
		value  string
		double int
		status int
		err    string
	}{
		{name: "double of two", value: "2", double: 4, status: http.StatusOK},
		{name: "missing value", value: "", status: http.StatusBadRequest, err: "missing value"},
		{name: "not a number", value: "x", status: http.StatusBadRequest, err: "not a number: x"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create a request: %v", err)
			}

			rec := httptest.NewRecorder()

			doubleHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != http.StatusBadRequest {
					t.Errorf("expected status Bad Request; got %v", res.StatusCode)
				}

				if msg := string(bytes.TrimSpace(b)); msg != tc.err {
					t.Errorf("expected message %s; got %s", tc.err, msg)
				}

				return
			}

			if res.StatusCode != tc.status {
				t.Errorf("expected status OK; got %v", res.StatusCode)
			}

			d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
			if err != nil {
				t.Fatalf("expected an integer; got %s", b)
			}

			if d != tc.double {
				t.Fatalf("expected %s to be %d; got %v", tc.name, tc.double, d)
			}
		})
	}
}
