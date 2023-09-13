package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Не указан путь к csv-файлу")
	}

	filePath := os.Args[2]

	// Открытие csv-файла
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создание ридера CSV
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6
	reader.Comment = '#'
	reader.LazyQuotes = true

	fmt.Printf("Вы уверены, что хотите удалить файлы по записям из %s? (y/n)", filePath)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Text() == "y" {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			filePath := record[1]

			if _, err := os.Stat(filePath); err != nil {
				if os.IsNotExist(err) {
					log.Printf("Файл %s не существует\n", filePath)
				} else {
					panic(err)
				}
			} else {
				log.Printf("Файл %s удален\n", filePath)
				os.Remove(filePath)

			}

		}
	}

}
