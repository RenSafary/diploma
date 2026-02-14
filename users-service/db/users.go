package db

import (
	"database/sql"
	"diploma/auth-service/utils"
	userspb "diploma/proto/users"
	"log"
	"strconv"
)

type Users struct {
	DB *sql.DB
}

func UsersInit(db *sql.DB) *Users {
	return &Users{DB: db}
}

type User struct {
	Id        int
	Username  string
	Password  string
	Email     string
	FirstName string
	LastName  string
	Sex       string
	Admin     bool
}

func (c *Users) CreateUser(username, password, firstname, lastname, email, sex, age string) (int32, error) {
	age_int, err := strconv.Atoi(age)
	if err != nil {
		return 0, err
	}

	// making hashed password
	hashedPass, err := utils.MakeHashed(password)
	if err != nil {
		return 0, err
	}

	var userId int32
	err = c.DB.QueryRow(
		`INSERT INTO users (Username, Passwd, FirstName, LastName, Email, Sex, Age, Adm)
         VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id`,
		username, hashedPass, firstname, lastname, email, sex, age_int, false,
	).Scan(&userId)
	if err != nil {
		log.Println("Error inserting user:", err)
		return 0, err
	}

	return userId, nil
}

func (r *Users) GetAllUsers() ([]*userspb.User, error) {
	rows, err := r.DB.Query("SELECT id, email, firstname, lastname, age, sex, adm FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*userspb.User, 0)

	for rows.Next() {
		var id int
		var email, firstname, lastname, sexStr string
		var age int32
		var admBool bool

		err := rows.Scan(&id, &email, &firstname, &lastname, &age, &sexStr, &admBool)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}

		var sex userspb.Sex
		switch sexStr {
		case "MALE", "М":
			sex = userspb.Sex_MALE
		case "FEMALE", "Ж":
			sex = userspb.Sex_FEMALE
		default:
			sex = userspb.Sex_UNKNOWN
		}

		u := &userspb.User{
			Id:        int32(id),
			Email:     email,
			Firstname: firstname,
			Lastname:  lastname,
			Age:       age,
			Sex:       sex,
			Adm:       admBool,
		}

		users = append(users, u)
	}

	log.Println("Found users:", len(users))
	return users, nil
}
