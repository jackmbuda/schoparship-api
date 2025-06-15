package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// ScholarshipApplication represents the data for a single scholarship application.
type ScholarshipApplication struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	Essay       string
	SubmittedAt time.Time
}

var db *sql.DB
var tmpl *template.Template

func main() {
	var err error
	// Initialize the database connection
	db, err = initDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Parse all template files in the 'templates' directory
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}

	// Serve static files (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define application routes
	http.HandleFunc("/", applyHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/applications", applicationsHandler)
	http.HandleFunc("/health", healthCheckHandler)

	// Start the web server
	fmt.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// initDB initializes the SQLite database and creates the 'applications' table.
// It now stores the database in the /data directory for persistence on Fly.io.
func initDB() (*sql.DB, error) {
	// The database file will be located at /data/scholarships.db
	// This path corresponds to the volume mount point in fly.toml
	dbPath := "/data/scholarships.db"

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not open database at %s: %w", dbPath, err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS applications (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        firstname TEXT NOT NULL,
        lastname TEXT NOT NULL,
        email TEXT NOT NULL,
        essay TEXT NOT NULL,
        submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	if _, err = database.Exec(createTable); err != nil {
		return nil, fmt.Errorf("could not create table: %w", err)
	}

	return database, nil
}

// applyHandler serves the scholarship application form.
func applyHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "apply.html", nil); err != nil {
		http.Error(w, "Failed to render application form", http.StatusInternalServerError)
	}
}

// submitHandler processes the scholarship application form submission.
func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Prepare the SQL statement for inserting a new application
	stmt, err := db.Prepare("INSERT INTO applications (firstname, lastname, email, essay) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Execute the statement with form values
	_, err = stmt.Exec(r.FormValue("firstName"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("essay"))
	if err != nil {
		http.Error(w, "Failed to save application", http.StatusInternalServerError)
		return
	}

	// Redirect to a success page
	http.Redirect(w, r, "/?status=success", http.StatusSeeOther)
}

// applicationsHandler displays all submitted scholarship applications.
func applicationsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, firstname, lastname, email, essay, submitted_at FROM applications ORDER BY submitted_at DESC")
	if err != nil {
		http.Error(w, "Failed to retrieve applications", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var apps []ScholarshipApplication
	for rows.Next() {
		var app ScholarshipApplication
		if err := rows.Scan(&app.ID, &app.FirstName, &app.LastName, &app.Email, &app.Essay, &app.SubmittedAt); err != nil {
			http.Error(w, "Failed to scan application data", http.StatusInternalServerError)
			continue
		}
		apps = append(apps, app)
	}

	if err := tmpl.ExecuteTemplate(w, "applications.html", apps); err != nil {
		http.Error(w, "Failed to render applications list", http.StatusInternalServerError)
	}
}

// healthCheckHandler provides a simple health check endpoint.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Scholarship API is running")
}
