package server

import (
	"encoding/json"
	"log"
	"net/http"

	"tech-support-platform/internal/models"
)

func (s *Server) handleGetRequests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requests := make([]models.Request, 0)

		query := "SELECT * FROM Requests"

		rows, err := s.database.Query(query)
		if err != nil {
			http.Error(w, "Failed to retrieve data from database", http.StatusInternalServerError)
			log.Println("Retrieveing error:", err)
		}

		for rows.Next() {
			request := models.Request{}

			err := rows.Scan(&request.ID, &request.Title, &request.Description,
				&request.Status)

			if err != nil {
				log.Fatal("Error getting the data from database row: ", err)
			}

			requests = append(requests, request)
		}
	}
}

func (s *Server) handleCreateRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Adding new element to database")

		w.Header().Set("Content-Type", "application/json")
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Ivalid form data", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		description := r.FormValue("description")

		if title == "" || description == "" {
			http.Error(w, "Missiong form fields", http.StatusBadRequest)
			return
		}

		status := 0

		query := "INSERT INTO requests (title, description, status) VALUES (?, ?, ?)"
		_, err = s.database.Exec(query, title, description, status)
		if err != nil {
			http.Error(w, "Failed to insert data into database", http.StatusInternalServerError)
			log.Println("Insert error:", err)
			return
		} else {
			log.Println("New element added to database")
		}

		json.NewEncoder(w).Encode(map[string]string{"status": "Request added successfully"})
	}
}
