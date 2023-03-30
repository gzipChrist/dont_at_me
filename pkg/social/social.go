package social

import (
	"context"
	"github.com/gzipchrist/dont_at_me/pkg/style"
	"io"
	"net/http"
	"strings"
	"time"
)

type Platform int

const (
	Instagram Platform = iota + 1
	TikTok
	GitHub
	Snapchat
	Twitch
	YouTube
	Mastodon
)

var Platforms = []Platform{Instagram, TikTok, GitHub, Snapchat, Twitch, YouTube, Mastodon}

var PlatformStrings = map[Platform]string{
	Instagram: "Instagram",
	TikTok:    "TikTok",
	GitHub:    "GitHub",
	Snapchat:  "Snapchat",
	Twitch:    "Twitch",
	YouTube:   "YouTube",
	Mastodon:  "Mastodon",
}

var PlatformBaseUrls = map[Platform]string{
	Instagram: "https://instagram.com/",
	TikTok:    "https://us.tiktok.com/@",
	GitHub:    "https://github.com/",
	Snapchat:  "https://www.snapchat.com/add/",
	Twitch:    "https://www.twitch.tv/",
	YouTube:   "https://youtube.com/@",
	Mastodon:  "https://mastodon.social/@",
}

func (p Platform) String() string {
	return PlatformStrings[p]
}

func (p Platform) BaseUrl() string {
	return PlatformBaseUrls[p]
}

func (p Platform) Spacer() int {
	return style.MaxCharWidth - len(p.String())
}

type Status int

const (
	Unavailable Status = iota - 1
	Unknown
	Available
)

var StatusMessages = map[Status]string{
	Unavailable: style.Red.Colorize("✖️"),
	Unknown:     style.Red.Colorize("?"),
	Available:   style.Green.Colorize("✔️"),
}

func (s Status) String() string {
	return StatusMessages[s]
}

func (p Platform) GetAvailability(username string) Status {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	url := p.BaseUrl() + username

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return Unknown
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Unknown
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		switch p {
		case GitHub, YouTube:
			return Available
		}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Unknown
	}

	switch p {
	case Instagram:
		if strings.Contains(string(body), "<title>Instagram</title>") {
			return Available
		}
	case TikTok:
		if strings.Contains(string(body), "Watch the latest video from .") {
			return Available
		}
	case Snapchat:
		if strings.Contains(string(body), "content=\"Not_Found\"") {
			return Available
		}
	case Twitch:
		if strings.Contains(string(body), "content='Twitch is the world") {
			return Available
		}
	case YouTube:
		if strings.Contains(string(body), "<title>404 Not Found</title>") {
			return Available
		}
	case Mastodon:
		if strings.Contains(string(body), "<title>The page you are looking for") || strings.Contains(string(body), "<title>The page you were looking for") {
			return Available
		}
	}

	return Unavailable
}
