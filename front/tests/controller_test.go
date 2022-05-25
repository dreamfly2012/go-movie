package tests

import (
	"net/http"
	"testing"
)

//TestIndex request .
func TestIndex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test index page",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := http.Get("http://127.0.0.1:8080")
			if err != nil {
				t.Errorf("request failed")
			}
			//			if resp.StatusCode != 200 {
			//				t.Error("access index page failed")
			//			}
		})
	}
}
