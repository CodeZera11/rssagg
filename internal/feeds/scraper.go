package feeds

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/codezera11/rssagg/internal/database"
)

func FeedScraper(dbQueries *database.Queries, concurrency int, interval time.Duration) {

	log.Printf("Collecting feeds every %v seconds on %v routines\n", interval.Seconds(), concurrency)
	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {
		feeds, err := dbQueries.GetNextFeedsToFetch(context.Background(), int32(concurrency))

		if err != nil {
			log.Printf("Error fetching feeds")
			continue
		}

		log.Printf("Found %v feeds to fetch.", len(feeds))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go ScrapeFeed(*dbQueries, wg, feed)
		}

		wg.Wait()
	}
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Link        string `xml:"link"`
}

func ScrapeFeed(dbQueries database.Queries, wg *sync.WaitGroup, feed database.Feed) {

	defer wg.Done()

	rss, err := FetchFeed("https://blog.boot.dev/index.xml")

	if err != nil {
		log.Println("Error fetching feed data")
		return
	}

	for _, item := range rss.Channel.Items {
		fmt.Println("Found post:", item.Title)
	}

	_, err = dbQueries.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error marking feed as fetched")
		return
	}

	fmt.Printf("Feed collected %v with %v items\n", rss.Channel.Title, len(rss.Channel.Items))
}

func FetchFeed(url string) (*Rss, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://blog.boot.dev/index.xml", nil)
	if err != nil {
		log.Println("Error creating request!")
		return &Rss{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request")
		return &Rss{}, err
	}

	defer resp.Body.Close()

	rss := Rss{}
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		log.Println("Error decoding params")
		return &Rss{}, err
	}

	return &rss, nil
}
