package config

type TokenConfig struct {
	AccessLifeTime  uint64
	AccessKey       string
	RefreshLifeTime uint64
	RefreshKey      string
}
