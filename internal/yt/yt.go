package yt

import (
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"go.uber.org/zap"
)

type YtClient struct {
	browserCookies bool
	browserName    *string
	log            *zap.SugaredLogger
}

type ytDlpOutput struct {
	exitCode int
	stdout   string
	stderr   string
}

// TODO: refactor this to use ytDlpOutput
func NewYtClient(browserName *string, log *zap.SugaredLogger) (*YtClient, error) {
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

	var browserCookies bool = false

	if browserName != nil {
		browserCookies = true
	}

	return &YtClient{
		browserName:    browserName,
		browserCookies: browserCookies,
		log:            log,
	}, nil
}

// TODO: refactor this to use ytDlpOutput
func (c *YtClient) GetVideoIds(channelId string) ([]string, error) {
	link := fmt.Sprintf("https://www.youtube.com/@%s/videos", channelId)

	result, err := c.executeYdlp("--flat-playlist", "--print", "id", link)

	if err != nil {
		return nil, fmt.Errorf("failed to GetVideoIds: %w", err)
	}

	lines := strings.Split(result.stdout, "\n")
	lines = lines[:len(lines)-1]

	return lines, nil
}

// TODO: refactor this to properly return ytDlpOutput
func (c *YtClient) executeYdlp(args ...string) (*ytDlpOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	if c.browserCookies {
		args = append(args, "--cookies-from-browser")
		args = append(args, *c.browserName)
	}

	c.log.Debugf("yt-dlp %s", strings.Join(args, " "))

	output, err := exec.CommandContext(ctx, "yt-dlp", args...).Output()

	if err != nil {
		return nil, fmt.Errorf("failed to execute ytdlp command: %w", err)
	}

	// TODO: what is exitCode doing here?
	// if exitCode other than 0 will be handled here and not returned
	// then it doen't make sense to have it in ytDlpOutput
	return &ytDlpOutput{
		exitCode: 0,
		stdout:   string(output),
	}, nil
}
