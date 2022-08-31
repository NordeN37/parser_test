package bl

import "parser_test/internal/store/mongo"

type BL struct {
	Parser IParseHtml
}

func NewBL(mr *mongo.MongoRepo) *BL {
	return &BL{
		Parser: NewParseHtml(mr),
	}
}
