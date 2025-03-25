package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type App struct {
	sync.Mutex
	items  map[int]Item
	nextID int
}

func NewApp() *App {
	app := &App{
		items:  make(map[int]Item),
		nextID: 1,
	}
	app.seedData()
	return app
}

func (a *App) seedData() {
	a.items[a.nextID] = Item{ID: a.nextID, Name: "Apple", Price: 1}
	a.nextID++
	a.items[a.nextID] = Item{ID: a.nextID, Name: "Banana", Price: 2}
	a.nextID++
	a.items[a.nextID] = Item{ID: a.nextID, Name: "Orange", Price: 3}
	a.nextID++
}

func (a *App) getItems(w http.ResponseWriter, r *http.Request) {
	a.Lock()
	defer a.Unlock()

	var list []Item
	for _, item := range a.items {
		list = append(list, item)
	}
	writeJSON(w, list)
}

func (a *App) createItem(w http.ResponseWriter, r *http.Request) {
	a.Lock()
	defer a.Unlock()

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.ID = a.nextID
	a.nextID++
	a.items[item.ID] = item
	writeJSON(w, item)
}

func (a *App) updateItem(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	a.Lock()
	defer a.Unlock()

	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := a.items[id]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	item.ID = id
	a.items[id] = item
	writeJSON(w, item)
}

func (a *App) deleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	a.Lock()
	defer a.Unlock()

	if _, exists := a.items[id]; !exists {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}
	delete(a.items, id)
	w.WriteHeader(http.StatusNoContent)
}

func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid path")
	}
	return strconv.Atoi(parts[2])
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	app := NewApp()
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	}))

	// Routes
	mux.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			app.getItems(w, r)
		case http.MethodPost:
			app.createItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPatch:
			app.updateItem(w, r)
		case http.MethodDelete:
			app.deleteItem(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server running at http://localhost:8000 with seeded data")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
