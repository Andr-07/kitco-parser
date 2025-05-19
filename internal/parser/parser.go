package parser

import (
	"encoding/xml"
	"io"
	"kitco-parser/pkg/utils"
	"net/http"
	"strings"
	"time"
)

type Sitemap struct {
	URLs []URLItem `xml:"url"`
}

type URLItem struct {
	Loc  string     `xml:"loc"`
	News NewsDetail `xml:"news"`
}

type NewsDetail struct {
	Title       string      `xml:"title"`
	PubDate     string      `xml:"publication_date"`
	Publication Publication `xml:"publication"`
}

type Publication struct {
	Name     string `xml:"name"`
	Language string `xml:"language"`
}

type ParsedNews struct {
	Title     string
	TitleHash []byte
	URL       string
	Source    string
	Lang      string
	Published time.Time
}

func FetchKitcoNews() ([]ParsedNews, error) {
	resp, err := http.Get("https://www.kitco.com/static-sitemaps/news.xml")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sitemap Sitemap
	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.DefaultSpace = "news"
	if err := decoder.Decode(&sitemap); err != nil {
		return nil, err
	}

	var result []ParsedNews
	// TODO: remove slice
	for _, item := range sitemap.URLs[:3] {
		t, err := time.Parse(time.RFC3339, item.News.PubDate)
		if err != nil {
			continue
		}

		result = append(result, ParsedNews{
			Title:     item.News.Title,
			TitleHash: utils.HashMD5(item.News.Title),
			URL:       item.Loc,
			Source:    item.News.Publication.Name,
			Lang:      item.News.Publication.Language,
			Published: t,
		})
	}

	return result, nil
}
