package handlers

type UsecaseURL interface {
	Shorten(url string) (string, error)
	Expand(url string) (string, error)
}

type Manager struct {
	usecase UsecaseURL
}

func NewManager(usecase UsecaseURL) *Manager {
	return &Manager{
		usecase: usecase,
	}
}
