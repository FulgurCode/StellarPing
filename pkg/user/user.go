package user

type User struct {
	Email    string `json:"email" bson:"email" form:"email"`
	Password string `json:"password" bson:"password" form:"password"`
	Timezone string `json:"timezone" bson:"timezone" form:"timezone"`
	Country  string `json:"country" bson:"country" form:"country"`
}
