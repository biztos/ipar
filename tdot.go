package ipar

// TDot represents the "dot object" made available to templates, i.e. a
// "Template Dot."
type TDot struct {
	Site    *Site    // The Site itself.
	Pageset *Pageset // The Pageset, if any.
	Page    *Page    // The current Page, if any.
}
