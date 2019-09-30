package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/oklog/ulid"
)

type Stats struct {
	Group     string `json:"group"`
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	URL       string `json:"url"`
}

func main() {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	ulid := ulid.MustNew(ulid.Timestamp(t), entropy).String()
	fmt.Println(ulid)

	group := os.Getenv("GROUP")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s\n", ulid)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		stat := Stats{
			Group:     group,
			ID:        ulid,
			Timestamp: time.Now().UnixNano(),
			URL:       r.URL.Path,
		}

		js, err := json.Marshal(stat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(js)

	})

	_ = http.ListenAndServe(":8181", nil)
}
