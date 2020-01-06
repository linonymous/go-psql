package geolocation

type GeolocationService struct {
	config     *Config
	repository *GeolocationRepository
}

func NewGeolocationService(config *Config, repository *GeolocationRepository) *GeolocationService {
	return &GeolocationService{
		config:     config,
		repository: repository,
	}
}

func (g *GeolocationService) Query() {
	if !g.config.Enabled {
		return
	}
	g.repository.Query()
}
