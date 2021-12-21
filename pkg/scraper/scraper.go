package scraper

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/url"
	"regexp"
	"strings"
)

var nondigit = regexp.MustCompile(`\D`)

var ErrDomainNotSupported = errors.New("error: domain is not supported")

//go:embed scraper-config.json
var config []byte

type Scraper interface {
	ScrapeURL(host, productURL string) (string, string, string, string, error)
}

func NewScraper() (Scraper, error) {
	var s scraper
	err := json.Unmarshal(config, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

type scraper map[string]struct {
	Title       map[string]string `json:"title"`
	Image       map[string]string `json:"image"`
	Price       map[string]string `json:"price"`
	Description map[string]string `json:"description"`
}

func (s *scraper) ScrapeURL(host, productURL string) (string, string, string, string, error) {
	element, ok := (*s)[host]
	if !ok {
		return "", "", "", "", ErrDomainNotSupported
	}

	c := colly.NewCollector(
		colly.DetectCharset(),
	)

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
	})

	var title, image, price, description string

	registerCallback(c, element.Title, &title)
	registerCallback(c, element.Image, &image)
	registerCallback(c, element.Price, &price)
	registerCallback(c, element.Description, &description)

	err := c.Visit(productURL)
	if err != nil {
		return "", "", "", "", err
	}

	err = sanitizeOutput(&title, &image, &price, &description)
	if err != nil {
		return "", "", "", "", err
	}

	return title, image, price, description, nil
}

func registerCallback(c *colly.Collector, element map[string]string, value *string) {
	c.OnHTML(element["tag"], func(e *colly.HTMLElement) {
		if attr := element["attr"]; attr != "" {
			*value = e.Attr(attr)
		} else if child := element["child"]; child != "" {
			*value = e.ChildText(child)
		} else {
			*value = e.Text
		}
	})
}

func sanitizeOutput(title, image, price, description *string) error {
	parsedImageURL, err := url.Parse(*image)
	if err != nil {
		return err
	}

	if parsedImageURL.Scheme == "" || parsedImageURL.Scheme == "http" {
		parsedImageURL.Scheme = "https"
	}

	*title = strings.TrimSpace(*title)
	*image = fmt.Sprintf("%s://%s%s", parsedImageURL.Scheme, parsedImageURL.Host, parsedImageURL.Path)
	if parsedImageURL.RawQuery != "" {
		*image += fmt.Sprintf("?%s", parsedImageURL.RawQuery)
	}
	*price = nondigit.ReplaceAllString(*price, "")
	*description = strings.TrimSpace(*description)

	return nil
}
