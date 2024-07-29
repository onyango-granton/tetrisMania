#!/bin/bash

git add finalProject/utils/byteToInt.go


Refactor: Improved error handling and added documentation for byteToInt function in utils package.

Refactor: Added byteToInt function to convert byte characters to integers. Introduced error handling for unsupported characters and included a debug print statement for error context.

Refactor ErrorHandling function to validate command-line arguments and filename format

feat(utils): Implement Tetromino placement and grid completion functions

Refactor: Improved InitGrid function to initialize a square grid with default value "*"

feat(utils): Implement stringToIntSlice function for converting 4-char strings to int slices

feat: Enhance Tetromino processing and validation in utils package
- Added `sliceIsEmpty` function to check if a slice of integers is empty.
- Improved `TetroGroupFunc` to handle text file reading, processing, and validation of Tetromino shapes.
- Implemented Windows-specific newline handling for cross-platform compatibility.
- Introduced map-based storage for Tetromino shapes and names during processing.
- Enhanced error handling and logging for file reading and shape validation.
- Calculated grid size required for Tetrominoes using square root and ceiling functions.
- Refactored return statement to use a trimming function for final Tetromino list.


Refactor tetromino trimming functions to improve readability and efficiency.
    Simplified trimTetrominoListFunc by removing commented-out code and unnecessary variable declarations.
    Enhanced trimTetromino function to handle tetrominos of dimensions 3xN and 2xN more efficiently by removing empty rows and columns.
    Ensured consistent trimming logic across different tetromino dimensions for better performance and maintainability.


Refactor Go utility functions for tetromino validation:

    Introduced allOne function to check if two integers are both 1, simplifying the logic in isSurroundedByOnes.
    Enhanced isSurroundedByOnes to check if an element in a 2D array is surrounded by 1s horizontally and vertically.
    Implemented fullyConnected to verify if a tetromino shape (2D slice) is fully connected with each '1' directly connected to another '1' horizontally or vertically.
    Created isValidTetro to validate a tetromino shape, ensuring it has exactly 4 '1's, is fully connected, and does not have more than 4 borders surrounded by '1's.
