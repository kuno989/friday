package pkg

import (
	"errors"
	"fmt"
	"github.com/google/wire"
	"github.com/hillu/go-yara/v4"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var (
	DefaultYaraConfig = YaraConfig{
		RulePath: "",
	}
	YaraProviderSet = wire.NewSet(NewYara, ProvideYaraConfig)
)

type Rule struct {
	Namespace string
	Filename  string
}

type YaraConfig struct {
	RulePath string `mapstructure:"rule_path"`
}

func ProvideYaraConfig(cfg *viper.Viper) (YaraConfig, error) {
	yara := DefaultYaraConfig
	err := cfg.UnmarshalKey("yara", &yara)
	return yara, err
}

type Yara struct {
	Config YaraConfig
	Load   *yara.Rules
}

func NewYara(cfg YaraConfig) (*Yara, error) {
	yaraRuleFolders, err := ioutil.ReadDir(cfg.RulePath)
	c, err := yara.NewCompiler()
	if err != nil {
		msg := fmt.Sprintf("yara rule path not found %s", err)
		return nil, errors.New(msg)
	}
	for _, yaraFolder := range yaraRuleFolders {
		if yaraFolder.IsDir() {
			ruleFolderPath := fmt.Sprintf("%s/%s", cfg.RulePath, yaraFolder.Name())
			rules, _ := ioutil.ReadDir(ruleFolderPath)
			for _, rule := range rules {
				if !rule.IsDir() && rule.Name() != ".DS_Store" {
					ruleFile := fmt.Sprintf("%s/%s/%s", cfg.RulePath, yaraFolder.Name(), rule.Name())
					nameSpace := strings.Split(ruleFile, ".")[0]
					f, err := os.Open(ruleFile)
					if err != nil {
						msg := fmt.Sprintf("couldn't open rule file %s: %s", rule.Name(), err)
						return nil, errors.New(msg)
					}
					err = c.AddFile(f, nameSpace)
					f.Close()
					if err != nil {
						msg := fmt.Sprintf("could not parse rule file %s: %s", rule.Name(), err)
						return nil, errors.New(msg)
					}
				}
			}
		}
	}
	r, err := c.GetRules()
	if err != nil {
		msg := fmt.Sprintf("failed to compile rules: %s", err)
		return nil, errors.New(msg)
	}
	return &Yara{
		Config: cfg,
		Load:   r,
	}, nil
}
