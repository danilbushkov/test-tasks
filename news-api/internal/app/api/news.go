package api

func (api *Api) regNewsApi() {
	api.registrar().Get("/list", api.Handlers.News.List)
	api.registrar().Post("/edit", api.Handlers.News.Edit)
	api.registrar().Post("/add", api.Handlers.News.Add)
}
