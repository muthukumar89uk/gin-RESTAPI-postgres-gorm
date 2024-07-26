package helper

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Configure(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		fmt.Printf("error at loading %s file", filename)
		return err
	}

	return nil
}