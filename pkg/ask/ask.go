//==============================================================================
//
// Title:		ask.go
// Purpose:		asks user for the game's inputs(name, age and guessed numbers)
//
// Created on:	7/26/2016 at 5:37:47 PM by Avelino Puello
// Copyright:	. All Rights Reserved.
//
//==============================================================================
/*

-------------------------------------------------------------------------
*/

package ask

import(
	"fmt"
	"bufio"
    "os"
    print "../print"
)
//this function returns 2 values(name string and age int)
//this func asks player for name and age
//it makes sure that age is an int
//***TODO: It wouldn't be a bad idea to check name is string type also
func Ask()(string, int){
	readerName := bufio.NewReader(os.Stdin)
	print.PrintWhatsYourName()
	name, _ := readerName.ReadString('\n')
	
	var age int
 	print.PrintHowOldAreYou()
 	//Here we make sure user is entering a number(we could do something similar with the name above)
 	if _, err := fmt.Scanf("%d", &age); err != nil {
 		//if user enter a text we return "errAgeNotNumer", 0 and terminate the game becuase we felt offended by this :)
   		return "errAgeNotNumber", 0
 	}
	
	
	return name, age
	
}
