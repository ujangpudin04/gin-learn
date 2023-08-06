package upload

type ImageData struct {
	ID       int    `json:"id"`
	Filename string `json:"filename" binding:"required" form:"filename"`
}

type ImageDataRequest struct {
	Filename string `json:"filename" binding:"required" form:"filename" validate:"required"`
}