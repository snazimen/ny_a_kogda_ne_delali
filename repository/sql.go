package repository

const (
	SQLCreateMessage = `
	
	CREATE TABLE USERS  (
	ID INTEGER PRIMARY KEY,
	FIO VARCHAR(128) NOT NULL,
	EMAIL VARCHAR(128) NOT NULL,
	STATUS VARCAR(128) NOT NULL,

	)`
)
