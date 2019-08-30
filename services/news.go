package services

import (
	"context"
	"encoding/json"
	"github.com/karldoenitz/Tigo/logger"
	"github.com/olivere/elastic"
	"polar/global/config"
	"polar/models"
)

func GetNews(city string, category string, page int64, size int64) (result []models.News, count int64, err error) {
	client := config.GetESClient()
	cityMatchQuery := elastic.NewMatchQuery("cities", city)
	categoryMatchQuery := elastic.NewMatchQuery("categories", category)
	boolQuery := elastic.NewBoolQuery().Must(cityMatchQuery, categoryMatchQuery)
	ctx := context.Background()
	searchResult, err := client.Search().
		Index("media").
		Type("news").
		Query(boolQuery).
		Sort("weight", false).
		From(int(page)).Size(int(size)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return
	}
	count = searchResult.Hits.TotalHits.Value
	if count > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var news models.News
			err := json.Unmarshal(hit.Source, &news)
			if err != nil {
				logger.Error.Println(err.Error())
				continue
			}
			result = append(result, news)
		}
	}
	return
}
