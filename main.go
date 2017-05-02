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
		if v, err := eval(e); err == nil {
			fmt.Println(v)
			return nil
		}
		return err
	}
	return errors.New("string is not *ast.BinaryExpr")
}

func eval(expr *ast.BinaryExpr) (constant.Value, error) {
	xLit, ok := expr.X.(*ast.BasicLit)
	if !ok {
		return constant.MakeUnknown(), errors.New("left operand faild")
	}

	yLit, ok := expr.Y.(*ast.BasicLit)
	if !ok {
		return constant.MakeUnknown(), errors.New("right operand faild")
	}

	x := evalBasicLit(xLit)
	y := evalBasicLit(yLit)
	return constant.BinaryOp(x, expr.Op, y), nil
}

func evalBasicLit(expr *ast.BasicLit) constant.Value {
	return constant.MakeFromLiteral(expr.Value, expr.Kind, 0)
}
