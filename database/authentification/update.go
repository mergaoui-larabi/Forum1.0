package authdatabase

import (
	"forum/database"
	"log"
)

func UpdateUserSession(id int, token string) error {
	updateQuery := `UPDATE sessions SET expires_at = DATETIME('now', '+1 hour'), session_token = ? WHERE user_id = ?`
	insertQuery := `INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, DATETIME('now', '+1 hour'))`

	tx, err := database.DB.Begin()
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return err
	}

	res, err := tx.Exec(updateQuery, token, id)
	if err != nil {
		log.Println("Update error:", err)
		tx.Rollback()
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		tx.Rollback()
		return err
	}

	if count > 0 {
		return tx.Commit() // if update succed
	}

	_, err = tx.Exec(insertQuery, id, token)
	if err != nil {
		log.Println("Insert error:", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func ResetUserSession(session_id string) (bool, error) {
	query := `UPDATE sessions SET session_token = NULL WHERE session_token = ?`
	_, err := database.DB.Exec(query, session_id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdatePassword(id int, password string) error {
	query := `UPDATE users SET updated_at = DATETIME('now'), password_hash = ? WHERE id = ?`
	res, err := database.DB.Exec(query, password, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

func UpdateUsernmae(id int, username string) error {
	query := `UPDATE users SET updated_at = DATETIME('now') , username = ? WHERE id = ?`
	res, err := database.DB.Exec(query, username, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

func UpdateEmail(id int, email string) error {
	query := `UPDATE users SET updated_at = DATETIME('now') , email = ? WHERE id = ?`
	res, err := database.DB.Exec(query, email, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err

}
