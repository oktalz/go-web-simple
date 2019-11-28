package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/oklog/ulid"
)

type Stats struct {
	Group         string `json:"group"`
	ID            string `json:"id"`
	Timestamp     int64  `json:"timestamp"`
	URL           string `json:"url"`
	RequestsCount int64  `json:"requests-num"`
}

var requestsCount int64 = 0

func requestsCountInc() int64 {
	return atomic.AddInt64(&requestsCount, 1)
}

func main() {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	ulid := ulid.MustNew(ulid.Timestamp(t), entropy).String()
	ulidShort := ulid[len(ulid)-3:]
	fmt.Println(ulid)

	group := os.Getenv("GROUP")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		requestsCountInc()
		fmt.Fprintf(w, "%s\n", ulid)
	})

	http.HandleFunc("/gid", func(w http.ResponseWriter, r *http.Request) {
		requestsCountInc()
		fmt.Fprintf(w, "%s-%s\n", group, ulid)
	})

	http.HandleFunc("/gidc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s-%s-%d\n", group, ulid, requestsCountInc())
	})

	http.HandleFunc("/gids", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s-%s-%d\n", group, ulidShort, requestsCountInc())
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter := requestsCountInc()
		stat := Stats{
			Group:         group,
			ID:            ulid,
			Timestamp:     time.Now().UnixNano(),
			URL:           r.URL.Path,
			RequestsCount: counter,
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
