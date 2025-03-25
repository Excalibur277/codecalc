package listener

import (
	"fmt"

	"codecalc.com/m/v2/ast"
	"codecalc.com/m/v2/parsing"
	"github.com/antlr4-go/antlr/v4"
)

func ParseStream(is antlr.CharStream) ast.Module {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	lexer := parsing.NewCodeCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parsing.NewCodeCalcParser(stream)

	listener := NewListener()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Module())

	listener.Pop(1)
	module := Dequeue[ast.Module](listener)
	return module
}
