package database

import (
	"context"

	"github.com/i3-MS/post-management/internal/audio"
	"github.com/i3-MS/post-management/internal/banner"
	"github.com/i3-MS/post-management/internal/category"
	"github.com/i3-MS/post-management/internal/genre"
	"github.com/i3-MS/post-management/internal/image_post"
	"github.com/i3-MS/post-management/internal/playlist"
	"github.com/i3-MS/post-management/internal/post"
	"github.com/i3-MS/post-management/internal/repost"
	"github.com/i3-MS/post-management/internal/story"
	"github.com/i3-MS/post-management/internal/tag"
	"github.com/i3-MS/post-management/internal/video_post"
)

type Transaction func(dataStore PSQLDBStore) error

type PSQLDBStore interface {
	WithTransaction(ctx context.Context, txFn Transaction) error
	GenreRepo() genre.GenreRepository
	TagRepo() tag.TagRepository
	CategoryRepo() category.CategoryRepository
	BannerRepo() banner.BannerRepository
	RepostRepo() repost.RepostRepository
	AudioRepo() audio.AudioRepository
	VideoPostRepo() video_post.VideoPostRepository
	PostRepo() post.PostRepository
	ImagePostRepo() image_post.ImagePostRepository
	PlaylistRepo() playlist.PlaylistRepository
	StoryRepo() story.StoryRepository
}
