package authdatabase

import (
	"forum/database"
)

func DeleteUser(user_id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := database.DB.Exec(query, user_id)
	if err != nil {
		return err
	}
	return nil
}
