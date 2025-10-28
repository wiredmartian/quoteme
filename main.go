package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Service struct {
	ID          string
	Name        string
	Price       int
	Description string
}

type Category struct {
	Category string
	Services []Service
}

// Global variables to store data
var categories []Category
var serviceMap map[string]Service

func main() {
	// Load once
	data, _ := os.ReadFile("services.json")
	json.Unmarshal(data, &categories)

	// Create a map for quick service lookup by name
	serviceMap = make(map[string]Service)
	for _, category := range categories {
		for _, service := range category.Services {
			serviceMap[service.Name] = service
		}
	}

	http.HandleFunc("/", estimatePage)
	http.HandleFunc("/update", updateTotal)
	http.ListenAndServe(":8080", nil)
}

func estimatePage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, categories)
}

func updateTotal(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	total := 0
	for serviceName := range r.Form {
		if serviceName == "submit" { // Skip the submit button
			continue
		}
		if service, exists := serviceMap[serviceName]; exists {
			total += service.Price
		}
	}
	w.Write([]byte("R" + strconv.Itoa(total)))
}
