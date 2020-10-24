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
	ch := make(chan []Result)
	wg := sync.WaitGroup{}
	for i,file := range files{
		wg.Add(1)
		go func(ctx context.Context, fl string, i int, ch chan<- []Result){
			defer wg.Done()
			result:= []Result{}
			fileData,err := ioutil.ReadFile(fl)
			if err != nil{
				fmt.Print("cant't read file: " + fl)
			}
			lines := strings.Split(string(fileData),"\n")
			for index, line:= range lines {
				if strings.Contains(line,phrase){
					result = append(result, Result{Phrase: phrase, Line: line, LineNum: int64(index+1),ColNum:  int64(strings.Index(line, phrase)) + 1})
					
				}
			}
			if len(result)>0{
				ch <- result
			}
			
		}(ctx, file, i, ch)
	}
	
	go func() {
		defer close(ch)
		wg.Wait()

	}()
	return ch
}