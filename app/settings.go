package app

import(
  "github.com/kataras/config"
)

templateConfig := config.Template {
        Engine:        iris.HandlebarsEngine,
        Gzip:          false,
        IsDevelopment: false,
        Directory:     "templates",
        Extensions:    []string{".html"},
        ContentType:   "text/html",
        Charset:       "UTF-8",
        Layout:        "", // currently this is the only config which not working for pongo2 yet but I will find a way
        HTMLTemplate:  HTMLTemplate{Left: "{{", Right: "}}", Funcs: make(map[string]interface{}, 0), LayoutFuncs: make(map[string]interface{}, 0)},
        Jade:          Jade{Left: "{{", Right: "}}", Funcs: make(map[string]interface{}, 0), LayoutFuncs: make(map[string]interface{}, 0)},
        Pongo:         Pongo{Filters: make(map[string]pongo2.FilterFunction, 0), Globals: make(map[string]interface{}, 0)},
        Markdown:      Markdown{Sanitize: false},
        Amber:         Amber{Funcs: template.FuncMap{}},
        Handlebars:    Handlebars{Helpers: make(map[string]interface{}, 0)},
}
