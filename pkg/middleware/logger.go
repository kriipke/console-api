package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

				 c1 := color.New(color.FgMagenta, color.Bold)
				 c2 := color.New(color.FgWhite)
				 c3 := color.New(color.FgYellow)
				 c4 := color.New(color.FgCyan, color.Bold)
				 c5 := color.New(color.FgWhite, color.Bold)

        c.Next()

				latency := time.Since(t)

        c4.Printf("\n  %s", c.Request.Method)
				c5.Printf("%s  ", c.Request.RequestURI)
				c1.Printf("%s  ", c.Request.Proto)
				c2.Printf("%s\n",  latency)

				for name, values := range c.Request.Header {
						for _, value := range values {
							c3.Printf("   %-20s:", name)
							c2.Printf("%-40s\n", value)
						}
				}
    }
}

func ResponseLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

        c.Next()

				//c1 := color.New(color.FgMagenta, color.Bold)
				c2 := color.New(color.FgWhite)
				c3 := color.New(color.FgYellow)
			  c4 := color.New(color.FgCyan, color.Bold)
			  c5 := color.New(color.FgWhite, color.Bold)

        c4.Printf("\n   %s", c.Writer.Status())
				c4.Printf("%s  ", c.Request.Method)
				c5.Printf("%s\n",  c.Request.RequestURI)

				for name, values := range c.Writer.Header() {
						for _, value := range values {
							c3.Printf("    %-40s:", name)
							c2.Printf("%-80s\n", value)
						}
				}

    }
}
