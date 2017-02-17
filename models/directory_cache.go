package models

import (
	"time"

	"github.com/0xdeafcafe/go-xbdm/models"
)

type DirectoryCache struct {
	RefreshedAt time.Time
	Content     []*models.DirectoryItem
}

func NewDirectoryCache(content []*models.DirectoryItem) *DirectoryCache {
	return &DirectoryCache{
		RefreshedAt: time.Now().UTC(),
		Content:     content,
	}
}
