package figma

import (
	"context"
	"encoding/json"
)

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID      string `json:"id"`
	FileKey string `json:"file_key"`
}

func (c *Client) GetComments(ctx context.Context, fileKey string) (*Comments, error) {
	req, _ := c.newRequest(ctx, "GET", "/v1/files/"+fileKey+"/comments", nil)
	res, _ := c.send(ctx, req)

	var comments Comments
	err := json.Unmarshal(res, &comments)

	if err != nil {
		return nil, err
	}

	return &comments, nil
}
