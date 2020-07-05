package services

import (
	"context"
	"errors"

	"github.com/rylio/ytdl"
)

//GetYoutubeDownloadURL return youtube download url
func GetYoutubeDownloadURL(link string) (string, error) {
	ctx := context.Background()
	client := ytdl.DefaultClient
	videoInfo, err := client.GetVideoInfo(ctx, link)
	if err != nil {
		return "", err
	}
	for _, format := range videoInfo.Formats {
		if format.AudioEncoding == "opus" || format.AudioEncoding == "aac" || format.AudioEncoding == "vorbis" {
			data, err := client.GetDownloadURL(ctx, videoInfo, format)
			if err != nil {
				return "", err
			}
			return data.String(), nil
		}
	}
	return "", errors.New("Audio format not found")
}
