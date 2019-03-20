package db

type User struct{
	Name string `json:"name"`
	Value string `json:"value"`
}
type Users struct{
	element []User
}

var db = make(map[string]string)
func Init(){
	db["test"] = "prueba"
	db["roberto"] = "Backend"
	db["user1"] = "Frontend"
	db["user2"] = "QA"
	
}

func GetDB() map[string]string{
	return db
}