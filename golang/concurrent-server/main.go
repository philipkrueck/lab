package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32 // don't do this

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database struct {
	sync.Mutex
	db map[string]dollars
}

func parseItem(r *http.Request) (string, error) {
	item := r.URL.Query().Get("item")
	if item == "" {
		return "", errors.New("item is empty")
	}
	return item, nil
}

func parseItemPrice(r *http.Request) (string, dollars, error) {
	item, err := parseItem(r)
	if err != nil {
		return "", 0, err
	}

	price := r.URL.Query().Get("price")
	if price == "" {
		return "", 0, errors.New("price is empty")
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		return "", 0, fmt.Errorf("invalid price: %q", price)
	}

	return item, dollars(p), nil
}

func (d *database) list(w http.ResponseWriter, r *http.Request) {
	d.Lock()
	b, err := json.Marshal(d.db)
	d.Unlock()

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func (d *database) create(w http.ResponseWriter, r *http.Request) {
	item, price, err := parseItemPrice(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d.Lock()
	if _, ok := d.db[item]; ok {
		d.Unlock()
		msg := fmt.Sprintf("item %q already exists", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	d.db[item] = price
	d.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "item added"})
}

func (d *database) update(w http.ResponseWriter, r *http.Request) {
	item, price, err := parseItemPrice(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d.Lock()
	if _, ok := d.db[item]; !ok {
		d.Unlock()
		msg := fmt.Sprintf("item %q doesn't exist", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	d.db[item] = price
	d.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "item updated"})
}

func (d *database) delete(w http.ResponseWriter, r *http.Request) {
	item, err := parseItem(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d.Lock()
	if _, ok := d.db[item]; !ok {
		d.Unlock()
		msg := fmt.Sprintf("item %q doesn't exist", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	delete(d.db, item)
	d.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func (d *database) get(w http.ResponseWriter, r *http.Request) {
	item, err := parseItem(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	d.Lock()
	price, ok := d.db[item]
	d.Unlock()

	if !ok {
		msg := fmt.Sprintf("item %q doesn't exist", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"price": price.String()})
}

func runServer() {
	log.Println("Starting server...")

	db := database{
		db: map[string]dollars{
			"cup":  4,
			"sofa": 449,
		},
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/get", db.get)

	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func main() {
	runServer()
}
