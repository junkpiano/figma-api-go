# Figma API Client for Golang

## Installation

```
go get -u github.com/junkpiano/figma-api-go
```

## Usage

```go
import (
    "fmt"
    "context"
    _ "github.com/junkpiano/figma-api-go"
)

c := NewClient("token")
ctx := context.Background()

files, _ := c.GetProjectFiles(ctx, "team_id")
fmt.Println(files)
```

## Go Version

1.12.5
