package main

import(
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main(){
//	name := "jocker"
	tmpl, err := template.New("test").Parse("hello,{{.}}")
	if err != nil {
		panic(err)
	}
	err =  tmpl.Execute(os.Stdout,"")
	if err != nil {
		panic(err)
	}
}