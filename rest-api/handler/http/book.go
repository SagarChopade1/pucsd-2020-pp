package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pucsd2020-pp/rest-api/handler"
	"github.com/pucsd2020-pp/rest-api/model"
	"github.com/pucsd2020-pp/rest-api/repository"
	"github.com/pucsd2020-pp/rest-api/repository/book"
)

type Books struct {
	handler.HTTPHandler
	repo repository.IRepository
}

func NewBooksHandler(conn *sql.DB) *Books {
	return &Books{
		repo: book.NewBooksRepository(conn),
	}
}

func (book *Books) GetHTTPHandler() []*handler.HTTPHandler {
	return []*handler.HTTPHandler{
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "books/{id}", Func: book.GetByID},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPost, Path: "books", Func: book.Create},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodPut, Path: "books/{id}", Func: book.Update},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodDelete, Path: "books/{id}", Func: book.Delete},
		&handler.HTTPHandler{Authenticated: true, Method: http.MethodGet, Path: "books", Func: book.GetAll},
		///write path =books/{id}
	}
}

func (book *Books) GetByID(w http.ResponseWriter, r *http.Request) {
	var usr interface{}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		usr, err = book.repo.GetByID(r.Context(), id)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (book *Books) Create(w http.ResponseWriter, r *http.Request) {
	var usr model.Books
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}

		_, err = book.repo.Create(r.Context(), usr)
		break
	}
	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (book *Books) Update(w http.ResponseWriter, r *http.Request) {
	var iUsr interface{}
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	usr := model.Books{}
	err := json.NewDecoder(r.Body).Decode(&usr)
	for {
		if nil != err {
			break
		}
		usr.Id = id
		if nil != err {
			break
		}

		// set logged in Book id for tracking update
		//usr.UpdatedBy = 0

		iUsr, err = book.repo.Update(r.Context(), usr)
		if nil != err {
			break
		}
		usr = iUsr.(model.Books)
		break
	}

	handler.WriteJSONResponse(w, r, usr, http.StatusOK, err)
}

func (book *Books) Delete(w http.ResponseWriter, r *http.Request) {
	var payload string
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	for {
		if nil != err {
			break
		}

		err = book.repo.Delete(r.Context(), id)
		if nil != err {
			break
		}
		payload = "Book deleted successfully"
		break
	}

	handler.WriteJSONResponse(w, r, payload, http.StatusOK, err)
}

func (book *Books) GetAll(w http.ResponseWriter, r *http.Request) {
	usrs, err := book.repo.GetAll(r.Context())
	handler.WriteJSONResponse(w, r, usrs, http.StatusOK, err)
}
