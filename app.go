package main

import (
	"fmt"
	"net/http"
)
 
func main(){
	db:=database {"shoes":50,"socks":5}
	http.HandleFunc("/list",db.list)
	http.HandleFunc("/price",db.price)
	http.ListenAndServe(":80", nil)
}

type dollars float32

func (d dollars) String() string {return fmt.Sprintf("$%.2f",d)}

type database map[string]dollars

func (db database) ServeHTTP (w http.ResponseWriter, req *http.Request){
	switch req.URL.Path{
	case "/list":
		for item,price:=range db{
		fmt.Fprintf(w,"%s:%s\n", item, price)
	}
	case "/price":
	item:=req.URL.Query().Get("item")
	price, ok:=db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)//404
		fmt.Fprintf(w,"нет товара: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
	
	default:
		w.WriteHeader(http.StatusNotFound)//404
		fmt.Fprintf(w, "page is not found: %s\n",req.URL)

	}	
}

func (db database) list (w http.ResponseWriter, req *http.Request){
	for item,price :=range db{
		fmt.Fprintf(w, "%s:,%s\n", item,price)
	}
}

func(db database) price(w http.ResponseWriter, req *http.Request){
	item:=req.URL.Query().Get("item")
	price, ok:=db[item]

	if!ok{
		w.WriteHeader(http.StatusNotFound)//404
		fmt.Fprintf(w,"no such item:%q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n",price)
}