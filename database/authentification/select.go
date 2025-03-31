package authdatabase

import (
	"database/sql"
	"errors"
	"forum/database"
	"forum/models"
)

func AddNewUser(username, email, hashedPass string) error {
	query := `INSERT INTO users (username, email, password_hash) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, username, email, hashedPass)
	return err
}

func SelectUserSession(session_id string) (int, bool, error) {
	var user_id int

	query := `SELECT user_id FROM sessions WHERE session_token = ?`
	err := database.DB.QueryRow(query, session_id).Scan(&user_id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return user_id, false, nil // not logged in
		}
		return user_id, false, err // err in database
	}
	return user_id, true, nil // logged in
}

func AlreadyExists(username, email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = ? OR email = ?`
	err := database.DB.QueryRow(query, username, email).Scan(&count)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}
	return count > 0, nil
}

func GetUserHashByUsername(username string) (int, string, error) {
	var hash string
	var id int
	query := `SELECT id,password_hash FROM users WHERE username = ? OR email = ?`
	err := database.DB.QueryRow(query, username, username).Scan(&id, &hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return id, hash, nil
		}
		return id, hash, err
	}
	return id, hash, nil
}

func GetUserHashById(id int) (string, error) {
	var hash string
	query := `SELECT password_hash FROM users WHERE id = ?`
	err := database.DB.QueryRow(query, id).Scan(&hash)
	if err != nil {
		return hash, err
	}
	return hash, nil
}

func DupplicatedUsername(username string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = ?`
	err := database.DB.QueryRow(query, username).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func DupplicatedEmail(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = ?`
	err := database.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetUserInfo(user_id int) (models.User, error) {
	var user models.User

	user_query := `SELECT * FROM users WHERE id = ?`
	err := database.DB.QueryRow(user_query, user_id).Scan(&user.Id, &user.Username, &user.Email, &user.Password_hash, &user.Created_at, &user.Updated_at)
	if err != nil {
		return user, err
	}
	return user, nil
}
