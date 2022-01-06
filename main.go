package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
)

func main() {
	//fmt.Println("Hello World")

	readData()

}

func GetESClient() (*elastic.Client, error) {

	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func readData() {
	/* Creating context for API calls */
	ctx := context.Background()
	/* Fetching elastic Search Client */
	esclient, err := GetESClient()

	/* If err is present */
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	/* creating stundents array(slice) to store received student */
	var employees []Employee

	/* Creating Query */
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery("name.first_name", "Rahul"))

	// /* this block will basically print out the es query */
	// queryStr, err1 := searchSource.Source()
	// queryJs, err2 := json.Marshal(queryStr)

	// if err1 != nil || err2 != nil {
	// 	fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	// }
	// /* Generated Query */
	// fmt.Println("[esclient]Final ESQuery=\n", string(queryJs))

	/* searching with query */
	searchService := esclient.Search().Index("employee_catalogue").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}

	/* populating slice with the results */
	for _, hit := range searchResult.Hits.Hits {
		var employee Employee
		err := json.Unmarshal(hit.Source, &employee)
		if err != nil {
			fmt.Println("[Getting Students][Unmarshal] Err=", err)
		}
		employees = append(employees, employee)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, s := range employees {
			fmt.Printf("Student found Name: %s \n", s.Name)
		}
	}

}
