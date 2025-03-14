package handlers

import (
	"fmt"
	"log"
	"net/http"
	// "strings"
	"text/template"
	"database/sql"
    _ "github.com/mattn/go-sqlite3"

	"forum/database"
)


func ForumHandler(w http.ResponseWriter, r *http.Request) {

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
        return
    }

    tmpl, err := template.ParseGlob("./static/templates/*")
    if err != nil {
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        return
    }
	// data, err := database.Db.Query("SELECT post.id, post.content, comments.id, comments.post_id, comments.comment FROM post LEFT JOIN comments ON post.id = comments.post_id ORDER BY post.id")
	// if err != nil{
	// 	fmt.Println("Error opening or getting the posts data : ", err)
	// 	return 
	// }
	// defer data.Close()

	// contents := []string{}
	// for data.Next() {
	// 	var temp string
	// 	if err := data.Scan(&temp); err != nil {
	// 		log.Println("Error scanning row:", err)
	// 		continue
	// 	}
	// 	// fmt.Println(content)
	// 	contents = append(contents, temp)
	// }
	// comments, err := database.Db.Exec("SELECT comments.id, comments.post_id, comments.comment FROM comments")
	// if err != nil{
	// 	log.Print("Failed to get data from Comment s table", err)
	// }
	// fmt.Println(comments)

	type Comment struct {
		ID     int
		PostId int
		Text   string
	}
	
	type Post struct {
		ID       int
		Content  string
		Comments []Comment
		Likes int
		Dislikes int
	}
	type Like struct{
		ID int
		PostId int
		likeOrDislike bool
	}
	
	postMap := make(map[int]Post)
	
	data, err := database.Db.Query(`
		SELECT post.id, post.content, 
    COUNT(CASE WHEN likes_dislikes.is_like = 1 THEN 1 END) AS like_count,
    COUNT(CASE WHEN likes_dislikes.is_like = 0 THEN 1 END) AS dislike_count,
    comments.post_id, comments.comment 
	FROM post
	LEFT JOIN comments ON post.id = comments.post_id
	LEFT JOIN likes_dislikes ON post.id = likes_dislikes.post_id
	GROUP BY post.id
	ORDER BY post.id;
	`)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer data.Close()
	
	for data.Next() {
		var PostId int
		var postContent string
		var commentID, commentPostId sql.NullInt64
		var commentText sql.NullString
		var likes, dislikes int
		
		if err := data.Scan(&PostId, &postContent, &likes, &dislikes, &commentID, &commentText); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		// fmt.Println("Likes are :", likes)
		// fmt.Println("DisLikes are :", dislikes)
		post, exists := postMap[PostId]
		if !exists {
			post = Post{
				ID:       PostId,
				Content:  postContent,
				Comments: []Comment{},
				Likes:    likes,     
				Dislikes: dislikes,
			}
		}
		
		if commentID.Valid {
			comment := Comment{
				ID:     int(commentID.Int64),
				PostId: int(commentPostId.Int64),
				Text:   commentText.String,
			}
			post.Comments = append(post.Comments, comment)
		}
		
		postMap[PostId] = post
	}

    // username := r.FormValue("username")
    // password := r.FormValue("password")
    // fmt.Println("Username:", username, "Password:", password)
    // var storedHashedPassword string
    // err = database.Db.QueryRow("SELECT hashed_password FROM users WHERE name = ?", username).Scan(&storedHashedPassword)
    // if err == sql.ErrNoRows {
    //     http.Error(w, "User not found", http.StatusUnauthorized)
    //     return
    // } else if err != nil {
	// 	fmt.Println(err)
    //     http.Error(w, "Database error", http.StatusInternalServerError)
    //     return
    // }
    // err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
    // if err != nil {
    //     http.Error(w, "Incorrect password", http.StatusUnauthorized)
    //     return
    // }

	var posts []Post
	for _, post := range postMap {
		posts = append(posts, post)
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Posts": posts,
	})

}


// func Login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		tmpl, err := template.ParseGlob("./static/templates/*")
// 		if err != nil {
// 			http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		tmpl.ExecuteTemplate(w, "login.html", nil)
// 		return
// 	}

// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	var storedHashedPassword string

// 	err := database.Db.QueryRow("SELECT hashed_password FROM users WHERE username = ?", username).Scan(&storedHashedPassword)
// 	if err == sql.ErrNoRows {
// 		http.Error(w, "User not found", http.StatusUnauthorized)
// 		return
// 	} else if err != nil {
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(password))
// 	if err != nil {
// 		http.Error(w, "Incorrect password", http.StatusUnauthorized)
// 		return
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)
// }

// func Regist(w http.ResponseWriter, r *http.Request) {
// 	tmpl, err := template.ParseGlob("./static/templates/*")
// 	if err != nil {
// 		http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl.ExecuteTemplate(w, "regist.html", nil)
// }

// func hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {

// 		username := r.FormValue("username")
// 		email := r.FormValue("email")
// 		password := r.FormValue("password")
// 		confirmPassword := r.FormValue("confirmPassword")

// 		if password != confirmPassword {
// 			http.Error(w, "Passwords do not match", http.StatusBadRequest)
// 			return
// 		}

// 		hashedPassword, err := hashPassword(password)
// 		if err != nil {
// 			http.Error(w, "Error hashing password", http.StatusInternalServerError)
// 			return
// 		}

// 		_, err = database.Db.Exec("INSERT INTO users (name, email, hashed_password) VALUES (?, ?, ?)", username, email, hashedPassword)
// 		if err != nil {
// 			http.Error(w, "Error creating user", http.StatusInternalServerError)
// 			return
// 		}

// 		http.Redirect(w, r, "/login", http.StatusSeeOther)
// 	} else {
// 		http.ServeFile(w, r, "templates/register.html")
// 	}
// }
