package yt

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"time"
)

type YtClient struct {
	browserCookies bool
	browserName    string
}

func NewYtClient(browserName string) (*YtClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	output, err := exec.CommandContext(ctx, "yt-dlp", "--version").Output()

	if err != nil {
		return nil, fmt.Errorf("failed to create yt client: %w", err)
	}

	var expectedFormat = regexp.MustCompile(`^\d{4}\.\d{2}\.\d{2}\r?\n?$`)

	if !expectedFormat.MatchString(string(output)) {
		return nil, fmt.Errorf("yt-dlp --version returned unexpected output: %s", string(output))
	}

	return &YtClient{}, nil
}

func (c *YtClient) GetVideoIds(channelId string) []string {
	// TODO
	return []string{}
}
