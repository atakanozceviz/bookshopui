package model

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	KAdi   string
	KSifre string
}

func openDb() *sql.DB {
	db, err := sql.Open("sqlite3", "model/db.db")
	if err != nil {
		log.Println(err)
	}
	return db
}

func (User) CheckUser(u User) int {
	db := openDb()
	// query
	rows, err := db.Query("SELECT id, kSifre FROM users WHERE kAdi = ?", u.KAdi)
	if err != nil {
		log.Println(err)
	}

	var id int
	var hp string
	for rows.Next() {
		err = rows.Scan(&id, &hp)
		if err != nil {
			log.Println(err)
		}
	}
	rows.Close() //good habit to close
	err = decrypt(hp, u.KSifre)

	if err != nil {
		log.Println(err)
		return 0
	} else {
		return id
	}
}

func (User) InsertUser(u User) (int64, error) {
	db := openDb()
	stmt, err := db.Prepare("INSERT INTO users(kAdi, kSifre) VALUES (?,?); ")
	if err != nil {
		log.Println(err)
	}

	hp, err := encrypt(u.KSifre)
	if err != nil {
		log.Println(err)
	}
	res, err := stmt.Exec(u.KAdi, hp)
	if err != nil {
		log.Println(err)
	}

	return res.LastInsertId()
}

func encrypt(password string) ([]byte, error) {
	p := []byte(password)
	return bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
}

func decrypt(hashedPassword, password string) error {
	hp := []byte(hashedPassword)
	p := []byte(password)
	return bcrypt.CompareHashAndPassword(hp, p)
}
