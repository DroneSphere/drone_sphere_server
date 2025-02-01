package user_app

import (
	"context"
	platformapp "drone_sphere_server/internal/domain/platform/app"
	"drone_sphere_server/internal/domain/user"
	"drone_sphere_server/internal/domain/user/repo"
	"drone_sphere_server/internal/infra/eventbus"
	"drone_sphere_server/pkg/token"
	"errors"
)

type Application struct {
	repo repo.IRepository
	bus  *eventbus.EventBus
}

func New(repo repo.IRepository, bus *eventbus.EventBus) *Application {
	return &Application{
		repo: repo,
		bus:  bus,
	}
}

func (a *Application) Register(ctx context.Context, c RegisterCommand) (*LoginResult, error) {
	// 创建领域对象
	u := &user.User{
		Username: c.Username,
		Password: c.Password,
	}

	// 调用领域对象的充血模型方法
	if err := u.HashPassword(); err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 调用仓储接口
	if err := a.repo.Store(u); err != nil {
		return nil, errors.New("failed to store u")
	}

	// 生成 JWT token
	// 调用外部包（相当于调用外部服务），不能放在领域对象/领域服务中
	jwt, err := token.GenerateJWT(u.Username, []byte("secret"))
	if err != nil {
		return nil, errors.New("failed to generate JWT")
	}

	// 返回结果
	u.Password = ""
	return &LoginResult{
		User:  u,
		Token: jwt,
	}, nil
}

// Login 登录
func (a *Application) Login(ctx context.Context, c LoginCommand) (*LoginResult, error) {
	// 创建领域对象
	u, err := a.repo.FindByUsername(c.Username)
	if err != nil {
		return nil, errors.New("u not found")
	}

	// 调用领域对象的充血模型方法
	if !u.Authenticate(c.Password) {
		return nil, errors.New("password is incorrect")
	}

	// 发布登录事件
	eventbus.Publish(a.bus, LoginSuccessEvent{
		User: u,
		SN:   c.SN,
	})

	// 生成 JWT token
	// 调用外部包（相当于调用外部服务），不能放在领域对象/领域服务中
	jwt, err := token.GenerateJWT(u.Username, []byte("secret"))
	if err != nil {
		return nil, errors.New("failed to generate JWT")
	}

	u.Password = ""
	app := platformapp.New()
	pl, _ := app.GetPlatform()
	return &LoginResult{
		User:  u,
		Token: jwt,
		Info: platformapp.InfoResult{
			Platform:  pl.Name,
			Workspace: pl.DefaultWorkspace().Name,
			Desc:      pl.DefaultWorkspace().Description,
		},
		Params: struct {
			MQTT platformapp.MQTTParam `json:"mqtt"`
			HTTP platformapp.HTTPParam `json:"http"`
		}{
			MQTT: platformapp.MQTTParam{
				Host:     "tcp://47.245.40.222:1883",
				Username: "drone",
				Password: "drone",
			},
			HTTP: platformapp.HTTPParam{
				Host:  "http://",
				Token: jwt,
			},
		},
	}, nil
}

func (a *Application) GetUserStatus(ctx context.Context, id int64) (*user.User, error) {
	u, err := a.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("u not found")
	}
	return u, nil
}
