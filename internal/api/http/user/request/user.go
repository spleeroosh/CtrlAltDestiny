package request

type User struct {
	Age    uint64 `json:"age" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Social string `json:"social" binding:"required"`
}
