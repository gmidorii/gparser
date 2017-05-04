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
		v, err := parse("10" + v + "2")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v)
	}
}

func parse(str string) (string, error) {
	fmt.Println(str)
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return "", err
	}

	v, err := evalExpr(expr)
	if err != nil {
		return "", err
	}
	return v.String(), nil
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

func evalUnary(expr *ast.UnaryExpr) (constant.Value, error) {
	x, err := evalExpr(expr.X)
	if err != nil {
		return constant.MakeUnknown(), err
	}

	return constant.UnaryOp(expr.Op, x, 0), nil
}

func evalExpr(expr ast.Expr) (constant.Value, error) {
	switch e := expr.(type) {
	case *ast.ParenExpr:
		return evalExpr(e.X)
	case *ast.BinaryExpr:
		return evalBinary(e)
	case *ast.UnaryExpr:
		return evalUnary(e)
	case *ast.BasicLit:
		return constant.MakeFromLiteral(e.Value, e.Kind, 0), nil
	}
	return constant.MakeUnknown(), errors.New("unknown node")
}
