/*
* @Author: souravray
* @Date:   2014-10-20 16:21:45
* @Last Modified by:   souravray
* @Last Modified time: 2014-10-21 10:02:41
 */

package router

import (
	"github.com/gorilla/mux"
	"github.com/souravray/gadgetWish/controllers"
	"net/http"
)

func Routes(rtr *mux.Router) {
	//static assets
	rtr.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("./static/js"))))
	rtr.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("./static/img"))))
	rtr.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./static/css"))))
	rtr.PathPrefix("/libs/").Handler(http.StripPrefix("/libs/", http.FileServer(http.Dir("./static/libs"))))
	rtr.PathPrefix("/partials/").Handler(http.StripPrefix("/partials/", http.FileServer(http.Dir("./static/partials"))))

	//Home page
	rtr.HandleFunc("/", controllers.Landing).Methods("GET").Name("Homepage")

	// API routes
	apiSubrtr := rtr.PathPrefix("/api").Subrouter()
	apiSubrtr.HandleFunc("/user", controllers.SignIn).Methods("POST").Name("LogIn")
	apiSubrtr.HandleFunc("/user", controllers.UserInfo).Methods("GET").Name("GetUser")
	apiSubrtr.HandleFunc("/user", controllers.SignOut).Methods("DELETE").Name("LogOut")
	apiSubrtr.HandleFunc("/products", controllers.ProductList).Methods("GET").Name("ProductList")
	apiSubrtr.HandleFunc("/bucket", controllers.WishList).Methods("GET").Name("WishList")
	apiSubrtr.HandleFunc("/bucket", controllers.AddWish).Methods("POST").Name("AddToWishList")
	apiSubrtr.HandleFunc("/bucket", controllers.RemoveWish).Methods("DELETE").Name("RemoveFromWishList")
}
