// config.go

package ipar

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	// Third-party packages:
	"gopkg.in/yaml.v2"
)

// Config defines the configuration of a Site.
type Config struct {
	Dir         string                 // Root directory for the site.
	Name        string                 // Public-facing name of the site.
	Base        string                 // Base URL for generating links.
	Owner       string                 // Public-facing name of the owner
	Insecure    bool                   // Run insecure, i.e. without TLS?
	CertFile    string                 // TLS Cert file if not Insecure.
	KeyFile     string                 // TLS Cert file if not Insecure.
	Port        int                    // Port on which to listen.
	Data        map[string]interface{} // Arbitrary data.
	Strict      bool                   // Strict mode: all errors are errors.
	SmartAssets bool                   // Intelligent asset file handling.
	Future      bool                   // Show pages published in the Future.
	Drafts      bool                   // Show pages marked as Draft.
	Reload      bool                   // Watch for file changes and reload.

}

// LoadConfig loads a Config from the given file and sets its defaults as:
//
//   	Dir: the directory of the file.
//   	Port: 8086
//   	Name: "Ipar Web Site"
//   	Owner: "Exceptionally Discerning Personage" (obviously)
//   	CertFile: "server_certificate.pem" in Dir UNLESS Insecure is true.
//   	KeyFile: "server_key.pem" in Dir UNLESS Insecure is true.
//      Data: empty map
//
func LoadConfig(file string) (*Config, error) {

	if file == "" {
		return nil, errors.New("No config file specified")
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	if cfg.Dir == "" {
		cfg.Dir = filepath.Dir(file)
	}
	if cfg.Name == "" {
		cfg.Name = "Ipar Web Site"
	}
	if cfg.Port == 0 {
		cfg.Port = 8086 // $NOSTAGLIA++
	}
	if cfg.Owner == "" {
		cfg.Owner = "Exceptionally Discerning Personage"
	}
	if cfg.Data == nil {
		cfg.Data = map[string]interface{}{}
	}
	if cfg.Insecure == false {
		if cfg.CertFile == "" {
			cfg.CertFile = filepath.Join(cfg.Dir, "server_certificate.pem")
		}
		if cfg.KeyFile == "" {
			cfg.KeyFile = filepath.Join(cfg.Dir, "server_key.pem")
		}
	}

	return cfg, nil

}

// Check tests that all required files are present and returns an error on the
// first file found to be missing or not a plain file.  The master directory
// is also checked.
func (c *Config) Check() error {

	dirInfo, err := os.Stat(c.Dir)
	if err != nil {
		return err
	}
	if !dirInfo.IsDir() {
		return errors.New("not a directory: " + c.Dir)
	}

	// We only care about TLS files if we are running securely.
	// For now we just check the file's existence, for the super-edgy edge
	// case of it being unreadable or a directory we just let it fail later.
	if c.Insecure == false {
		_, err = os.Stat(c.CertFile)
		if err != nil {
			return err
		}
		_, err = os.Stat(c.KeyFile)
		if err != nil {
			return err
		}
	}
	return nil

}
