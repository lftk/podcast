package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/4396/podcast"
)

func main() {
	c := podcast.Channel{
		Title:     "All About Everything",
		Link:      "http://www.example.com/podcasts/everything/index.html",
		Language:  "en-us",
		Copyright: "&#x2117; &amp; &#xA9; 2014 John Doe &amp; Family",
		Subtitle:  "A show about everything",
		Author:    "John Doe",
		Summary: &podcast.Summary{
			Text: "All About Everything is a show about everything. Each week we dive into any subject known to man and talk about it as much as we can. Look for our podcast in the Podcasts app or in the iTunes Store",
		},
		Owner: &podcast.Owner{
			Name:  "John Doe",
			Email: "john.doe@example.com",
		},
		Image: &podcast.Image{
			Href: "http://example.com/podcasts/everything/AllAboutEverything.jpg",
		},
		Explicit:    podcast.ExplicitNo,
		Description: "hello description...",
	}

	c.AddCategory("Technology", "Gadgets")
	c.AddCategory("TV &amp; Film", "")
	c.AddCategory("Arts", "Food")

	c.AddItem(&podcast.Item{
		Title:    "Shake Shake Shake Your Spices",
		Author:   "John Doe",
		Subtitle: "A short primer on table spices",
		Summary: &podcast.Summary{
			Text: `This week we talk about <a href="https://itunes/apple.com/us/book/antique-trader-salt-pepper/id429691295?mt=11">salt and pepper shakers</a>, comparing and contrasting pour rates, construction materials, and overall aesthetics. Come and join the party!`,
		},
		Image: &podcast.Image{
			Href: "http://example.com/podcasts/everything/AllAboutEverything/Episode1.jpg",
		},
		Link: "http://www.example.com/podcasts/everything/index.html",
		Enclosure: &podcast.Enclosure{
			URL:    "http://example.com/podcasts/everything/AllAboutEverythingEpisode3.m4a",
			Length: 8727310,
			Type:   podcast.EnclosureMP3,
		},
		GUID: "http://example.com/podcasts/archive/aae20140615.m4a",
		PubDate: &podcast.PubData{
			Time: time.Now(),
		},
		Duration: "07:04",
		Explicit: podcast.ExplicitNo,
	})

	buf := bytes.NewBuffer(nil)
	_, err := c.WriteTo(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())
}
