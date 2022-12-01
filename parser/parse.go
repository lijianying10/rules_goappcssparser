package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
)

type Parse struct {
	strs []string
}

func (parse *Parse) IO(inputFileList []string, output string) error {
	for _, fname := range inputFileList {
		f, err := ioutil.ReadFile(fname)
		if err != nil {
			return errors.New("during open file " + fname + " " + err.Error())
		}
		err = parse.AppendCodeGragment(string(f))
		if err != nil {
			return errors.New("during parse file " + fname + " " + err.Error())
		}
	}
	html, err := parse.Generate()
	if err != nil {
		return errors.New("during generate: " + err.Error())
	}
	err = ioutil.WriteFile(output, []byte(html), 0644)
	if err != nil {
		return errors.New("during write result: " + err.Error())
	}
	return nil
}

// AppendCodeGragment append golang code fragment for go parser
func (parse *Parse) AppendCodeGragment(src string) error {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", src, parser.ParseComments)
	if err != nil {
		return err
	}

	ast.Inspect(f, func(n ast.Node) bool {
		// Called recursively.
		//ast.Print(fset, n)
		if _, ok := n.(*ast.ImportSpec); ok {
			return false
		}
		if val, ok := n.(*ast.BasicLit); ok {
			if val.Kind == token.STRING {
				parse.strs = append(parse.strs, val.Value)
			}
		}
		return true
	})
	return nil
}

// Generate gen html code for tailwind post css
func (parse *Parse) Generate() (string, error) {
	divs := ""
	for _, v := range parse.strs {
		divs += fmt.Sprintf(`<div class="%s"></div>`+"\n", v)
	}
	return fmt.Sprintf(`<html>
    <body>
      %s
    </body>
</html>
`, divs), nil
}
