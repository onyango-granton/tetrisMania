# Tetris Optimizer
## Description
This is  a program that receives only one argument, a path to a text file which will contain a list of [tetrominoes](https://en.wikipedia.org/wiki/Tetromino) and assembles them in order to create the smallest square possible.

### Example of a text File
```
#...
#...
#...
#...

....
....
..##
..##
```
### Expected Output for the above Example
```
ABB.
ABB.
A...
A...
```
## Guidelines for tetromino Text files Format

1. The file **must** have a .txt extension e.g text.txt, mug.txt
2. The tetrominos **must** have a length of 4 i.e
```
...# 1
...# 2
...# 4
...# 4
```
3. They **must** be separated by an empty line i.e
```
...#
...#
...#
...#

...#
...#
...#
...#
```
4. The length of each string **must** be 4
5. There **must** only be 4 hashes(#) to represent a tetromino


## Installation

- This application requires Go (golang) 1.18 or higher to run. You can get it [here](https://go.dev/doc/install)

- To clone and run this program, you'll need **Git** installed on your computer.

- From the **command line**,

```Bash
# clone this repository
$ git clone https://learn.zone01kisumu.ke/git/aosindo/tetris-optimizer.git

# go into the repository
$ cd tetris-optimizer

# open a code editor like VS Code
$ code .
```
## Usage
- Once the program has been installed, add your desired text file with the valid format to the directory containing ``main.go``

- Run the program using the name of the text file added, like this:

    `$ go run . text.txt`

## Notion

[Fillit: Solving for the Smallest Square of Tetrominoes](https://medium.com/@bethnenniger/fillit-solving-for-the-smallest-square-of-tetrominos-c6316004f909)

## Enjoy!


## Authors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://github.com/andyosyndoh>
            <img src=https://lh3.googleusercontent.com/a/ACg8ocLUKAW3QwBqLDqDcmkFTC3wmCPq0dd25wVFn3CPEkCfhQQme9Lx=s288-c-no width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=Andrew/>
            <br />
            <sub style="font-size:14px"><b>Andrew Osindo</b></sub>
        </a>
    </td>
    
</tr>
</table>