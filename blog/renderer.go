package blog

import (
	"embed"
	"html/template"
	"io"
)

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	var views []PostViewModel
	for _, post := range posts {
		views = append(views, *NewPostViewModel(post))
	}
	return r.templ.ExecuteTemplate(w, "index.gohtml", views)
}
