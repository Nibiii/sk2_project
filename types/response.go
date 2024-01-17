package types

type Response struct {
	HTTPVersion  string          `json:"Http_Version"`
	StatusCode   int             `json:"Status-Code"`
	ReasonPhrase string          `json:"Reason-Phrase"`
	Header       ResponseHeaders `json:"Header,omitempty"`
	Body         string          `json:"Body,omitempty"`
}
