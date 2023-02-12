package expr

import (
	"fmt"
	"math"
)

// Expr is one expression arithmetic
type Expr interface {
	// Eval returns the result of expression Expr on environment env
	Eval(env Env) float64
}

// Var identify one variable, for example: x
type Var string

// Eval returns the result of expression Expr on environment env
func (v Var) Eval(env Env) float64 {
	return env[v]
}

// literal is a numeric constant, for example, 3.141
type literal float64

// Eval returns the result of expression Expr on environment env
func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

// unary is an expression with unary operator like -x
type unary struct {
	op rune // value between '+' and '-'
	x  Expr
}

// Eval returns the result of expression Expr on environment env
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

// binary represents expression with binary operator, like x+y
type binary struct {
	op   rune // value between '+', '-', '*', '/'
	x, y Expr
}

// Eval returns the result of expression Expr on environment env
func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

type call struct {
	fn   string // values like: "pow", "sin", "sqrt"
	args []Expr
}

// Eval returns the result of expression Expr on environment env
func (c *call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

type Env map[Var]float64
