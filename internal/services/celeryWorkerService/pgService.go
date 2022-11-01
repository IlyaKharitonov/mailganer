package celeryWorkerService

import (
	"context"
	"fmt"

	//"github.com/jackc/pgx/pgxpool"
	"github.com/jmoiron/sqlx"
)

type postgres struct {
	//*pgxpool.Pool
	*sqlx.DB
}

func NewPg(pg *sqlx.DB)*postgres{
	return &postgres{pg}
}

func(p *postgres)GetTemplate(id uint) (string, error){
	var (
		query = `SELECT template FROM template WHERE id = $1`
		blankTemplate string
	)

	row := p.QueryRow(query, id)
	if err := row.Scan(&blankTemplate); err != nil{
		return "", fmt.Errorf("(p *postgres)GetTemplate #1, Error: %s", err.Error())
	}

	return blankTemplate, nil
}

func (p *postgres)AddTemplate(ctx context.Context, template string) error{

	return nil
}

