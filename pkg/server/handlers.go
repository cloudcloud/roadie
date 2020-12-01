package server

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudcloud/roadie/pkg/data"
	"github.com/gin-gonic/gin"
)

func destinations(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return d.GetDestinations(), []string{}
	})
}

func execute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "execute"})
}

func historical(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return d.GetHistories(), []string{}
	})
}

func index(c *gin.Context) {
	r := bytes.NewReader(MustAsset("index.html"))
	c.DataFromReader(
		http.StatusOK,
		int64(len(MustAssetString("index.html"))),
		"text/html",
		r,
		map[string]string{},
	)
}

func source(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return gin.H{
			"source":  d.GetSource(ctx.Param("name")),
			"entries": d.GetSourceRefs(ctx.Param("name")),
		}, []string{}
	})
}

func sources(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return d.GetSources(), []string{}
	})
}

func wrap(c *gin.Context, f func(*gin.Context, *data.Data) (interface{}, []string)) {
	begin := time.Now()
	d := c.MustGet("data").(*data.Data)

	out, errs := f(c, d)
	latency := time.Since(begin)

	c.JSON(http.StatusOK, gin.H{
		"items":  out,
		"errors": errs,
		"meta": map[string]interface{}{
			"latency": fmt.Sprintf("%v", latency),
			"errors":  len(errs),
		},
	})
}
