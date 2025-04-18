package migrations

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	_ "github.com/go-sql-driver/mysql"
)

// Define metadata structure matching frontmatter
type Metadata struct {
	Title string   `yaml:"title"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}

func main() {
	// Read file
	content, err := ioutil.ReadFile("posts/001-python-regex.md")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Extract frontmatter
	var meta Metadata
	body, err := frontmatter.Parse(strings.NewReader(string(content)), &meta)
	if err != nil {
		log.Fatalf("Error parsing frontmatter: %v", err)
	}

	// Connect to MariaDB
	dsn := "user:password@tcp(localhost:3306)/your_database"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	// Parse date
	createdOn, err := time.Parse("2006-01-02", meta.Date)
	if err != nil {
		createdOn = time.Now() // fallback
	}

	// Insert into BlogPosts
	_, err = db.Exec(`
		INSERT INTO BlogPosts (title, author, content, tags, subject, created_on)
		VALUES (?, ?, ?, ?, ?, ?)`,
		meta.Title,
		"Keith Thomson",
		string(body),
		strings.Join(meta.Tags, ","),
		"Regex", // Subject: manually from filename or logic
		createdOn,
	)
	if err != nil {
		log.Fatalf("Error inserting into DB: %v", err)
	}

	fmt.Println("Blog post inserted successfully!")
}
