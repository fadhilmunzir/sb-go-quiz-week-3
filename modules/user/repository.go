package user

import "database/sql"

func GetUserByIdFrom(db *sql.DB, targetUser User) (user User, err error) {
	sql := "SELECT * FROM categories WHERE id = $1"

	rows, err := db.Query(sql, targetUser.Id)
	if err != nil {
		return
	}

	defer rows.Close()
	err = rows.Scan(&user.Id,
		&user.Username,
		&user.Password,
		&user.Created_at,
		&user.Created_by,
		&user.Modified_at,
	)
	if err != nil {
		return
	}
	return
}
