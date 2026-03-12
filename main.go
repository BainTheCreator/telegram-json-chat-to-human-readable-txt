package main

import (
	"bufio"
	helpers "chatreader/helpers"
	types "chatreader/types"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Открываем и читаем JSON файл
	jsonFile, err := os.Open("data/result.json")
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer jsonFile.Close()

	var chat types.Chat

	err = json.NewDecoder(jsonFile).Decode(&chat)
	if err != nil {
		log.Fatalf("Ошибка потокого парсинга JSON: %v", err)
		return
	}

	// Обрабатываем список сообщений
	messageList := helpers.ToMessageList(chat.Messages, types.ToMessageListOptions{})

	if len(messageList) == 0 {
		log.Println("Сообщений нет, файл не будет создан")
		return
	}

	// Запись файла
	filePath := filepath.Join("result", "result.txt")
	dirPath := filepath.Dir(filePath)

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Fatalf("Ошибка создания директории: %v", err)
		return
	}

	txtFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Ошибка создания файла: %v", err)
		return
	}
	defer txtFile.Close()

	writer := bufio.NewWriter(txtFile)

	for _, msg := range messageList {
		_, err := writer.WriteString(msg + "\n")
		if err != nil {
			log.Fatalf("Ошибка записи в файл: %v", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Fatalf("Ошибка сброса буфера: %v", err)
	}

	log.Printf("Файл успешно создан: %s", filePath)
}
