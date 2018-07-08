// cmd/ipar/main.go -- the "ipar" executable.

package main

import (
	"github.com/biztos/ipar"
	"github.com/docopt/docopt-go"
	"log"
	"os"
)

// Options contains the app options set via docopt.
type Options struct {
	Dir    string
	Config string
	Future bool
	Drafts bool
}

// Version is the official app version for the executable.
var Version = "0.1.0"

// Usage is the docopt usage/help text.
var Usage = `ipar - a Markdown-centric web server.

Usage:
  ipar [options] DIR
  ipar [options] --config=FILE
  ipar -h | --help
  ipar --version

Options:
  --config=FILE Configuration file if not "config.yaml" in DIR.
  --future      Display pages published in the Future.
  --drafts      Display pages marked as Draft.
  -h --help     Show this screen.
  --version     Show version.

Note that if --config is specified then DIR must not be; and in this case the
directory will be that defined in the config file or the directory of that
file by default.  This option is generally used for testing alternative
config files.

Description:

The ipar web server loads and evaluates content from the provided DIR and
serves it at the configured port.

For more information see the project page:

https://github.com/biztos/ipar

`

func main() {
	err := Start(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

// Start starts the application, returning an error if anything goes wrong.
// If successful, the application will block (and log) until terminated.
func Start(argv []string) error {

	o := &Options{}

	opts, err := docopt.ParseArgs(Usage, argv, "ipar version "+Version)
	if err != nil {
		return err
	}
	err = opts.Bind(o)
	if err != nil {
		return err
	}

	var site *ipar.Site
	if o.Config != "" {
		site, err = ipar.NewSiteFromFile(o.Config)
	} else {
		site, err = ipar.NewSiteFromDir(o.Dir)
	}

	if err != nil {
		return err
	}

	// For our two flags we override the config when set to true.  There seems
	// to be no easy way to do this but hey, maybe on godoc... or maybe write
	// one of our own...
	if o.Future {
		site.Config.Future = true
	}
	if o.Drafts {
		site.Config.Drafts = true
	}

	log.Println(site.Config.Dir)
	log.Println(site.Config.Future)
	log.Println(site.Config.Drafts)
	return nil
}
