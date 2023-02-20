package routes

import (
	"book-store/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	routes.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	routes.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("Delete")
}
