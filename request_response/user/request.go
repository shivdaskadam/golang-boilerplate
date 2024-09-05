package user

type UserReq struct {
	Id   int    `json:"id"`
	Name string `json:"id"`
}

type GetUserReq struct {
	Id int `json:"id"`
	// UserReq
}

type PutUerReq struct {
	UserReq
	Age int `json:"age"`
}
