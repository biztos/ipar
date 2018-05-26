// siteconfig_test.go

package ipar_test

import (
	// Standard:
	"os"
	"path/filepath"
	"testing"

	// Helpful:
	"github.com/stretchr/testify/assert"

	// Under test:
	"github.com/biztos/ipar"
)

func Test_LoadSiteConfig_ErrorNoFile(t *testing.T) {

	assert := assert.New(t)

	_, err := ipar.LoadSiteConfig("")
	if assert.Error(err) {
		assert.Equal("No config file specified", err.Error(),
			"error is useful")
	}

}

func Test_LoadSiteConfig_ErrorFileNotExist(t *testing.T) {

	assert := assert.New(t)

	_, err := ipar.LoadSiteConfig("nonesuch.yaml")
	if assert.Error(err) {
		assert.True(os.IsNotExist(err), "error isa IsNotExist")
		assert.Regexp("nonesuch.yaml", err.Error(),
			"error is useful")
	}

}

func Test_LoadSiteConfig_ErrorBadYAML(t *testing.T) {

	assert := assert.New(t)

	file := filepath.Join("testdata", "config-broken.yaml")
	_, err := ipar.LoadSiteConfig(file)
	if assert.Error(err) {
		assert.False(os.IsNotExist(err), "error nota IsNotExist")
		assert.Regexp("yaml", err.Error(), "error is useful")
	}

}

func Test_LoadSiteConfig_Success(t *testing.T) {

	assert := assert.New(t)

	file := filepath.Join("testdata", "config-good.yaml")
	cfg, err := ipar.LoadSiteConfig(file)
	if !assert.Nil(err, "no error") {
		t.Fatal(err)
	}
	assert.Equal("example.com", cfg.Host, "Host set")
	assert.Equal(8080, cfg.Port, "Port set")
	assert.Equal("Ipar Test Site", cfg.Name, "Name set")
	assert.Equal("Son of Kisipar", cfg.Owner, "Owner set")
	assert.Equal("/some/path/to/cert", cfg.CertFile, "CertFile set")
	assert.Equal("/some/path/to/key", cfg.KeyFile, "KeyFile set")
	assert.True(cfg.Insecure, "Insecure set")
	assert.Equal("/some/path/to/dir", cfg.Dir, "Dir set")

	// Data is an arbitrary map useful for putting, well, anything in your
	// templates.  Nothing is normalized, so anything more complicated than
	// a simple set of strings might be pretty hard to support in your
	// templates; however it should be possible, in particular for stuff
	// like rendering an arbitrary chunk of JSON.  That will be tested
	// separately.
	expData := map[string]interface{}{
		"anything": map[interface{}]interface{}{"could_be_here": true},
		"numbers":  []interface{}{1, 1.23, 4.545454545454545e+25},
	}
	assert.Equal(expData, cfg.Data, "Data set")

}

func Test_LoadSiteConfig_SuccessWithDefaults(t *testing.T) {

	assert := assert.New(t)

	file := filepath.Join("testdata", "config-defaults.yaml")
	expDir, _ := filepath.Abs("testdata")
	expCert := filepath.Join(expDir, "server_certificate.pem")
	expKey := filepath.Join(expDir, "server_key.pem")
	cfg, err := ipar.LoadSiteConfig(file)
	if !assert.Nil(err, "no error") {
		t.Fatal(err)
	}
	assert.Equal("localhost:8086", cfg.Host, "Host set")
	assert.Equal(8086, cfg.Port, "Port set")
	assert.Equal("Ipar Web Site", cfg.Name, "Name defaults")
	assert.Equal("Exceptionally Discerning Personage", cfg.Owner, "Owner defaults")
	assert.Equal(expCert, cfg.CertFile, "CertFile defaults")
	assert.Equal(expKey, cfg.KeyFile, "KeyFile defaults")
	assert.False(cfg.Insecure, "Insecure false")
	assert.Equal(expDir, cfg.Dir, "Dir set")
	assert.Nil(cfg.Data, "Data empty")

}

func Test_Check_ErrorNoDir(t *testing.T) {

	assert := assert.New(t)

	cfg := &ipar.SiteConfig{Dir: "/no/such/dir"}
	err := cfg.Check()
	if assert.Error(err) {
		assert.Equal("stat /no/such/dir: no such file or directory",
			err.Error(), "error is useful")
	}

}

func Test_Check_ErrorDirNotDir(t *testing.T) {

	assert := assert.New(t)

	cfg := &ipar.SiteConfig{Dir: "siteconfig_test.go"}
	err := cfg.Check()
	if assert.Error(err) {
		assert.Equal("not a directory: siteconfig_test.go",
			err.Error(), "error is useful")
	}

}

func Test_Check_ErrorMissingCert(t *testing.T) {

	assert := assert.New(t)

	cfg := &ipar.SiteConfig{Dir: "testdata", CertFile: "/no/such/cert"}
	err := cfg.Check()
	if assert.Error(err) {
		assert.Equal("stat /no/such/cert: no such file or directory",
			err.Error(), "error is useful")
	}

}

func Test_Check_ErrorMissingKey(t *testing.T) {

	assert := assert.New(t)

	cfg := &ipar.SiteConfig{
		Dir:      "testdata",
		CertFile: "siteconfig_test.go",
		KeyFile:  "/no/such/key",
	}
	err := cfg.Check()
	if assert.Error(err) {
		assert.Equal("stat /no/such/key: no such file or directory",
			err.Error(), "error is useful")
	}

}

func Test_Check_SuccessSecure(t *testing.T) {

	assert := assert.New(t)

	cfg := &ipar.SiteConfig{
		Dir:      "testdata",
		CertFile: "siteconfig_test.go",
		KeyFile:  "siteconfig_test.go",
	}
	assert.Nil(cfg.Check(), "no error on Check")

}

func Test_Check_SuccessInsecure(t *testing.T) {

	assert := assert.New(t)

	cfg := &ipar.SiteConfig{
		Dir:      "testdata",
		Insecure: true,
		CertFile: "/no/such/cert",
		KeyFile:  "/no/such/key",
	}
	assert.Nil(cfg.Check(), "no error on Check")

}
