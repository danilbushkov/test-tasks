package api

func (api *Api) regNewsApi() {
	api.registrar().Get("/list", api.Handlers.News.List)
	api.registrar().Post("/edit", api.Handlers.News.Edit)
	api.registrar().Get("/add", api.Handlers.News.Add)
}
