package blog

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is post a",
			Description: "This is desc",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		err := Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

	})
}
