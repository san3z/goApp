package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func Handling() {
	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/FindEmployee", findEmployee)
	r.HandleFunc("/promo", promo)

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("frontend/assets/"))))

	fmt.Println("Server is running on port 80")
	http.ListenAndServe(":80", r)
}

func renderHTML(w http.ResponseWriter, r *http.Request, filename string) {
	tmpl, err := template.ParseFiles("frontend/pages/" + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, r)
}

func home(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, r, "Home.html")
}

func findEmployee(w http.ResponseWriter, r *http.Request) {
	renderHTML(w, r, "FindEmployee.html")
}

type PromoRequest struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

type PromoResponse struct {
	Promos []string `json:"promos"`
}

func generatePromos(category string, count int) []string {
	promos := make([]string, count)
	for i := 0; i < count; i++ {
		promos[i] = fmt.Sprintf("PROMO_%s_%d", category, i+1)
	}
	return promos
}

func promo(w http.ResponseWriter, r *http.Request) {
	initLogger()
	logger.Println("------------")
	defer logger.Println("------------")
	logger.Printf("promo function called with method %v", r.Method)
	if r.Method == http.MethodGet {
		logger.Println("rendering HTML 62")
		renderHTML(w, r, "promo.html")
	} else if r.Method == http.MethodPost {
		logger.Println("Processing POST request 65")
		var promoReq PromoRequest
		err := json.NewDecoder(r.Body).Decode(&promoReq)
		if err != nil {
			logger.Println("Error decoding request", err, "69")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		logger.Println("Received category: ", promoReq.Category, "and count: ", promoReq.Count, "73")
		promos := generatePromos(promoReq.Category, promoReq.Count)
		response := PromoResponse{Promos: promos}
		logger.Println("Sending response: ", response)
		json.NewEncoder(w).Encode(response)
	} else {
		logger.Println("Method not allowed. 79")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
