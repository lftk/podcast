package podcast

import (
	"encoding/xml"
	"io"
	"time"
)

type EnclosureType string

const (
	EnclosureM4A  EnclosureType = "audio/x-m4a"
	EnclosureMP3                = "audio/mpeg"
	EnclosureMOV                = "video/quicktime"
	EnclosureMP4                = "video/mp4"
	EnclosureM4V                = "video/x-m4v"
	EnclosurePDF                = "application/pdf"
	EnclosureEPUB               = "document/x-epub"
)

type PubData struct {
	time.Time
}

func (p PubData) MarshalText() (text []byte, err error) {
	t := p.Time
	if t.IsZero() {
		t = time.Now().UTC()
	}
	text = []byte(t.Format(time.RFC1123Z))
	return
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Summary struct {
	Text string `xml:",cdata"`
}

type Item struct {
	Title     string       `xml:"title"`
	Link      string       `xml:"link,omitempty"`
	Author    string       `xml:"itunes:author"`
	Subtitle  string       `xml:"itunes:subtitle"`
	Summary   *Summary     `xml:"itunes:summary"`
	Image     *Image       `xml:"itunes:image"`
	Enclosure *Enclosure   `xml:"enclosure"`
	GUID      string       `xml:"guid"`
	PubDate   *PubData     `xml:"pubDate,omitempty"`
	Duration  string       `xml:"itunes:duration,omitempty"`
	Explicit  ExplicitType `xml:"itunes:explicit,omitempty"`
}

type Owner struct {
	Name  string `xml:"itunes:name"`
	Email string `xml:"itunes:email"`
}

type Image struct {
	Href string `xml:"href,attr"`
}

type Category struct {
	Text        string    `xml:"text,attr"`
	Subcategory *Category `xml:"itunes:category,omitempty"`
}

type ExplicitType string

const (
	ExplicitYes   ExplicitType = "yes"
	ExplicitNo                 = "no"
	ExplicitClean              = "clean"
)

type Channel struct {
	Title       string       `xml:"title"`
	Link        string       `xml:"link,omitempty"`
	Language    string       `xml:"language"`
	Copyright   string       `xml:"copyright,omitempty"`
	Subtitle    string       `xml:"itunes:subtitle"`
	Author      string       `xml:"itunes:author"`
	Summary     *Summary     `xml:"itunes:summary"`
	Description string       `xml:"description"`
	Owner       *Owner       `xml:"itunes:owner,omitempty"`
	Image       *Image       `xml:"itunes:image,omitempty"`
	Categories  []*Category  `xml:"itunes:category,omitempty"`
	Explicit    ExplicitType `xml:"itunes:explicit,omitempty"`
	Items       []*Item      `xml:"item,omitempty"`
}

func (c *Channel) AddCategory(category, subcategory string) {
	if len(category) > 0 {
		ca := &Category{Text: category}
		if len(subcategory) > 0 {
			ca.Subcategory = &Category{Text: subcategory}
		}
		c.Categories = append(c.Categories, ca)
	}
}

func (c *Channel) AddItem(item *Item) {
	c.Items = append(c.Items, item)
}

func (c *Channel) WriteTo(w io.Writer) (n int64, err error) {
	v := struct {
		XMLName xml.Name `xml:"rss"`
		Version string   `xml:"version,attr"`
		Itunes  string   `xml:"xmlns:itunes,attr"`
		Channel *Channel `xml:"channel"`
	}{
		Version: "2.0",
		Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Channel: c,
	}

	n1, err := w.Write([]byte(xml.Header))
	if err != nil {
		return
	}

	b, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return
	}

	n2, err := w.Write(b)
	if err != nil {
		return
	}

	n = int64(n1) + int64(n2)
	return
}
