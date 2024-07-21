package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Data   string `json:"data"`
	Msg    string `json:"msg"`
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	// پارس کردن فرم
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// دریافت مقادیر فرم
	password := r.FormValue("password")

	// دیباگ: چاپ مقادیر فرم
	log.Println("Username:", r.FormValue("username"), "Password:", password)

	// تنظیم و ارسال پاسخ
	response := Response{
		Status: "true",
		Msg:    "login",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/login", LoginController).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // یا می‌توانید دامنه‌های خاص را مشخص کنید
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(route)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
