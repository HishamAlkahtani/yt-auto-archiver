package yt

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
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

func (c *YtClient) GetVideoIds(channelId string) ([]string, error) {
	link := fmt.Sprintf("https://www.youtube.com/@%s/videos", channelId)

	result, err := c.executeYdlp("--flat-playlist", "--print", "id", link)

	if err != nil {
		return nil, fmt.Errorf("failed to GetVideoIds: %w", err)
	}

	lines := strings.Split(result, "\n")
	lines = lines[:len(lines)-1]

	return lines, nil
}

func (c *YtClient) executeYdlp(args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	if c.browserCookies {
		args = append(args, "--cookies-from-browser")
		args = append(args, c.browserName)
	}

	// TODO: before execution, log the whole command

	output, err := exec.CommandContext(ctx, "yt-dlp", args...).Output()

	if err != nil {
		return "", fmt.Errorf("failed to execute ytdlp command: %w", err)
	}

	return string(output), nil
}
