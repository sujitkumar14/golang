//Data types

package main
import (
	"fmt"
	"os"
	"bufio"
)

func main(){

	var i int64 = 4
	var d float64 = 4.0
	var s string = "HackerRank"
	var x int64;
	var y float64;
	var z string;

	reader := bufio.NewReader(os.Stdin)
	fmt.Scanf("%d",&x);
	fmt.Scanf("%f",&y);
	z,_ = reader.ReadString('\n')

	fmt.Printf("%d\n%0.1f\n%s",x+i,d+y,s+z)
}