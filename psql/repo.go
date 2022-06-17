package psql

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/topheruk-go/util/sql/pg"
	bc "go.topheruk.bookcatalogue"
	"go.topheruk.bookcatalogue/pkg/isbn"
)

type Repo struct {
	db *pgx.Conn
}

func NewRepo(conn *pgx.Conn) *Repo {
	return &Repo{conn}
}

var findBookQry string

func (r *Repo) FindBookContext(ctx context.Context, id isbn.ISBN) (book *bc.Book, err error) {
	err = pg.QueryRow(r.db, findBookQry, func(row pgx.Row) error {
		return row.Scan(
			&book.Isbn,
			&book.Title,
			&book.Author,
			&book.CreatedAt)
	})

	return
}

func (r *Repo) FindBook(id isbn.ISBN) (*bc.Book, error) {
	return r.FindBookContext(context.Background(), id)
}

var listBooksQry string

func (r *Repo) ListBooksContext(ctx context.Context, args ...any) ([]bc.Book, error) {
	return pg.Query(r.db, listBooksQry, func(rows pgx.Rows, book *bc.Book) error {
		return rows.Scan(
			&book.Isbn,
			&book.Title,
			&book.Author,
			&book.CreatedAt)
	})
}
