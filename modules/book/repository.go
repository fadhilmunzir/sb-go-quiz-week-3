package book

import (
	"database/sql"
)

func GetAllBookFrom(db *sql.DB) (result []Book, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book Book

		err = rows.Scan(&book.Id,
			&book.Title,
			&book.Description,
			&book.Image_url,
			&book.Release_year,
			&book.Price,
			&book.Total_page,
			&book.Thickness,
			&book.Category_id,
			&book.Created_at,
			&book.Created_by,
			&book.Modified_at,
			&book.Modified_by,
		)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func InsertBookTo(db *sql.DB, book Book) (err error) {
	sql := `
		INSERT INTO books(title, description, image_url, release_year, price, total_page, thickness, category_id, created_by, modified_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) 
		`

	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Category_id, book.Created_by, book.Modified_by)

	return errs.Err()
}

func GetBookByIdFrom(db *sql.DB, targetBook Book) (book Book, err error) {
	sql := "SELECT * FROM books WHERE id = $1"

	rows, err := db.Query(sql, targetBook.Id)
	if err != nil {
		return
	}

	defer rows.Close()
	err = rows.Scan(&book.Id,
		&book.Title,
		&book.Description,
		&book.Image_url,
		&book.Release_year,
		&book.Price,
		&book.Total_page,
		&book.Thickness,
		&book.Category_id,
		&book.Created_at,
		&book.Created_by,
		&book.Modified_at,
		&book.Modified_by,
	)
	if err != nil {
		return
	}
	return
}

func DeleteBookFrom(db *sql.DB, book Book) (err error) {
	sql := "DELETE FROM books WHERE id = $1"

	errs := db.QueryRow(sql, book.Id)
	return errs.Err()
}
