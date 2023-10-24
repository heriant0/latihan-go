package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
}

func main() {
	now := time.Now()
	asynchornous()

	log.Println("Done in", time.Since(now).Seconds(), "s")
}
func asynchornous() {
	userCh, _ := readFileConcurrent("./data.json")
	done := make(chan bool)

	writeToFileConcurrent(userCh, done)

	if <-done {
		log.Println("Done")
	}
}

func readFileConcurrent(filename string) (<-chan User, error) {
	userCh := make(chan User)
	now := time.Now()
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	users := []User{}

	err = json.Unmarshal(dataByte, &users)
	if err != nil {
		return nil, err
	}

	go func() {
		for _, user := range users {
			userCh <- user
		}

		close(userCh)
	}()

	log.Println("success read data in", time.Since(now).Seconds(), "s")
	return userCh, nil
}

func writeToFileConcurrent(dataCh <-chan User, done chan bool) {
	wg := sync.WaitGroup{}
	folderPath := "./filesGenerated/"
	_ = checkFolderPath(folderPath)

	for data := range dataCh {
		wg.Add(1)
		go func(data User) {
			// convert USD to IDR
			newData := data
			newData.Salary = newData.Salary * 15_000
			data = newData

			// Write data to json file
			user, _ := json.Marshal(data)
			err := os.WriteFile(folderPath+data.Name+".json", user, 0666)
			if err != nil {
				log.Println("error when try to write file", err)
			}
			wg.Done()
		}(data)
	}

	go func() {
		wg.Wait()
		done <- true
	}()
}

func checkFolderPath(folderPath string) error {
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		fmt.Printf("The folder %s does not exist, creating it...\n", folderPath)
		err := os.MkdirAll(folderPath, 0755)
		if err != nil {
			return err
		}
		fmt.Printf("Folder %s has been created.\n", folderPath)
	}
	return nil
}
