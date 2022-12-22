// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package post

//go:generate mockgen -destination mock_post.go -package post github.com/marmotedu/miniblog/internal/miniblog/biz/post PostBiz

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/marmotedu/miniblog/internal/miniblog/store"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	"github.com/marmotedu/miniblog/internal/pkg/log"
	"github.com/marmotedu/miniblog/internal/pkg/model"
	v1 "github.com/marmotedu/miniblog/pkg/api/miniblog/v1"
)

// PostBiz defines functions used to handle post request.
type PostBiz interface {
	Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error)
	Update(ctx context.Context, username, postID string, r *v1.UpdatePostRequest) error
	Delete(ctx context.Context, username, postID string) error
	DeleteCollection(ctx context.Context, username string, postIDs []string) error
	Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error)
	List(ctx context.Context, username string, offset, limit int) (*v1.ListPostResponse, error)
}

// The implementation of PostBiz interface.
type postBiz struct {
	ds store.IStore
}

// Make sure that postBiz implements the PostBiz interface.
// We can find this problem in the compile stage with the following assignment statement.
var _ PostBiz = (*postBiz)(nil)

func New(ds store.IStore) *postBiz {
	return &postBiz{ds: ds}
}

// Create is the implementation of the `Create` method in PostBiz interface.
func (b *postBiz) Create(ctx context.Context, username string, r *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	var postM model.PostM
	_ = copier.Copy(&postM, r)
	postM.Username = username

	if err := b.ds.Posts().Create(ctx, &postM); err != nil {
		return nil, err
	}

	return &v1.CreatePostResponse{PostID: postM.PostID}, nil
}

// Delete is the implementation of the `Delete` method in PostBiz interface.
func (b *postBiz) Delete(ctx context.Context, username, postID string) error {
	if err := b.ds.Posts().Delete(ctx, username, []string{postID}); err != nil {
		return err
	}

	return nil
}

// DeleteCollection is the implementation of the `DeleteCollection` method in PostBiz interface.
func (b *postBiz) DeleteCollection(ctx context.Context, username string, postIDs []string) error {
	if err := b.ds.Posts().Delete(ctx, username, postIDs); err != nil {
		return err
	}

	return nil
}

// Get is the implementation of the `Get` method in PostBiz interface.
func (b *postBiz) Get(ctx context.Context, username, postID string) (*v1.GetPostResponse, error) {
	post, err := b.ds.Posts().Get(ctx, username, postID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}

		return nil, err
	}

	var resp v1.GetPostResponse
	_ = copier.Copy(&resp, post)

	resp.CreatedAt = post.CreatedAt.Format("2006-01-02 15:04:05")
	resp.UpdatedAt = post.UpdatedAt.Format("2006-01-02 15:04:05")

	return &resp, nil
}

// Update is the implementation of the `Update` method in PostBiz interface.
func (b *postBiz) Update(ctx context.Context, username, postID string, r *v1.UpdatePostRequest) error {
	postM, err := b.ds.Posts().Get(ctx, username, postID)
	if err != nil {
		return err
	}

	if r.Title != nil {
		postM.Title = *r.Title
	}

	if r.Content != nil {
		postM.Content = *r.Content
	}

	if err := b.ds.Posts().Update(ctx, postM); err != nil {
		return err
	}

	return nil
}

// List is the implementation of the `List` method in PostBiz interface.
func (b *postBiz) List(ctx context.Context, username string, offset, limit int) (*v1.ListPostResponse, error) {
	count, list, err := b.ds.Posts().List(ctx, username, offset, limit)
	if err != nil {
		log.C(ctx).Errorw("Failed to list posts from storage", "err", err)
		return nil, err
	}

	posts := make([]*v1.PostInfo, 0, len(list))
	for _, item := range list {
		post := item
		posts = append(posts, &v1.PostInfo{
			Username:  post.Username,
			PostID:    post.PostID,
			Title:     post.Title,
			Content:   post.Content,
			CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: post.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &v1.ListPostResponse{TotalCount: count, Posts: posts}, nil
}
