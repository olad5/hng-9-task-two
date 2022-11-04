package router

import (
	"fmt"
	"github.com/olad5/hng-9-task-two/handlers"
	"github.com/olad5/hng-9-task-two/middleware"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		{
			handlers.GetUserProfile(w, r)
		}
	case http.MethodPost:
		{
			handlers.PerformArithemticOperation(w, r)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func Initialize() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	return middleware.Json(mux)
}
