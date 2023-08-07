package Database

import(
	"chat/Utils"
	b64 "encoding/base64"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"errors"
)

func Register(Username, Password string) bool {	
	if !FindUserByName(Username){
		Token := Utils.GenerateToken(31)
		ID := Utils.GenerateToken(26)
		NewPassword := b64.URLEncoding.EncodeToString([]byte(Password))
		query := "INSERT INTO users (Token, ID, Name, Password, Perm) VALUES ($1,$2,$3,$4,$5)"
		_,err := db.Exec(query,Token,ID,Username,NewPassword,"User")
		if err != nil {
			fmt.Println(err)
			return false
		}
		
		return true
	} else {
		fmt.Println("Bu Kullanıcı zaten Kullanılıyor")
		return false
	}
}

func Login(Name, Password string) (string, error) {
	query := "SELECT token, password FROM users WHERE Name = $1"
	row := db.QueryRow(query, Name)

	var token, hashedPassword string
	row.Scan(&token, &hashedPassword)
	RealPassword, _ := b64.URLEncoding.DecodeString(hashedPassword)

	if(string(RealPassword) == Password){
		return token, nil
	}

	return "Kullanıcı Bulunamadı", errors.New("Kullanıcı adı ya da şifre yanlış")
}

func FindUserByToken(Token string) (User, error){
    query := "SELECT * FROM users WHERE token = $1"
    row := db.QueryRow(query, Token)

    var user User
    err := row.Scan(&user.Token,&user.ID,&user.Name,&user.Password,&user.Perm)
    if err != nil {
        if err == sql.ErrNoRows {
            return User{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return User{}, err
    }
    return user, nil
}

func FindUserByName(Name string) bool {
    query := "SELECT COUNT(*) FROM users WHERE Email = $1"
    row := db.QueryRow(query, Name)

    var count int
    err := row.Scan(&count)
    if err != nil {
        return false
    }

    if count == 0 {
        return false
    }

    return true
}