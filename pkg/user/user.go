package user

type User struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Timezone string `json:"timezone" bson:"timezone"`
	Country  string `json:"country" bson:"country"`
}
