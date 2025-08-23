package yt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_YTDLPSetup(t *testing.T) {
	log, _ := zap.NewDevelopment()
	_, err := NewYtClient(nil, log.Sugar())
	if err != nil {
		t.Error(err)
	}
}

func Test_getChannelIds(t *testing.T) {
	log, _ := zap.NewDevelopment()
	client, _ := NewYtClient(nil, log.Sugar())

	// idek what this channel is, let's just hope
	// they keep their promise or this test will break.
	vids, err := client.GetVideoIds("YumYumRainbows")

	assert.Nil(t, err)
	assert.NotNil(t, vids)
	assert.Equal(t, 41, len(vids))
	assert.Equal(t, "qjD3xYfYTRk", vids[0])
	assert.Equal(t, "OQ3r8JBXlpI", vids[40])
}
