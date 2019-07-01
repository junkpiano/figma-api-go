package figma

import (
	"context"
	"encoding/json"
	"fmt"
)

type User struct {
	Handle string `json:"handle"`
	ImgURL string `json:"img_url"`
	ID     string `json:"id"`
}

func (c *Client) GetMe(ctx context.Context) (*User, error) {
	req, _ := c.newRequest(ctx, "GET", fmt.Sprintf("/v1/me"), nil)
	res, _ := c.send(ctx, req)
	fmt.Println(res)

	var result User
	err := json.Unmarshal(res, &result)
	if err != nil {
		return nil, nil
	}

	return &result, nil
}
