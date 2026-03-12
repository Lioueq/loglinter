package analyzer

import (
	"regexp"
	"strings"
)

var sensitivePatterns []string

type Config struct {
	SensitivePatterns []string `json:"sensitive_patterns"`
}

func SetSensitivePatterns(patterns []string) {
	sensitivePatterns = append([]string(nil), patterns...)
}

func buildSensitivePattern() (*regexp.Regexp, error) {
	patterns := make([]string, 0, len(sensitivePatterns))
	for _, pattern := range sensitivePatterns {
		if strings.TrimSpace(pattern) == "" {
			continue
		}
		patterns = append(patterns, pattern)
	}
	if len(patterns) == 0 {
		return nil, nil
	}

	return regexp.Compile(`(?i)(` + strings.Join(patterns, "|") + `)`)
}
