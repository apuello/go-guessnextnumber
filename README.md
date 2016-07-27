**#GuessNextNumber**
 is a game idea created by Avelino Puello for the only purpose of learning/practising the [GO](https://golang.org/) programming language.

The Game consists in guessing if the next random number thrown by the system is lower or greater than 5


Specs for Version: 0.0.1

1. When game starts it should introduced the rules of the game mentioned below
2. should ask for your name and age before player starts guessing numbers 
3. you start with 2 lives
4. iF you guess right, you go to next level.
5. if your guess is wrong you lose one live.
6. if you match the random number, you get a live and also skip a level
7. if you lose all your lives then the game is over
8. once you reach the last level(less keep it to 5 just to save time when testing) you get a congrats message stating your score
9. Should give you output on every guess, so that players can see the random number, level, lives they have and see their progress.
10. Once game is over, it gives you an JSON output format like below
	{"name":"avel", "age":32, "level":5, "lives":5, "points":30},
		

I developed this game after reading taking the [Tour of Go](https://tour.golang.org) -  **I highly recommend it!**

1. In this v0.0.1 you can find the use of: 
	1. Packages 
	2. Variables 
	3. Functions 
	4. Flow Control - For(ever) 
	5. Structs 
	6. Pointers 
	7. Arrays 
	8. Functions 
	9. Slices
	10. Packages
	11. and more! 

**#TODO: Version: 0.0.2**

1. Save Player's info in a DB(Probably PostgreSQL)
2. Create a web interface
