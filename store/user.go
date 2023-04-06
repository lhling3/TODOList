package store

type User struct {
	Id       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

// UserMap 合法用户全集
var UserMap = map[uint64]*User{
	1: &User{
		Id:       1,
		Name:     "admin",
		Password: "admin666",
	},
	2: &User{
		Id:       2,
		Name:     "mary",
		Password: "mary123",
	},
	3: &User{
		Id:       3,
		Name:     "jack",
		Password: "jack123",
	},
}
