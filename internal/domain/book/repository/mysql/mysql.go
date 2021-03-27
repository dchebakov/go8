package mysql

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/gmhafiz/go8/internal/domain/book"
	"github.com/gmhafiz/go8/internal/models"
	"github.com/gmhafiz/go8/internal/utility/message"
)

type repository struct {
	db *sqlx.DB
}

const (
	InsertIntoBooks         = "INSERT INTO books (title, published_date, image_url, description) VALUES (?,?,?,?)"
	SelectFromBooks         = "SELECT * FROM books ORDER BY created_at DESC"
	SelectFromBooksPaginate = "SELECT * FROM books ORDER BY created_at DESC LIMIT ? OFFSET ?"
	SelectBookByID          = "SELECT * FROM books where book_id = ?"
	UpdateBook              = "UPDATE books set title = ?, description = ?, published_date = ?, image_url = ?, updated_at = ? where book_id = ?"
	DeleteByID              = "DELETE FROM books where book_id = (?)"
	SearchBooks             = "SELECT * FROM books where title like '%' || ? || '%' and description like '%'|| ? || '%' ORDER BY published_date DESC"
	SearchBooksPaginate     = "SELECT * FROM books where title like '%' || ? || '%' and description like '%'|| ? || '%' ORDER BY published_date DESC LIMIT ? OFFSET ?"
)

func New(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(ctx context.Context, book *models.Book) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx, InsertIntoBooks)

	if err != nil {
		return 0, err
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	result, err := stmt.ExecContext(ctx, book.Title, book.PublishedDate, book.ImageURL, book.Description)
	if err != nil {
		return 0, err
	}

	bookID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	return bookID, nil
}

func (r *repository) List(ctx context.Context, f *book.Filter) ([]*models.Book, error) {
	if f.Base.DisablePaging {
		var books []*models.Book
		err := r.db.SelectContext(ctx, &books, SelectFromBooks)
		if err != nil {
			return nil, message.ErrFetchingBook
		}

		return books, nil
	} else {
		var books []*models.Book
		err := r.db.SelectContext(ctx, &books, SelectFromBooksPaginate, f.Base.Size, f.Base.Page)
		if err != nil {
			return nil, message.ErrFetchingBook
		}

		return books, nil
	}
}

func (r *repository) Read(ctx context.Context, bookID int64) (*models.Book, error) {
	stmt, err := r.db.Prepare(SelectBookByID)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var b models.Book
	err = r.db.GetContext(ctx, &b, SelectBookByID, bookID)
	if err != nil {
		return nil, err
	}

	return &b, err
}

func (r *repository) Update(ctx context.Context, book *models.Book) error {
	now := time.Now()

	_, err := r.db.ExecContext(ctx, UpdateBook, book.Title, book.Description,
		book.PublishedDate, book.ImageURL, now, book.BookID)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, bookID int64) error {
	_, err := r.db.ExecContext(ctx, DeleteByID, bookID)

	return err
}

func (r *repository) Search(ctx context.Context, f *book.Filter) ([]*models.Book, error) {
	var books []*models.Book
	err := r.db.SelectContext(ctx, &books, SearchBooksPaginate, f.Title, f.Description,
		f.Base.Size,
		f.Base.Page)
	if err != nil {
		return nil, err
	}

	return books, nil
}
