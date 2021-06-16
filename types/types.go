package types

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Items    []interface{}
	Requests []Request
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
