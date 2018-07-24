package main

import (
	"encoding/json"
	"net/http"

	"github.com/acoshift/hime"
)

func main() {
	http.Handle("/", hime.H(index))
	http.Handle("/script.js", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "script.js")
	}))

	app := hime.New()

	app.TemplateFunc("json", func(v interface{}) string {
		b, _ := json.Marshal(v)
		return string(b)
	})

	app.Template().
		Dir("view").
		Root("root").
		Parse("index", "index.html")

	app.Handler(http.DefaultServeMux).
		Address(":8080").
		ListenAndServe()
}

func index(ctx *hime.Context) hime.Result {
	return ctx.View("index", map[string]interface{}{
		"Username": "acoshift",
		"List": []string{
			"Go",
			"Vue",
			"PostgreSQL",
		},
	})
}
