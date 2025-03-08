package handlers

import (
	// "database/sql"
	"fmt"
	"net/http"
	"text/template"
	"log"

	"forum/database"

	// "golang.org/x/crypto/bcrypt"
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
	data, err := database.Db.Query("SELECT content FROM post")
	if err != nil{
		fmt.Println("Error opening or getting the posts data : ", err)
		return 
	}
	defer data.Close()
	contents := []string{}
	for data.Next() {
		var temp string
		if err := data.Scan(&temp); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		// fmt.Println(content)
		contents = append(contents, temp)
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

    data1 := struct {
        IsLoggedIn bool
		Posts []string
    }{
        IsLoggedIn: true,
		Posts: contents,
    }

    // fmt.Println("User authenticated:", data.IsLoggedIn)
	tmpl, err1 := template.ParseFiles("./static/templates/index.html")
    if err1 != nil {
		fmt.Println("hnaa")
        log.Fatal(err1)
    }

    // if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
    //     http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
    //     return
    // }
	err = tmpl.Execute(w, data1)
	if err != nil{
		log.Println("Error executing template:", err)
	}

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
