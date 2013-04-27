// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mapper

import (
	"fmt"
	"github.com/andreaskoch/allmark/parser"
	"github.com/andreaskoch/allmark/path"
	"github.com/andreaskoch/allmark/view"
	"regexp"
	"time"
)

// Pattern which matches all HTML/XML tags
var HtmlTagPattern = regexp.MustCompile(`\<[^\>]*\>`)

func createMessageMapperFunc(parsedItem *parser.Result, pathProvider *path.Provider, targetFormat string) view.Model {

	return view.Model{
		Path:        pathProvider.GetWebRoute(parsedItem),
		Title:       getTitle(parsedItem),
		Description: getDescription(parsedItem),
		Content:     parsedItem.ConvertedContent,
		LanguageTag: getTwoLetterLanguageCode(parsedItem.MetaData.Language),
		Type:        parsedItem.Type,
	}

}

func getDescription(parsedItem *parser.Result) string {
	return parsedItem.MetaData.Date.Format(time.RFC850)
}

func getTitle(parsedItem *parser.Result) string {
	text := HtmlTagPattern.ReplaceAllString(parsedItem.ConvertedContent, "")
	excerpt := getTextExcerpt(text, 30)
	time := parsedItem.MetaData.Date.Format(time.RFC850)

	return fmt.Sprintf("%s: %s", time, excerpt)
}

func getTextExcerpt(text string, length int) string {

	if len(text) <= length {
		return text
	}

	return text[0:length] + " ..."
}
