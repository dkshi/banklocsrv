package banklocsrv

type User struct {
	Id       string `json:"-" bson:"-"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
