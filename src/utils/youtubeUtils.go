package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"segmenta/src/models"
)

type VideoMetadata struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Channel           string `json:"channel"`
	ChannelLink       string `json:"channel_url"`
	VideoLink         string `json:"webpage_url"`
	ThumbnailImageURL string `json:"thumbnail"`
}

// Error implements [error].
func (v *VideoMetadata) Error() string {
	panic("unimplemented")
}

func FetchVideoMetadata(videoLink string) (*VideoMetadata, error) {
	// Clean the video link (remove extra params like &t=639s&pp=...)
	cleanLink := cleanVideoURL(videoLink)

	// Use --dump-json for reliable JSON output (handles special chars in description)
	cmd := exec.Command("yt-dlp", "--dump-json", "--skip-download", "--no-playlist", cleanLink)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	output, errorHandler := cmd.Output()
	if errorHandler != nil {
		return nil, fmt.Errorf("yt-dlp error: %w, stderr: %s", errorHandler, stderr.String())
	}

	var metadata VideoMetadata
	if errorHandler := json.Unmarshal(output, &metadata); errorHandler != nil {
		return nil, fmt.Errorf("json parse error: %w", errorHandler)
	}

	return &metadata, nil
}

func cleanVideoURL(videoLink string) string {
	if idx := strings.Index(videoLink, "&"); idx != -1 {
		if strings.Contains(videoLink[:idx], "watch?v=") {
			return videoLink[:idx]
		}
	}
	return videoLink
}

func AutoUpdateMetadata(course *models.Course, videoLink string) *VideoMetadata {
	metadata, errorHandler := FetchVideoMetadata(videoLink)
	if errorHandler != nil {
		return nil
	}

	course.Title = metadata.Title
	course.Description = &metadata.Description
	course.Channel = &metadata.Channel
	course.ChannelLink = &metadata.ChannelLink
	course.VideoLink = &videoLink
	course.ThumbnailImageURL = &metadata.ThumbnailImageURL

	return metadata
}

func IsDurationInDescription(description string) bool {
	pattern := regexp.MustCompile(`\b\d{1,2}:\d{2}(?::\d{2})?\b`)
	return pattern.MatchString(description)
}

func ChapterMakerFromDescription(description string, courseID int) []models.Chapter {
	lines := strings.Split(description, "\n")

	var chapters []models.Chapter
	// Matches: optional non-digit/paren prefix (emoji etc), then optional '(', timestamp, optional ')', then title
	timestampPattern := regexp.MustCompile(`\(?(\d{1,2}:\d{2}(?::\d{2})?)\)?\s*[-–—]?\s*(.+)`)

	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		match := timestampPattern.FindStringSubmatch(line)
		if len(match) < 3 {
			continue
		}

		// Make sure the line actually starts close to the timestamp (not just any line with a time-like pattern)
		// Find where the timestamp starts in the original line
		tsIdx := strings.Index(line, match[1])
		if tsIdx == -1 {
			continue
		}
		// Only allow up to 10 runes of prefix before the timestamp (emoji + space)
		prefix := []rune(line[:tsIdx])
		if len(prefix) > 10 {
			continue
		}

		timestamp := strings.TrimSpace(match[1])
		title := strings.TrimSpace(match[2])
		if title == "" {
			continue
		}

		secondTimestamp, errorHandler := ConvertTimestampToSeconds(timestamp)
		if errorHandler != nil || secondTimestamp < 0 {
			continue
		}

		tsStr := strconv.Itoa(secondTimestamp)
		chapter := models.Chapter{
			CourseID:       courseID,
			VideoTimestamp: &tsStr,
			Title:          title,
			Position:       i + 1,
		}
		chapters = append(chapters, chapter)
	}

	// Re-number positions sequentially
	for i := range chapters {
		chapters[i].Position = i + 1
	}

	return chapters
}

func ConvertTimestampToSeconds(timestamp string) (int, error) {
	var totalSeconds int
	parts := strings.Split(timestamp, ":")
	if len(parts) == 2 {
		minutes, errorHandler := strconv.Atoi(parts[0])
		if errorHandler != nil {
			return 0, fmt.Errorf("invalid minutes in timestamp: %w", errorHandler)
		}
		seconds, errorHandler := strconv.Atoi(parts[1])
		if errorHandler != nil {
			return 0, fmt.Errorf("invalid seconds in timestamp: %w", errorHandler)
		}
		totalSeconds = minutes*60 + seconds
	} else if len(parts) == 3 {
		hours, errorHandler := strconv.Atoi(parts[0])
		if errorHandler != nil {
			return 0, fmt.Errorf("invalid hours in timestamp: %w", errorHandler)
		}
		minutes, errorHandler := strconv.Atoi(parts[1])
		if errorHandler != nil {
			return 0, fmt.Errorf("invalid minutes in timestamp: %w", errorHandler)
		}
		seconds, errorHandler := strconv.Atoi(parts[2])
		if errorHandler != nil {
			return 0, fmt.Errorf("invalid seconds in timestamp: %w", errorHandler)
		}
		totalSeconds = hours*3600 + minutes*60 + seconds
	} else {
		return 0, fmt.Errorf("invalid timestamp format")
	}
	return totalSeconds, nil
}
