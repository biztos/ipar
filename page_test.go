// page_test.go

package ipar_test

import (
	// Standard:
	"testing"

	// Helpful:
	"github.com/stretchr/testify/assert"

	// Under test:
	"github.com/biztos/ipar"
)

func Test_Page_Title_UpperCase(t *testing.T) {

	assert := assert.New(t)

	p := &ipar.Page{Meta: map[string]interface{}{
		"TITLE": "this",
	}}
	assert.Equal("this", p.Title())

}

func Test_Page_Title_MixedCase(t *testing.T) {

	assert := assert.New(t)

	p := &ipar.Page{Meta: map[string]interface{}{
		"Title": "this",
	}}
	assert.Equal("this", p.Title())

}

func Test_Page_Title_Float(t *testing.T) {

	assert := assert.New(t)

	p := &ipar.Page{Meta: map[string]interface{}{
		"title": 1.2345,
	}}
	assert.Equal("1.2345", p.Title())

}

func Test_Page_Title_Bool(t *testing.T) {

	assert := assert.New(t)

	p := &ipar.Page{Meta: map[string]interface{}{
		"title": true,
	}}
	assert.Equal("true", p.Title())

}

func Test_Page_Title_NoMatch(t *testing.T) {

	assert := assert.New(t)

	p := &ipar.Page{Meta: map[string]interface{}{
		"notTitle": "this",
	}}
	assert.Equal("", p.Title())

}
