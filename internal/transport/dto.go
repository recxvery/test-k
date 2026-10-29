package transport

import "kaspersky-testovoe/internal/db"

type ThreatCreateRequest struct {
	MD5         string `json:"md5"`
	Name        string `json:"name"`
	Description string `json:"threat_description"`
	Severity    string `json:"threat_severity"`
	ThreatType  string `json:"threat_type"`
}

type ThreatUpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"threat_description"`
}

func ToThreatModel(dto ThreatCreateRequest) db.ThreatModel {
	return db.ThreatModel{
		MD5:         dto.MD5,
		Name:        dto.Name,
		Description: dto.Description,
		Severity:    dto.Severity,
		ThreatType:  dto.ThreatType,
	}
}

// var md5Regex = regexp.MustCompile(`^[a-fA-F0-9]{32}$`)

// func ValidateToCreate(threatDTO ThreatCreateRequest) (db.ThreatModel, error) {
// 	threatDTO.MD5 = strings.TrimSpace(threatDTO.MD5)
// 	threatDTO.Name = strings.TrimSpace(threatDTO.Name)
// 	threatDTO.Description = strings.TrimSpace(threatDTO.Description)
// 	threatDTO.Severity = strings.TrimSpace(threatDTO.Severity)
// 	threatDTO.ThreatType = strings.TrimSpace(threatDTO.ThreatType)

// 	switch {
// 	case len(threatDTO.MD5) != 32:
// 		return db.ThreatModel{}, core.IncorrectMD5Length

// 	case !md5Regex.MatchString(threatDTO.MD5):
// 		return db.ThreatModel{}, core.IncorrectMD5Format

// 	case threatDTO.Name == "":
// 		return db.ThreatModel{}, core.IncorrectName

// 	case threatDTO.Description == "":
// 		return db.ThreatModel{}, core.IncorrectDescription

// 	case !isValidSeverity(threatDTO.Severity):
// 		return db.ThreatModel{}, core.IncorrectSeverity

// 	case threatDTO.ThreatType == "":
// 		return db.ThreatModel{}, core.IncorrectThreatType
// 	}

// 	return db.ThreatModel{
// 		MD5:         threatDTO.MD5,
// 		Name:        threatDTO.Name,
// 		Description: threatDTO.Description,
// 		Severity:    threatDTO.Severity,
// 		ThreatType:  threatDTO.ThreatType,
// 		CreatedAt:   time.Now(),
// 	}, nil
// }

// func isValidSeverity(severity string) bool {
// 	switch severity {
// 	case "low", "medium", "high", "critical":
// 		return true
// 	default:
// 		return false
// 	}
// }

// func ValidteThreatUpdate(threatDTO ThreatUpdateRequest) (string, error) {
// 	if threatDTO.Description == "" {
// 		return "", core.IncorrectDescription
// 	}

// 	return threatDTO.Description, nil
// }
