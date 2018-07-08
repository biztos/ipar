// page.go

package ipar

import (
	"fmt"
	"strings"
)

// Page represents a single web page to be rendered using a template.
type Page struct {
	Path    string
	Link    string
	Meta    map[string]interface{}
	Content string
}

// Title returns the case-insensitive "title" value of the Page's Meta.
// If the "title" value is not a string, its stringfied value ("%v") is
// returned. If no "title" value exists the empty string is returned.
// If more than one key renders as "title" in lowercase the behavior is
// undefined.
func (p *Page) Title() string {

	for k, v := range p.Meta {
		if strings.ToLower(k) == "title" {
			s, ok := v.(string)
			if ok {
				return s
			}
			return fmt.Sprintf("%v", v)
		}
	}
	return ""
}

// Live
