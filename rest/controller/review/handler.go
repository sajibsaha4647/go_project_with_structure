package review

import "github.com/jmoiron/sqlx"

type Handler struct{
	
}

func NewHandler(dbConn *sqlx.DB) *Handler {
	return &Handler{}
}