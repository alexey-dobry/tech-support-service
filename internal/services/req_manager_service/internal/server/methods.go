package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexey-dobry/tech-support-platform/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

func (s *Server) handleGetClientData() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientID := c.Param("client_id")

		var activeSession models.Manager

		query := fmt.Sprintf("SELECT * FROM sessions WHERE client_id=%s", ClientID)

		data, err := s.database.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			log.Printf("error: could not querry data from database, errormsg: %s", err)
			return
		} else {
			data.Next()
			err := data.Scan(&activeSession.ManagerID, &activeSession.IsFree, &activeSession.ClientID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
				return
			} else {
				c.JSON(http.StatusOK, activeSession.ManagerID)
			}
		}
	}
}

func (s *Server) handleGetManagerData() gin.HandlerFunc {
	return func(c *gin.Context) {
		managerID := c.Param("manager_id")

		var activeSession models.Manager

		query := fmt.Sprintf("SELECT * FROM sessions WHERE manager_id=%s", managerID)

		data, err := s.database.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			log.Printf("error: could not querry data from database, errormsg: %s", err)
			return
		} else {
			data.Next()
			err := data.Scan(&activeSession.ManagerID, &activeSession.IsFree, &activeSession.ClientID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
				return
			} else {
				if activeSession.ClientID == 0 {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
				} else {
					c.JSON(http.StatusOK, activeSession.ClientID)
				}
			}
		}
	}
}

func (s *Server) handleAddNewManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		var manager models.EndRequest
		if err := c.BindJSON(&manager); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		query := "INSERT INTO sessions (manager_id, is_free, client_id) VALUES ( ?, ?, ?)"

		_, err := s.database.Exec(query, manager.ManagerID, true, 0)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error writing data to database"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}

func (s *Server) handleAssingnManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.Request{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			log.Println("error: Could not bind json")
			return
		}

		query := "SELECT * FROM sessions WHERE is_free='true'"

		data, err := s.database.Query(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error retreaving data from database"})
			log.Printf("error: could not querry data from database, errormsg: %s", err)
			return
		}

		DataFromDB := models.Manager{}

		data.Next()
		err = data.Scan(&DataFromDB.ManagerID, &DataFromDB.IsFree, &DataFromDB.ClientID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding data from database"})
			log.Printf("error: %s", err)
			return
		}

		query = fmt.Sprintf("UPDATE sessions SET client_id = '%d' WHERE manager_id=%d", DataFromBot.ClientID, DataFromDB.ManagerID)

		_, err = s.database.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error executing command"})
			return
		} else {
			c.JSON(http.StatusOK, DataFromDB.ManagerID)
		}
	}
}

func (s *Server) handleFreeManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		DataFromBot := models.EndRequest{}
		if err := c.BindJSON(&DataFromBot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		query := fmt.Sprintf("UPDATE sessions SET is_free = true, client_id=0 WHERE manager_id=%d", DataFromBot.ManagerID)

		_, err := s.database.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error executing command"})
			log.Println("Insert error:", err)
			return
		} else {
			log.Print("Successfully freed manager")
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}

func (s *Server) handleEndSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		managerID := c.Param("manager_id")

		query := fmt.Sprintf("DELETE FROM sessions WHERE manager_id=%s", managerID)

		_, err := s.database.Exec(query)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error executing command"})
			log.Println("Insert error:", err)
			return
		} else {
			log.Print("Successfully deleted session")
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		}
	}
}
