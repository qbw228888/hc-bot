package entity

type Sender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Card     string `json:"card"`  //群名片，群备注
	Area     string `json:"area"`  //地区
	Level    string `json:"level"` //等级
	Role     string `json:"role"`  // 角色
	Title    string `json:"title"` //专属头衔
}
