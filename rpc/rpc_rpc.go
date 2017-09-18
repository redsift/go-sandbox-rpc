package rpc

import "net/textproto"

/*
   From https://tools.ietf.org/html/rfc2616
   Request (section 5) and Response (section 6) messages use the generic
   message format of RFC 822 [9] for transferring entities (the payload
   of the message). Both types of message consist of a start-line, zero
   or more header fields (also known as "headers"), an empty line (i.e.,
   a line with nothing preceding the CRLF) indicating the end of the
   header fields, and possibly a message-body.
        generic-message = start-line
                          *(message-header CRLF)
                          CRLF
                          [ message-body ]
        start-line      = Request-Line | Status-Line
        Request-Line   = Method SP Request-URI SP HTTP-Version CRLF
        Status-Line = HTTP-Version SP Status-Code SP Reason-Phrase CRLF
*/

type Request struct {
	RemoteAddr string `json:"remote_addr"` // e.g. "IP:Port"
	Method     string `json:"method"`
	// RequestURI is the unmodified Request-URI of the
	// Request-Line (RFC 2616, Section 5.1) as sent by the client
	// to a server.
	RequestURI string               `json:"request_uri"` // encoded path
	Header     textproto.MIMEHeader `json:"header"`
	Body       []byte               `json:"body"` // []byte encodes as a base64-encoded string
}

type Response struct {
	// our codes:
	// 500 - We messed up;
	// TODO 502 - Sift messed up;
	// 504 - Sift too slow;
	// TODO 413 - We/Sift can't eat more;
	// 401 - Who are you?;
	// TODO 408 - You're too slow
	StatusCode int                  `json:"status_code"`
	Header     textproto.MIMEHeader `json:"header"`
	Body       []byte               `json:"body"`
}
