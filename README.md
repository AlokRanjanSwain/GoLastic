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
            "high_level":"true",
            "low_level":"false",
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
6. You can run go file using
```
go run ./     
OR
go run main.go employee_struct.go
```



