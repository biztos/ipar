// site.go

package ipar

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Site is a website in the ipar application.
type Site struct {
	Config  *Config
	pathMap map[string]interface{}
}

// NewSite instantiates a Site with the provided configuration.  If the
// Config fails its Check() then an error is returned.
func NewSite(cfg *Config) (*Site, error) {

	if err := cfg.Check(); err != nil {
		return nil, err
	}
	site := &Site{
		Config:  cfg,
		pathMap: map[string]interface{}{},
	}
	return site, nil
}

// NewSiteFromDir instantiates a Site at the provided directory.  There must
// be a configuration file named "config.yaml" in the directory.
func NewSiteFromDir(dir string) (*Site, error) {

	return NewSiteFromFile(filepath.Join(dir, "config.yaml"))

}

// NewSiteFromFile instantiates a Site from the given config file.
func NewSiteFromFile(file string) (*Site, error) {

	cfg, err := LoadConfig(file)
	if err != nil {
		return nil, err
	}
	return NewSite(cfg)

}

// Init prepares a Site for serving: assets are catalogued, pages rendered,
// and templates compiled. If the site's configuration specifies Strict,
// recoverable error conditions will return errors; otherwise they will result
// in warnings to standard error.
func (site *Site) Init() error {

	// We keep track of all known paths so we can do superfast 404's:
	// (Obviously this *might* change someday to support huge sites and/or
	// dynamic content, but by then we're not serving a set of regular pages
	// anyway and have either gone back towards the full-on original flexible
	// scope-creepy server project, or have moved on entirely.)

	// Take note of everything in static as the start of our URL mapping.
	if err := site.prepStatic(); err != nil {
		return err
	}
	if err := site.prepContent(); err != nil {
		return err
	}

	// Load and compile all templates (this is easy-ish).
	return errors.New("NOT YET")
}

// Load all the static stuff into the path map.
func (site *Site) prepStatic() error {

	dir := filepath.Join(site.Config.Dir, "static")

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	walker := func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		// It might be a link to a dir, or something missing...
		realInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if realInfo.IsDir() {
			return nil
		}

		// Looks good, so let's keep it as a path, respecting the "index.html"
		// behavior of net/http.
		reqPath := filepath.ToSlash(strings.TrimPrefix(path, dir))
		reqPath = strings.TrimSuffix(reqPath, "/index.html")
		site.pathMap[reqPath] = path

		return nil

	}
	err = filepath.Walk(dir, walker)
	if err != nil {
		return fmt.Errorf("Error walking %s: %v", dir, err)
	}
	return nil

}

// Load all the content stuff into the path map (pages and files).
func (site *Site) prepContent() error {

	dir := filepath.Join(site.Config.Dir, "content")

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	walker := func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		// It might be a link to a dir, or something missing...
		realInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if realInfo.IsDir() {
			return nil
		}

		// First, determine its path.
		reqPath := filepath.ToSlash(strings.TrimPrefix(path, dir))
		reqPath = strings.TrimSuffix(reqPath, "/index.html")
		reqPath = strings.TrimSuffix(reqPath, "/index.md")

		// Duplicates should ideally NOT occur, but there's no way to control
		// for that since explicit indexes can conflict with implicit ones,
		// e.g. /foo/bar.md vs /foo/bar/index.md.  If we ever support multiple
		// content directories per site this will have to get a bit more
		// nuanced.
		if site.pathMap[reqPath] != nil {
			msg := fmt.Sprintf("Duplicate content at %s: %s vs %s",
				reqPath, site.pathMap[reqPath], path)
			if site.Config.Strict {
				return errors.New(msg)
			}
			log.Println(msg)

		}

		// If it's a Markdown file we turn it into a Page.
		log.Fatal("CREATE PAGE ETC")
		site.pathMap[reqPath] = path

		return nil

	}
	err = filepath.Walk(dir, walker)
	if err != nil {
		return fmt.Errorf("Error walking %s: %v", dir, err)
	}
	return nil

}
