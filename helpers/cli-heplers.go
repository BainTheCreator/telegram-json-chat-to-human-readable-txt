package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func GetSelectedJsonFile(selectedJsonFile *string, numberOfJsonFiles int, entries []os.DirEntry) {
	switch numberOfJsonFiles {
	case 0:
		color.Red("В директории data нет json файлов")
		ExitProgram()
	case 1:
		*selectedJsonFile = strconv.Itoa(numberOfJsonFiles)
	default:
		color.White("Выберите json файл от %d до %d:\n", 1, numberOfJsonFiles)
		for index, entry := range entries {
			if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
				color.White("%d: %s\n", index+1, entry.Name())
			}
		}
		fmt.Scanln(selectedJsonFile)
		if *selectedJsonFile < "1" || *selectedJsonFile > strconv.Itoa(numberOfJsonFiles) {
			color.Red("Неверный выбор файла: %s", *selectedJsonFile)
			ExitProgram()
		}
	}
}

func GetName(name *string, text string) {
	fmt.Printf("%s: ", text)
	fmt.Scanln(name)
}

func GetJsonFileByNumber(number int, entries []os.DirEntry) (string, bool) {
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			count++
			if count == number {
				return strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name())), true
			}
		}
	}

	return "", false
}

func ExitProgram(status ...int) {
	code := 1
	if len(status) > 0 {
		code = status[0]
	}

	color.White("Для выхода нажмите Enter...")
	fmt.Scanln()
	os.Exit(code)
}
