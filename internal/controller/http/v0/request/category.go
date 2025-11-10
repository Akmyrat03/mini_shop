package request

type CreateCategoryReq struct {
	NameTK string `json:"name_tk"`
	NameEN string `json:"name_en"`
	NameRU string `json:"name_ru"`
}
