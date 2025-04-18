// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1

package types

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code int    `json:"code"`
	Data string `json:data`
	Msg  string `json:"msg"`
}

type UserInfo struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
}

type UserResponse struct {
	Code int      `json:"code"`
	Data UserInfo `json:"data"`
	Msg  string   `json:"msg"`
}
