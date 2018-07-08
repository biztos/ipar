// doc.go - overview documentation

// Package ipar defines an opinionated, Markdown-centric web server.
//
// It is normally used via the "ipar" command found under the "cmd"
// subdirectory. The documentation here is for developers wishing to customize
// or extend the software.  For user documenation please see the repository
// page:
//
// https://github.com/biztos/ipar
//
// STATUS
//
// This alpha software and its API should not be considered stable until a
// 1.0 tag is set and this notice is removed.
//
// Opinionated
//
// This package presents an opinion of how dynamically rendered static content
// sites should be organized.  Deviating from this view is of course possible,
// and it is likely that future versions will support more flexibility, but
// as of this writing there is an "ipar way" which is:
//
//  * One directory per site.
//  * TLS (HTTPS) by default.
//  * Three subdirectories:
//    + content
//    + templates
//    + static
//  * The content directory may also contain static files.
//  * Markdown files (.md) are rendered, all else are static files.
//
// Markdown-Centric
//
// Pages under the content directory are stored as Markdown files, to be
// parsed and rendered with the "frostedmd" package, which is based on
// the excellent "blackfriday" package -- by default this is a well-extended
// version of Markdown.
//
// At present only Markdown files are rendered; other types may be available
// in future versions.
//
// Web Server
//
// This software is designed to run as a public-facing web server in UNIX
// under "systemd" -- but for anything very popular, important, or sensitive
// the author STRONGLY suggests putting a more mature web server such as
// "nginx" in front of it.
//
// The Hierarchy of Structures
//
// A Site has a Pageset which is made of Pages.  The Site also knows about
// static assets and templates; when a template is rendered it is given a
// Dot, which contains the Site and the relevant (sub-)Pageset together
// with a Page, which may be nil.
//
// The App is used to serve a Site; the standard executable command is a
// thin wrapper around that.
//
// Future Plans
//
// Please see the project roadmap for planned and speculative future
// development:
//
// https://github.com/biztos/ipar
package ipar
