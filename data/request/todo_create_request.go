package request

type TodoCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
