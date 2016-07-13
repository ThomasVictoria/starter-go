package app

import(
  "github.com/kataras/iris"
)

func Router () {

  iris.Get("/", index)

  iris.Listen(":8000")

}
