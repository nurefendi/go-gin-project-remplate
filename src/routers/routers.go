package routers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleRouter() {
	log.Info("Start routing")
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "up",
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "sys_admin.html", nil)
	})

	router.Use(static.Serve("/components", static.LocalFile("./src/templates/assets/component", true)))
	router.HTMLRender = loadTemplates("./src/templates")
	router.Static("/css", "./src/templates/assets/css")
	router.Static("/js", "./src/templates/assets/js")
	router.Static("/images", "./src/templates/assets/img")
	router.Static("/font", "./src/templates/assets/font")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Info("Routed to port " + port)
	log.Error(router.Run(":" + port))
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/content/**/*")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
