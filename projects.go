package figma

import (
	"context"
	"encoding/json"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectFiles struct {
	Files []ProjectFile `json:"files"`
}

type ProjectFile struct {
	Key          string `json:"key"`
	Name         string `json:"name"`
	ThumbnailURL string `json:"thumbnail_url"`
	LastModified string `json:"last_modified"`
}

func (c *Client) GetTeamProjects(ctx context.Context, teamID string) (*Projects, error) {
	req, _ := c.newRequest(ctx, "GET", "/v1/teams/"+teamID+"/projects", nil)
	res, _ := c.send(ctx, req)

	var projects Projects
	err := json.Unmarshal(res, &projects)

	return &projects, err
}

func (c *Client) GetProjectFiles(ctx context.Context, projectID string) (*ProjectFiles, error) {
	req, _ := c.newRequest(ctx, "GET", "/v1/projects/"+projectID+"/files", nil)
	res, _ := c.send(ctx, req)

	var files ProjectFiles
	err := json.Unmarshal(res, &files)
	return &files, err
}
