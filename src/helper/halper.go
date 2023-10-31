package helper

import (
	"strconv"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)


func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func ParseStringToInt(value string) int {
	if value == "" {
		return 0
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		log.Error(err)
		return 0
	}
	return i
}

func LoadTemplates(templatesDir string) multitemplate.Renderer {
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