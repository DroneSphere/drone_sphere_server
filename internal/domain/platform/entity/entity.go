package entity

type Platform struct {
	Name       string
	Workspaces []Workspace
}

type Workspace struct {
	Name        string
	Description string
}

func (e *Platform) DefaultWorkspace() *Workspace {
	return &Workspace{
		Name:        "default",
		Description: "演示环境",
	}
}
