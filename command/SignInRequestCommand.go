package command

type SignInRequestCommand struct {
	Mail     string `json:"mail"`
	Password string `json:"pass"`
}
