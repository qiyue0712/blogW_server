package main

import (
	"blogW_server/core"
	"blogW_server/flags"
	"blogW_server/global"
	"blogW_server/models"
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func create() {
	var article = models.ArticleModel{
		Model: models.Model{
			ID: 1,
		},
		Title:   "王",
		Content: "这是内容",
		UserID:  1,
		Status:  1,
	}

	indexResponse, err := global.ESClient.Index().Index(article.Index()).BodyJson(article).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", indexResponse)
}

func list() {
	limit := 2
	page := 1
	from := (page - 1) * limit

	query := elastic.NewBoolQuery()
	res, err := global.ESClient.Search(models.ArticleModel{}.Index()).Query(query).From(from).Size(limit).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	count := res.Hits.TotalHits.Value // 总数
	fmt.Println(count)
	fmt.Println(res.Hits)
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}

func DocDelete() {

	deleteResponse, err := global.ESClient.Delete().
		Index(models.ArticleModel{}.Index()).Id("P81S05YBb86QLo2LPb6q").Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deleteResponse)
}

func update() {
	updateResponse, err := global.ESClient.Update().Index(models.ArticleModel{}.Index()).Refresh("true").
		Id("Qs0w3ZYBb86QLo2Lw76l").
		Doc(map[string]any{
			"content": "冬天111",
		}).Do(context.Background())
	fmt.Println(updateResponse, err)
}

func main() {
	flags.Parse()                   // 环境变量参数
	global.Config = core.Readconf() // 读配置文件
	core.InitLogrus()
	global.ESClient = core.EsConnect()
	//create()
	//list()
	//DocDelete()
	update()
}
