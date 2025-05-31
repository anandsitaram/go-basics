package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFile(fileName string, balance float64) {
	writeFileTxt := fmt.Sprintf("The balance is %v", balance)
	//0644 - > read and write file for owner and others only  read file
	os.WriteFile(fileName, []byte(writeFileTxt), 0644)

}

// Custom error
// Functions/consts/variable name should start with upper case then only it can be used in other project
func ReadFile(fileName string) (float64, error) {
	//Here ReadFile returns byte[] and error  and error not used now so we can declare '_'
	data, err1 := os.ReadFile(fileName)
	if err1 != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceText := string(data)
	//strconv - string conversion
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse text")
	}
	return balance, nil

}
