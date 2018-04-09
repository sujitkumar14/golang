//arithmetic operations

package main


import "fmt"

func main(){
	

	var option int;
	var firstNumber float64;
	var secNumber float64;
	var result float64;

	fmt.Println("Choose an option to perform arithmetic operations.\n 1. add \n 2. subtract \n 3. multiply \n 4. divide \n");

	fmt.Scanf("%d",&option);

	fmt.Println("Please Enter you number\n");
	fmt.Println("First Number: ");
	fmt.Scanf("%f",&firstNumber);
	fmt.Println("sec Number");
	fmt.Scanf("%f",&secNumber);

	if(option == 1){

		result = firstNumber+secNumber;
	}else if(option == 2){
		result = firstNumber - secNumber;
	}else if(option == 3){
		result = firstNumber * secNumber;

	}else{
		result  = firstNumber / secNumber;
	}


	fmt.Printf("result : %f",result);

}