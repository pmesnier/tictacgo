# tictacgo
another variation of tic-tac-toe implemented in go. My goal is to use thi as a way to become familiar with the language feaatureas and idions. the actual game play is a secondary concern. I intend to add the means to publish a web interface so that i can support multiple concurrent games, historical data archiving and retrieval and whatever lese I can think of.

##building theprogram
I am still learning the correct ay to manage the layout and so I am leaning on the goland tool to help me with the build process. 

##playing the game. 
Running the main package from tictacgo/main, yields a request for two player names then alternates quiries for input from player 1 (the X player) and player 2 (the O player) unoccupied cells are identified by a number, 1 to 9. 


'
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
---+---+---
'

paly continues until a player achieves a win or all cells are occupied.
