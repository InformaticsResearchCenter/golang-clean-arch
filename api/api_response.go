package api

type ResponseAPI struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
