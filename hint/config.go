package hint

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var defaultCommonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}

var defaultBadReceiverNames = map[string]bool{
	"me":   true,
	"this": true,
	"self": true,
}

// Config defines configuration options for linter
type Config struct {
	Package            bool `json:"package"`
	Imports            bool `json:"imports"`
	Names              bool `json:"names"`
	Exported           bool `json:"exported"`
	VarDecls           bool `json:"var-decls"`
	Elses              bool `json:"elses"`
	MakeSlice          bool `json:"make-slice"`
	ErrorReturn        bool `json:"error-return"`
	IgnoredReturn      bool `json:"ignored-return"`
	PackageUnderscore  bool `json:"package-underscore"`
	NamedReturn        bool `json:"named-return"`
	PackagePrefixNames bool `json:"package-prefix-names"`
	Ranges             bool `json:"ranges"`
	ReceiverNames      bool `json:"receiver-names"`
	Errorf             bool `json:"errorf"`
	Errors             bool `json:"errors"`
	ErrorStrings       bool `json:"error-strings"`
	IncDec             bool `json:"inc-dec"`

	MinConfidence float64 `json:"min-confidence"`

	IgnoreFiles    []string `json:"ignore-files"`
	ignoreFilesMap map[string]bool

	IgnorePackages    []string `json:"ignore-packages"`
	ignorePackagesMap map[string]bool

	IgnoreTypes    []string `json:"ignore-types"`
	ignoreTypesMap map[string]bool

	Initialisms      map[string]bool `json:"initialisms"`
	BadReceiverNames map[string]bool `json:"bad-receivers"`
}

// NewDefaultConfig creates linter config with predefined options
func NewDefaultConfig() *Config {
	return &Config{
		Package:            false,
		Imports:            false,
		Names:              false,
		Exported:           false,
		VarDecls:           false,
		Elses:              false,
		MakeSlice:          false,
		ErrorReturn:        false,
		IgnoredReturn:      false,
		PackageUnderscore:  false,
		NamedReturn:        false,
		PackagePrefixNames: false,
		Ranges:             false,
		ReceiverNames:      false,
		Errorf:             false,
		Errors:             false,
		ErrorStrings:       false,
		IncDec:             false,

		MinConfidence:    0.8,
		Initialisms:      defaultCommonInitialisms,
		BadReceiverNames: defaultBadReceiverNames,
	}
}

// NewConfig reads config from given file. If filename is empty, default config will be returned
func NewConfig(file string) (*Config, error) {
	c := NewDefaultConfig()

	if file != "" {
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("could not read config file %s: %s", file, err.Error())

		}

		if err := json.Unmarshal(contents, c); err != nil {
			return nil, fmt.Errorf("could not parse configuration from %s: %s", file, err.Error())
		}

		sliceToMapBool := func(slice []string) map[string]bool {
			res := map[string]bool{}
			for _, v := range slice {
				res[v] = true
			}

			return res
		}
		c.ignoreFilesMap = sliceToMapBool(c.IgnoreFiles)
		c.ignorePackagesMap = sliceToMapBool(c.IgnorePackages)
		c.ignoreTypesMap = sliceToMapBool(c.IgnoreTypes)
	}

	return c, nil
}
