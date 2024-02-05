package web_service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "path/filepath"
	"html/template"

	cacher "github.com/nats-io/go-nats-streaming/services/cacher"
)

var cache *cacher.Cache
var tmpl *template.Template

func WebService(myCache *cacher.Cache) {
	cache = myCache
	tmpl = template.Must(template.ParseFiles("web/templates/index.html"))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/order", getOrderHandler)
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	order, ok := cache.Get(orderID)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
