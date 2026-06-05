package internal

type IConfigProvider interface {
	Load() (map[string]any, error)
}
