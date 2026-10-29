package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func (r *ThreatRepositoryDB) InsertThreat(ctx context.Context, threat ThreatModel) (ThreatModel, error) {
	exists, err := r.RecordExists(ctx, threat.MD5)

	if err != nil {
		return ThreatModel{}, err
	}

	if exists {
		return ThreatModel{}, ErrThreatAlreadyExists
	}

	sqlQuery := `
		INSERT INTO threats (md5, threat_name, threat_description, severity, threat_type, created_at)
		VALUES($1, $2, $3, $4, $5, $6)
		RETURNING md5, threat_name, threat_description, severity, threat_type, created_at, updated_at;
	`

	var outputThreat ThreatModel

	err = r.pool.QueryRow(
		ctx,
		sqlQuery,
		threat.MD5,
		threat.Name,
		threat.Description,
		threat.Severity,
		threat.ThreatType,
		threat.CreatedAt,
	).Scan(
		&outputThreat.MD5,
		&outputThreat.Name,
		&outputThreat.Description,
		&outputThreat.Severity,
		&outputThreat.ThreatType,
		&outputThreat.CreatedAt,
		&outputThreat.UpdatedAt,
	)

	if err != nil {
		log.Println(err)
		return ThreatModel{}, err
	}

	return outputThreat, nil
}

func (r *ThreatRepositoryDB) GetThreat(ctx context.Context, md5 string) (ThreatModel, error) {

	sqlQuery := `
		SELECT md5, threat_name, threat_description, severity, threat_type, created_at, updated_at FROM threats 
		WHERE md5 = $1;
	`

	var outputThreat ThreatModel

	err := r.pool.QueryRow(ctx, sqlQuery, md5).Scan(
		&outputThreat.MD5,
		&outputThreat.Name,
		&outputThreat.Description,
		&outputThreat.Severity,
		&outputThreat.ThreatType,
		&outputThreat.CreatedAt,
		&outputThreat.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return ThreatModel{}, ErrThreatNotFound
	}

	if err != nil {
		log.Println(err)
		return ThreatModel{}, fmt.Errorf("get request err: %w", err)
	}

	return outputThreat, nil
}

func (r *ThreatRepositoryDB) UpdateThreat(ctx context.Context, md5 string, description string) (ThreatModel, error) {
	sqlQuery := `
		UPDATE threats
		SET threat_description=$1, updated_at=$2
		WHERE md5=$3
		RETURNING md5, threat_name, threat_description, severity, threat_type, created_at, updated_at; 
	`

	var outputThreat ThreatModel

	err := r.pool.QueryRow(ctx, sqlQuery, description, time.Now(), md5).Scan(
		&outputThreat.MD5,
		&outputThreat.Name,
		&outputThreat.Description,
		&outputThreat.Severity,
		&outputThreat.ThreatType,
		&outputThreat.CreatedAt,
		&outputThreat.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return ThreatModel{}, ErrThreatNotFound
	}

	if err != nil {
		log.Println(err)
		return ThreatModel{}, fmt.Errorf("update err req: %w", err)
	}

	return outputThreat, nil
}

func (r *ThreatRepositoryDB) RecordExists(ctx context.Context, md5 string) (bool, error) {
	var exists bool

	sqlQuery := `SELECT EXISTS(SELECT 1 FROM threats WHERE md5=$1)`

	err := r.pool.QueryRow(ctx, sqlQuery, md5).Scan(
		&exists,
	)

	if err != nil {
		return false, fmt.Errorf("record threat existence err: %w", err)
	}

	return exists, nil
}
