package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "internal/scripts/kubernetes/deploy.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Comments:")
	// for _, c := range node.Comments {
	// 	fmt.Print(c.Text())
	// }

	// fmt.Println("Functions:")

	// for _, f := range node.Decls {
	// 	fn, ok := f.(*ast.FuncDecl)
	// 	if !ok {
	// 		continue
	// 	}
	// 	fmt.Println(fn.Name.Name)
	// 	fmt.Println(fn.)
	// }

	// spew.Dump(f)

	cmap := ast.NewCommentMap(fset, f, f.Comments)

	// fmt.Printf("%#v\n", cmap)

	v := visitor{commentMap: cmap}
	ast.Walk(v, f)
}

type visitor struct {
	commentMap ast.CommentMap
}

func (v visitor) Visit(n ast.Node) ast.Visitor {
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
		fmt.Println(stmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value)

		if comments := v.commentMap.Filter(n).Comments(); len(comments) > 0 {
			for _, c := range comments {
				fmt.Println(c.Text())
			}
		}
		// fmt.Println(stmt.Rhs[1].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value)
	} else if functionName == "RequiredConfigFetch" {
		fmt.Println(stmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value)

		if comments := v.commentMap.Filter(n).Comments(); len(comments) > 0 {
			for _, c := range comments {
				fmt.Println(c.Text())
			}
		}

	}
	// fmt.Println(stmt.Rhs[0].(*ast.CallExpr).Fun.(*ast.SelectorExpr).X)
	// fmt.Println(stmt.Rhs[0].(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.String())
	// if (stmt.Rhs[0].(*ast.CallExpr).Fun.(*ast.SelectorExpr).Sel.String())
	// // spew.Dump(v.commentMap[n])
	// // spew.Dump(stmt)
	// if comments := v.commentMap.Filter(n).Comments(); len(comments) > 0 {
	// 	for _, c := range comments {
	// 		fmt.Println(c.Text())
	// 	}
	// }

	// fmt.Println(stmt.Rhs[0].(*ast.CallExpr).Args[0].(*ast.BasicLit).Value)

	return v
}
