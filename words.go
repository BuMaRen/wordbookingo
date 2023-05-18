package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Word struct {
	Name        string
	Decs        string
	Correct     float64
	Fault       float64
	CorrectRate float64
	TouchTime   int64
}
type WordBook struct {
	Words []Word
}

var wordBook *WordBook = nil

func ReadWordBook(path string) []Word {
	if wordBook != nil {
		return wordBook.Words
	}
	fd, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err is", err)
		return nil
	}
	defer fd.Close()
	fi, err := fd.Stat()
	if err != nil {
		fmt.Println("read file stat err is", err)
		return nil
	}
	buffer := make([]byte, fi.Size())
	if _, err = fd.Read(buffer); err != nil {
		fmt.Println("read file failed err is", err)
		return nil
	}
	wordBook = &WordBook{}
	if err := json.Unmarshal(buffer, wordBook); err != nil {
		fmt.Println("unmarshal failed, err is", err)
		return nil
	}
	return wordBook.Words
}
