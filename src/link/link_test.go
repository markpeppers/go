package link_test

import (
	"fmt"
	"io/ioutil"
	"link"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLink(t *testing.T) {

	type test struct {
		testFile string
		want     []link.Link
	}
	tests := []test{
		{
			testFile: "testfiles/ex1.html",
			want: []link.Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			testFile: "testfiles/ex2.html",
			want: []link.Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on Github!",
				},
			},
		},
		{
			testFile: "testfiles/ex3.html",
			want: []link.Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			testFile: "testfiles/ex4.html",
			want: []link.Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				},
			},
		},
	}

	for _, tc := range tests {
		bytes, err := ioutil.ReadFile(tc.testFile)
		assert.NoError(t, err, fmt.Sprintf("Error reading file %s", tc.testFile))
		links := link.ParseLinks(string(bytes))
		assert.Equal(t, tc.want, links, fmt.Sprintf("Failed on file %s", tc.testFile))
	}
}
