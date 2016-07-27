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
    playerStruct "./structs/player"
	gameStruct "./structs/game"
	game "./pkg/game"
	ask "./pkg/ask"
	print "./pkg/print"
)


func main() {
	
	//creating game struct and assigning values to the fields
	var g = gameStruct.Game{true, 5, 0}
	
	//printing game header
	print.PrintHeader(g.NumberToGuess)

	//asking name/age of player
	name, age := ask.Ask()
	if name == "errAgeNotNumber"{ //if player try to play with us entering a string instead of int we exit the game!
		return
	} 
	
	//creating player struct once we have name and age, assigning these values to the repective fields
	var p = playerStruct.Player{name, age, 0, 2}
	
	//Passing the structs as pointers so that we can modify them in the play game
	game.Play(&p, &g)




}