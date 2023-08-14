package configs

import (
	_ "embed"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"usbkill-go/devices"
)

// can embed it too
const configFilename = "usbkill.yml"

type Config struct {
	DryRun      bool           `yaml:"dry-run"`
	Action      string         `yaml:"action"`
	Whitelisted devices.Device `yaml:"whitelisted"`

	/* todo:
	BurnFiles   []string       `yaml:"burn-files"`
	*/
	Commands map[string][]string `yaml:"commands"`

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

func (c *Config) Read(fs fs.FS) {
	if !c.readFS(fs) {
		c.readOS()
	}
}

func (c *Config) readOS() {
	if _, err := os.Stat(configFilename); os.IsNotExist(err) {
		c.save()
	}

	b, err := os.ReadFile(configFilename)
	if err != nil {
		log.Panicln(err)
	}
	c.unmarshal(b)
}

func (c *Config) readFS(f fs.FS) bool {
	content, err := fs.ReadFile(f, configFilename)
	if err != nil {
		return false
	}

	c.unmarshal(content)
	return true
}

func (c Config) save() {
	if err := os.WriteFile(configFilename, c.marshal(), 0644); err != nil {
		log.Panicln(err)
	}
}
