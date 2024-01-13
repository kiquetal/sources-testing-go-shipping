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
		{Word: "hello", Language: "english", Translation: "hello"},
		{Word: "hello", Language: "spanish", Translation: "hola"},
		{Word: "hello", Language: "german", Translation: "hallo"},
	}
	for _, tc := range tt {
		res := Translation(tc.Word, tc.Language)
		if res != tc.Translation {
			t.Errorf("Translation(%s, %s) = obtained -> %s; wanted-> %s", tc.Word, tc.Language, res, tc.Translation)
		}
	}
}
