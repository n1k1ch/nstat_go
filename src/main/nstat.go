package nstat

import (
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
	http.HandleFunc("/api/v1/entries", Entries)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./web/dist"))))
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
	log.Println(entry)
	entry.Date = time.Now()

	key := datastore.NewIncompleteKey(ctx, "Entry", nil)
	if _, err := datastore.Put(ctx, key, &entry); err != nil {
		_handleError(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Location", "/api/v1/entries/" + key.StringID())
}

func retreiveEntries(w http.ResponseWriter, r *http.Request, ctx context.Context) {
	entries := make([]Entry, 0, 100)

	q := datastore.NewQuery("Entry").Limit(100)
	if _, err := q.GetAll(ctx, &entries); err != nil {
		_handleError(w, err, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&entries)
}

func _handleError(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	w.Write([]byte(err.Error()))
}