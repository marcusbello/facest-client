package facest

import (
	"fmt"
	"net/http"
	"time"
)

// FacesListOptions .
type FacesListOptions struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

// FacesList .
type FacesList struct {
	Count      int    `json:"count"`
	PagesCount int    `json:"pages_count"`
	Faces      []Face `json:"faces"`
}

// Face .
type Face struct {
	FaceToken  string      `json:"face_token"`
	FaceID     string      `json:"face_id"`
	FaceImages []FaceImage `json:"face_images"`
	CreatedAt  time.Time   `json:"created_at"`
}

// FaceImage .
type FaceImage struct {
	ImageToken string    `json:"image_token"`
	ImageURL   string    `json:"image_url"`
	CreatedAt  time.Time `json:"created_at"`
}

func (c *Client) GetFaces(options *FacesListOptions) (*FacesList, error) {
	limit := 100
	page := 1
	if options != nil {
		limit = options.Limit
		page = options.Page
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/faces?limit=%d&page=%d", c.BaseURL, limit, page), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res := FacesList{}

	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
