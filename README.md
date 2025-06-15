# Go Scholarship Application Web App

A simple, self-contained web application built with Go for submitting and reviewing scholarship applications. This project uses a SQLite database, making it easy to run and deploy without external database dependencies.

## Description

This web application provides two main functionalities:
1.  A public-facing form where students can submit their scholarship applications, including personal details and an essay.
2.  An administrative view where all submitted applications can be reviewed in a clean, chronological list.

The project is designed to be a straightforward, practical example of a full-stack web application using Go's standard library for the backend and simple HTML/CSS for the frontend.

## Features

* **Submit Applications:** A clean and simple form for applicants.
* **View Submissions:** A private page to review all applications, sorted by submission date.
* **Self-Contained:** Uses a file-based SQLite database, requiring no external database setup.
* **Easy to Run:** Compiles to a single binary for easy execution.
* **Deployable:** Includes a `Dockerfile` and is ready for deployment on platforms like [Fly.io](https://fly.io/).

## Technologies Used

* **Backend:** [Go (Golang)](https://go.dev/)
* **Database:** [SQLite](https://www.sqlite.org/index.html) (via `github.com/mattn/go-sqlite3`)
* **Frontend:** HTML5, CSS3
* **Containerization:** [Docker](https://www.docker.com/)

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing.

### Prerequisites

* [Go](https://go.dev/doc/install) (version 1.20 or newer)
* [Git](https://git-scm.com/downloads)
* (Optional) [Docker](https://www.docker.com/products/docker-desktop/) for containerized deployment.

### Installation

1.  **Clone the repository:**
    ```sh
    git clone [https://github.com/your-username/go-scholarship-app.git](https://github.com/your-username/go-scholarship-app.git)
    cd go-scholarship-app
    ```
    *(Replace `your-username` with your actual GitHub username)*

2.  **Install dependencies:**
    Go will automatically handle dependencies when you build or run the project. To explicitly install them, run:
    ```sh
    go mod tidy
    ```

3.  **Run the application:**
    ```sh
    go run main.go
    ```

4.  **Access the web app:**
    The server will start on `localhost:8080`.
    * **Application Form:** Open your web browser to [http://localhost:8080/](http://localhost:8080/)
    * **View Submissions:** Navigate to [http://localhost:8080/applications](http://localhost:8080/applications)
    * **Health Check:** Check the API status at [http://localhost:8080/health](http://localhost:8080/health)

A `scholarships.db` file will be automatically created in your project directory upon the first run.

## File Structure

/
|-- main.go                 # Main application logic, routing, and handlers
|-- go.mod                  # Go module definitions
|-- go.sum                  # Go module checksums
|-- Dockerfile              # Instructions for building a Docker container
|-- README.md               # This file
|-- scholarships.db         # The SQLite database file (created on run)
|-- .gitignore              # Files for Git to ignore
|
|-- static/
|   -- style.css # CSS stylesheets |-- templates/
|-- apply.html          # HTML template for the application form
`-- applications.html   # HTML template for viewing submissions

## Deployment

This application is ready to be deployed. It includes a `Dockerfile` for building a containerized version of the app, which can be hosted on any platform that supports Docker containers.

For a detailed guide on deploying to [Fly.io](https://fly.io/) with a persistent volume for the SQLite database, you can create a `DEPLOYMENT.md` file in your repository.

## Contributing

Contributions are welcome! If you have ideas for improvements or find a bug, please feel free to open an issue or submit a pull request.

1.  Fork the Project
2.  Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3.  Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4.  Push to the Branch (`git push origin feature/AmazingFeature`)
5.  Open a Pull Request
