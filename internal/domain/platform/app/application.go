package platform_app

import "drone_sphere_server/internal/domain/platform/entity"

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) GetPlatform() (*platform_entity.Platform, error) {
	pl := &platform_entity.Platform{
		Name: "Drone Sphere",
	}
	return pl, nil
}
