// siteconfig.go

package ipar

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	// Third-party packages:
	"gopkg.in/yaml.v2"
)

// SiteConfig defines the configuration of a Site.
type SiteConfig struct {
	Name     string                 // Public-facing name of the site.
	Owner    string                 // Public-facing name of the owner
	Insecure bool                   // Run insecure, i.e. without TLS?
	CertFile string                 // TLS Cert file if not default.
	KeyFile  string                 // TLS Cert file if not default.
	Host     string                 // Only serve requests for this host.
	Port     int                    // Port on which to listen.
	Dir      string                 // Root directory for the site.
	Data     map[string]interface{} // Arbitrary data.
}

// LoadSiteConfig loads a single Config from a YAML file.  The first error
// encountered is returned.  Note that the keys in the YAML file should be
// in lowercase.
//
// Default values:
//   	Dir: the absolute path of the provided file's directory.
//   	Host: "localhost:" + Port
//   	Port: 8086
//   	Name: "Ipar Web Site"
//   	Insecure: false
//   	Owner: "Exceptionally Discerning Personage" (obviously)
//   	CertFile: "server_certificate.pem" in the Dir.
//   	KeyFile: "server_key.pem" in the Dir.
//
// Note that files are not checked here; cf. the Check function.
func LoadSiteConfig(file string) (*SiteConfig, error) {

	if file == "" {
		return nil, errors.New("No config file specified")
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	// NOTE: under what condition would Abs return an error?
	// If we ever figure that out we can test for it.
	dir, _ := filepath.Abs(filepath.Dir(file))
	cfg := &SiteConfig{
		Name:  "Ipar Web Site",
		Owner: "Exceptionally Discerning Personage",
		Dir:   dir,
		Port:  8086, // $NOSTAGLIA++
	}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		return nil, err
	}

	// Conditional defaults:
	// (And maybe someday for content/static/templates as well, so we can
	// use the same static/templates across multiple content sites.  But
	// not yet.  Focus Mr Frost!)
	if cfg.Host == "" {
		cfg.Host = fmt.Sprintf("localhost:%d", cfg.Port)
	}
	if cfg.CertFile == "" {
		cfg.CertFile = filepath.Join(cfg.Dir, "server_certificate.pem")
	}
	if cfg.KeyFile == "" {
		cfg.KeyFile = filepath.Join(cfg.Dir, "server_key.pem")
	}

	return cfg, nil
}

// Check tests that all required files are present and returns an error on the
// first file found to be missing or not a plain file.  The master directory
// is also checked.
func (cfg *SiteConfig) Check() error {

	dirInfo, err := os.Stat(cfg.Dir)
	if err != nil {
		return err
	}
	if !dirInfo.IsDir() {
		return errors.New("not a directory: " + cfg.Dir)
	}

	// We only care about TLS files if we are running securely.
	// For now we just check the file's existence, for the super-edgy edge
	// case of it being unreadable or a directory we just let it fail later.
	if cfg.Insecure == false {
		_, err = os.Stat(cfg.CertFile)
		if err != nil {
			return err
		}
		_, err = os.Stat(cfg.KeyFile)
		if err != nil {
			return err
		}
	}
	return nil

}
