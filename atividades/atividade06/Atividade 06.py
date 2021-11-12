"""
Nome: Erick Correia Silva 
Data 12/11/2021

Rotinas de Análise Sintática Descendente Recursiva em Python
"""
import funcoes_atividade_05 as funcoes #importa para usar as funçoes da atividade passada (funçao lex)

# Constantes 
LETTER = 0
DIGIT = 1
UNKNOWN = 99
INT_LIT = 10
IDENT = 11
ASSIGN_OP = 20
ADD_OP = 21
SUB_OP = 22
MULT_OP = 23
DIV_OP = 24
LEFT_PAREN = 25
RIGHT_PAREN = 26
EOF = -1 #arquivo

# Variaveis globais do código inicializadas 
charClass = 'a'  #Categorias (char ou digitos )
lexeme = [] #Vetor com lexeme
nextChar = 'a'
lexLen = 0 #Comprimento
token = 0
nextToken = 0


def expr():
#Analisa sintaticamente cadeias na linguagem gerada pela regra:
# <expr> -> <term> {(+ | -) <term>}
    print("Enter <expr> \n")
    term() #Analisa sintaticamente o primeiro termo

    while (nextToken == ADD_OP | nextToken == SUB_OP):
        funcoes.lex()
        term()

    print("Exit <expr>\n")


def term():
#Analisa sintaticamente cadeias na linguagem gerada pela regra:
# <term> -> <factor> {(* | /) <factor>)
    print("Enter <term> \n")
    factor()

    while (nextToken == MULT_OP | nextToken == DIV_OP):
        funcoes.lex()
        factor()

    print("Exit <term>\n")


def factor():
#Analisa sintaticamente cadeias na linguagem gerada pela regra:
# <factor> -> id | int_constant | (<expr),
    print("Enter <factor>\n")

    if (nextToken == IDENT | nextToken == INT_LIT):
    #Obtém o próximo token *
        funcoes.lex()
    else:
        if (nextToken == LEFT_PAREN):
            funcoes.lex()
            expr()
            if (nextToken == RIGHT_PAREN):
                funcoes.lex()
            else:
                print("ERRO")
        else:
            print("ERRO")


        
