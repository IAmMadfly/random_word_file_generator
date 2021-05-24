package main

// import "github.com/sendgrid/rest"
import (
	"fmt"
	"strconv"
	"github.com/sendgrid/rest"
	"encoding/json"
)


func main() {
	const host = "https://random-word-api.herokuapp.com/word";
	param := "number";
	number_words := strconv.Itoa(500);
	end_point := "?" + param + "=" + number_words;
	baseURL := host + end_point;
	request := rest.Request {
		Method: rest.Get,
		BaseURL: baseURL,
	}

	res, err := rest.Send(request);
	

	if err != nil {
		fmt.Println("Failed to get request!");
		return;
	} else {

		var data_format []string;

		_ = json.Unmarshal([]byte(res.Body), &data_format);
		fmt.Printf("Response: %s", data_format[0]);
	}

}