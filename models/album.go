package models

//ImgReq :
type ImgReq struct {
	Data  string `json:"data"`
	Name  string `json:"name"`
	Album string `json:"album"`
}

//GetFileName :
func GetFileName(img, album string) string {
	return img + album
}
