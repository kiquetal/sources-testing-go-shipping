package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateApi(t *testing.T) {

	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint: "/hello",

			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:   "/hello?language=german",
			StatusCode: http.StatusOK,

			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedTranslation: "",
			ExpectedLanguage:    "",
		},
	}

	handler := http.HandlerFunc(TranslateHandler)

	for _, tc := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", tc.Endpoint, nil)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != tc.StatusCode {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, tc.StatusCode)
		}
		var resp Resp
		json.Unmarshal(rr.Body.Bytes(), &resp)
		if resp.Language != tc.ExpectedLanguage {
			t.Errorf("handler returned unexpected language: got %v want %v", resp.Language, tc.ExpectedLanguage)
		}
		if resp.Translation != tc.ExpectedTranslation {
			t.Errorf("handler returned unexpected translation: got %v want %v", resp.Translation, tc.ExpectedTranslation)
		}

	}

}
