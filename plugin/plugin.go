package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/lioueq/loglinter/analyzer"
	"golang.org/x/tools/go/analysis"
)

type Settings struct {
	SensitivePatterns []string `json:"sensitive_patterns"`
}

type LinterPlugin struct {
	settings Settings
}

func New(conf any) (register.LinterPlugin, error) {
	settings, err := register.DecodeSettings[Settings](conf)
	if err != nil {
		return nil, err
	}

	return &LinterPlugin{settings: settings}, nil
}

func (p *LinterPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	analyzer.SetSensitivePatterns(p.settings.SensitivePatterns)
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (*LinterPlugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func init() {
	register.Plugin("loglinter", New)
}
