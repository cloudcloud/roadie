package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cloudcloud/roadie/pkg/data"
	dest "github.com/cloudcloud/roadie/pkg/destinations"
	sour "github.com/cloudcloud/roadie/pkg/sources"
	"github.com/cloudcloud/roadie/pkg/types"
	"github.com/gin-gonic/gin"
)

func config(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (interface{}, []string) {
		return gin.H{
			"destinations": d.GetDestinationsWithDetails(),
			"sources":      d.GetSources(),
		}, []string{}
	})
}

func configAdd(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (any, []string) {
		t := ctx.Param("type")

		_ = t

		return gin.H{}, []string{}
	})
}

func configEdit(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (any, []string) {
		t, n := ctx.Param("type"), ctx.Param("name")

		switch t {
		case "source":
			body := types.Source{}
			ctx.BindJSON(&body)

			body.Name = sour.FromURL(n)
			if err := d.UpdateSource(body); err != nil {
				return gin.H{}, []string{err.Error()}
			}
			return gin.H{"source": body}, []string{}

		case "destination":
			// TODO: Use types.ActionDestination
			// types.Destination is the wrapper, not the core of it

			body := types.Destination{}
			ctx.BindJSON(&body)
			body.Name = dest.FromURL(n)

			// load the existing destination, override values and store
			destination := d.GetDestination(n)
			if destination.Name == "" {
				return gin.H{}, []string{fmt.Sprintf("Unable to find the '%s' destination", n)}
			}

			if err := d.UpdateDestination(body); err != nil {
				return gin.H{}, []string{err.Error()}
			}
			return gin.H{"destination": body}, []string{}

		}

		return gin.H{}, []string{"unknown type provided"}
	})
}

func configRemove(c *gin.Context) {
	wrap(c, func(ctx *gin.Context, d *data.Data) (any, []string) {
		t, n := ctx.Param("type"), ctx.Param("name")

		switch t {
		case "source":
			if err := d.RemoveSource(sour.FromURL(n)); err != nil {
				return gin.H{}, []string{err.Error()}
			}

		case "destination":
			if err := d.RemoveDestination(dest.FromURL(n)); err != nil {
				return gin.H{}, []string{err.Error()}
			}

		}

		return gin.H{"remove": "success"}, []string{}
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
