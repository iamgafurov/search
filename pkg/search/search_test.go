package search

import (
	"testing"
	"context"
	"fmt"
	)
		
func TestSearch_success(t *testing.T){
	ctx := context.Background()
	result := Any(ctx, "1111", []string{"./file1.txt","./file2.txt","./file3.txt"})
	/*if len(result) == 0 {
		t.Errorf("Channel must not be empty")
	}*/
	res:= <-result
	fmt.Print(res,"\n")
	res= <-result
	fmt.Print(res,"\n")
	res= <-result
	fmt.Print(res,"\n")
	
}