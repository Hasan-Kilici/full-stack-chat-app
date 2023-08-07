package Database

import(
	"chat/Utils"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func CreateRoom(Name,OwnerID string) (string, error) {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO rooms (Token,Name,OwnerID) VALUES ($1, $2, $3)"
	_,err := db.Exec(query,Token,Name,OwnerID)
	if err != nil {
		return "oda oluşturulamadı",err
	}

	return Token, nil
}

func DeleteRoom(Token string) error {
	query := "DELETE FROM rooms WHERE = $1"
	_,err := db.Exec(query,Token)
	if err != nil {
		return err
	}

	return nil
}

func FindRoom(Token string) (Room, error) {
	query := "SELECT * FROM rooms WHERE token = $1"
    row := db.QueryRow(query, Token)

    var room Room
    err := row.Scan(&room.Token,&room.Name,&room.OwnerID)
    if err != nil {
        if err == sql.ErrNoRows {
            return Room{}, fmt.Errorf("kullanıcı bulunamadı")
        }
        return Room{}, err
    }
    return room, nil
}

func ListRoom() ([]Room, error) {
	query := "SELECT * FROM rooms"
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    rooms := []Room{}
    for rows.Next() {
        var room Room
        err := rows.Scan(&room.Token,&room.Name,&room.OwnerID)
        if err != nil {
            return nil, err
        }
        rooms = append(rooms, room)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return rooms, nil
}