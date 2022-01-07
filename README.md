# GoLastic [Golang+ElasticSearch]
1. Install elastic search, Kibana, and Golang for your OS. Its pretty easy to do with help of google
2. Run elastic search and kibana on local host
3. To use kibana dev tools. log on to http://localhost:5601/app/dev_tools#/console
4. Create employee_catalogue [index] on elastic search using Kibana Console with mappings provided in empoyee_map.json
```
PUT employee_catalogue/
{
    "settings":{
        "number_of_shards": 1,
    	"number_of_replicas": 1
    },
    "mappings": {
        "properties": {
            "name": {
                "properties": {
                    "first_name": {
                        "type":"text"
                    },
                    "last_name": {
                        "type":"text"
                    }
                }
            },
            "department":{
                "properties": {
                    "department_name": {
                        "type":"text"
                    }
                }
            },
            "salary" :{
                "properties" : {
                    "level" : {
                        "properties":{
                            "high_level": {
                                "type":"boolean"
                            },
                                "low_level": {
                                "type":"boolean"
                            },
                                "amount" : {
                                "type":"long"
                            }
                        }    
                            
                    }
                }
            },
            "address" :{
                "properties":{
                    "street_name": {
                        "properties":{
                            "house_no":{
                                "type":"text"
                            },
                            "locality":{
                                "type":"text"
                            },
                            "landmark":{
                                "type":"text"
                            }
                        }
                    },
                    "city": {
                        "type":"text"
                    }
                }
            },
            "family_members" : {
                "type":"text"
            }
        }

    }
    
}

```
4. Post data inside the index
```
POST employee_catalogue/_doc
{
    "name": {
        "first_name":"Rahul",
        "last_name":"Baba"
    },
    "department": {
        "department_name":"HR"
    },
    "salary": {
        "level": {
            "high_level":true,
            "low_level":false,
            "amount":23000
        }
    },
    "address": {
        "street_name": {
            "house_no":"12/A",
            "locality":"IIT",
            "landmark":"Rajya Sabha"
        },
        "city":"Delhi"
    },
    "family_members": [
        "Parents",
        "Spouse",
        "Son"
    ]
    
}

```
6. Initialize module for GO project
 ```
go mod init module_name
```
7. For integrating Golang with Elastic, the mentioned package can be used 
    [Olivere Elastic](https://github.com/olivere/elastic)
8. To import the module use on CLI [command line interface] on the project folder

```
go get "github.com/olivere/elastic/v7"
```
9. You can run go file using
```
go run ./     
OR
go run main.go employee_struct.go
```
## Notes 
###### Insert
Data is inserted by reading Json Files. There are commented lines showing how the same can be done by using structure for the data [Need to assign value and Marshall it to json data]

###### Read
There are various type of queries (range, match, term ,etc.) which could be used depending upon the situation.\
**Reminder** In match query, there is no exact matching for string data type.

We either need to
* Define mappings with field assigned to type **Keyword**
* Use term query 

###### Update 
In the demo, update by query is shown. But if document ID is known,the same can be done by Update Service.\
The Kibana console command for the same is:
```
POST employee_catalogue/_update_by_query
{
  "script": {
    "source": "ctx._source.name.first_name=params.name",
    "lang": "painless",
    "params": {
      "name":"new_first"
    }
  }
  , "query": {
      "match": {
        "name.first_name": "Alok"
    }
  }
}
```
###### Delete
Similiar to update operation, delete is also done by query in the demo. Other alternative is to do it with delete Service.\
The Kibana console command for the same is:
```
POST employee_catalogue/_delete_by_query
{
  "query":{
    "match":{
      "name.first_name":"new_first"
    }
  }
}
```

## Purpose

Both **Golang** and **Elastic Search** is new and trending technology, So it could be used heavily in near future.
* With **NoSQL** Databases merging as everyone's favorite in DB technology.\It is good to work on some NoSQL technology which is open source, highly scalable and easy to understand.
* **Golang** gives better concurrency capabilities, which provide light Apps to work more faster. 

I am open to feedbacks, Please let me know if any.
## Happy Learning  



