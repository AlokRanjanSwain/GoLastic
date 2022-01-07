package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/olivere/elastic/v7"
)

func main() {
	//fmt.Println("Hello World")

	insertData()
	//readData()
	//updateData()

}

func GetESClient() (*elastic.Client, context.Context) {
	/* Creating context for API calls */
	ctx := context.Background()
	/* Fetching elastic Search Client */
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	/* If err is present */
	if err != nil {
		fmt.Println("Error initializing : ", err)
		panic("Client fail ")
	}

	return client, ctx

}

func readData() {

	esclient, ctx := GetESClient()

	/* creating stundents array(slice) to store received student */
	var employees []Employee

	/* Creating Query */
	searchSource := amntQuer()

	/* Print the Query */
	printQuer(*searchSource)
	// searchSource := elastic.NewSearchSource()
	// searchSource.Query(elastic.NewMatchQuery("name.first_name", "Rahul"))

	/* searching with query */
	searchService := esclient.Search().Index("employee_catalogue").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return
	}
	/* Showing the total hits */
	fmt.Println("Number of employees meeting the desired criteria: ", searchResult.TotalHits())
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

func printQuer(searchSource elastic.SearchSource) {
	/* Getting the raw source of the query */
	queryStr, err1 := searchSource.Source()
	/*encoding it to json format */
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[esclient][GetResponse]err during query marshal=", err1, err2)
	}
	fmt.Println("Final ESQuery=\n", string(queryJs))
}

func insertData() {
	elsClient, ctx := GetESClient()

	/* copying raw json data */

	data, _ := ioutil.ReadFile("employee_dat2.json")

	//fmt.Println("data_rec", string(data))

	/* casting it to employee structure : Not Necessary Though */
	var employee Employee
	_ = json.Unmarshal(data, &employee)
	fmt.Println("employee name: ", employee.Name)

	/* for stringify the json of the defined structure */
	//dataJSON, _ := json.Marshal(employee)
	//js := string(dataJSON)

	js := string(data)
	// fmt.Println("string data", js)

	_, err := elsClient.Index().
		Index("employee_catalogue").
		BodyJson(js).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(" Data Inserted Successfully")

}

func updateData() {
	elsClient, ctx := GetESClient()

	/* Creating Script */
	script := elastic.NewScriptInline("ctx._source.name.first_name = params.name").Lang("painless").Param("name", "new_first")

	/* Creating Match Query : does not do exact matching for String type*/
	query := elastic.NewMatchQuery("name.first_name", "Alok")

	/*Initialize the services */
	updateService := elastic.NewUpdateByQueryService(elsClient).Index("employee_catalogue").Script(script).Query(query)

	/* Starting the Update operation */
	response, err := updateService.Do(ctx)

	if err != nil {
		panic(err)
	}

	if response.Updated > 0 {
		fmt.Printf("Updates successfully done on [ %d ] documents \n", response.Updated)
	} else {
		fmt.Println("No Documents matched with the query. Updated document: ", response.Updated)
	}

}

func deleteData() {

}
