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
			log.Println("error: Could not bind json")
			return
		}

		query := fmt.Sprintf("SELECT * FROM managers WHERE username='%s'", DataFromBot.Username)

		data, err := s.database.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			log.Printf("error: could not querry data from database, errormsg: %s", err)
			return
		}

		DataFromDB := models.LoginData{}

		data.Next()
		err = data.Scan(&DataFromDB.Username, &DataFromDB.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding data from DB"})
			log.Printf("error: %s", err)
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
