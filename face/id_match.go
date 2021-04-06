package face

const (
	URLIDMatch = "/rest/2.0/face/v3/person/idmatch"
)

func (f *FaceClient) IDMatch(name, idCardNumber string) *AuthResponseJSON {
	resp := &AuthResponseJSON{}
	postData := make(map[string]interface{})
	postData["name"] = name
	postData["id_card_number"] = idCardNumber
	f.httpClient.AuthPOST(URLIDMatch, nil, postData, resp)
	return resp
}
