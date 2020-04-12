package book

import (
	"context"
	"database/sql"

	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/model"
)

type bookRepository struct {
	conn *sql.DB
}

func NewBooksRepository(conn *sql.DB) *bookRepository {
	return &bookRepository{conn: conn}
}

func (book *bookRepository) GetByID(cntx context.Context, id int64) (interface{}, error) {
	obj := new(model.Books)
	return driver.GetById(book.conn, obj, id)
}

func (book *bookRepository) Create(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.Books)
	result, err := driver.Create(book.conn, &usr)
	if nil != err {
		return 0, err
	}

	id, _ := result.LastInsertId()
	usr.Id = id
	return id, nil
}

func (book *bookRepository) Update(cntx context.Context, obj interface{}) (interface{}, error) {
	usr := obj.(model.Books)
	err := driver.UpdateById(book.conn, &usr)
	return obj, err
}

func (book *bookRepository) Delete(cntx context.Context, id int64) error {
	obj := &model.Books{Id: id}
	return driver.SoftDeleteById(book.conn, obj, id)
}

func (book *bookRepository) GetAll(cntx context.Context) ([]interface{}, error) {
	obj := &model.Books{}
	return driver.GetAll(book.conn, obj, 0, 0)
}
