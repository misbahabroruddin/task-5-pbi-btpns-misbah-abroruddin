package app

type PhotoRequestBody struct{
	Title string `json:"title" binding:"required"`
	Caption string `json:"caption" binding:"required"`
	PhotoUrl string `json:"photo_url" binding:"required"`
}