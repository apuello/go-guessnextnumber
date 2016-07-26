//==============================================================================
//
// Title:		main.go
// Purpose:		Guess if random number is greater or less than 5
//
// Created on:	7/25/2016 at 3:56:36 PM by Avelino Puello
// Copyright:	. All Rights Reserved.
//
//==============================================================================
/*

---------------------------------------------------------------------------

Project : Simple Game Design
Autthor : Avelino Puello
Date : July 25 2016

Description: This Game consiste in 
- Guess  the  range 
- the secrect number is lower or higher than 5
- win 1 point for each correct guess or lose 1 point wrong guess
- s command print your info (name age lives)
- q exit game
- logging with player name

--------------------------------------------------------------------------
*/

package main 

import(
	"fmt"
	"bufio"
    "os"
	"math/rand"
)

//Globals

// var MIDDLE_NUM int 
// var NUM_CHANCES int

//strucs

type Player struct {
   name  string 
   age, pointsWon, lives int
}

type Game struct {
   active bool
   numberToGuess, level int
}



/* Here we print the intro of the game */

func printHeader(numberToGuess int){
 fmt.Println("====================================================================")
 fmt.Println("======================~~ Avelino RULLETE ~~=========================")
 fmt.Println("Welcome to the game of Avelino and become millonaire in one game")
 fmt.Printf("Guess if number is greater or less than %d and win instantly ! \n", numberToGuess )
 fmt.Println("====================================================================")
}


//this function returns 2 values(name string and age int)
//this func asks player for name and age
//it makes sure that age is an int
//***TODO: It wouldn't be a bad idea to check name is string type also
func ask()(string, int){
	readerName := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? : ")
	name, _ := readerName.ReadString('\n')
	
	var age int
 	fmt.Print("How old are you? : ")
 	//Here we make sure user is entering a number(we could do something similar with the name above)
 	if _, err := fmt.Scanf("%d", &age); err != nil {
 		//if user enter a text we return "errAgeNotNumer", 0 and terminate the game becuase we felt offended by this :)
   		return "errAgeNotNumber", 0
 	}
	
	
	return name, age
	
}

//I wasn't able to use the rand method and return random numbers
//what i did for now was to create an array of 10 int(s) from 1 to 0
//using the r.Perm(int) function
//example [8,6,0,1,7,2,5,4,9,5]
//finally, I return the []int
func giveMeTheRandoms() []int{
	r := rand.New(rand.NewSource(99))
	ranNumber := r.Perm(10)
	//Slicing the first 5 random numbers
	var takeFirstFive []int = ranNumber[:6]
	return takeFirstFive
}

//play func is hte heart of this program
//here we receive 2 pointers(we need pointers as we will be modifying the properties of these)
//we will also return these 2 pointers as we will be saving the info in a DB
//maybe postgreSQL?
func play(p *Player, g *Game) (*Player, *Game){
	var guessedNumber int
	randoms := giveMeTheRandoms() //getting random number []int
	for i := 0; i < g.numberToGuess; i++ { //looping through game max number to guess
		
	 	fmt.Print("Guess Next Number?: ")
	 	//Here we make sure user is entering a number(we could do something similar with the name above)
	 	if _, err := fmt.Scanf("%d", &guessedNumber); err != nil {
	   		fmt.Println("Slacker! That's not a numner ;) ")
	   		
	 	}else{
	 			if p.lives <= 0 { //if not more lives we terminate the game
	 				fmt.Printf("GAME OVER! You don't have more lives to play\n")
	 				return p, g
	 			}else if guessedNumber ==  randoms[i] { //if guessedNumber nad random are equals
	 				p.lives += 1 //incresing live by 1
	 				g.level += 2 //incresing level by 2
	 				p.pointsWon += 10 //incresing points by 10
	 				fmt.Printf("PERFECT!, You guessed it! Random#:%d, your#: %d, skipped 1 level {lives:%d, level:%d, points:%d} \n", randoms[i], guessedNumber, p.lives, g.level, p.pointsWon)
	 				//making sure that guessNumber and random are less or greater then numberToGuess
	 			}else if guessedNumber <= g.numberToGuess && randoms[i] <= g.numberToGuess || guessedNumber >= g.numberToGuess && randoms[i] >= g.numberToGuess{
	 				g.level += 1 // incresing level by 1
	 				p.pointsWon += 5 //incresing points by 5
	 				fmt.Printf("Great!, Random#:%d, your#: %d, {lives:%d, level:%d, points:%d} \n", randoms[i], guessedNumber, p.lives, g.level, p.pointsWon)
	 			}else{ //if player didn't match we then
	 				p.lives -= 1  //remove 1 live
	 				if p.lives <= 0 { //if 0 lives ->GAME OVER<-
		 				fmt.Printf("=============================================\n")
		 				fmt.Printf("GAME OVER! You don't have more lives to play\n")
		 				fmt.Printf("=============================================\n")
		 				fmt.Printf("================ ~~~SCORE~~~ ===============\n")
		 				fmt.Printf("{name: %s, age: %d, pointsWon:%d, lives: %d, level:%d} \n", p.name, p.age, p.pointsWon, p.lives, g.level)
		 				fmt.Printf("================ ~~~SCORE~~~ ===============\n")
		 				return p, g
		 			}
	 				fmt.Printf("Sorry!, Random#:%d, your#: %d, You just lost one live, you now have %d live(s)\n", randoms[i], guessedNumber, p.lives)
	 			} 

	 	}
	}

	fmt.Printf("=============================================\n")
	fmt.Printf("================ ~~~SCORE~~~ ===============\n")
	fmt.Printf("{name: %s, age: %d, pointsWon:%d, lives: %d, level:%d} \n", p.name, p.age, p.pointsWon, p.lives, g.level)
	fmt.Printf("================ ~~~SCORE~~~ ===============\n")
	return p, g  //we return p and g as we want to save these record in a database.
}




func main() {
	
	//creating game struct and assigning values to the fields
	var game = Game{true, 5, 0}
	
	//printing game header
	printHeader(game.numberToGuess)

	//asking name/age of player
	name, age := ask()
	if name == "errAgeNotNumber"{
		return
	} 
	
	//creating player struct once we have name and age, assigning these values to the repective fields
	var player = Player{name, age, 0, 2}
	
	//Passing the structs as pointers so that we can modify them in the play game
	play(&player, &game)




}