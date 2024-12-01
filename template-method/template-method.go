package main

import (
	"fmt"
	"strings"
)

type DBOperation interface {
	connect()
	query()
	insert()
	update()
	delete()
}

func DailyOperation(db DBOperation) {
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("Daily Operation Start")
	db.connect()
	db.query()
	db.insert()
	db.update()
	db.delete()
	fmt.Println("Daily Operation End")
	fmt.Println(strings.Repeat("=", 50))
}

type DB struct {
	DBOperation
}

func (db *DB) connect() {
	fmt.Println("DB connect")
}

func (db *DB) query() {
	fmt.Println("DB query")
}

func (db *DB) insert() {
	fmt.Println("DB insert")
}

func (db *DB) update() {
	fmt.Println("DB update")
}

func (db *DB) delete() {
	fmt.Println("DB delete")
}

type MongoDB struct {
	DB
}

func (mdb *MongoDB) connect() {
	fmt.Println("MongoDB connect")
}

func (mdb *MongoDB) update() {
	fmt.Println("MongoDB update")
}

type PostgresDB struct {
	DB
}

func (pdb *PostgresDB) connect() {
	fmt.Println("PostgresDB connect")
}

func (pdb *PostgresDB) query() {
	fmt.Println("PostgresDB query")
}
