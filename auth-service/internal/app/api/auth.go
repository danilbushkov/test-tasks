package api

func (api *Api) regAuthApi() {
	api.registrar().Post("/api/auth/get", api.Handlers.Auth.Get)
	api.registrar().Post("/api/auth/refresh", api.Handlers.Auth.Refresh)
}
