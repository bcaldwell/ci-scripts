package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type configValue struct {
	parameter    string
	description  string
	required     bool
	defaultValue string
}

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "internal/scripts/kubernetes/deploy.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)

	v := &visitor{commentMap: cmap, configValues: make([]configValue, 0)}
	ast.Walk(v, f)

	v.Table(os.Stdout)
}

type visitor struct {
	commentMap   ast.CommentMap
	configValues []configValue
}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	stmt, ok := n.(*ast.AssignStmt)
	if !ok {
		return v
	}

	rhs, ok := stmt.Rhs[0].(*ast.CallExpr)
	if !ok {
		return v
	}

	expr, ok := rhs.Fun.(*ast.SelectorExpr)
	if !ok {
		return v
	}

	functionName := expr.Sel.String()
	if functionName == "ConfigFetch" {
		c := configValue{
			parameter: stmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value,
			required:  false,
		}

		if len(stmt.Rhs[0].(*ast.CallExpr).Args) >= 2 {
			// spew.Dump(stmt)
			c.defaultValue = stmt.Rhs[0].(*ast.CallExpr).Args[1].(*ast.BasicLit).Value
		}

		if comments := v.commentMap.Filter(n).Comments(); len(comments) > 0 {
			for _, com := range comments {
				c.description = strings.TrimSpace(com.Text())
			}
		}

		v.configValues = append(v.configValues, c)
	} else if functionName == "RequiredConfigFetch" {
		c := configValue{
			parameter: stmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value,
			required:  true,
		}

		if comments := v.commentMap.Filter(n).Comments(); len(comments) > 0 {
			for _, com := range comments {
				c.description = strings.TrimSpace(com.Text())
			}
		}

		v.configValues = append(v.configValues, c)
	}

	return v
}

func (v *visitor) Table(w io.Writer) {
	if len(v.configValues) == 0 {
		return
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(v.configValues, func(i, j int) bool {
		return v.configValues[i].parameter < v.configValues[j].parameter
	})

	for _, v := range v.configValues {
		fmt.Println(v.parameter, v.description, v.required, v.defaultValue)
	}
}
