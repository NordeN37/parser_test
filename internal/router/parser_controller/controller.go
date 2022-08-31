package parser_controller

import "parser_test/internal/bl"

type ParserController struct {
	bl *bl.BL
}

func NewParserController(bl *bl.BL) *ParserController {
	return &ParserController{
		bl: bl,
	}
}
