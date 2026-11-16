package blog

import (
	"fmt"
	"io"
)

func Render(w io.Writer, p Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
