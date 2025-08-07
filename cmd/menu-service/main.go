package menuservice

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/javy99/oms/internal/server"
)

func main() {
	// 1. All the common setup is not handled by our package.
	srv := server.New(":8080")

	// 2. This service only needs to know about its own specific routes.
	srv.AddRoutes(registerMenuRoutes)

	// 3. Start the server
	srv.Start()
}

// registerMenuRoutes contains logic specific to the menu service.
func registerMenuRoutes(r *mux.Router) {
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Menu service is up!"))
	}).Methods(http.MethodGet)
	// r.HandleFunc("/items", getItemsHandler).Methods(http.MethodGet)
}
