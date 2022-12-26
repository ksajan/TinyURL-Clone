package main

import (
	"fmt"
	service "tinyURL/services/url_generator"
)

func main() {
	fmt.Println("Hello, playground")
	url := "https://tinyurl.com/api-create.php?url=https%"
	res := service.GenerateURL(url, "https://tinyurl.com")
	fmt.Println(res)
}
