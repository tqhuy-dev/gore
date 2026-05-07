package templates

import (
	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Repo struct {
	conn        *pgx.Conn
	mongoClient *mongo.Client
	nameTable   string
}

func NewRepo(
	conn *pgx.Conn,
	mongoClient *mongo.Client) *Repo {
	return &Repo{
		conn:        conn,
		mongoClient: mongoClient,
	}
}
