package handlers

import ( // "fmt"
	// "text/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func LikeHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	return
}

// func AddPostHandler(w http.ResponseWriter, r *http.Request) {
// 	user_id := r.Context().Value(userIDKey).(int)
// 	if r.Method != "POST" {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	content := r.FormValue("content")
// 	if content == "" {
// 		return
// 	}

// 	err := database.InsertPost(user_id, content)
// 	if err != nil {
// 		log.Print(err)
// 		http.Error(w, "Failed to insert post ", http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/page", http.StatusSeeOther)
// }

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
