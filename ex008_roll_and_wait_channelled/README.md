Change your solution for `ex007_roll_and_wait` to cater for the following:

* Create a timer that keeps track of when time is up (don't use `NewTimer` or `time.After`) 
* Players send the number they rolled to the main program, which keeps a track of each player's total
* Players print out the status after each roll e.g. "rolled a x, waiting y seconds..."
* When the timer finishes, print out the totals and declare the winner
