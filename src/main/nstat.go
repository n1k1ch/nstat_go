package nstat

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
	"time"
	"golang.org/x/net/context"
	"errors"
)

type Entry struct {
	Description string
	Amount      float32
	Currency    string
	Date        time.Time
}

func init() {
	log.Printf("%q", "Hallo")
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/v1/entries", Entries)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%q", "handler")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hie")
}

func Entries(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	switch r.Method {
		case http.MethodGet: retreiveEntries(w, r, c)
		case http.MethodPut: createEntry(w, r, c)
		default:
			_handleError(w, errors.New("Unsupported method"), http.StatusBadRequest)
	}
}

func createEntry(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	var entry Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err == io.EOF {
		_handleError(w, err, http.StatusBadRequest)
	}

	key := datastore.NewIncompleteKey(ctx, "Entry", nil)
	if _, err := datastore.Put(ctx, key, entry); err != nil {
		_handleError(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Location", "/api/v1/entries/" + key.StringID())
}

func retreiveEntries(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	/*if jentries, err := ; err != nil {
		_handleError(w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jentries)
	}*/

	//entries := make([]Entry, 0)

	//q := datastore.NewQuery("Entry")
	//q.GetAll(ctx, &entries)
}

func _handleError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}