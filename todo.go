package TODO_app

type TodoList struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desciption string `json:"desciption"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Desciption string `json:"desciption"`
	Done       bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
