package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/format"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed curry.go.tpl
var codeTemplateContents string
var codeTemplate = template.Must(template.New("main").Parse(codeTemplateContents))

type Data struct {
	PackageName string
	Funcs       []Func
}

var numbers = [...]string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

type Func struct {
	Name              string
	CurriedArgIndexes []int
	ArgIndexes        []int
	RetIndexes        []int
	Variadic          bool
	InputArguments    int
}

func main() {
	var (
		rootPath string
	)
	flag.StringVar(&rootPath, "root-path", "", "Root path from where packages have to be created")
	flag.Parse()

	if rootPath == "" {
		log.Fatalf("-root-path has to be provided.")
	}

	if _, err := os.ReadDir(rootPath); err != nil {
		log.Fatalf("Error reading -root-path=%s: %s", rootPath, err)
	}

	for curried := 1; curried <= 5; curried++ {
		data := Data{PackageName: "curry" + strings.ToLower(numbers[curried])}
		for args := curried; args <= 5; args++ {
			for rets := 0; rets <= 5; rets++ {
				name := fmt.Sprintf("From%sArg%s", numbers[args], plural(args))
				variadicName := name + "Variadic"
				if rets > 0 {
					suffix := fmt.Sprintf("%sRet%s", numbers[rets], plural(rets))
					name += suffix
					variadicName += suffix
				}
				data.Funcs = append(data.Funcs, Func{
					Name:              name,
					CurriedArgIndexes: indexes(curried),
					ArgIndexes:        indexes(args - curried),
					RetIndexes:        indexes(rets),
					InputArguments:    args,
				})

				data.Funcs = append(data.Funcs, Func{
					Name:              variadicName,
					CurriedArgIndexes: indexes(curried),
					ArgIndexes:        indexes(args - curried),
					RetIndexes:        indexes(rets),
					Variadic:          true,
					InputArguments:    args,
				})
			}
		}

		buf := bytes.NewBuffer(nil)
		if err := codeTemplate.Execute(buf, data); err != nil {
			log.Fatalf("Can't execute template: %s", err)
		}
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			log.Printf("Can't format output:")
			for i, line := range strings.Split(string(buf.Bytes()), "\n") {
				fmt.Fprintf(os.Stderr, "Line %3d: %s\n", i+1, line)
			}
			log.Fatalf("Error was: %s", err)
		}

		dir := filepath.Join(rootPath, data.PackageName)
		if err := os.MkdirAll(dir, fs.ModePerm); err != nil {
			log.Fatalf("Error creating path: %s", err)
		}
		filename := filepath.Join(dir, "curry.go")
		if err := os.WriteFile(filename, formatted, 0600); err != nil {
			log.Fatalf("Error creating file: %s", err)
		}
	}

}

func plural(n int) string {
	if n > 1 {
		return "s"
	}
	return ""
}

func indexes(n int) []int {
	ret := make([]int, n)
	for i := range ret {
		ret[i] = i + 1
	}
	return ret
}
