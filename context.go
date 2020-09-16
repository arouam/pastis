package pastis

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
)

type Context struct {
	Params params
	Queries queries
	Request events.ALBTargetGroupRequest
	Response events.ALBTargetGroupResponse
}


func NewContext(params params, request events.ALBTargetGroupRequest) *Context{
	return &Context{
		Params:  params,
		Queries: request.QueryStringParameters,
		Request: request,
		Response: events.ALBTargetGroupResponse{},
	}
}

func (c *Context)JSON(statusCode int,i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		c.Response = events.ALBTargetGroupResponse{
			StatusCode:        http.StatusInternalServerError,
			StatusDescription: http.StatusText(http.StatusInternalServerError),
			Headers: map[string]string{"content-type": "application/json"},
			MultiValueHeaders: nil,
			Body:              "",
			IsBase64Encoded:   false,
		}
		return
	}
	c.Response = events.ALBTargetGroupResponse{
		StatusCode:        statusCode,
		StatusDescription: http.StatusText(statusCode),
		Headers: map[string]string{"content-type": "application/json"},
		MultiValueHeaders: nil,
		Body:           string(b)   ,
		IsBase64Encoded:   false,
	}
}

func (c *Context)Param(key string) string {
	return c.Params.get(key)
}

func (c *Context)Query(key string) string {
	return c.Queries.get(key)
}