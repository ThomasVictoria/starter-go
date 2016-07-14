package app

import(
  "github.com/kataras/iris"
)

func Router () {

  iris.Config.Render.Template.Extensions = []string{".jade"}
iris.Config.Render.Template.Layout = "layouts/layout.jade"
  iris.Config.Render.Template.Engine = iris.JadeEngine

  iris.Get("/index", index)

  iris.Listen(":8080")

}
