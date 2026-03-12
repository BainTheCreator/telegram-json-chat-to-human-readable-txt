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
		color.Red("Не удалось открыть файл «%s». Проверьте путь и права доступа: %v", jsonPath, err)
		return
	}
	defer jsonFile.Close()

	var chat types.Chat

	err = json.NewDecoder(jsonFile).Decode(&chat)
	if err != nil {
		color.Red("Не удалось прочитать данные: файл повреждён или это не диалог в формате JSON. Подробности: %v", err)
		return
	}

	// Обрабатываем список сообщений
	messageList := helpers.ToMessageList(chat.Messages, types.ToMessageListOptions{})

	if len(messageList) == 0 {
		color.Red("В выбранном диалоге нет сообщений. Текстовый файл не создан.")
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
		color.Red("Не удалось создать папку для результата «%s»: %v", dirPath, err)
		return
	}

	txtFile, err := os.Create(filePath)
	if err != nil {
		color.Red("Не удалось создать файл результата «%s». Проверьте права доступа: %v", filePath, err)
		return
	}
	defer txtFile.Close()

	writer := bufio.NewWriter(txtFile)

	for _, msg := range messageList {
		_, err := writer.WriteString(msg + "\n")
		if err != nil {
			color.Red("Ошибка при записи сообщений в файл «%s»: %v", filePath, err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		color.Red("Не удалось сохранить данные в файл «%s»: %v", filePath, err)
	}

	color.Green("Файл успешно создан: %s", filePath)
}
