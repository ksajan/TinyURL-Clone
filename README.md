# Tiny URL Clone

A Tiny URL Clone is a web application that allows users to create short, unique URLs (also known as "tiny URLs") for long, cumbersome URLs. This can be useful for a variety of purposes, such as sharing links on social media, sending links via email or messaging, or simply for the convenience of having a shorter, easier-to-remember URL.

## Getting Started

To get started with the Tiny URL Clone, you will need to have Go installed on your computer. You can download Go from the [official Go website](https://golang.org/dl/).

Once you have Go installed, you can clone the Tiny URL Clone repository from GitHub:

```git clone https://github.com/ksajan/tinyURL-clone.git```

This will create a new directory called `tinyURL-clone` in your current directory. You can then navigate into the `tinyURL-clone` directory and run the application:

```cd tinyURL-clone```
```go mod init tinyURL```
```go mod tidy```
```go run main.go```

The application will start running on port 8000. You can then visit `http://localhost:8000` in your browser to view the application.

## Usage

The Tiny URL Clone application allows users to create short URLs for long URLs. To create a short URL, users can simply enter the long URL into the application and click the "Shorten URL" button. The application will then generate a short URL that redirects to the original URL.

Open postman or any other API testing tool and send a POST request to `http://localhost:8000/shorten` with the following JSON body:

```json
{
  "url": "https://www.google.com"
}
```

The response will be a JSON object containing the original URL and the shortened URL:

```json
{
  "originalURL": "https://www.google.com",
  "shortURL": "http://localhost:8000/1"
}
```

You can then visit the shortened URL in your browser to be redirected to the original URL.

## Technical Details

The Tiny URL Clone application is written in Go and uses the [fiber, aerospike, and uuid]

Please refer to the application's architecture and design below:



`Note: This section will be updated as the application is developed.`



## Testing

This application is under development. Tests are coming soon!

## Contributing

Bug reports and pull requests are welcome on GitHub

## License

The application is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).


