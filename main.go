package main

import (
	"net/http"
	"log"
	"golang.org/x/net/html"
	"fmt"
	"flag"
	"github.com/skratchdot/open-golang/open"
)


var url string 
func main() {
	flag.Parse()
	latestQues:= GrabQuestions()
	for{
		newQues := GrabQuestions()
		if latestQues!=newQues{
			fmt.Println("New Question:",newQues)
			open.Run(url)
			latestQues = newQues
		}
	}
	fmt.Println(latestQues)

}

//GrabQuestions visits http://stackoverflow.com/questions/tagged/x and if there is a new question with tag x, then returns the question as a string.  
func GrabQuestions() string {
	url = "http://stackoverflow.com/questions/tagged/" + flag.Args()[0]
	resp,err := http.Get(url)
	if err!=nil{
		log.Fatal(err)
	}

	z := html.NewTokenizer(resp.Body)

	for{
		tt := z.Next()

		if tt == html.ErrorToken {
			return "No question found"
		}else if tt == html.StartTagToken {
			i:= z.Token()

			if i.Data=="a"{
				for _,a := range i.Attr{
					if a.Key=="class" && a.Val=="question-hyperlink"{
						tt = z.Next()
						return z.Token().Data
						 
					}
				}
			}
		}
	}
}