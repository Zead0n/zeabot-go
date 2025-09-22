package music

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"time"
)

const MAX_SEARCH_QUERIES = 5

type YtdlpEntry struct {
	Thumbnail  string `json:"thumbnail"`
	Title      string `json:"title"`
	Url        string `json:"webpage_url"`
	Channel    string `json:"channel"`
	ChannelUrl string `json:"channel_url"`
}

func (ye *YtdlpEntry) FormatString() string {
	return fmt.Sprintf("[%s](<%s>) by [%s](<%s>)", ye.Title, ye.Url, ye.Channel, ye.ChannelUrl)
}

type YtdlpResponse struct {
	Thumbnails []struct {
		Url string `json:"url"`
	} `json:"thumbnails"`
	Title      string `json:"title"`
	Url        string `json:"webpage_url"`
	Channel    string `json:"channel"`
	ChannelUrl string `json:"channel_url"`
}

func (yr *YtdlpResponse) toEntry() *YtdlpEntry {
	var entry YtdlpEntry

	entry.Thumbnail = ""
	if len(yr.Thumbnails) > 0 {
		entry.Thumbnail = yr.Thumbnails[len(yr.Thumbnails)-1].Url
	}
	entry.Title = yr.Title
	entry.Url = yr.Url
	entry.Channel = yr.Channel
	entry.ChannelUrl = yr.ChannelUrl

	return &entry
}

func QueryYoutube(query string, search bool) ([]*YtdlpEntry, bool) {
	var fixedQuery string
	if search {
		fixedQuery = fmt.Sprintf("ytsearch%d:%s", MAX_SEARCH_QUERIES, query)
	} else {
		fixedQuery = query
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	out, err := exec.CommandContext(ctx, "yt-dlp", "-f", "ba", "-I", "1:5", "--skip-download", "--print-json", fixedQuery).
		Output()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed yt-dlp query: %s", err.Error()))
		return nil, false
	}

	var entries []*YtdlpEntry
	for _, line := range strings.Split(string(out), "\n") {
		if len(line) <= 0 {
			continue
		}

		var response YtdlpResponse
		if err = json.Unmarshal([]byte(line), &response); err != nil {
			slog.Error(fmt.Sprintf("Failed to parse youtube entry: %s", err.Error()))
			continue
		}

		entries = append(entries, response.toEntry())
	}

	return entries, true
}
