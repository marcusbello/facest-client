package facest

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const (
	faceID = "integration_face_id"
)

// will be set by TestGetFaces
var faceToken string

func TestGetFaces(t *testing.T) {
	c := NewClient(os.Getenv("FACEST_INTEGRATION_API_KEY"))

	res, err := c.GetFaces(nil)
	assert.Nil(t, err, "Expecting Nil Error")
	require.Nil(t, err, "Expecting Nil Error")
	assert.NotNil(t, res, "expecting non-nil result")
	if res != nil {
		assert.Equal(t, 1, res.Count, "expecting 1 face found")
		assert.Equal(t, 1, res.PagesCount, "expecting 1 PAGE found")

		if res.Count > 0 {
			assert.Equal(t, faceID, res.Faces[0].FaceID, "expecting correct face_id")
			assert.NotEmpty(t, res.Faces[0].FaceToken, "expecting non-empty face_token")
			assert.Greater(t, len(res.Faces[0].FaceImages), 0, "expecting non-empty face_images")

			faceToken = res.Faces[0].FaceToken
		}
	}
}
