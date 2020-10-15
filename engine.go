package pastis

import (
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

const (
	pathParamIndicator  = ":"
	componentsSeparator = "/"
)

type Handler func(*Context)

type Engine struct {
	BasePath       string
	tree           *node
	DefaultHandler Handler
}

func (e *Engine) GET(path string, handler Handler) {
	e.handle(http.MethodGet, path, handler)
}

func (e *Engine) POST(path string, handler Handler) {
	e.handle(http.MethodPost, path, handler)
}

func (e *Engine) PATCH(path string, handler Handler) {
	e.handle(http.MethodPatch, path, handler)
}

func (e *Engine) PUT(path string, handler Handler) {
	e.handle(http.MethodPut, path, handler)
}

func (e *Engine) DELETE(path string, handler Handler) {
	e.handle(http.MethodDelete, path, handler)
}

func (e *Engine) handle(method, path string, handler Handler) {
	if path[0] != '/' {
		panic("Path has to start with a /.")
	}
	e.tree.addNode(method, path, handler)
}

func (e *Engine) Run(event events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	params := make(params)
	node, _ := e.tree.traverse(strings.Split(event.Path, "/")[1:], params)
	context := NewContext(params, event)
	if handler := node.methods[event.HTTPMethod]; handler != nil {
		handler(context)
	} else if e.DefaultHandler != nil {
		e.DefaultHandler(context)
	}
	return context.Response, nil
}

func New() *Engine {
	return &Engine{
		tree: &node{
			component:    "/",
			isNamedParam: false,
			methods:      make(map[string]Handler),
		},
		DefaultHandler: func(c *Context) {
			c.Response = events.ALBTargetGroupResponse{
				StatusCode:        http.StatusNotFound,
				StatusDescription: "Not found",
				Headers:           nil,
				MultiValueHeaders: nil,
				Body:              "",
				IsBase64Encoded:   false,
			}
		},
	}
}
