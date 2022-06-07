package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func callAPI(t *testing.T, method, url string, body []byte) (int, []byte) {
	t.Helper()

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if err := res.Body.Close(); err != nil {
		t.Fatal(err)
	}

	return res.StatusCode, resBody
}

func TestIbanValidatorAPI(t *testing.T) {
	t.Run("Iban Validator", func(t *testing.T) {
		t.Run("POST /validators/iban", func(t *testing.T) {
			iban := "invalid-number"
			body := []byte(fmt.Sprintf(`{"iban":"%s"}`, iban))
			status, _ := callAPI(t, "POST", "http://localhost:9090/validators/iban", body)
			if status != 400 {
				t.Errorf("expected status 400 but got %d", status)
			}
		})
	})
}
