package main

import (
  "html/template"
  "os"
  "strings"
  "math"
)

func GetCategories(products []Product) (categories []string) {
  catMap := map[string]string{}
  for _, p := range products {
    if (catMap[p.Category] == "") {
      catMap[p.Category] = p.Category
      categories = append(categories, p.Category)
    }
  }
  return
}

func Exec(t *template.Template) error {
  productMap := map[string]Product {}
  for _, p := range Products {
    productMap[p.Name] = p
  }
  return t.Execute(os.Stdout, &productMap)
}

func main() {
  allTemplates := template.New("allTemplates")
  allTemplates.Funcs(map[string]interface{} {
    "getCats": GetCategories,
    "lower": strings.ToLower,
    "sqrt": math.Sqrt,
  })
  allTemplates, err := allTemplates.ParseGlob("templates/*.html")

  if (err == nil) {
    selectedTemplated := allTemplates.Lookup("mainTemplate")
    err = Exec(selectedTemplated)
  }
  if (err != nil) {
    Printfln("Error: %v", err.Error())
  }
}
