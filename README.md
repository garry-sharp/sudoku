# Recursive (Backtracking) sudoku solver in golang

The solution is a solution to the problem posed in [Leetcode Problem](https://leetcode.com/problems/sudoku-solver) on their site.

The challenge is to take a 2d byte array representing a sudoku and solve it.

The rules of sudoku are relatively straight forward. A partially filled 9x9 grid of numbers is presented and the goal is to fill in the grid such that _all_ of the following contraints hold true.

- Each column of 9 numbers must not have any single repreating number.
- Each row of 9 numbers must not have any single repeating number.
- Each 3x3 mini grid must not have any single repeating number. A 3x3 mini grid is set out from top corner. Such that 9 mini grids exist within the larger 9x9 grid.

# Approach

At first as a fan of Sudoku puzzles I decided to try and build some strategies into the sudoku game and go through the grid cell by cell and reduce the number of possibilities found. For instance if there are 2 cells in a row with only possible values of 2 and 4, then even if you don't know which is which, you do know that no other cell in that row can be either of those 2 values.

There are many techniques and strategies which are often used in conjuction with one another in Sudoku to come up with an answer.

This however was too cumbersome for a coding exercise and each strategy would involve iterating through the grid in a different way each time. So I decided to take a more logical approach but with one timesaver.

The idea is to as far as reasonably (and easily) possible, go through the grid once and apply simple logic to fill in some squares. For instance, if there is only 1 possible value but it wasn't filled in, we will fill it in.

Afterwards we hit the problem by using backtrack recursion. For an unknown square we will get a list of possible values and "guess" the value of the cell and then move onto the next cell. This continues recursively until one of the following happens.

1. The recursion terminates (singalling a successful solution)
2. No more solutions are found.

In the latter case the calling parent function will try its next possible value. If all possibilities are exhaused then the functions continue upwards with possibilities being recursed through and deeper cells being reexamined as new possibilities are added.

# Style

I love working with objects as much as I can. I turned a 2d byte array into a struct with Cells, Grids, Columns and Rows all pointing at each other. I find it makes the code more readable.

Additionally there is a print function which I used to visually represent a sudoku grid.

Due to it being a Leetcode problem all code is held within 1 file. To try it on Leetcode itself copy everything except main.

# Error Handling

The code is not tested with solutions that are presented invalid or with puzzles that have multiple valid soltions.
