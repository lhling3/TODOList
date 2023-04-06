package store

import "time"

type TodoList struct {
	Id         uint      `json:"id" form:"id"`
	Title      string    `json:"title" form:"title"`
	Body       string    `json:"body" form:"body"`
	UserID     uint64    `json:"user_id" form:"user_id"`
	CreateTime time.Time `json:"create_time" form:"create_time"`
	UpdateTime time.Time `json:"update_time" form:"update_time"`
}

var TODOListMap = map[uint64]map[uint]TodoList{
	1: {
		1: {
			Id:         1,
			Title:      "用户1的代办事项1",
			Body:       "早上九点半开会",
			UserID:     1,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		3: {
			Id:         3,
			Title:      "用户1的代办事项2",
			Body:       "回复邮件",
			UserID:     1,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	},
	2: {
		2: {
			Id:         2,
			Title:      "用户2的代办事项1",
			Body:       "提交文件",
			UserID:     2,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
		5: {
			Id:         5,
			Title:      "用户2的代办事项2",
			Body:       "下午3点开会",
			UserID:     2,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	},
	3: {
		3: {
			Id:         3,
			Title:      "用户3的代办事项1",
			Body:       "提交文件",
			UserID:     3,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	},
}
