package models 

import (
	"database/sql"
	"log"
)


type BlogPost struct{
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author string `json:"author"`
	Subtitle  string `json:"subtitle"`
	Content string `json:"content"`
	
}


func GetPosts(db *sql.DB) ([]BlogPost, error) {
	var posts []BlogPost
	rows, err := db.Query("SELECT id, title, author, content FROM blog_posts")
	if err != nil {
		log.Println("Error querying blog posts:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post BlogPost
		if err := rows.Scan(&post.ID, &post.Title, &post.Author, &post.Content); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error with rows:", err)
		return nil, err
	}

	return posts, nil
}


func CreatePost(db *sql.DB, b BlogPost) error {
	_, err := db.Exec("INSERT INTO blog (title, author, subtitle, content) VALUES (?, ?, ?, ?)", b.Title, 
		b.Author, b.Subtitle, b.Content)

	return err
}
