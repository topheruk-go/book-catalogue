package googleapi

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/topheruk-go/util/http/fetch"
	"gotest.tools/v3/assert"
)

func TestGoogleBooksAPI(t *testing.T) {
	// https://www.googleapis.com/books/v1/volumes?q=isbn:9780141349978
	isbn := "9780141349978"

	var books Books
	err := fetch.Fetch("https://www.googleapis.com/books/v1/volumes?q=isbn:"+isbn, func(resp *http.Response) error {
		return json.NewDecoder(resp.Body).Decode(&books)
	}, fetch.DefaultOptions)
	assert.Assert(t, err)

	assert.Equal(t, books.Items[0].VolumeInfo.Title, "Fantastic Mr Fox")
}
