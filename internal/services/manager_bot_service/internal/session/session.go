package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetActiveClientForManager(managerID int64) (int64, error) {
	url := fmt.Sprintf("http://localhost:8070/sessions/manager/%d", managerID)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("failed to get active client: %s", string(body))
	}

	var response struct {
		ClientID int64 `json:"client_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response.ClientID); err != nil {
		return 0, err
	}

	return response.ClientID, nil
}

func GetAssignedManager(clientID int64) (int64, error) {
	url := fmt.Sprintf("http://localhost:8070/sessions/client/%d", clientID)

	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		return 0, nil
	}
	defer resp.Body.Close()

	var response struct {
		ManagerID int64 `json:"manager_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response.ManagerID); err != nil {
		return 0, nil
	}

	return response.ManagerID, nil
}

func AssignClientToManager(clientID int64) (int64, error) {
	url := "http://localhost:8070/assign"
	requestBody, _ := json.Marshal(map[string]int64{"client_id": clientID})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Print("error1")
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Print("error2")
		return 0, fmt.Errorf("failed to assign manager: %s", string(body))
	}

	var response struct {
		ManagerID int64 `json:"manager_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response.ManagerID); err != nil {
		log.Print("error3")
		return 0, nil
	}

	return response.ManagerID, nil
}

func IsAuthorized(senderId int64) bool {
	url := fmt.Sprintf("http://localhost:8070/sessions/manager/%d", senderId)

	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()
	return true
}

func AddNewManager(managerID int64) error {
	url := "http://localhost:8070/create"
	requestBody, _ := json.Marshal(map[string]int64{"manager_id": managerID})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to assign manager: %s", string(body))
	}

	var response struct {
		ManagerID int64 `json:"manager_id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response.ManagerID); err != nil {
		return err
	}

	return nil
}

func DeauthorizeManager(managerID int64) error {
	url := fmt.Sprintf("http://localhost:8070/delete/%d", managerID)
	requestBody, _ := json.Marshal(map[string]int64{"manager_id": managerID})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to assign manager: %s", string(body))
	}

	return nil
}

func FreeManager(managerID int64) error {
	url := "http://localhost:8070/end"
	requestBody, _ := json.Marshal(map[string]int64{"manager_id": managerID})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to free manager: %s", string(body))
	}

	return nil
}
