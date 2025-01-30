package user_app

import (
	mockuser "drone_sphere_server/internal/domain/user/repo"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestApplication_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockuser.NewMockIRepository(ctrl)
	repo.EXPECT().Store(gomock.Any()).Return(nil).AnyTimes()

	command := RegisterCommand{
		Username: "test",
		Password: "test",
	}
	app := NewApplication(repo)

	res, err := app.Register(nil, command)
	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}
	assert.NotEmpty(t, res.Token)
	assert.Empty(t, res.User.Password)
}
