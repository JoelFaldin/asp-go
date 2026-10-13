# asp-go (Web Scraper)

A simple web scraper built in [Golang](https://go.dev/), focused on chilean sites. Process the public HTML of each site and extracts useful information like tech stack, cdns and others.

> [!NOTE]  
> The project is currently under construction. Expect bugs, unfinished code and non-passing tests.

## Important

This project is for educational purposes only. It does not collect personal information and has no intentions on doing so.

## Features

* Real-time feedback: Prints to the console the results of processing each url.
* Multi-threading processing: Usage of multiple goroutines in batches to avoid server overloading and improved performance.

### Future plans

* Postgres in docker: Add docker compose file to build a postgres db locally.
* Full docker implementation: Add dockerfile to fully build and prepare the application.
* Web frontend: Expose the results live on a web frontend, built using modern technologies.
* Desktop application: Run the app as a native desktop application.
* Fully customizable: Add controls to change internal behaviour of the application.
* CLI-first: Run the core application using the CLI.

## Project structure

The project follows a traditional Golang folder structure:

### `cmd/`

Entry point of the application. Initializes logging, loads data file, and measure the total operation time.

### `internal/`

The core functionality. Configuration files, logger, fetcher, etc.

* `config/`: Loads the data file and unmarshals the urls.
* `crawler/`: Launches a batch of goroutines that listen a channel of urls.
* `detector/`: Rules that detect each technology.
* `fetcher/`: Sends a `GET` request into a singular URL.
* `logger/`: Helper functions to print information into the console.

## Running the application locally

1. Make sure you have [Golang 1.27](https://go.dev/doc/devel/release) installed on your machine.
2. Install dependencies:
```
go mod tidy
```

3. Run the project:
```
go run ./cmd/scraper/main.go
```

## License

This project uses an *MIT* license. Read it [here](LICENSE)!

## Thanks for visiting!
