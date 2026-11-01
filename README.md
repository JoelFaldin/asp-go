# asp-go (Web Scraper)

A simple web scraper built in [Golang](https://go.dev/), focused on chilean sites. Parses the public HTML of each site and extracts useful information like tech stack, cdns and others.

> [!NOTE]  
> The project is currently under construction. Expect bugs, unfinished code and non-passing tests.

## Important

This project is for educational purposes only. It does not collect personal information and has no intentions in doing so.

## Features

* Real-time feedback: Prints to the console the results of processing each url.
* Multi-threading processing: Usage of multiple goroutines in batches to avoid server overloading and improved performance.
* Postgres in docker: Add docker compose file to build a postgres db locally.
* Fully customizable: Add controls to change internal behaviour of the application.
* CLI-first: Run the core application using the CLI.

### Future plans

* Full docker implementation: Add dockerfile to fully build and prepare the application.
* Web frontend: Expose the results live on a web frontend, built using modern technologies.
* Desktop application: Run the app as a native desktop application.

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
* `storage/`: Configures and exposes functions to interact with the database.

## Running the application locally

1. Make sure you have [Golang 1.27](https://go.dev/doc/devel/release) installed on your machine.

<details>
<summary>2. Manage docker:</summary>

The project includes `compose.yml`, a docker configuration file that starts a Postgres database.

Follow these steps to start and use it:

Start the container:
```
docker compose up -d postgres
```

Check container id and save the first 2 characters:
```
docker ps
```

Enter the container using bash (windows 11) (_aa_ is the id of the container):
```
docker exec -it aa bash
```

Enter the database:
```
psql -U scraper -d scraper_data
```

And you are in! You can now execute queries, like `select * from sites` or `select * from scans`.
</details>

3. Install dependencies:
```
go mod tidy
```

4. Run the project:
```
go run ./cmd/scraper/
```

## License

This project uses an *MIT* license. Read it [here](LICENSE)!

## Thanks for visiting!
