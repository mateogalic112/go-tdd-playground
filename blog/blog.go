// Package blog provides posts
package blog

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags                                     []string
	HTMLBody                                 template.HTML
}

func NewPostViewModel(p Post) *PostViewModel {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	mdParser := parser.NewWithExtensions(extensions)
	htmlBody := template.HTML(markdown.ToHTML([]byte(p.Body), mdParser, nil))

	return &PostViewModel{
		Title:          p.Title,
		SanitisedTitle: sanitiseTitle(p),
		Description:    p.Description,
		Body:           p.Body,
		Tags:           p.Tags,
		HTMLBody:       htmlBody,
	}
}

func sanitiseTitle(post Post) string {
	return strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

	scanner.Scan() // ignore dashed line
	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: tags, Body: body}, nil
}

func readBody(scanner *bufio.Scanner) string {
	var b strings.Builder
	for scanner.Scan() {
		fmt.Fprintln(&b, scanner.Text())
	}
	return strings.TrimSuffix(b.String(), "\n")
}
