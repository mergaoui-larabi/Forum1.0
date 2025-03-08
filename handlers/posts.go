package handlers

import (
	"log"
	"net/http"
	// "fmt"

	"forum/database"
	// "text/template"
	_ "github.com/mattn/go-sqlite3"
)


func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	UserId := 1
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		return
	}

	_, err := database.Db.Exec("INSERT INTO post (user_id, content) VALUES (?, ?)", UserId, content)
	if err != nil {
		log.Print(err)
		http.Error(w, "Failed to insert post ", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/page", http.StatusSeeOther)
}

// func ShowPosts(w http.ResponseWriter, r *http.Request){
// 	data, err := database.Db.Query("SELECT content FROM post")
// 	if err != nil{
// 		fmt.Println("Error opening or getting the posts data : ", err)
// 		return 
// 	}
// 	defer data.Close()
// 	contents := []string{}
// 	for data.Next() {
// 		var content string
// 		if err := data.Scan(&content); err != nil {
// 			log.Println("Error scanning row:", err)
// 			continue
// 		}
// 		// fmt.Println(content)
// 		contents = append(contents, content)
// 	}
// 	// return contents
// 	// type Post struct {
// 	// 	ID      int
// 	// 	Content string
// 	// }
// 	// data1 := struct {
// 	// 	Posts []Post
// 	// }{
// 	// 	Posts : contents,
// 	// }
// 	// tmpl.Execute(w, contents)

// }