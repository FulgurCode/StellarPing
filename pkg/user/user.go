package user

type User struct {
	Id       string `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email" form:"email"`
	Password string `json:"password" bson:"password" form:"password"`
	Timezone string `json:"timezone" bson:"timezone" form:"timezone"`
	Country  string `json:"country" bson:"country" form:"country"`
}
