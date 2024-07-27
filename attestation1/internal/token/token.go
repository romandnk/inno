package token

import "inno/attestation1/internal/entity"

type WhiteList map[string]struct{}

func NewWhiteList(cfg Config) WhiteList {
	whiteList := make(WhiteList, len(cfg.WhiteTokens))
	for _, token := range cfg.WhiteTokens {
		whiteList[token] = struct{}{}
	}
	return whiteList
}

func (t WhiteList) checkToken(tok string) bool {
	_, ok := (t)[tok]
	return ok
}

func (t WhiteList) ValidateToken(msg entity.Message, fn func(msg entity.Message)) {
	if !t.checkToken(msg.Token) {
		return
	}
	fn(msg)
}
