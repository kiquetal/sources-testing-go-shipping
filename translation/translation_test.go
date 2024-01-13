package translation

import (
	"testing"
)

func TestTranslation(t *testing.T) {
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{Word: "hello", Language: "english", Translation: "Hello"},
		{Word: "hello", Language: "spanish", Translation: "Hola"},
		{Word: "hello", Language: "german", Translation: "Hallo"},
	}
	for _, tc := range tt {
		res := Translation(tc.Word, tc.Language)
		if res != tc.Translation {
			t.Errorf("Translation(%s, %s) = obtained -> %s; wanted-> %s", tc.Word, tc.Language, res, tc.Translation)
		}
	}
}
