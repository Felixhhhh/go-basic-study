package main

//func main() {
//	c, err := redis.Dial("tcp", "localhost:6379")
//	if err != nil {
//		fmt.Println("conn redis failed,", err)
//		return
//	}
//
//	defer c.Close()
//	_, err = c.Do("MSet", "qq", 11, "ww", 22)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	r, err := redis.Ints(c.Do("MGet", "qq",  "ww"))
//	for _, v := range r {
//		fmt.Println(v)
//	}
//
//	_, err = c.Do("expire", "qq", 10)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	//_, err = c.Do("Set", "abc", 100)
//	//if err != nil {
//	//	fmt.Println(err)
//	//	return
//	//}
//	//
//	//r, err := redis.Int(c.Do("Get", "abc"))
//	//if err != nil {
//	//	fmt.Println("get abc failed,", err)
//	//	return
//	//}
//	//
//	//fmt.Println(r)
//}
