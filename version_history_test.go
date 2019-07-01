package figma

import (
	"context"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestGetFileVersions(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	mockprojects := Versions{
		[]Version{
			Version{
				"11111",
				"2019-04-11T04:39:11Z",
				nil,
				nil,
				User{
					"test",
					"test",
					"test",
				},
				"testtest",
			},
		},
		Pagination{
			"testurl",
		},
	}

	gock.New("https://api.figma.com").
		Get("/v1/files/fake/versions").
		Reply(200).
		JSON(mockprojects)

	c := NewClient("fakefake")
	ctx := context.Background()
	result, err := c.GetFileVersions(ctx, "fake")

	if err != nil {
		t.Errorf("shouldn't fail")
	}

	if len(result.Versions) != 1 {
		t.Errorf("should be 1")
	}
}
