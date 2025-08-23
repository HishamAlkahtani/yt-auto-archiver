package yt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_YTDLPSetup(t *testing.T) {
	_, err := NewYtClient("firefox")
	if err != nil {
		t.Error(err)
	}
}

func Test_getChannelIds(t *testing.T) {
	client, _ := NewYtClient("firefox")

	// idek what this channel is, let's just hope
	// they keep their promise or this test will break.
	vids, err := client.GetVideoIds("YumYumRainbows")

	assert.Nil(t, err)
	assert.NotNil(t, vids)
	assert.Equal(t, 41, len(vids))
	assert.Equal(t, "qjD3xYfYTRk", vids[0])
	assert.Equal(t, "OQ3r8JBXlpI", vids[40])
}
