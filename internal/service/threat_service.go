package service

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"kaspersky-testovoe/internal/core"
	"kaspersky-testovoe/internal/db"
)

var md5Regex = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

type ThreatRepository interface {
	GetThreat(ctx context.Context, md5 string) (db.ThreatModel, error)
	InsertThreat(ctx context.Context, threat db.ThreatModel) (db.ThreatModel, error)
	UpdateThreat(ctx context.Context, md5 string, name string) (db.ThreatModel, error)
}

type ThreatService struct {
	repo ThreatRepository
}

func NewThreatService(repo ThreatRepository) *ThreatService {
	return &ThreatService{
		repo: repo,
	}
}

func (s *ThreatService) GetThreat(
	ctx context.Context,
	md5 string,
) (db.ThreatModel, error) {
	md5 = strings.TrimSpace(md5)

	if len(md5) != 32 {
		return db.ThreatModel{}, core.IncorrectMD5Length
	}

	if !md5Regex.MatchString(md5) {
		return db.ThreatModel{}, core.IncorrectMD5Format
	}

	return s.repo.GetThreat(ctx, md5)
}

func (s *ThreatService) CreateThreat(
	ctx context.Context,
	threat db.ThreatModel,
) (db.ThreatModel, error) {
	threat.MD5 = strings.TrimSpace(threat.MD5)
	threat.Name = strings.TrimSpace(threat.Name)
	threat.Description = strings.TrimSpace(threat.Description)
	threat.Severity = strings.TrimSpace(threat.Severity)
	threat.ThreatType = strings.TrimSpace(threat.ThreatType)

	switch {
	case len(threat.MD5) != 32:
		return db.ThreatModel{}, core.IncorrectMD5Length

	case !md5Regex.MatchString(threat.MD5):
		return db.ThreatModel{}, core.IncorrectMD5Format

	case threat.Name == "":
		return db.ThreatModel{}, core.IncorrectName

	case threat.Description == "":
		return db.ThreatModel{}, core.IncorrectDescription

	case !isValidSeverity(threat.Severity):
		return db.ThreatModel{}, core.IncorrectSeverity

	case threat.ThreatType == "":
		return db.ThreatModel{}, core.IncorrectThreatType
	}

	threat.CreatedAt = time.Now()

	return s.repo.InsertThreat(ctx, threat)
}

func (s *ThreatService) UpdateThreat(
	ctx context.Context,
	md5 string,
	name string,
) (db.ThreatModel, error) {
	md5 = strings.TrimSpace(md5)
	name = strings.TrimSpace(name)

	if len(md5) != 32 {
		return db.ThreatModel{}, core.IncorrectMD5Length
	}

	if !md5Regex.MatchString(md5) {
		return db.ThreatModel{}, core.IncorrectMD5Format
	}

	if name == "" {
		return db.ThreatModel{}, core.IncorrectName
	}

	return s.repo.UpdateThreat(ctx, md5, name)
}

func isValidSeverity(severity string) bool {
	switch severity {
	case "low", "medium", "high", "critical":
		return true
	default:
		return false
	}
}

var ErrInvalidOperation = errors.New("invalid operation")
