package testutils

type Token struct {
	Token string
}

func NewToken() Token {
	return Token{"test-token"}
}

func (t Token) Generate() string {
	return t.Token
}
