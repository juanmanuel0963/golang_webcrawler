<!-- BEGIN_TF_DOCS -->

# Web Crawler in Go

This is a simple web crawler written in Go that generates a sitemap tree of URLs starting from a given URL. It allows you to specify the depth level to retrieve.

## Table of Contents

- [Getting Started](#getting-started)
- [Usage](#usage)
- [Example](#example)
- [Documentation](#documentation)
- [Author](#author)

## Getting Started

To get started with this web crawler, follow these steps:

1. **Clone the Repository**

   Clone this repository to your local machine:

   ```sh
   git clone https://github.com/juanmanuel0963/webcrawler.git
   cd webcrawler

2. **Install Go**

Make sure you have Go installed on your system. If not, you can download and install it from https://golang.org/dl/.

3. **Install Dependencies**

Install the project dependencies by running the following command:

go mod download

## Usage

1. **Run the Crawler**

Run the web crawler with the following command:

go run webcrawler.go <starting_url> <deep_level_to_retrieve>

Replace <starting_url> with the URL from which you want to start crawling, and <deep_level_to_retrieve> with the maximum depth level you want to retrieve.

## Example

Here's an example of how to use the web crawler:

go run webcrawler.go https://parserdigital.com 2

This command will start crawling from https://parserdigital.com and retrieve URLs up to a depth level of 2.

## Documentation

You can access the documentation for this code by following these steps:

1. **Install godoc Tool**

Install the godoc tool using the following command:

go install golang.org/x/tools/cmd/godoc@latest

2. **Run godoc**

Run godoc with the following command:

godoc -http :6060

2. **View Documentation**

Open your web browser and navigate to http://localhost:6060/pkg/webcrawler/ to view the documentation.

## Author

Juan Diaz

Feel free to use and modify this code for your own purposes.

Happy crawling!


