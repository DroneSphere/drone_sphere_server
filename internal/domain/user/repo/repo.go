package repo

import (
	"drone_sphere_server/internal/domain/user"
	"drone_sphere_server/internal/infra/rdb"
	"drone_sphere_server/pkg/log"
)

type IRepository interface {
	Store(user *user.User) error
	FindByID(id int64) (*user.User, error)
	FindByUsername(username string) (*user.User, error)
}

type Repository struct {
	rdb *rdb.RDB
}

func NewRepository(rdb *rdb.RDB) *Repository {
	logger := log.GetLogger()
	logger.Info("auto migrate user table")
	err := rdb.DB.AutoMigrate(&user.User{})
	if err != nil {
		panic(err)
	}

	return &Repository{
		rdb: rdb,
	}
}

func (r *Repository) Store(user *user.User) error {
	// 如果 ID 为 0，说明是新建用户
	if user.ID == 0 {
		return r.rdb.DB.Create(user).Error
	}
	// 如果 ID 不为 0，说明是更新用户
	return r.rdb.DB.Save(user).Error
}

func (r *Repository) FindByID(id int64) (*user.User, error) {
	u := &user.User{}
	err := r.rdb.DB.First(u, id).Error
	return u, err
}

func (r *Repository) FindByUsername(username string) (*user.User, error) {
	u := &user.User{}
	err := r.rdb.DB.Where("username = ?", username).First(u).Error
	return u, err
}
