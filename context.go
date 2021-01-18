package beegin

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type (
	H map[string]interface{}
)

type Context struct {
	Writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) ExBody() []byte {
	body, err := ioutil.ReadAll(c.Req.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Status(statusCode int) {
	c.StatusCode = statusCode
	c.Writer.WriteHeader(statusCode)
}

func (c *Context) String(httpStatusCode int, reObj interface{}) error {
	var t = reflect.TypeOf(reObj)
	var s string

	switch t.Kind() {
	case reflect.String:
		s = reObj.(string)
	case reflect.TypeOf((*error)(nil)).Kind():
		s = reObj.(error).Error()
	default:
		log.Fatal(fmt.Errorf("error reObj arg: %v", reObj))
	}

	c.SetHeader("Context-Type", "text/plain")
	c.Status(httpStatusCode)
	_, err := c.Writer.Write([]byte(s))
	return err
}
