//==============================================================================
//
// Title:		game.go
// Purpose:		game.go is the heart of the game, here is where everything happens.
//
// Created on:	7/26/2016 at 5:37:47 PM by Avelino Puello
// Copyright:	. All Rights Reserved.
//
//==============================================================================
/*

-------------------------------------------------------------------------
*/

package game 

import(
	"fmt"
	"math/rand"
	player "../../structs/player" //using alias - player
	game "../../structs/game"   //using alias - game
	print "../print"
)


//play func is hte heart of this program
//here we receive 2 pointers(we need pointers as we will be modifying the properties of these)
//we will also return these 2 pointers as we will be saving the info in a DB
//maybe postgreSQL?
func Play(p *player.Player, g *game.Game) (*player.Player, *game.Game){
	var guessedNumber int
	randoms := giveMeTheRandoms() //getting random number []int
	for i := 0; i < g.NumberToGuess; i++ { //looping through game max number to guess
		
	 	print.PrintNextNumber()
	 	//Here we make sure user is entering a number(we could do something similar with the name above)
	 	if _, err := fmt.Scanf("%d", &guessedNumber); err != nil {
	   		print.PrintSlacker()
	 	}else{
	 			if p.Lives <= 0 { //if not more lives we terminate the game
	 				print.PrintGameOver()
	 				return p, g
	 			}else if guessedNumber ==  randoms[i] { //if guessedNumber nad random are equals
	 				p.Lives += 1 //incresing live by 1
	 				g.Level += 2 //incresing level by 2
	 				p.PointsWon += 10 //incresing points by 10
	 				print.PrintPerfect(randoms[i], guessedNumber, p.Lives, g.Level, p.PointsWon)
	 				//making sure that guessNumber and random are less or greater then numberToGuess
	 			}else if guessedNumber <= g.NumberToGuess && randoms[i] <= g.NumberToGuess || guessedNumber >= g.NumberToGuess && randoms[i] >= g.NumberToGuess{
	 				g.Level += 1 // incresing level by 1
	 				p.PointsWon += 5 //incresing points by 5
	 				print.PrintGreat(randoms[i], guessedNumber, p.Lives, g.Level, p.PointsWon)
	 			}else{ //if player didn't match we then
	 				p.Lives -= 1  //remove 1 live
	 				if p.Lives <= 0 { //if 0 lives ->GAME OVER<-
		 				print.PrintEndOfGame(p.Name, p.Age, p.PointsWon, p.Lives, g.Level)
		 				return p, g
		 			}
	 				print.PrintSorry(randoms[i], guessedNumber, p.Lives)
	 			} 

	 	}
	}

	print.PrintScore(p.Name, p.Age, p.PointsWon, p.Lives, g.Level)
	return p, g  //we return p and g as we want to save these record in a database.
}


//I wasn't able to use the rand method and return random numbers
//what i did for now was to create an array of 10 int(s) from 1 to 0
//using the r.Perm(int) function
//example [8,6,0,1,7,2,5,4,9,5]
//finally, I return the []int
//*****Notice how this function is not capitalized - this means we don't want to export it.
func giveMeTheRandoms() []int{
	r := rand.New(rand.NewSource(99))
	ranNumber := r.Perm(10)
	//Slicing the first 5 random numbers
	var takeFirstFive []int = ranNumber[:6]
	return takeFirstFive
}
