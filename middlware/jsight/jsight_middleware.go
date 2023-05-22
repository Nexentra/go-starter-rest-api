package jsight

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"bytes"
	"os"
	"regexp"
)

var jSight JSight

type bodyLogWriter struct {
    gin.ResponseWriter
    body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
    w.body.Write(b)
    return w.ResponseWriter.Write(b)
}

func Validator() gin.HandlerFunc {
	return func(c *gin.Context) {
		if jSight == nil {
			jSight = NewJSight("./middlware/jsight/jsightplugin-alpine.so") // For Alpine
			// jSight = NewJSight("./middlware/jsight/jsightplugin.so") // For other linuxes
			fmt.Println("JSight validator enabled")
			fmt.Print(jSight.Stat())
		}
		
		// before request

		jsightSpecPath := "./jsight/api-spec.jst"
		reqBody, _ := io.ReadAll(c.Request.Body)

		jSight.ClearCache() // Comment this line in production to gain performance!!!

		// validate request
		err := jSight.ValidateHTTPRequest(
			jsightSpecPath,
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Header,
			reqBody,
		)

		if err != nil {
			c.Header("Content-Type", "application/json")
			c.String(400, err.ToJSON())
			return
		}

		// check, if the jsight spec was requested
		matched, _ := regexp.MatchString(`.*jsight/?`, c.Request.RequestURI)
		if matched {
			jsightCode, _ := os.ReadFile(jsightSpecPath)
			c.Writer.WriteHeader(200)
			c.Writer.Write(jsightCode)
			return
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		// before response

		// validate response
		err = jSight.ValidateHTTPResponse(
			jsightSpecPath,
			c.Request.Method,
			c.Request.RequestURI,
			c.Writer.Status(),
			c.Writer.Header(),
			blw.body.Bytes(),
		)
	
		if err != nil {
			c.Writer.WriteHeader(500)
			c.Writer.Write([]byte("\n\nRESPONSE ERROR:\n\n"))
			c.Writer.Write([]byte(err.ToJSON()))
			return
		}
	}
}