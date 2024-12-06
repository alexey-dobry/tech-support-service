package middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func Register(hashKey, username, password string) error {

	if hashKey != "place_to_append_hash" {
		return errors.New("неверное значение ключа")
	}

	payload := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("Ошибка сериализации данных:", err)
		return errors.New("ошибка сериализации данных")
	}

	resp, err := http.Post("http://localhost:8080/reg", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Ошибка запроса к микросервису:", err)
		return errors.New("ошибка запроса к микросервису")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Ошибка аутентификации:", resp.Status)
		return errors.New("ошибка аутентификации")
	}

	// Парсим ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return errors.New("ошибка чтения ответа")
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Println("Ошибка парсинга ответа:", err)
		return errors.New("ошибка парсинга ответа")
	}

	return nil
}
