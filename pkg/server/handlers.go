package server

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cloudcloud/roadie/pkg/data"
	dest "github.com/cloudcloud/roadie/pkg/destinations"
	"github.com/cloudcloud/roadie/pkg/info"
	sour "github.com/cloudcloud/roadie/pkg/sources"
	"github.com/cloudcloud/roadie/pkg/types"
	"github.com/gin-gonic/gin"
)

func config(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		paths := []string{}

		for _, p := range d.GetDestinations() {
			if p.Type() == dest.DestinationLocalPath {
				paths = append(paths, p.GetLocation())
			}
		}

		for _, p := range d.GetSources() {
			if p.Type() == sour.SourceLocalPath {
				paths = append(paths, p.GetLocation())
			}
		}

		i := info.DiskDetails(d.Content)

		return gin.H{
			"disk_info": i,
		}, []string{}
	})
}

func destination(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		n := dest.FromURL(ctx.Param("name"))
		return gin.H{
			"destination": d.GetDestination(n),
			"entries":     d.GetDestinationRefs(n),
		}, []string{}
	})
}

func destinations(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return dest.PrepareList(d.GetDestinations()), []string{}
	})
}

func execute(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		body := types.ExecutePayload{}
		ctx.BindJSON(&body)

		body.SourceName = sour.FromURL(body.SourceName)
		body.DestinationName = dest.FromURL(body.DestinationName)

		return d.Copy(body), []string{}
	})
}

func historical(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		h := d.GetHistories()
		for a, x := range h {
			h[a].Destination = dest.PrepareDestination(x.Destination)
			h[a].Source = sour.PrepareSource(x.Source)
		}

		return h, []string{}
	})
}

func index(c *gin.Context) {
	s := strings.Replace(
		MustAssetString("index.html"),
		"<head>",
		"<head><script id=\"config\">{\"hostname\":\""+c.MustGet("config").(types.Configer).GetHostname()+"\"}</script>",
		1,
	)

	r := bytes.NewReader([]byte(s))
	c.DataFromReader(
		http.StatusOK,
		int64(len(s)),
		"text/html",
		r,
		map[string]string{},
	)
}

func remove(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		body := types.RemovePayload{}
		ctx.BindJSON(&body)

		body.DestinationName = dest.FromURL(body.DestinationName)
		return d.RemoveFile(body), []string{}
	})
}

func source(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		n := sour.FromURL(ctx.Param("name"))
		return gin.H{
			"source":  sour.PrepareSource(d.GetSource(n)),
			"entries": d.GetSourceRefs(n),
		}, []string{}
	})
}

func sources(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return sour.PrepareList(d.GetSources()), []string{}
	})
}

func subSource(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		n := sour.FromURL(ctx.Param("name"))
		s := sour.FromURL(ctx.Param("sub"))

		return gin.H{
			"source":  sour.PrepareSource(d.GetSource(n)),
			"entries": d.GetSubSourceRefs(n, s),
		}, []string{}
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
