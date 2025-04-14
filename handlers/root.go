package handlers

import (
	"fmt"
	"net/http"

	"forum/config"
	"forum/database"
)

type Post struct {
	Content   string
	Interest  string
	Username  string
	CreatedAt string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]interface{}{"Authenticated": false})
		return
	}

	userID, exists := database.GetUserBySession(sessionCookie.Value)
	if !exists {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]interface{}{"Authenticated": false})
		return
	}

	user := database.GetUserInfo(userID)

	if r.Method == http.MethodPost {
		content := r.FormValue("content")
		interest := r.FormValue("interest")

		if content != "" {
			CreatePost(userID, content, interest)
		}
	}

	posts, err := GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}
	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]interface{}{
		"Authenticated": true,
		"Username":      user.Username,
		"Posts":         posts,
	})
}

func CreatePost(userID int, content, interest string) {
	stmt, err := database.DB.Prepare("INSERT INTO posts(user_id, content, interest) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, content, interest)
	if err != nil {
		fmt.Println("Error executing statement:", err)
		return
	}
	
	fmt.Println("Post created:", content, interest)
}
func GetAllPosts() ([]Post, error) {
	rows, err := database.DB.Query(`
		SELECT users.username, posts.content, posts.interest, posts.created_at
		FROM posts
		JOIN users ON posts.user_id = users.id
		ORDER BY posts.created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Username, &post.Content, &post.Interest, &post.CreatedAt)
		if err != nil {
			continue
		}
		posts = append(posts, post)
	}

	fmt.Println("Retrieved posts:", posts)
	return posts, nil
}
