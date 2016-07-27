//==============================================================================
//
// Title:		print.go
// Purpose:		prints all the interactions to the user
//
// Created on:	7/26/2016 at 5:37:47 PM by Avelino Puello
// Copyright:	. All Rights Reserved.
//
//==============================================================================
/*

-------------------------------------------------------------------------
*/

package print

import(
	"fmt"
)

/* Here we print the intro of the game */

func PrintHeader(numberToGuess int){
 fmt.Println("====================================================================")
 fmt.Println("======================~~ Avelino RULLETE ~~=========================")
 fmt.Println("Welcome to the game of Avelino and become millonaire in one game")
 fmt.Printf("Guess if number is greater or less than %d and win instantly ! \n", numberToGuess )
 fmt.Println("====================================================================")
}

func PrintWhatsYourName(){
	fmt.Print("What is your name? : ")
}

func PrintHowOldAreYou(){
	fmt.Print("How old are you? : ")
}
func PrintNextNumber(){
	fmt.Print("Guess Next Number?: ")
}

func PrintSlacker(){
	fmt.Println("Slacker! That's not a numner ;) ")
}

func PrintGameOver(){
	fmt.Printf("GAME OVER! You don't have more lives to play\n")
}

func PrintPerfect(randomNumber , guessedNumber, lives, level, pointsWon int){
	fmt.Printf("PERFECT!, You guessed it! Random#:%d, your#: %d, skipped 1 level {lives:%d, level:%d, points:%d} \n", randomNumber, guessedNumber, lives, level, pointsWon)
}


func PrintGreat(randomNumber, guessedNumber, lives, level, pointsWon int){
	fmt.Printf("Great!, Random#:%d, your#: %d, {lives:%d, level:%d, points:%d} \n", randomNumber, guessedNumber, lives, level, pointsWon)
}

func PrintEndOfGame(name string, age, pointsWon, lives, level int){
	fmt.Printf("=============================================\n")
	fmt.Printf("GAME OVER! You don't have more lives to play\n")
	fmt.Printf("=============================================\n")
	fmt.Printf("================ ~~~SCORE~~~ ===============\n")
	fmt.Printf("{name: %s, age: %d, pointsWon:%d, lives: %d, level:%d} \n", name, age, pointsWon, lives, level)
	fmt.Printf("================ ~~~SCORE~~~ ===============\n")
}

func PrintSorry(randomNumber, guessedNumber, lives int){
	fmt.Printf("Sorry!, Random#:%d, your#: %d, You just lost one live, you now have %d live(s)\n", randomNumber, guessedNumber, lives)
}

func PrintScore(name string, age, pointsWon, lives, level int){
	fmt.Printf("=============================================\n")
	fmt.Printf("================ ~~~SCORE~~~ ===============\n")
	fmt.Printf("{name: %s, age: %d, pointsWon:%d, lives: %d, level:%d} \n", name, age, pointsWon, lives, level)
	fmt.Printf("================ ~~~SCORE~~~ ===============\n")
}