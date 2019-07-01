package figma

import (
	"context"
	"encoding/json"
	"fmt"
)

type Versions struct {
	Versions   []Version  `json:"versions"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	PrevPage string `json:"prev_page"`
}

type Version struct {
	ID           string      `json:"id"`
	CreatedAt    string      `json:"created_at"`
	Label        interface{} `json:"label"`
	Description  interface{} `json:"description"`
	User         User        `json:"user"`
	ThumbnailURL string      `json:"thumbnail_url"`
}

func (c *Client) GetFileVersions(ctx context.Context, fileKey string) (*Versions, error) {
	req, _ := c.newRequest(ctx, "GET", fmt.Sprintf("/v1/files/%s/versions", fileKey), nil)
	res, _ := c.send(ctx, req)

	var result Versions
	err := json.Unmarshal(res, &result)
	if err != nil {
		return nil, nil
	}

	return &result, nil
}
