package figma

import (
	"context"
	"encoding/json"
)

func UnmarshalFile(data []byte) (*File, error) {
	var r File
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *File) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFileNodes(data []byte) (*FileNodes, error) {
	var r FileNodes
	err := json.Unmarshal(data, &r)
	return &r, err
}

func (r *FileNodes) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type File struct {
	Document      Document              `json:"document"`
	Components    Components            `json:"components"`
	SchemaVersion int64                 `json:"schemaVersion"`
	Styles        map[string]StyleValue `json:"styles"`
	Name          string                `json:"name"`
	LastModified  string                `json:"lastModified"`
	ThumbnailURL  string                `json:"thumbnailUrl"`
	Version       string                `json:"version"`
}

type FileNode struct {
	Document      Document              `json:"document"`
	Components    Components            `json:"components"`
	SchemaVersion int64                 `json:"schemaVersion"`
	Styles        map[string]StyleValue `json:"styles"`
}

type Components struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Color struct {
	R float64 `json:"r"`
	G float64 `json:"g"`
	B float64 `json:"b"`
	A int64   `json:"a"`
}

type Document struct {
	ID                      string              `json:"id"`
	Name                    string              `json:"name"`
	Type                    Type                `json:"type"`
	BlendMode               BlendMode           `json:"blendMode"`
	Opacity                 *float64            `json:"opacity,omitempty"`
	AbsoluteBoundingBox     AbsoluteBoundingBox `json:"absoluteBoundingBox"`
	Constraints             Constraints         `json:"constraints"`
	Fills                   []FillElement       `json:"fills"`
	Strokes                 []interface{}       `json:"strokes"`
	StrokeWeight            int64               `json:"strokeWeight"`
	StrokeAlign             StrokeAlign         `json:"strokeAlign"`
	ExportSettings          []interface{}       `json:"exportSettings"`
	Effects                 []interface{}       `json:"effects"`
	Characters              *string             `json:"characters,omitempty"`
	Style                   *Style              `json:"style,omitempty"`
	CharacterStyleOverrides []interface{}       `json:"characterStyleOverrides"`
	StyleOverrideTable      *Components         `json:"styleOverrideTable,omitempty"`
	StrokeJoin              *StrokeJoin         `json:"strokeJoin,omitempty"`
	StrokeMiterAngle        interface{}         `json:"strokeMiterAngle"`
	Styles                  *Styles             `json:"styles,omitempty"`
	CornerRadius            *int64              `json:"cornerRadius,omitempty"`
	Children                []Document          `json:"children,omitempty"`
	BackgroundColor         Color               `json:"backgroundColor"`
	PrototypeStartNodeID    interface{}         `json:"prototypeStartNodeID"`
}

type AbsoluteBoundingBox struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type Constraints struct {
	Vertical   Vertical   `json:"vertical"`
	Horizontal Horizontal `json:"horizontal"`
}

type FillElement struct {
	BlendMode BlendMode `json:"blendMode"`
	Type      FillType  `json:"type"`
	Color     Color     `json:"color"`
}

type Style struct {
	FontFamily                string      `json:"fontFamily"`
	FontPostScriptName        interface{} `json:"fontPostScriptName"`
	FontWeight                int64       `json:"fontWeight"`
	FontSize                  int64       `json:"fontSize"`
	TextAlignHorizontal       Horizontal  `json:"textAlignHorizontal"`
	TextAlignVertical         Vertical    `json:"textAlignVertical"`
	LetterSpacing             int64       `json:"letterSpacing"`
	LineHeightPx              float64     `json:"lineHeightPx"`
	LineHeightPercent         int64       `json:"lineHeightPercent"`
	LineHeightPercentFontSize float64     `json:"lineHeightPercentFontSize"`
	LineHeightUnit            string      `json:"lineHeightUnit"`
}

type Styles struct {
	Fill string `json:"fill"`
}

type StyleValue struct {
	Key       string    `json:"key"`
	Name      string    `json:"name"`
	StyleType StyleType `json:"styleType"`
}

type BlendMode string

const (
	PassThrough BlendMode = "PASS_THROUGH"
	Normal      BlendMode = "Normal"
)

type Horizontal string

const (
	Left Horizontal = "LEFT"
)

type Vertical string

const (
	Top Vertical = "TOP"
)

type FillType string

const (
	Solid FillType = "SOLID"
)

type StrokeAlign string

const (
	Center  StrokeAlign = "CENTER"
	Outside StrokeAlign = "OUTSIDE"
)

type StrokeJoin string

const (
	Bevel StrokeJoin = "BEVEL"
)

type Type string

const (
	Rectangle Type = "RECTANGLE"
	Text      Type = "TEXT"
)

type StyleType string

const (
	Fill StyleType = "FILL"
)

type FileNodes struct {
	Name         string              `json:"name"`
	LastModified string              `json:"lastModified"`
	ThumbnailURL string              `json:"thumbnailUrl"`
	Version      string              `json:"version"`
	Nodes        map[string]FileNode `json:"nodes,omitempty"`
}

func (c *Client) GetFile(ctx context.Context, fileKey string) (*File, error) {
	req, _ := c.newRequest(ctx, "GET", "/v1/files/"+fileKey, nil)
	res, err := c.send(ctx, req)

	if err != nil {
		return nil, err
	}

	return UnmarshalFile(res)
}

func (c *Client) GetFileNode(ctx context.Context, fileKey, ids string) (*FileNodes, error) {
	req, _ := c.newRequest(ctx, "GET", "v1/files/"+fileKey+"/nodes", nil)

	q := req.URL.Query()
	q.Add("ids", ids)
	req.URL.RawQuery = q.Encode()

	res, err := c.send(ctx, req)

	if err != nil {
		return nil, err
	}

	return UnmarshalFileNodes(res)
}
