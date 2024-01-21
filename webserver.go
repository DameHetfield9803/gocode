package main

import (
	"fmt"
	"os"
)

type Page struct {
	// Title is a string
	Title string
	// Body is a byte slice
	Body []byte
}

func (p *Page) save() error {
	// file name is the title of the page
	filename := p.Title + ".txt"
	// write the body of the page to a file
	// 0600 is the Unix permission mode
	// The octal integer literal 0600, passed as the third parameter to WriteFile,
	// indicates that the file should be created with read-write permissions for the current user only.
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	// ReadFile returns a byte slice and an error

	body, error := os.ReadFile(filename)
	// writes the body of the page to a file

	if error != nil {
		return nil, error
	}
	// The & operator generates a pointer to its operand.
	return &Page{Title: title, Body: body}, nil
}

func mainWeb() {
	// create a new page called TestPage
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// Saves the page to a file
	p1.save()

	// load the page from the file
	p2, _ := loadPage("TestPage")

	fmt.Println(string(p2.Body))
}
