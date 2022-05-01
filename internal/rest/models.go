package rest

type AddItemRequest struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

type GetItemRequest struct {
	Key string `uri:"key" binding:"required"`
}

type ItemResponse struct {
	Key string `json:"key"`
	Val string `json:"val"`
}
