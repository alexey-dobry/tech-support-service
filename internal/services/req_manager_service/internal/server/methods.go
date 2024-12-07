package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.Session{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			log.Println("error: Could not bind json")
			return
		}

		query := "SELECT * FROM sessions WHERE clientid=0"

		data, err := s.database.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			log.Printf("error: could not querry data from database, errormsg: %s", err)
			return
		}

		DataFromDB := models.Session{}

		data.Next()
		err = data.Scan(&DataFromDB.ManagerID, &DataFromDB.ClientID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding data from DB"})
			log.Printf("error: %s", err)
			return
		}

		query = fmt.Sprintf("UPDATE sessions SET clientid = '%d' WHERE managerid='%d'", DataFromBot.ClientID, DataFromDB.ManagerID)

		_, err = s.database.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error writing data to database"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}

	}
}

func (s *Server) handleCreateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.Session{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		query := "INSERT INTO requests (username, password) VALUES (?, ?)"
		_, err := s.database.Exec(query, DataFromBot.ManagerID, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding session data to database"})
			log.Println("Insert error:", err)
			return
		} else {
			log.Print("Successfully created new session")
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}
