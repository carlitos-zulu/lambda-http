package entities

type (
	RequestData struct {
		UserID string `json:"userId" form:"userId" binding:"required"`
	}

	User struct {
		ID     string `json:"_id"`
		Name   string `json:"full_name"`
		Method string `json:"method"`
	}
)
