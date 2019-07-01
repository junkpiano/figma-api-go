package figma

import (
	"context"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

func TestGetTeamProjects(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	mockprojects := Projects{
		[]Project{
			Project{
				"11111",
				"golang",
			},
		},
	}
	gock.New("https://api.figma.com").
		Get("/v1/teams/fake/projects").
		Reply(200).
		JSON(mockprojects)

	c := NewClient("testtesttest")
	ctx := context.Background()
	projects, _ := c.GetTeamProjects(ctx, "fake")

	if len(projects.Projects) != 1 {
		t.Errorf("projects must be 1.")
	}
}

func TestNoTeamProjects(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://api.figma.com").
		Get("/v1/teams/not_existing_team_id/projects").
		Reply(404)

	c := NewClient("testtesttest")
	ctx := context.Background()
	_, err := c.GetTeamProjects(ctx, "not_existing_team_id")

	if err == nil {
		t.Errorf("error must not be nil.")
	}
}

func TestGetProjectFiles(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	mockfiles := ProjectFiles{
		[]ProjectFile{
			ProjectFile{
				"ababa", "test", "teststes", "testset",
			},
		},
	}
	gock.New("https://api.figma.com").
		Get("/v1/projects/fake/files").
		Reply(200).
		JSON(mockfiles)

	c := NewClient("testtesttest")
	ctx := context.Background()
	files, _ := c.GetProjectFiles(ctx, "fake")
	if len(files.Files) != 1 {
		t.Errorf("files must be 1.")
	}
}
