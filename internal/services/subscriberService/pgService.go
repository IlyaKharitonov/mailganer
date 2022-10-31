package subscriberService

import (
	"mailganer/internal/entities"
	//"github.com/jackc/pgx/pgxpool"
	"github.com/jmoiron/sqlx"
)

type postgres struct {
	pool *sqlx.DB
}

func NewPg(pg *sqlx.DB)*postgres{
	return &postgres{pg}
}

func (p *postgres)GetList(IDs[]uint)([]*entities.Subscriber, error){
	return nil, nil
}

func (p *postgres)Add(subscribers []*entities.Subscriber)error{
	return nil
}



