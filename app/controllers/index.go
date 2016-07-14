package controllers

import(
  "github.com/kataras/iris"
)

type Title struct{
  First string
  Second []*Subtitle
}

type Subtitle struct{
  Title string
  Description string
}

func Index_index (ctx *iris.Context) {

  first := Subtitle{Title: "First Subtitle", Description: "I'm the first subtitle"}
  second := Subtitle{Title: "Second Subtitle", Description: "I'm the second subtitle"}

  data := Title{
    First: "I am the big one",
    Second: []*Subtitle{&first, &second},
  }

  ctx.MustRender("index_index.jade", data)
}
