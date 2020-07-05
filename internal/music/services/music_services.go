package services

import (
	"context"
	"errors"

	"github.com/Planxnx/discordBot-Golang/internal/music/model"
	"github.com/rylio/ytdl"
)

//GetYoutubeDownloadURL return youtube download url
func GetYoutubeDownloadURL(link string) (*model.Song, error) {
	ctx := context.Background()
	client := ytdl.DefaultClient
	videoInfo, err := client.GetVideoInfo(ctx, link)
	if err != nil {
		return nil, err
	}

	for _, format := range videoInfo.Formats {
		if format.AudioEncoding == "opus" || format.AudioEncoding == "aac" || format.AudioEncoding == "vorbis" {
			data, err := client.GetDownloadURL(ctx, videoInfo, format)
			if err != nil {
				return nil, err
			}
			return &model.Song{
				Title:        videoInfo.Title,
				Link:         link,
				DownloadLink: data.String(),
				Duration:     videoInfo.Duration,
				Uploader:     videoInfo.Uploader,
				ThumbnailURL: videoInfo.GetThumbnailURL(ytdl.ThumbnailQualityHigh),
			}, nil
		}
	}
	return nil, errors.New("Audio format not found")
}
