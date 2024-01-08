package response

// type UserResponse struct {
// 	Status  int       `json:"status"`
// 	Message string    `json:"message"`
// 	Data    *echo.Map `json:"data"`
// }

type ServiceResponse struct {
	ID          string `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Statement   string `json:"statement,omitempty" bson:"statement,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Input       string `json:"input,omitempty" bson:"input,omitempty"`
	Output      string `json:"output,omitempty" bson:"output,omitempty"`
	Image       string `json:"image,omitempty" bson:"image,omitempty"`
}
