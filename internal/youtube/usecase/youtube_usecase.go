package usecase

import (
	"context"
	"fmt"

	"github.com/Planxnx/discordBot-Golang/internal/music/model"
	"github.com/rylio/ytdl"
)

//Usecase interface
type Usecase interface {
	GetYoutubeDownloadURL(string) (*model.Song, error)
}

type youtubeUsecase struct {
	ytdlClient *ytdl.Client
}

//NewYoutubeUsecase new message delivery
func NewYoutubeUsecase() Usecase {
	client := ytdl.DefaultClient

	return &youtubeUsecase{
		ytdlClient: client,
	}
}

//GetYoutubeDownloadURL return youtube download url
func (yu youtubeUsecase) GetYoutubeDownloadURL(link string) (*model.Song, error) {
	ctx := context.Background()
	client := yu.ytdlClient
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
	return nil, fmt.Errorf("Audio format not found")
}
