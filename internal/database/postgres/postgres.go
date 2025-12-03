package postgres

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/i3-MS/post-management/internal/audio"
	audioRepo "github.com/i3-MS/post-management/internal/audio/repository/postgres"
	"github.com/i3-MS/post-management/internal/banner"
	bannerRepo "github.com/i3-MS/post-management/internal/banner/repository/postgres"
	"github.com/i3-MS/post-management/internal/category"
	categoryRepo "github.com/i3-MS/post-management/internal/category/repository/postgres"
	genreRepo "github.com/i3-MS/post-management/internal/genre/repository/postgres"
	"github.com/i3-MS/post-management/internal/image_post"
	"github.com/i3-MS/post-management/internal/playlist"
	"github.com/i3-MS/post-management/internal/post"
	postRepo "github.com/i3-MS/post-management/internal/post/repository/postgres"
	"github.com/i3-MS/post-management/internal/repost"
	repostRepo "github.com/i3-MS/post-management/internal/repost/repository/postgres"
	"github.com/i3-MS/post-management/internal/story"
	"github.com/i3-MS/post-management/internal/tag"
	tagRepo "github.com/i3-MS/post-management/internal/tag/repository/postgres"
	"github.com/i3-MS/post-management/internal/video_post"
	videoPostRepo "github.com/i3-MS/post-management/internal/video_post/repository/postgres"

	playlistRepo "github.com/i3-MS/post-management/internal/playlist/repository/postgres"

	storyRepo "github.com/i3-MS/post-management/internal/story/repository/postgres"

	imagePostRepo "github.com/i3-MS/post-management/internal/image_post/repository/postgres"

	"github.com/i3-MS/post-management/internal/database"
	"github.com/i3-MS/post-management/internal/genre"
	"github.com/i3-MS/post-management/pkg/connection"
	"github.com/jackc/pgx/v5"

	"go.uber.org/zap"
)

// Ensuring DataStore implements database.PSQLDBStore.
var _ database.PSQLDBStore = (*DataStore)(nil)

type DataStore struct {
	db            connection.DB
	genre         genre.GenreRepository
	genreInit     sync.Once
	tag           tag.TagRepository
	tagInit       sync.Once
	category      category.CategoryRepository
	categoryInit  sync.Once
	banner        banner.BannerRepository
	bannerInit    sync.Once
	repost        repost.RepostRepository
	repostInit    sync.Once
	audio         audio.AudioRepository
	audioInit     sync.Once
	videoPost     video_post.VideoPostRepository
	videoPostInit sync.Once
	post          post.PostRepository
	postInit      sync.Once
	playlist      playlist.PlaylistRepository
	playlistInit  sync.Once
	imagePost     image_post.ImagePostRepository
	imagePostInit sync.Once
	story         story.StoryRepository
	storyInit     sync.Once
}

// NewPSQLDBStore creates and returns a new instance of DataStore.
func NewPSQLDBStore(db connection.DBOps) database.PSQLDBStore {
	return &DataStore{
		db: db,
	}
}

// WithTransaction method is a transaction method for performing multitasks.
func (d *DataStore) WithTransaction(ctx context.Context, transactionFn database.Transaction) error {
	// Assert that the db implements DBOps for transaction capabilities.
	db, ok := d.db.(connection.DBOps)
	if !ok {
		return errors.New("[postgres.WithTransaction]: err in type assertion")
	}

	// begin transaction.
	tx, err := db.Begin(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("[postgres.WithTransaction]: err in db.Begin: %w", err)
	}

	// Ensure the transaction is rolled back if an error occurs.
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				zap.Any("rbErr in tx.Rollback", rbErr)
			}
			zap.Any("transaction rolled back due to error", err)
		}
	}()

	// Wrap the database in the transactional context.
	transactionalDB := &DataStore{db: tx}

	// Run the transaction function.
	err = transactionFn(transactionalDB)
	if err != nil {
		return fmt.Errorf("error during transaction function execution: %w", err)
	}

	// Commit the transaction if no error occurred during execution.
	if cErr := tx.Commit(ctx); cErr != nil {
		return fmt.Errorf("error in committing transaction: %w", cErr)
	}

	return nil
}

func (d *DataStore) GenreRepo() genre.GenreRepository {
	d.genreInit.Do(func() {
		d.genre = genreRepo.NewGenreRepository(d.db)
	})

	return d.genre
}

func (d *DataStore) TagRepo() tag.TagRepository {
	d.tagInit.Do(func() {
		d.tag = tagRepo.NewTagRepository(d.db)
	})

	return d.tag
}

func (d *DataStore) CategoryRepo() category.CategoryRepository {
	d.categoryInit.Do(func() {
		d.category = categoryRepo.NewCategoryRepository(d.db)
	})

	return d.category
}

func (d *DataStore) BannerRepo() banner.BannerRepository {
	d.bannerInit.Do(func() {
		d.banner = bannerRepo.NewBannerRepository(d.db)
	})

	return d.banner
}

func (d *DataStore) RepostRepo() repost.RepostRepository {
	d.repostInit.Do(func() {
		d.repost = repostRepo.NewRepostRepository(d.db)
	})

	return d.repost
}

func (d *DataStore) AudioRepo() audio.AudioRepository {
	d.audioInit.Do(func() {
		d.audio = audioRepo.NewAudioRepository(d.db)
	})

	return d.audio
}

func (d *DataStore) VideoPostRepo() video_post.VideoPostRepository {
	d.videoPostInit.Do(func() {
		d.videoPost = videoPostRepo.NewVideoPostRepository(d.db)
	})

	return d.videoPost
}

func (d *DataStore) PostRepo() post.PostRepository {
	d.postInit.Do(func() {
		d.post = postRepo.NewPostRepository(d.db)
	})

	return d.post
}

func (d *DataStore) PlaylistRepo() playlist.PlaylistRepository {
	d.playlistInit.Do(func() {
		d.playlist = playlistRepo.NewPlaylistRepository(d.db)
	})

	return d.playlist
}

func (d *DataStore) ImagePostRepo() image_post.ImagePostRepository {
	d.imagePostInit.Do(func() {
		d.imagePost = imagePostRepo.NewImagePostRepository(d.db)
	})

	return d.imagePost
}

func (d *DataStore) StoryRepo() story.StoryRepository {
	d.storyInit.Do(func() {
		d.story = storyRepo.NewStoryRepository(d.db)
	})

	return d.story
}
