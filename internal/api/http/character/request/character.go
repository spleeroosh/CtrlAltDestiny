package request

type Character struct {
	Age        uint64 `json:"age" binding:"required"`
	UserID     int    `json:"user_id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Profession string `json:"profession" binding:"required"`
}
