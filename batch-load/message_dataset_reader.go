package batchload

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gokberkotlu/auto-messaging/entity"
)

func ReadCSV() []entity.Message {
	file, err := os.Open("./asset/message_dataset.csv")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return nil
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Failed to read CSV file:", err)
		return nil
	}

	var messages []entity.Message
	for _, row := range rows[1:] {
		messages = append(messages, entity.Message{
			To:      row[0],
			Content: row[1],
			Status:  entity.Active,
		})
	}

	return messages
}
