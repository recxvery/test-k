package http

import (
	"kaspersky-testovoe/db"
	"regexp"
	"strings"
	"time"
)

type ThreatCreateRequest struct {
	MD5         string `json:"md5"`
	Name        string `json:"name"`
	Description string `json:"threat_description"`
	Severity    string `json:"threat_severity"`
	ThreatType  string `json:"threat_type"`
}

var md5Regex = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

func ValidateToCreate(threatDTO ThreatCreateRequest) (db.ThreatModel, error) {
	threatDTO.MD5 = strings.TrimSpace(threatDTO.MD5)
	threatDTO.Name = strings.TrimSpace(threatDTO.Name)
	threatDTO.Description = strings.TrimSpace(threatDTO.Description)
	threatDTO.Severity = strings.TrimSpace(threatDTO.Severity)
	threatDTO.ThreatType = strings.TrimSpace(threatDTO.ThreatType)

	switch {
	case len(threatDTO.MD5) != 32:
		return db.ThreatModel{}, incorrectMD5Length

	case !md5Regex.MatchString(threatDTO.MD5):
		return db.ThreatModel{}, incorrectMD5Format

	case threatDTO.Name == "":
		return db.ThreatModel{}, incorrectName

	case threatDTO.Description == "":
		return db.ThreatModel{}, incorrectDescription

	case !isValidSeverity(threatDTO.Severity):
		return db.ThreatModel{}, incorrectSeverity

	case threatDTO.ThreatType == "":
		return db.ThreatModel{}, incorrectThreatType
	}

	return db.ThreatModel{
		MD5:         threatDTO.MD5,
		Name:        threatDTO.Name,
		Description: threatDTO.Description,
		Severity:    threatDTO.Severity,
		ThreatType:  threatDTO.ThreatType,
		CreatedAt:   time.Now(),
	}, nil
}

func isValidSeverity(severity string) bool {
	switch severity {
	case "low", "medium", "high", "critical":
		return true
	default:
		return false
	}
}

type ThreatUpdateRequest struct {
	Name string `json:"name"`
	Description string `json:"threat_description"`
}

func ValidteThreatUpdate(threatDTO ThreatUpdateRequest) (string, error) {
	if threatDTO.Description == "" {
		return "", incorrectDescription
	}

	return threatDTO.Description, nil
}
