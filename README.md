GuessNextNumber: is a game idea created by Avelino Puello for the only purpose of learning/practising the GO programming language.



Specs for Version: 0.0.1

1) The game: GuessNextNumber consists in guessing if the next random number thrown by the system is lower or greater than 5
	a) When game starts it should introduced the rules of the game mentioned below
	b) should ask for your name and age before player starts guessing numbers 
	c) you start with 2 lives
	d) iF you guess right, you go to next level.
	e) if your guess is wrong you lose one live.
	f) if you match the random number, you get a live and also skip a level
	g) if you lose all your lives then the game is over
	j) once you reach the last level(less keep it to 5 just to save time when testing) you get a congrats message stating your score
	k) Should give you output on every guess, so that players can see the random number, level, lives they have and see their progress.
	l) Once game is over, it gives you an JSON output format like below
		{"name":"avel", "age":32, "level":5, "lives":5, "points":30},
		
