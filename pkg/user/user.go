package user

type User struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Timezone string `bson:"timezone"`
	Country  string `bson:"country"`
}
