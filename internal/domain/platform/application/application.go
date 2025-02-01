package application

import "drone_sphere_server/internal/domain/platform/entity"

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) GetPlatform() (*entity.Platform, error) {
	pl := &entity.Platform{
		Name: "Drone Sphere",
	}
	return pl, nil
}
