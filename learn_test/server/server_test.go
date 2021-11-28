package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestDoubleHandler(t *testing.T) {
	testCases := []struct {
		name string
		intput string
		result int
		status int
		err string
	}{
		{name: "double of two", intput: "2", result: 4, status: http.StatusOK, err: ""},
		{name: "double of nine", intput: "9", result: 18, status: http.StatusOK, err: ""},
		{name: "double of nil", intput: "", status: http.StatusBadRequest, err: "missing value"},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t* testing.T){

			request, err := http.NewRequest(http.MethodGet, "/localhost:4000/double?v=" + testCase.intput, nil)
			if err != nil {
				t.Fatalf("could not create a new request: %v, err : %v", request, err)
			}

			rec := httptest.NewRecorder()
			doubleHandler(rec, request)
			res := rec.Result()

			if res.StatusCode != testCase.status {
				t.Errorf("received status code %d, expect %d", res.StatusCode, testCase.status)
				return
			}

			respBytes, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("cannot read all from the response body, err : %v", err)
			}
			defer res.Body.Close()

			trimedResult := strings.TrimSpace(string(respBytes))
			if res.StatusCode !=  http.StatusOK {
				if trimedResult != testCase.err {
					// check the error message
					t.Errorf("received err msg %s, expect %s", trimedResult, testCase.err)
					return
				}
			}
			// compare the returned value

			doubleVal, err := strconv.Atoi(trimedResult)
			if err != nil {
				t.Errorf("cannot convert response body to int, err : %v", err)
				return
			}

			if doubleVal != testCase.result {
				t.Errorf("received result %d, expect %d",doubleVal, testCase.result)
				return
			}
		})
	}

}
