package main

import (
	"chatreader/helpers"
	"chatreader/utils"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
)

func main() {
	var (
		numberOfJsonFiles      int
		selectedNumberJsonFile string
		dirName                string
		txtName                string
	)

	entries, err := os.ReadDir("data")
	if err != nil {
		color.Red("Ошибка чтения директории: %v", err)
		helpers.ExitProgram()
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			numberOfJsonFiles++
		}
	}

	helpers.GetSelectedJsonFile(&selectedNumberJsonFile, numberOfJsonFiles, entries)
	helpers.GetName(&dirName, "Введите название директории")
	helpers.GetName(&txtName, "Введите название txt файла")

	if txtName == "" {
		txtName = "result"
	}

	fileNumber, err := strconv.Atoi(selectedNumberJsonFile)
	if err != nil {
		color.Red("Ошибка конвертации строки в число: %v", err)
		helpers.ExitProgram()
	}

	jsonFile, canContinue := helpers.GetJsonFileByNumber(fileNumber, entries)
	if !canContinue {
		color.Red("Файл не найден")
		helpers.ExitProgram()
	}

	utils.ConvertJsonToText(jsonFile, dirName, txtName)
	helpers.ExitProgram(0)
}
