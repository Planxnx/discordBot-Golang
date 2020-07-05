package model

import (
	"net/url"
	"time"
)

type Song struct {
	Title        string
	Link         string
	DownloadLink string
	Uploader     string
	Duration     time.Duration
	ThumbnailURL *url.URL
}
