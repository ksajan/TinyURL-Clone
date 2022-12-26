package main

import (
	service "tinyURL/services/url_generator"
	logging "tinyURL/utils"
)

func main() {

	logger := logging.NewLogger()

	logger.Infoln("Hello, playground")
	url := "https://tinyurl.com/api-create.php?url=https%"
	res := service.GenerateURL(url, "https://tinyurl.com")
	logger.Infoln(res)
}
