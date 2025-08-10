package scraper

import (
	"context"
	"database/sql"
	"regexp"
	"strings"
	"unicode"

	_ "github.com/lib/pq"
)

var htmlTagRegex = regexp.MustCompile("<.*?>")

func stripHTML(s string) string {
    return htmlTagRegex.ReplaceAllString(s, "")
}

func cleanText(text string) string {

	text = stripHTML(text)
    
    text = strings.TrimSpace(text)
    text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	if len(text) > 0 {
		runes := []rune(text)
		runes[0] = unicode.ToUpper(runes[0])
		text = string(runes)
	}
	return text
}

func saveToDB(db *sql.DB, data []Article) error{
	ctx := context.Background()

	for _, item := range data{
		_, err := db.ExecContext(ctx, 
			`INSERT INTO Articles (title, content, src) VALUES ($1, $2, $3)`,
			item.Title, item.Content, item.Src,
		)

		if(err != nil){
			return err
		}
	}

	return nil
}