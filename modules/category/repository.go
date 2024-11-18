package category

import (
	"database/sql"
	"sb-go-quiz-week-3/modules/book"
)

func GetAllCategoryFrom(db *sql.DB) (result []Category, err error) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category Category

		err = rows.Scan(&category.Id,
			&category.Name,
			&category.Created_at,
			&category.Created_by,
			&category.Modified_at,
			&category.Modified_by,
		)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func InsertCategoryTo(db *sql.DB, category Category) (err error) {
	sql := `
		INSERT INTO categories(name, created_by, modified_by)
		VALUES ($1, $2, $3) 
		`

	errs := db.QueryRow(sql, category.Name, category.Created_by, category.Modified_by)

	return errs.Err()
}

func GetCategoryByIdFrom(db *sql.DB, targetCategory Category) (category Category, err error) {
	sql := "SELECT * FROM categories WHERE id = $1"

	rows, err := db.Query(sql, targetCategory.Id)
	if err != nil {
		return
	}

	defer rows.Close()
	err = rows.Scan(&category.Id,
		&category.Name,
		&category.Created_at,
		&category.Created_by,
		&category.Modified_at,
		&category.Modified_by,
	)
	if err != nil {
		return
	}
	return
}

func DeleteCategoryFrom(db *sql.DB, category Category) (err error) {
	sql := "DELETE FROM categories WHERE id = $1"

	errs := db.QueryRow(sql, category.Id)
	return errs.Err()
}

func GetAllBookInCategoryFrom(db *sql.DB, category Category) (result []book.Book, err error) {
	sql := "SELECT * FROM books WHERE category_id = $1"

	rows, err := db.Query(sql, category.Id)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book book.Book

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
