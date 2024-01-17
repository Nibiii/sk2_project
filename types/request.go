package types

import (
	"fmt"
	"net"
	"strings"
)

type Request struct {
	Method      string         `json:"Method"`
	URI         string         `json:"Uri"`
	HTTPVersion string         `json:"Http_Version"`
	Header      RequestHeaders `json:"Header,omitempty"`
	Body        string         `json:"Body,omitempty"`
}

func (r *Request) Handle(conn net.Conn, files *Files) *Response {
	res := new(Response)
	// remove leading forward slash from uri
	r.URI = strings.TrimPrefix(r.URI, "/")
	// handle empty uri
	if r.URI == "" {
		res.StatusCode = 406
		res.ReasonPhrase = "uri_is_empty"
		return res
	}
	// do not allow for creating files in subdirectories
	if strings.Contains(r.URI, "/") {
		res.StatusCode = 406
		res.ReasonPhrase = "subdirectories_not_allowed"
		return res
	}
	switch r.Method {
	case "GET":
		res = r.Get(conn, files)
	case "HEAD":
		res = r.Get(conn, files)
	case "POST":
		res = r.Post(conn, files)
	case "PUT":
		res = r.Put(conn, files)
	case "DELETE":
		res = r.Delete(conn, files)
	default:
		res.StatusCode = 405
		res.ReasonPhrase = "unsupported_method"
	}
	return res
}

func (r *Request) Delete(conn net.Conn, files *Files) *Response {
	res := new(Response)
	file, err := files.FindByName(&r.URI)
	if err != nil {
		res.StatusCode = 404
		res.ReasonPhrase = "file_not_found"
		return res
	}
	// delete file from file system
	err = file.Delete()
	if err != nil {
		res.StatusCode = 500
		res.ReasonPhrase = "internal_server_error"
		res.Body = fmt.Sprintf("Error while deleting file: %s", err)
		return res
	}
	// delete file's object
	files.DeleteByName(&r.URI)
	if err != nil {
		res.StatusCode = 500
		res.ReasonPhrase = "internal_server_error"
		res.Body = fmt.Sprintf("Error while deleting file: %s", err)
		return res
	}
	res.StatusCode = 200
	res.ReasonPhrase = "OK"
	return res
}

func (r *Request) Get(conn net.Conn, files *Files) *Response {
	res := new(Response)
	file, err := files.FindByName(&r.URI)
	if err != nil {
		res.StatusCode = 404
		res.ReasonPhrase = "file_not_found"
		return res
	}
	res.StatusCode = 200
	res.ReasonPhrase = "OK"
	res.Body = string(file.Content)
	return res
}

func (r *Request) Head(conn net.Conn, files *Files) *Response {
	res := new(Response)
	_, err := files.FindByName(&r.URI)
	if err != nil {
		res.StatusCode = 404
		res.ReasonPhrase = "file_not_found"
		return res
	}
	res.StatusCode = 200
	res.ReasonPhrase = "OK"
	return res
}

func (r *Request) Post(conn net.Conn, files *Files) *Response {
	res := new(Response)
	_, err := files.FindByName(&r.URI)
	if err == nil {
		res.StatusCode = 400
		res.ReasonPhrase = "file_exists"
		return res
	}
	err = files.New(&r.URI, []byte(r.Body))
	if err != nil {
		res.StatusCode = 500
		res.ReasonPhrase = "internal_server_error"
		res.Body = fmt.Sprintf("Unable to create new file: %s", err)
		return res
	}
	res.StatusCode = 201
	res.ReasonPhrase = "Created"
	return res
}

func (r *Request) Put(conn net.Conn, files *Files) *Response {
	res := new(Response)
	file, err := files.FindByName(&r.URI)
	if err != nil {
		res.StatusCode = 404
		res.ReasonPhrase = "file_not_found"
		res.Body = fmt.Sprintf("File not found: %s", err)
		return res
	}
	err = file.Modify([]byte(r.Body))
	if err != nil {
		res.StatusCode = 500
		res.ReasonPhrase = "internal_server_error"
		res.Body = fmt.Sprintf("Error while modifying file: %s", err)
		return res
	}
	res.StatusCode = 200
	res.ReasonPhrase = "OK"
	return res
}

func (r *Request) Respond(conn net.Conn, res *Response) {
	// create string builder obj
	var resString strings.Builder

	// prepare response's start line
	res.HTTPVersion = "HTTP/1.1"
	resString.WriteString(fmt.Sprintf("%s %d %s\r\n", res.HTTPVersion, res.StatusCode, res.ReasonPhrase))
	if length := len(res.Body); length == 0 {
		res.Body = res.ReasonPhrase
	}

	// prepare header parameters
	resString.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
	resString.WriteString("Content-Encoding: none\r\n")
	resString.WriteString(fmt.Sprintf("Content-Length: %d\r", len(res.Body)))
	resString.WriteString("\n\r\n")

	// prepare body but only in method is not head
	if r.Method != "HEAD" {
		resString.WriteString(res.Body)
	}

	// print response string for debug
	fmt.Println("\nResponse: {\n", strings.ReplaceAll(resString.String(), "\r", ""), "\n}")

	// send response
	_, err := conn.Write([]byte(resString.String()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
