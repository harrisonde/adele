package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) WebRoutes() http.Handler {

	r := chi.NewRouter()

	/*
	|--------------------------------------------------------------------------
	| Web Middleware
	|--------------------------------------------------------------------------
	|
	| Here is where you can add your Middleware for the web routes.
	| These middleware are called on each web route request.
	|
	*/
	r.Use(a.App.NoSurf)

	r.Route("/", func(mux chi.Router) {

		/*
		|--------------------------------------------------------------------------
		| Web Routes
		|--------------------------------------------------------------------------
		|
		| Here is where you can add your web routes for the application. These
		| routes are loaded by the router.
		|
		*/

		r.Get("/", a.Handlers.Home)

	})
	return r
}
