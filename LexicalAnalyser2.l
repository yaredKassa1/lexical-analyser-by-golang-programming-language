%{
#include <stdio.h>
#include <stdlib.h>

int yylineno = 1;  // Initialize line number to 1
FILE *input_file;   // Pointer to the input file
%}

DIGIT            [0-9]+
FLOAT            [0-9]+\.[0-9]+
OPERATOR         [+ - / % = += -= = /= == != > < >= <= && || ! & | ^ << >> - !]
IDENTIFIER       [a-zA-Z_][a-zA-Z0-9_]*
KEYWORD          (break|const|continue|crate|else|enum|extern|false|fn|mut|in|let|for|if|impl|loop|match|move|mod|while|where|use|trait|unsafe|true|type|struct|static)
DELIMITER        [(){}[];.::=>...?]
WHITESPACE       [ \t\n\r]+

%%

{KEYWORD}                 { printf("| %-16s | %-11s | %-12d |\n", "KEYWORD", yytext, yylineno); }
{DELIMITER}               { printf("| %-16s | %-11s | %-12d |\n", "STRING LITERAL", yytext, yylineno); }
{DIGIT}                   { printf("| %-16s | %-11s | %-12d |\n", "DIGIT", yytext, yylineno); }
{FLOAT}                   { printf("| %-16s | %-11s | %-12d |\n", "FLOAT", yytext, yylineno); }
{IDENTIFIER}              { printf("| %-16s | %-11s | %-12d |\n", "IDENTIFIER", yytext, yylineno); }
{OPERATOR}                { printf("| %-16s | %-11s | %-12d |\n", "OPERATOR", yytext, yylineno); }
{DIGIT}{IDENTIFIER}       { printf("| %-16s | %-11s | %-12d |\n", "INVALID EXPRESSION", yytext, yylineno); }
{WHITESPACE}              { if (yytext[0] == '\n') yylineno++; }  // Increment line number on newline
"//".*                    { /* Ignore single-line comments */ }
"/\\*"([^*]|\*+[^*/])*\*+"/"   { /* Ignore multi-line comments */ }

. { printf("| %-16s | %-11s | %-12d |\n", "UNKNOWN EXPRESSION", yytext, yylineno); }

%%

/* Main function that handles file reading and lexer initialization */
int main(int argc, char **argv) {
    if (argc != 2) {
        printf("Usage: %s <input_file>\n", argv[0]);
        return 1;
    }

    input_file = fopen(argv[1], "r");
    if (input_file == NULL) {
        perror("Error opening file");
        return 1;
    }

    yyin = input_file;

    printf("| %-16s | %-11s | %-12s |\n", "Token Type", "Value", "Line Number");
    printf("|---------------|-------------|-------------|\n");

    yylex();  // Call the lexer to process the file

    fclose(input_file);  // Close the input file
    return 0;
}

/* Custom wrap function for Flex */
int yywrap() {
    return 1;
}
