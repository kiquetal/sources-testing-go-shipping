package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubbedService struct{}

func (s *stubbedService) Translate(word string, language string) string {
	if word == "foo" {
		return "bar"
	}
	return ""
}

func TestTranslateApi(t *testing.T) {

	tt := []struct { // <1>
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/foo",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/foo?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/baz",
			StatusCode:          404,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
		{
			Endpoint:            "/foo?language=GerMan",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "bar",
		},
	}

	underTest := NewTranslateHandler(&stubbedService{})
	handler := http.HandlerFunc(underTest.TranslateHandler)

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
