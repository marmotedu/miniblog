// Copyright 2022 Innkeeper Belm(孔令飞) <nosbelm@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/marmotedu/miniblog.

package user

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"

	"github.com/marmotedu/miniblog/internal/miniblog/store"
	"github.com/marmotedu/miniblog/internal/pkg/errno"
	"github.com/marmotedu/miniblog/internal/pkg/model"
	v1 "github.com/marmotedu/miniblog/pkg/api/miniblog/v1"
)

func fakeUser(id int64) *model.UserM {
	return &model.UserM{
		ID:        id,
		Username:  fmt.Sprintf("belm%d", id),
		Password:  fmt.Sprintf("belm%d", id),
		Nickname:  fmt.Sprintf("belm%d", id),
		Email:     "nosbelm@qq.com",
		Phone:     "18188888xxx",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Test_New(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := store.NewMockIStore(ctrl)

	type args struct {
		ds store.IStore
	}
	tests := []struct {
		name string
		args args
		want *userBiz
	}{
		{name: "default", args: args{ds: mockStore}, want: &userBiz{ds: mockStore}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.ds)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_userBiz_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 构造期望的返回结果
	fakeUsers := []*model.UserM{fakeUser(1), fakeUser(2), fakeUser(3)}
	wantUsers := make([]*v1.UserInfo, 0, len(fakeUsers))
	for _, u := range fakeUsers {
		wantUsers = append(wantUsers, &v1.UserInfo{
			Username:  u.Username,
			Nickname:  u.Nickname,
			Email:     u.Email,
			Phone:     u.Email,
			PostCount: 10,
			CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(5), fakeUsers, nil).Times(1)

	mockPostStore := store.NewMockPostStore(ctrl)
	mockPostStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(10), nil, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).Times(1)
	mockStore.EXPECT().Posts().Return(mockPostStore).AnyTimes()

	tests := []struct {
		name    string
		want    *v1.ListUserResponse
		wantErr bool
	}{
		{name: "default", want: &v1.ListUserResponse{TotalCount: 5, Users: wantUsers}, wantErr: false},
	}

	ub := New(mockStore)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ub.List(context.Background(), 0, 10)
			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ds store.IStore
	}
	tests := []struct {
		name string
		args args
		want *userBiz
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New(tt.args.ds))
		})
	}
}

func Test_userBiz_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		r   *v1.CreateUserRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			assert.Nil(t, b.Create(tt.args.ctx, tt.args.r))
		})
	}
}

func Test_userBiz_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	var want v1.GetUserResponse
	_ = copier.Copy(&want, fakeUser)
	want.CreatedAt = fakeUser.CreatedAt.Format("2006-01-02 15:04:05")
	want.UpdatedAt = fakeUser.UpdatedAt.Format("2006-01-02 15:04:05")

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *v1.GetUserResponse
	}{
		{name: "default", fields: fields{ds: mockStore}, args: args{context.Background(), "belm"}, want: &want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			got, err := b.Get(tt.args.ctx, tt.args.username)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_userBiz_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	r := &v1.UpdateUserRequest{
		Email: pointer.ToString("belm@qq.com"),
		Phone: pointer.ToString("18866xxxxxx"),
	}
	wantedUser := *fakeUser
	wantedUser.Email = *r.Email
	wantedUser.Phone = *r.Phone

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()
	mockUserStore.EXPECT().Update(gomock.Any(), &wantedUser).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
		user     *v1.UpdateUserRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), "belm", r}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			err := b.Update(tt.args.ctx, tt.args.username, tt.args.user)
			assert.Nil(t, err)
		})
	}
}

func Test_userBiz_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "default", fields: fields{mockStore}, args: args{context.Background(), "belm"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			assert.Nil(t, b.Delete(tt.args.ctx, tt.args.username))
		})
	}
}

func Test_userBiz_ChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()
	mockUserStore.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx      context.Context
		username string
		r        *v1.ChangePasswordRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "default",
			fields: fields{mockStore},
			args:   args{context.Background(), "belm", &v1.ChangePasswordRequest{"miniblog1234", "miniblog12345"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			err := b.ChangePassword(tt.args.ctx, tt.args.username, tt.args.r)
			assert.Equal(t, errno.ErrPasswordIncorrect, err)
		})
	}
}

func Test_userBiz_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fakeUser := fakeUser(1)
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(fakeUser, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().AnyTimes().Return(mockUserStore)

	type fields struct {
		ds store.IStore
	}
	type args struct {
		ctx context.Context
		r   *v1.LoginRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *errno.Errno
	}{
		{
			name:   "default",
			fields: fields{mockStore},
			args:   args{context.Background(), &v1.LoginRequest{Username: "belm", Password: "miniblog1234"}},
			want:   errno.ErrPasswordIncorrect,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &userBiz{
				ds: tt.fields.ds,
			}
			got, err := b.Login(tt.args.ctx, tt.args.r)
			assert.Nil(t, got)
			assert.Equal(t, tt.want, err)
		})
	}
}

func BenchmarkListUser(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	// 构造期望的返回结果
	fakeUsers := []*model.UserM{fakeUser(1), fakeUser(2), fakeUser(3)}
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(5), fakeUsers, nil).AnyTimes()

	mockPostStore := store.NewMockPostStore(ctrl)
	mockPostStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(10), nil, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()
	mockStore.EXPECT().Posts().Return(mockPostStore).AnyTimes()

	ub := New(mockStore)
	for i := 0; i < b.N; i++ {
		_, _ = ub.List(context.TODO(), 0, 0)
	}
}

func BenchmarkListWithBadPerformance(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	// 构造期望的返回结果
	fakeUsers := []*model.UserM{fakeUser(1), fakeUser(2), fakeUser(3)}
	mockUserStore := store.NewMockUserStore(ctrl)
	mockUserStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(5), fakeUsers, nil).AnyTimes()

	mockPostStore := store.NewMockPostStore(ctrl)
	mockPostStore.EXPECT().List(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(int64(10), nil, nil).AnyTimes()

	mockStore := store.NewMockIStore(ctrl)
	mockStore.EXPECT().Users().Return(mockUserStore).AnyTimes()
	mockStore.EXPECT().Posts().Return(mockPostStore).AnyTimes()

	ub := New(mockStore)
	for i := 0; i < b.N; i++ {
		_, _ = ub.ListWithBadPerformance(context.TODO(), 0, 0)
	}
}
