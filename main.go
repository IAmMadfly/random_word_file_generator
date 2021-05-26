package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"

	"github.com/alexflint/go-arg"
	"github.com/sendgrid/rest"
)

func write_to_file(
	file_name string, 
	word_pool *[]string, 
	waiter *sync.WaitGroup, 
	word_count int,
	) {
	file, err := os.Create(file_name)

	if err != nil {
		fmt.Println("Failed to make file!")
		return
	}

	// writer := bufio.NewWriter(file)
	writer := bufio.NewWriterSize(file, 10000)
	var rand_index int
	for i := 0; i < word_count; i++ {
		rand_index = rand.Intn(len(*word_pool))
		writer.Write([]byte((*word_pool)[rand_index]+"\n"))
	}

	fmt.Println(file_name, "- Completed writing")
	writer.Flush()
	waiter.Done()
}

func main() {

	var args struct {
		Pool_count int
		Word_count int
		File_count int
	}
	args.Pool_count = 1000
	args.Word_count = 1000
	args.File_count = 10

	arg.MustParse(&args)

	fmt.Println(
		"Word pool:", args.Pool_count, 
		", Words per file:", args.Word_count, 
		", File count:", args.File_count,
	)

	word_pool := args.Pool_count
	word_count := args.Word_count

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
	}

	var wait_group sync.WaitGroup
	fmt.Println("Starting to write random word files!")
	for i := 0; i < args.File_count; i++ {
		file_name := "data_"+strconv.Itoa(i)+".txt"

		wait_group.Add(1)
		go write_to_file(file_name, &words, &wait_group, word_count)
	}

	wait_group.Wait()
	fmt.Println("Task Complete!")
}
