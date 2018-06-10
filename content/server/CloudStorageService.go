package server

import (
	"cloud.google.com/go/storage"
	"context"
	iterator2 "google.golang.org/api/iterator"
	"strings"
)

type cloudStorageService struct {
	client *storage.BucketHandle
	bucketName string
}

func (c *cloudStorageService) GetYears(ctx context.Context) ([]Year, error) {
	query := &storage.Query{
		Delimiter: "/",
		Versions: false,
	}
	iterator := c.client.Objects(ctx, query)
	var years []Year
	for {
		folder, err := iterator.Next()
		if err == iterator2.Done {
			break
		} else if err != nil {
			return nil, err
		}
		folderName := strings.TrimSuffix(folder.Prefix, "/")
		year := Year{Year: folderName, Name: folderName, Uri: "/movies/" + folderName}
		years = append(years, year)
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

