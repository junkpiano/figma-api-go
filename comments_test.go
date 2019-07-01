package figma

import (
	"context"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestGetComments(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	mockcomments := Comments{
		[]Comment{
			Comment{
				"11111",
				"golang",
			},
		},
	}
	gock.New("https://api.figma.com").
		Get("/v1/files/fake/comments").
		Reply(200).
		JSON(mockcomments)

	c := NewClient("testtesttest")
	ctx := context.Background()
	comments, _ := c.GetComments(ctx, "fake")

	if len(comments.Comments) != 1 {
		t.Errorf("the number of comments must be 1.")
	}
}
