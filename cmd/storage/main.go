package main

import (
	"fmt"
	"log"
)
import "github.com/mdvasilyev/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(st.Files)

	fmt.Println("it is restored", restoredFile)
}
