package controllers

import(
  "github.com/kataras/iris"
  "strings"
)

type Render struct {
  Param string
}

func Param_index (ctx *iris.Context){

  id := ctx.Param("id")
  id = strings.Replace(id, "/", "", 1)

  data := Render{
    Param: id,
  }

ctx.MustRender("param_index.jade", data)

}
