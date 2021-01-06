package server

import (
	"time"

	"github.com/cloudcloud/roadie/pkg/data"
	"github.com/cloudcloud/roadie/pkg/types"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	skip = map[string]struct{}{
		"/health": struct{}{},
	}
)

// Server defines an expected behaviour set for the interaction of
// handling HTTP requests.
type Server interface {
	Start() error
}

// Serve is a Server implementation to manage the handling of HTTP
// based requests.
type Serve struct {
	c types.Configer
	g *gin.Engine
}

// New will generate a new Server instance that will setup a HTTP
// server ready to begin handling requests.
func New(c types.Configer) Server {
	g := gin.New()
	g.Use(
		cors.New(cors.Config{
			AllowOrigins: d.Content.Domains,
			AllowMethods: []string{"GET", "POST", "PUT", "OPTIONS", "HEAD", "DELETE"},
			AllowHeaders: []string{"Origin", "X-Client", "Content-Type"},
		}),
		logger(c),
		push(c),
	)

	g.StaticFS("/js",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "js/",
		},
	)

	g.StaticFS("/css",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "css/",
		},
	)

	g.StaticFS("/fonts",
		&assetfs.AssetFS{
			Asset:     Asset,
			AssetDir:  AssetDir,
			AssetInfo: AssetInfo,
			Prefix:    "fonts/",
		},
	)

	g.GET("/", index)
	g.GET("/sources", index)
	g.GET("/destinations", index)
	g.GET("/config", index)

	api := g.Group("/api/v1")
	api.GET("/sources", sources)
	api.GET("/sources/:name", source)
	api.GET("/sources/:name/:sub", subSource)
	api.GET("/destinations", destinations)
	api.GET("/destinations/:name", destination)
	api.POST("/execute", execute)
	api.GET("/historical", historical)
	api.DELETE("/remove", remove)

	return &Serve{
		c: c,
		g: g,
	}
}

// Start will begin the HTTP request handling process.
func (s *Serve) Start() error {
	return s.g.Run(s.c.GetListener())
}

func push(c types.Configer) gin.HandlerFunc {
	d := data.New(c)

	return func(ctx *gin.Context) {
		ctx.Set("data", d)
		ctx.Set("config", c)
		ctx.Next()
	}
}

func logger(c types.Configer) gin.HandlerFunc {
	log := c.GetLogger()
	return func(ctx *gin.Context) {
		start := time.Now()
		l := log.With(
			"client_ip", ctx.ClientIP(),
			"method", ctx.Request.Method,
			"start", start,
		)

		p := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		if raw != "" {
			p = p + "?" + raw
		}
		l = l.With("path", p)

		ctx.Set("log", l)
		ctx.Next()

		if _, ok := skip[p]; !ok {
			e := time.Now()
			l.With(
				"body_size", ctx.Writer.Size(),
				"end", e,
				"latency", e.Sub(start),
				"status", ctx.Writer.Status(),
			).Info("access_log")
		}
	}
}
