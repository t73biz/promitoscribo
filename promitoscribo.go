// main.go
package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/t73biz/promitoscribo/conf"
)

func main() {
	devroot := conf.GetDevRoot()
	server := martini.Classic()
	server.Use(martini.Static(devroot + "/public"))
	server.Use(render.Renderer(render.Options{
		Directory: devroot + "/views",
		Layout:    "layout",
	}))

	server.Get("/", func(r render.Render) {
		newmap := map[string]interface{}{
			"metatitle": "Promito Scribo",
		}
		r.HTML(200, "hello", newmap)

	})

	server.Run()
}
