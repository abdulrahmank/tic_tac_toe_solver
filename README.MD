**Tic Tac Toe Solver**

This is a tic tac toe solver backend in Golang. It doesn't use Minimax algorithm or any backtracking algorithm for 
solving. It just uses the count of the characters in the board and decide which moves to take.

Note:
1. Player is always considered as "X" and CPU as "O".
2. Game status sent as a part of response reflects the status from the CPU point of view.
3. This works out of the box only for 3 X 3 board, 
changes have to be done to make it work for bigger boards.
