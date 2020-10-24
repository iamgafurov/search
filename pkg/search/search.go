package search
import (
	"context"
	"sync"
	"fmt"
	"strings"
	"io/ioutil"
	)

type Result struct {
	Phrase string
	Line string
	LineNum int64
	ColNum int64
}

func All(ctx context.Context, phrase string, files []string)<-chan []Result{
	ch := make(chan []Result,len(files))
	defer close(ch)
	
	wg := sync.WaitGroup{}
	for i,file := range files{
		wg.Add(1)
		println("gorutine %v start ",i)
		go func(){
			defer wg.Done()
			var result  []Result
			fileData,err := ioutil.ReadFile(file)
			if err != nil{
				fmt.Print("cant't read file: " + file)
			}
			lines := strings.Split(string(fileData),"\n")
			for index, line:= range lines {
				if strings.Contains(line,phrase){
					words:= strings.Split(line, " ")
					for col,word:=range words {
						if word == phrase{
							result = append(result, Result{Phrase: phrase, Line: line, LineNum: int64(index+1),ColNum: int64(col+1)})
							//fmt.Println(file,line)
							break;
						}
					}
				}
			}
			
			ch <- result
			
		}()
	}
	wg.Wait()
	
	return ch
}