//array2 in go


package main

import "fmt"


func main(){
	
	var array[10] int64;

	for i := 0;i<len(array);i++ {

		fmt.Scanf("%d",&array[i]);
	}


	for i:= 0;i<len(array);i++ {

		fmt.Println(array[i]);
	}

}