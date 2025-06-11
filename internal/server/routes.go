package server

import (
	"go-htmx/cmd/web"
	"net/http"

	"github.com/a-h/templ/examples/integration-gin/gintemplrenderer"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"io/fs"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.LoadHTMLFiles("cmd/templates/base.templ")

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	// Disable trusted proxy warning.
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/", func(c *gin.Context) {
		r := gintemplrenderer.New(c.Request.Context(), http.StatusOK, Base())
		c.Render(http.StatusOK, r)
	})

	staticFiles, _ := fs.Sub(web.Files, "assets")
	r.StaticFS("/assets", http.FS(staticFiles))

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}
