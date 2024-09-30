package responses

type ResponseData struct {
	Code        string      `json:"responseCode"`
	Description string      `json:"responseDescription"`
	Data        interface{} `json:"data"`
}
