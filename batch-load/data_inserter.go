package batchload

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/gokberkotlu/auto-messaging/entity"
	"github.com/gokberkotlu/auto-messaging/repository"
)

func ReadCSV() {
	// Open the CSV file
	file, err := os.Open("./asset/MOCK_DATA.csv")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all rows from the CSV
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Failed to read CSV file:", err)
		return
	}

	// fmt.Println(len(rows))
	// for _, row := range rows[1:] {
	// 	fmt.Printf("phone: %s | status: %s | content: %s\n", row[0], row[1], row[2])
	// }

	var messages []entity.Message
	for _, row := range rows[1:] {
		intValue, _ := strconv.Atoi(row[1])
		messages = append(messages, entity.Message{
			To:      row[0],
			Status:  entity.StatusType(intValue),
			Content: row[2],
		})
	}

	var messageRepository repository.IMessageRepository = repository.NewMessageRepository()
	messageRepository.BulkLoad(messages)
}
