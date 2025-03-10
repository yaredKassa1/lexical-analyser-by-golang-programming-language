# lexical-analyser-by-golang-programming-language

## Project Description  
This project implements a lexical analyzer (lexer) using the Go programming language.  
The lexer reads source code from a file, tokenizes it into meaningful components (keywords, operators, delimiters, identifiers, and literals), and reports any lexical errors encountered during the analysis.  
It serves as a foundational tool for building compilers or interpreters, helping to understand the structure and components of source code.  

## Instructions for Running the Program(lexicalAnalyser.go)  
Install Go: Ensure that you have Go installed on your machine. You can download it from the official Go website:[ golang.org.](https://go.dev/)  
Clone or Download the Project: Clone the repository or download the project files to your local machine.  
Prepare Input File: Create or modify a text file containing the Go source code that you want to analyze. Name it input.txt.  
## Run the Lexer:  
Open a terminal or command prompt.  
Navigate to the directory where the lexicalAnalyser.go and input.txt files are located.  
## Execute the following command:
bash  
go run lexicalAnalyser.go input.txt 
## Instruction for running the program(LexicalAnalyser2.l) 
### Generate the Lexer:  
Open a terminal and navigate to the directory containing the LexicalAnalyser2.l file.  
Run the following command to generate the C source file:  
flex LexicalAnalyser2.l  
### Compile the Generated C File:  
This will generate a file named lex.yy.c. Compile it using gcc:  
gcc lex.yy.c -o lexer  
### Run the Lexer:  
Now, run the lexer with the input.txt file as an argument:  
./lexer input.txt  
View Output: After running the command, the terminal will display the tokens identified in the source code along with their types, line numbers, and column numbers. Any errors will also be reported.  

## Group Members

| Name           |  ID |
|----------------|------------|
| Abel Assefa  | 1300419    |
| Betsegaw Muluneh     | 1300580    |
| Geleta Bekele  | 1301323    |
| Yared Kassa   | 1303033    |
| Zekarias Woreket   | 1306501    |
