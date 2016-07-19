package app

import(
  "github.com/ThomasVictoria/starter-go/app/controllers"
  "github.com/kataras/iris"
)

func Router () {

  iris.Config.Render.Template.Extensions = []string{".jade"}
  iris.Config.Render.Template.Layout = "layouts/layout.jade"
  iris.Config.Render.Template.Engine = iris.JadeEngine

  iris.Get("/index", controllers.Index_index)

  iris.Get("/param/*id", controllers.Param_index)

  iris.Listen(":8080")

}
