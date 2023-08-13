package configs

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"usbkill-go/devices"
)

// can embed it too
const configFileName = "usbkill.yml"

type Config struct {
	DryRun      bool           `yaml:"dry-run"`
	Action      string         `yaml:"action"`
	Whitelisted devices.Device `yaml:"whitelisted"`

	/* todo:
	BurnFiles   []string       `yaml:"burn-files"`
	Commands [][]string `yaml:"commands"`
	*/

	KillSwitch struct {
		Enabled   bool   `yaml:"enabled"`
		ProductId string `yaml:"product-id"`
		VendorId  string `yaml:"vendor-id"`
	} `yaml:"kill-switch"`
}

func (c Config) HasKillSwitch() bool {
	return c.KillSwitch.Enabled &&
		len(c.KillSwitch.VendorId) > 0 &&
		len(c.KillSwitch.ProductId) > 0
}
func (c Config) marshal() []byte {
	d, err := yaml.Marshal(c)
	if err != nil {
		log.Panicln(err)
	}
	return d
}

func (c *Config) unmarshal(b []byte) {
	err := yaml.Unmarshal(b, c)
	if err != nil {
		log.Panicln(err)
	}
}

func (c *Config) Read() {
	if _, err := os.Stat(configFileName); os.IsNotExist(err) {
		c.save()
	}

	b, err := os.ReadFile(configFileName)
	if err != nil {
		log.Panicln(err)
	}
	c.unmarshal(b)
}

func (c Config) save() {
	if err := os.WriteFile(configFileName, c.marshal(), 0644); err != nil {
		log.Panicln(err)
	}
}
