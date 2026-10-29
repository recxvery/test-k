package db

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBRepository interface {
	InsertThreat(context.Context, ThreatModel) (ThreatModel, error)
	GetThreat(context.Context, string) (ThreatModel, error)
	UpdateThreat(context.Context, string, string) (ThreatModel, error)
}

type ThreatModel struct {
	MD5         string     `json:"md5"`
	Name        string     `json:"name"`
	Description string     `json:"threat_description"`
	Severity    string     `json:"threat_severity"`
	ThreatType  string     `json:"threat_type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type ThreatRepositoryDB struct {
	pool *pgxpool.Pool
}

func NewRepostoryDB(pool_connection *pgxpool.Pool) *ThreatRepositoryDB {
	return &ThreatRepositoryDB{
		pool: pool_connection,
	}
}

func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {
	return pgxpool.New(ctx, os.Getenv("CONN_STRING"))
}

func CheckConnection(ctx context.Context, pool *pgxpool.Pool) {
	if err := pool.Ping(ctx); err != nil {
		log.Println(err)
		panic(err)
	}

	log.Println("Successfully connected")
}

var (
	ErrThreatAlreadyExists = errors.New("threat already exists")
	ErrThreatNotFound      = errors.New("threat not found")
)
