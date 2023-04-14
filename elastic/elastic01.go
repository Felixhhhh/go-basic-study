package main

//
//// Elasticsearch demo
//var client *elastic.Client
//
//type Person struct {
//	Name    string `json:"name"`
//	Age     int    `json:"age"`
//	Married bool   `json:"married"`
//}
//
//func main() {
//	urlOpt := elastic.SetURL("http://127.0.0.1:9200")
//	sniffOpt := elastic.SetSniff(false)
//	var err error
//	client, err = elastic.NewClient(urlOpt, sniffOpt)
//	if err != nil{
//		// Handle error
//		panic(err)
//	}
//	fmt.Println("connect to es success")
//	p1 := Person{Name: "lmh", Age: 18, Married: false}
//	put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())
//	if err != nil {
//		// Handle error
//		panic(err)
//	}
//	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
//}
