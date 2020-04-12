package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/pucsd2020-pp/rest-api/config"
	"github.com/pucsd2020-pp/rest-api/driver"
	"github.com/pucsd2020-pp/rest-api/handler"
	httpHandler "github.com/pucsd2020-pp/rest-api/handler/http"
)

var (
	handlers = []handler.IHTTPHandler{}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	dbConn, err := driver.NewMysqlConnection(config.Config().Database)
	if nil != err {
		log.Printf("Error while creating db connectiion:%s", err.Error())
		os.Exit(1)
	}

	handlers = []handler.IHTTPHandler{
		httpHandler.NewUserHandler(dbConn),
		httpHandler.NewBooksHandler(dbConn),
	}
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("The time is: "))
}

func createRouterGroup(router *chi.Mux) {
	th := &timeHandler{format: time.RFC1123}
	router.Group(func(r chi.Router) {
		r.Handle("/time", th)
		for _, hdlr := range handlers { // register all handlers
			for _, hlr := range hdlr.GetHTTPHandler() {
				path := fmt.Sprintf("/webapi/v1/%s", hlr.Path)
				switch hlr.Method {
				case http.MethodGet:
					r.Get(path, hlr.Func)
				case http.MethodPost:
					r.Post(path, hlr.Func)
				case http.MethodPut:
					r.Put(path, hlr.Func)
				case http.MethodDelete:
					r.Delete(path, hlr.Func)
				default:
					log.Println("Invalid method")
				}
			}
		}
	})
}

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	createRouterGroup(router)

	http.ListenAndServe(fmt.Sprintf("%s:%d",
		config.Config().Host, config.Config().Port), router)
}
