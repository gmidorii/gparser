package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func main() {
	if err := parse("10*2"); err != nil {
		fmt.Println(err)
	}
}

func parse(str string) error {
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return err
	}

	ast.Inspect(expr, func(n ast.Node) bool {
		if n != nil {
			fmt.Printf("%[1]v(%[1]T)\n", n)
		}
		return true
	})

	if e, ok := expr.(*ast.BinaryExpr); ok {
		if v, err := eval(e); err == nil {
			fmt.Println(v)
			return nil
		}
		return err
	}
	return errors.New("string is not *ast.BinaryExpr")
}

func eval(expr *ast.BinaryExpr) (int64, error) {
	xLit, ok := expr.X.(*ast.BasicLit)
	if !ok {
		return 0, errors.New("left operand faild")
	}

	yLit, ok := expr.Y.(*ast.BasicLit)
	if !ok {
		return 0, errors.New("right operand faild")
	}

	if expr.Op != token.MUL {
		return 0, errors.New("avaliable operator is multiple")
	}

	x, err := strconv.ParseInt(xLit.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	y, err := strconv.ParseInt(yLit.Value, 10, 64)
	if err != nil {
		return 0, err
	}

	return x * y, nil
}
