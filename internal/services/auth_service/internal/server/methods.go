package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetLoginData() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.LoginData{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		query := fmt.Sprintf("SELECT * FROM Users WHERE username='%s'", DataFromBot.Username)

		data, err := s.database.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			return
		}

		DataFromDB := models.LoginData{}

		err = data.Scan(&DataFromDB.Username, &DataFromDB.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding data from DB"})
			return
		}
		// Проверяем логин и пароль
		if DataFromBot.Password == DataFromDB.Password {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	}
}

func (s *Server) handlePostLoginData() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.LoginData{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		query := "INSERT INTO requests (username, password) VALUES (?, ?)"
		_, err := s.database.Exec(query, DataFromBot.Username, DataFromBot.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error adding user data to database"})
			log.Println("Insert error:", err)
			return
		} else {
			log.Print("Successfully added new user to the database")
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}
