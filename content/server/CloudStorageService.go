package server

import (
	"cloud.google.com/go/storage"
	"context"
	iterator2 "google.golang.org/api/iterator"
	"strings"
	"net/http"
	"github.com/DiTo04/spexflix/common/codecs"
	"errors"
)

type cloudStorageService struct {
	client *storage.BucketHandle
	bucketName string
}

func (c *cloudStorageService) GetYears(ctx context.Context) ([]Year, error) {
	query := &storage.Query{
		Versions: false,
	}
	iterator := c.client.Objects(ctx, query)
	yearChannel := make(chan *Year)
	nrOfItems := 0
	for {
		file, err := iterator.Next()
		if err == iterator2.Done {
			break
		} else if err != nil {
			return nil, err
		}
		if !strings.Contains(file.Name, "meta.json") {
			continue
		}
		go c.getYearData(file, yearChannel)
		nrOfItems += 1
	}
	return gatherYearsFromChannel(nrOfItems, yearChannel)
}

func (c *cloudStorageService) getYearData(file *storage.ObjectAttrs, channel chan <- *Year) {
	folderName := strings.TrimSuffix(file.Name, "/meta.json")
	posterUri := make(chan string)
	go c.getPosterUri(folderName, posterUri)
	rsp, err := http.Get(file.MediaLink)
	if err != nil {
		channel <- nil
	}
	defer rsp.Body.Close()
	var year Year
	codecs.JSON.Decode(rsp.Body, &year)
	year.Year = folderName
	year.Uri = "/movies/" + folderName
	year.PosterUri = <- posterUri
	channel <- &year
}

func (c *cloudStorageService) getPosterUri(folder string, out chan <- string) {
	poster := c.client.Object(folder + "/poster")
	attr, err := poster.Attrs(context.TODO())
	if err != nil {
		out <- ""
	}
	out <- attr.MediaLink
}

func gatherYearsFromChannel(nrOfItems int, channel <-chan *Year) ([]Year, error){
	years := make([]Year, nrOfItems)
	select {
	case year := <- channel:
		if year == nil {
			return nil, errors.New("could not get metadata")
		}
		years[nrOfItems - 1] = *year
		nrOfItems -= 1
		if nrOfItems == 0 {
			break
		}
	}
	return years, nil
}

func (c *cloudStorageService) GetContent(ctx context.Context, year string) ([]Movie, error) {
	query := &storage.Query{
		Versions: false,
		Prefix: year,
	}
	iterator := c.client.Objects(ctx, query)
	var movies []Movie
	for {
		folder, err := iterator.Next()
		if err == iterator2.Done {
			break
		} else if err != nil {
			return nil, err
		}
		if !strings.Contains(folder.ContentType, "video") {
			continue
		}
		split := strings.Split(folder.Name, "/")
		folderName := split[len(split) - 1]
		movie := Movie{Name: folderName, Uri: folder.MediaLink}
		movies = append(movies, movie)
	}
	return movies, nil
}

