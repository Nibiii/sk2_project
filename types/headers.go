package types

type GeneralHeaders struct {
	CacheControl     string `json:"Cache-Control,omitempty"`
	Connection       string `json:"Connection,omitempty"`
	Date             string `json:"Date,omitempty"`
	Pragma           string `json:"Pragma,omitempty"`
	Trailer          string `json:"Trailer,omitempty"`
	TransferEncoding string `json:"Transfer-Encoding,omitempty"`
	Upgrade          string `json:"Upgrade,omitempty"`
	Via              string `json:"Via,omitempty"`
	Warning          string `json:"Warning,omitempty"`
}

type RequestHeaders struct {
	GeneralHeaders
	Accept             string `json:"Accept,omitempty"`
	AcceptCharset      string `json:"Accept-Charset,omitempty"`
	AcceptEncoding     string `json:"Accept-Encoding,omitempty"`
	AcceptLanguage     string `json:"Accept-Language,omitempty"`
	Authorization      string `json:"Authorization,omitempty"`
	Expect             string `json:"Expect,omitempty"`
	From               string `json:"From,omitempty"`
	Host               string `json:"Host,omitempty"`
	IfMatch            string `json:"If-Match,omitempty"`
	IfModifiedSince    string `json:"If-Modified-Since,omitempty"`
	IfNoneMatch        string `json:"If-None-Match,omitempty"`
	IfRange            string `json:"If-Range,omitempty"`
	IfUnmodifiedSince  string `json:"If-Unmodified-Since,omitempty"`
	MaxForwards        string `json:"Max-Forwards,omitempty"`
	ProxyAuthorization string `json:"Proxy-Authorization,omitempty"`
	Range              string `json:"Range,omitempty"`
	Referer            string `json:"Referer,omitempty"`
	TE                 string `json:"TE,omitempty"`
	UserAgent          string `json:"User-Agent,omitempty"`
	EntityHeaders
}

type EntityHeaders struct {
	Allow           string `json:"Allow,omitempty"`
	ContentEncoding string `json:"Content-Encoding,omitempty"`
	ContentLanguage string `json:"Content-Language,omitempty"`
	ContentLenght   string `json:"Content-Length,omitempty"`
	ContentLocation string `json:"Content-Location,omitempty"`
	ContentMD5      string `json:"Content-MD5,omitempty"`
	ContentRange    string `json:"Content-Range,omitempty"`
	ContentType     string `json:"Content-Type,omitempty"`
	Expires         string `json:"Expires,omitempty"`
	LastModified    string `json:"Last-Modified,omitempty"`
}

type ResponseHeaders struct {
	GeneralHeaders
	AcceptRanges      string `json:"Accept-Ranges,omitempty"`
	Age               string `json:"Age,omitempty"`
	ETag              string `json:"ETag,omitempty"`
	Location          string `json:"Location,omitempty"`
	ProxyAuthenticate string `json:"Proxy-Authenticate,omitempty"`
	RetryAfter        string `json:"Retry-After,omitempty"`
	Server            string `json:"Server,omitempty"`
	Vary              string `json:"Vary,omitempty"`
	WWWAuthenticate   string `json:"WWW-Authenticate,omitempty"`
	EntityHeaders
}
