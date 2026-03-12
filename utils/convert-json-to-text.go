package utils

import (
	"bufio"
	helpers "chatreader/helpers"
	types "chatreader/types"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func ConvertJsonToText(jsonName string, dirName string, txtName string) {
	// Открываем и читаем JSON файл
	var jsonPath string = fmt.Sprintf("data/%s.json", jsonName)

	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		color.Red("Ошибка открытия файла: %v", err)
		return
	}
	defer jsonFile.Close()

	var chat types.Chat

	err = json.NewDecoder(jsonFile).Decode(&chat)
	if err != nil {
		color.Red("Ошибка потокого парсинга JSON: %v", err)
		return
	}

	// Обрабатываем список сообщений
	messageList := helpers.ToMessageList(chat.Messages, types.ToMessageListOptions{})

	if len(messageList) == 0 {
		color.Red("Сообщений нет, файл не будет создан")
		return
	}

	// Запись файла
	var filePath string
	txtNameWithExtension := fmt.Sprintf("%s.txt", txtName)

	if dirName != "" {
		filePath = filepath.Join("result", dirName, txtNameWithExtension)
	} else {
		filePath = filepath.Join("result", txtNameWithExtension)
	}
	dirPath := filepath.Dir(filePath)

	err = os.MkdirAll(dirPath, 0755)
	if err != nil {
		color.Red("Ошибка создания директории: %v", err)
		return
	}

	txtFile, err := os.Create(filePath)
	if err != nil {
		color.Red("Ошибка создания файла: %v", err)
		return
	}
	defer txtFile.Close()

	writer := bufio.NewWriter(txtFile)

	for _, msg := range messageList {
		_, err := writer.WriteString(msg + "\n")
		if err != nil {
			color.Red("Ошибка записи в файл: %v", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		color.Red("Ошибка сброса буфера: %v", err)
	}

	color.Green("Файл успешно создан: %s", filePath)
}
