package httpWorker

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/taxes/", handle)

	log.Printf("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
