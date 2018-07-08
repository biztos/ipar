# ipar web server roadmap

TBD, let's get the MVP up running a couple blogs.

For it:

* site/pageset/page hierarchies
* mapping of all urls (404 via map)
* render all pages
* default template
* rss


## FOR (much?) LATER

* themes, including contrib from github?
    * Theme is a combo of templates and css
    * Ergo allow multiple (site, template) directories
    * And have "new" option load whatever sample config and content from theme
* other renderable page types
    * YAML obviously, it's pretty easy.  Maybe include "content" --?
    * Make them still be Page(s) and maybe make that a configurable option.

## MAYBE EVENTUALLY

* Environment overrides for config.
    * PRO: standard way to do systemd
    * CON: config file is preferred if possible
* Base URL.
    * Any real use for this as opposed to relative links everywhere?
