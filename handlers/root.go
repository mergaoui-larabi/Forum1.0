package handlers

import (
	"forum/config"
	"forum/database"
	"net/http"
	"text/template"
	"fmt"
	"database/sql"
	"log"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("session_token")
	if err != nil || sessionCookie.Value == "" {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]interface{}{"Authenticated": false})
		return
	}

	user_id, exist := database.GetUserBySession(sessionCookie.Value)

	if !exist {
		config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]interface{}{"Authenticated": false})
		return
	}


	if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form: "+err.Error(), http.StatusBadRequest)
        return
    }

    tmpl, err := template.ParseGlob("./static/templates/*")
    if err != nil {
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        return
    }

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
	
	data, err := database.DB.Query(`
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

	var posts []Post
	for _, post := range postMap {
		posts = append(posts, post)
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Posts": posts,
	})


	user := database.GetUserInfo(user_id)

	config.GLOBAL_TEMPLATE.ExecuteTemplate(w, "index.html", map[string]interface{}{"Authenticated": true, "Username": user.Username})

}
