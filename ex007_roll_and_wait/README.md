"Roll and Wait" a game played with dice and a watch.

1. You start with a score of 0
2. Roll a 6-sided die
3. Let's say you rolled an 'n'. Add n to your score
4. Wait (6-n) seconds - e.g. 2 seconds if you rolled a 4
5. Roll again and keep playing

Write a program that implements the game for two players, Ed and Natasha, where the die roll is randomised.
The game should end 30 seconds after it starts.

**Do not use channels**

Expected output should look something like:

```
Ed (0) rolled a 5, waiting 1 sec
Natasha (0) rolled a 1, waiting 5 sec
Ed (5) rolled a 2, waiting 4 sec
Ed (7) rolled a 2, waiting 4 sec
Natasha (1) rolled a 5, waiting 1 sec
Natasha (6) rolled a 4, waiting 2 sec
Natasha (10) rolled a 2, waiting 4 sec
Ed (9) rolled a 3, waiting 3 sec
Ed (12) rolled a 5, waiting 1 sec
Natasha (12) rolled a 5, waiting 1 sec
Ed (17) rolled a 5, waiting 1 sec
Natasha (17) rolled a 2, waiting 4 sec
Ed (22) rolled a 4, waiting 2 sec
Ed (26) rolled a 6, rolling again immediately
Ed (32) rolled a 3, waiting 3 sec
Natasha (19) rolled a 5, waiting 1 sec
Natasha (24) rolled a 1, waiting 5 sec
Ed (35) rolled a 1, waiting 5 sec
Natasha (25) rolled a 5, waiting 1 sec
Ed (36) rolled a 3, waiting 3 sec
Natasha (30) rolled a 1, waiting 5 sec
Ed (39) rolled a 2, waiting 4 sec
Natasha (31) rolled a 2, waiting 4 sec
Time up!
```
