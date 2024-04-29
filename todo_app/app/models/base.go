package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
	tableNameSession = "sessions"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE, 
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME,
		FOREIGN KEY(user_id) REFERENCES %s(id))`, tableNameTodo, tableNameUser)
	Db.Exec(cmdT)

	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME,
		FOREIGN KEY(user_id) REFERENCES %s(id))`, tableNameSession, tableNameUser)
	Db.Exec(cmdS)
}

func createUUID() (uuidobj uuid.UUID, ) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plain string) (crypt string) {
	crypt = fmt.Sprintf("%x", sha1.Sum([]byte(plain)))
	return crypt
}

