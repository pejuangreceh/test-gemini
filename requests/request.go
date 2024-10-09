package requests

import "mime/multipart"

type TextRequest struct {
	Question string `json:"question"`
}
type FileRequest struct {
	File multipart.File `json:"image_upload"`
}
