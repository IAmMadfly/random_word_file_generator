package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/sendgrid/rest"
)

func write_to_file(file_name string, word_pool *[]string, waiter *sync.WaitGroup) {
	// file, err := os.Create("./Data_0.txt")

	// if err != nil {
	// 	fmt.Println("Failed to make file!")
	// 	return
	// }

	// writer := bufio.NewWriter(file)

	fmt.Println(len(*word_pool))
	waiter.Done()
}

func main() {

	var word_pool int
	word_pool = 10

	const host = "https://random-word-api.herokuapp.com/word"
	param := "number"
	number_words := strconv.Itoa(word_pool)
	end_point := "?" + param + "=" + number_words
	baseURL := host + end_point
	request := rest.Request{
		Method:  rest.Get,
		BaseURL: baseURL,
	}

	res, err := rest.Send(request)
	var words []string

	if err != nil {
		fmt.Println("Failed to get request!")
		return
	} else {
		err = json.Unmarshal([]byte(res.Body), &words)

		if err != nil {
			fmt.Println("Failed to decode JSON response!")
			return
		}
		fmt.Printf("Response: %s\n", words[0])
	}

	// linesToWrite := []string{"some kind", "of data to be", "Written to file"}

	// writer.WriteString(linesToWrite[0])
	var wait_group sync.WaitGroup
	for i := 0; i < 10; i++ {

		wait_group.Add(1)
		go write_to_file("data_"+strconv.Itoa(i), &words, &wait_group)
	}

	wait_group.Wait()
}
