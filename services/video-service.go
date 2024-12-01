package services

import (
	"errors"
)

type Video struct {
	ID        string
	VideoPath string
}

type VideoService interface {
	GetVideoByID(id string) (*Video, error)
}

type videoServiceImpl struct {
	videos map[string]*Video
}

func NewVideoService() VideoService {
	return &videoServiceImpl{
		videos: map[string]*Video{
			"123": {ID: "123", VideoPath: "video.mp4"},
		},
	}
}

func (vs *videoServiceImpl) GetVideoByID(id string) (*Video, error) {
	video, exists := vs.videos[id]
	if !exists {
		return nil, errors.New("video not found")
	}
	return video, nil
}
