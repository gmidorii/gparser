package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/constant"
	"go/parser"
)

func main() {
	ops := []string{"+", "-", "*", "/"}

	for _, v := range ops {
		if err := parse("10" + v + "2"); err != nil {
			fmt.Println(err)
		}
	}
	if err := parse(`"hoge"+"hoge"`); err != nil {
		fmt.Println(err)
	}
}

func parse(str string) error {
	fmt.Println(str)
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return err
	}

	// ast.Inspect(expr, func(n ast.Node) bool {
	// 	if n != nil {
	// 		fmt.Printf("%[1]v(%[1]T)\n", n)
	// 	}
	// 	return true
	// })

	if e, ok := expr.(*ast.BinaryExpr); ok {
		if v, err := evalBinary(e); err == nil {
			fmt.Println(v)
			return nil
		}
		return err
	}
	return errors.New("string is not *ast.BinaryExpr")
}

func evalBinary(expr *ast.BinaryExpr) (constant.Value, error) {
	x, err := evalExpr(expr.X)
	if err != nil {
		return constant.MakeUnknown(), errors.New("left operand faild")
	}
	y, err := evalExpr(expr.Y)
	if err != nil {
		return constant.MakeUnknown(), errors.New("right operand faild")
	}
	return constant.BinaryOp(x, expr.Op, y), nil
}

func evalExpr(expr ast.Expr) (constant.Value, error) {
	switch e := expr.(type) {
	case *ast.BinaryExpr:
		return evalBinary(e)
	case *ast.BasicLit:
		return constant.MakeFromLiteral(e.Value, e.Kind, 0), nil
	}
	return constant.MakeUnknown(), errors.New("unknown node")
}
