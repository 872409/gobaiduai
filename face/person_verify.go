package face

const URLPersonVerify = "/rest/2.0/face/v3/person/verify"

type (
	ControlLevel = string
	ImageType    = string
)

const (
	ControlLevelNONE   ControlLevel = "NONE"
	ControlLevelLOW    ControlLevel = "LOW"
	ControlLevelNORMAL ControlLevel = "NORMAL"
	ControlLevelHIGH   ControlLevel = "HIGH"
)

const (
	ImageTypeBASE64 ImageType = "BASE64"
	ImageTypeURL    ImageType = "URL"
)

type PersonVerifyResponse struct {
	AuthResponseJSON

	Result struct {
		Score float32 `json:"score"`
	} `json:"result"`
}

func (p *PersonVerifyResponse) IsMatch(thresholdValue ...float32) bool {
	threshold := float32(80)
	if len(thresholdValue) > 0 {
		threshold = thresholdValue[0]
	}
	return p.Succeed && p.Result.Score >= threshold
}

func (f *FaceClient) PersonVerifyDefault(image string, imageType ImageType, name, idCardNumber string) *PersonVerifyResponse {
	return f.PersonVerify(image, imageType, name, idCardNumber, ControlLevelNORMAL, ControlLevelNORMAL, ControlLevelNORMAL)
}

func (f *FaceClient) PersonVerify(image string, imageType ImageType, name, idCardNumber string, qualityControl, livenessControl, spoofingControl ControlLevel) *PersonVerifyResponse {

	postData := make(map[string]interface{})
	postData["image"] = image
	postData["image_type"] = imageType
	postData["id_card_number"] = idCardNumber
	postData["name"] = name
	postData["quality_control"] = qualityControl
	postData["liveness_control"] = livenessControl
	postData["spoofing_control"] = spoofingControl

	resp := &PersonVerifyResponse{}
	f.httpClient.AuthPOST(URLPersonVerify, nil, postData, resp)

	return resp
}
