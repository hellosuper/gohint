package lint

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

var defaultCommonMethods = map[string]bool{
	"Error":     true,
	"Read":      true,
	"ServeHTTP": true,
	"String":    true,
	"Write":     true,
}

var defaultBadReceiverNames = map[string]bool{
	"me":   true,
	"this": true,
	"self": true,
}

type Config struct {
	Package       bool `json:"package"`
	Imports       bool `json:"imports"`
	Names         bool `json:"names"`
	Exported      bool `json:"exported"`
	VarDecls      bool `json:"var-decls"`
	Elses         bool `json:"elses"`
	MakeSlice     bool `json:"make-slice"`
	ErrorReturn   bool `json:"error-return"`
	IgnoredReturn bool `json:"ignored-return"`

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

func NewDefaultConfig() *Config {
	return &Config{
		Package:       true,
		Imports:       true,
		Names:         true,
		Exported:      true,
		VarDecls:      true,
		Elses:         true,
		MakeSlice:     true,
		ErrorReturn:   true,
		IgnoredReturn: true,

		MinConfidence:    0.8,
		Initialisms:      defaultCommonInitialisms,
		BadReceiverNames: defaultBadReceiverNames,

		//		IgnoreFiles:      []string{}, // TODO: for future use
		//		IgnorePackages:   []string{}, // TODO: for future use
		//		IgnoreTypes:      []string{}, // TODO: for future use
	}
}
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

		fmt.Printf("after: %#v\n", defaultCommonInitialisms)

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

// TODO: for future use
//func (c *Config) IsPackageIgnored(packageName string) (ok bool) {
//	_, ok = c.ignorePackagesMap[packageName]
//	return
//}
//
//func (c *Config) IsFileIgnored(fileName string) (ok bool) {
//	_, ok = c.ignoreFilesMap[fileName]
//	return
//}
//
//func (c *Config) IsTypeIgnored(typeName string) (ok bool) {
//	_, ok = c.ignoreTypesMap[typeName]
//	return
//}
