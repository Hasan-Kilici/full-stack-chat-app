package Database

type User struct {
	Token			string
	ID				string
	Name			string
	Password		string
	Perm			string
}

type Room struct {
	Token			string
	Name			string
	OwnerID			string
}