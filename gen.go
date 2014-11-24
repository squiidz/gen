package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/squiidz/fur/middle"
)

type App struct {
	Name    string
	Version string
	Routes  []Route
}

type Route struct {
	Url      string
	Method   string
	Template string
}

func NewGen(filename string) *App {
	app := &App{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(file, app)
	if err != nil {
		fmt.Println(err)
	}
	return app
}

func (a *App) Run(port string) {
	router := mux.NewRouter()

	for _, r := range a.Routes {
		file, err := ioutil.ReadFile(r.Template)

		if err != nil {
			fmt.Println("[x]", r.Template, "is not valid.")
			continue
		}

		router.HandleFunc(r.Url, func(rw http.ResponseWriter, req *http.Request) {
			rw.Write(file)
		}).Methods(r.Method)
	}

	fmt.Println("[+] Server Listening on", port[1:])
	http.ListenAndServe(port, middle.Logger(router))

}
