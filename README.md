## Tetris-Optimizer
This project is designed to read a list of tetrominoes from a text file and assemble them into the smallest possible square. The program is written in Go and utilizes a set of utility functions to handle file reading, tetromino validation, and grid manipulation.
```
Usage: go run . <filename.txt>
Example: go run . tetris.txt

```
## NOTE
1. This program is part of the **[Zone01 Cirriculum](https://github.com/01-edu/public)**
## Usage
To use the program, follow these steps:
1. Clone this repository to your local machine using the comand below.
``` bash
git clone https://learn.zone01kisumu.ke/git/gonyango/tetris-optimizer.git
```

2. Navigate to the directory where the repository is cloned.
```bash
cd tetris-optimizer
```

3. Initialize your module to get the required dependancies :Run the command below to initialize your module
```bash
go mod init tetris-optimizer
```
4. Run the program from the command line, providing a string as an argument and file to store art justify


## Example 1
- Sample text file format
```
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....
```



## Example 2

```console
go run . sample.txt
```
- Output
```
C C C * D D 
F F C D D A 
* F F G G A 
* H H H G A 
E E H * G A 
E E B B B B 
```

## Dependencies
This program requires Go (Golang) to be installed on your system. You can download and install it from the [official Go website](https://golang.org/dl/).

## Contributing
Contributions to this project are welcome! If you'd like to contribute, please fork the repository and submit a pull request with your changes.

## Contributors
<body>
<div style="display: flex !important; justify-content: center !important;">
    <div style="margin: 10px;">
        <img src="images/gonyango.png" style="border-radius: 50% !important; width: 150px !important; height: 150px; !important" alt="Granton">
        <p style="text-align: center;">Granton</p>
    </div>
</div>
</body>


## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.



