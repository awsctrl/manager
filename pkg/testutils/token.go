package testutils

type Token struct {
	Token string
}

func NewToken() Token {
	return Token{}
}

func (t Token) Generate() string {
	return t.Token
}
