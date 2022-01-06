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
	searchSource := amntQuer()
	// searchSource := elastic.NewSearchSource()
	// searchSource.Query(elastic.NewMatchQuery("name.first_name", "Rahul"))

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
			fmt.Println("[Getting Employees][Unmarshal] Err=", err)
		}

		employees = append(employees, employee)
	}

	if err != nil {
		fmt.Println("Fetching student fail: ", err)
	} else {
		for _, emp := range employees {
			fmt.Printf("Employee Name: %s \n", emp.Name)
			//fmt.Printf("salary details : %t , %t, %d ", emp.Salary.Level.HighLevel, emp.Salary.Level.LowLevel, emp.Salary.Level.Amount)
		}
	}

}

func amntQuer() *elastic.SearchSource {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewRangeQuery("salary.level.amount").Gte(20000))
	return searchSource
}
