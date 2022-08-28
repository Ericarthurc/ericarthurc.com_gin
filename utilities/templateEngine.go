package utilities

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type RenderOptions struct {
	TemplateDir      string
	TemplateExtenson string
	ContentType      string
}

type Pongo2Render struct {
	Options  *RenderOptions
	Template *pongo2.Template
	Context  pongo2.Context
}

func NewPongo(options RenderOptions) *Pongo2Render {
	return &Pongo2Render{
		Options: &options,
	}
}

func (p Pongo2Render) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	filename := path.Join(p.Options.TemplateDir, name)
	if gin.Mode() == "debug" {
		template = pongo2.Must(pongo2.FromFile(filename + "." + p.Options.TemplateExtenson))
	} else {
		template = pongo2.Must(pongo2.FromCache(filename + "." + p.Options.TemplateExtenson))
	}

	return Pongo2Render{
		Template: template,
		Context:  data.(pongo2.Context),
		Options:  p.Options,
	}
}

func (p Pongo2Render) Render(w http.ResponseWriter) error {
	p.WriteContentType(w)
	err := p.Template.ExecuteWriter(p.Context, w)
	return err
}

func (p Pongo2Render) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{p.Options.ContentType}
	}
}
