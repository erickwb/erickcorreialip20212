# Correção: 1,5
# Parece estar correto, mas quando executo não produz saída, nem mensagem de erro.

'''
Nome: Erick Correia Silva 
Dia: 04/11/2021

analisador léxico da seção 4.2 do livro texto
'''

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

# Funçoes 
#função para pesquisar operadores e parênteses e devolver o token 
def lookup(ch):
    if ch == '(':
        addChar()
        nextToken = LEFT_PAREN
    elif ch == ')':
        addChar()
        nextToken = RIGHT_PAREN
    elif ch == '+' :
        addChar()
        nextToken = ADD_OP
    elif ch == '-' :
        addChar()
        nextToken = SUB_OP
    elif ch == '*' :
        addChar()
        nextToken = MULT_OP
    elif ch == '/' :
        addChar()
        nextToken = DIV_OP
    else:
        addChar()
        nextToken = EOF

#função para adicionar nextChar ao lexema
def addChar():
    if (lexLen <= 98):
        lexeme.append(nextChar)
        lexeme.append(0)
    else:
        print("Error - lexeme is too long \n")


#função para obter o próximo caractere de inserir e determinar sua classe de caractere
def getChar(in_fp):
    #
    nextChar = in_fp.read(1)
    if (nextChar != EOF):
        if (nextChar.isalpha() == True):
            charClass = LETTER
        elif (nextChar.isdigit() == True):
            charClass = DIGIT
        else:
            charClass = UNKNOWN
    else:
        charClass = EOF

#função para chamar getChar até que retorna um caractere diferente de espaço em branco
def getNonBlank():
    while (nextChar.isspace()):
        getChar()

#um analisador léxico simples para aritmética expressões
def lex(): 
    lexLen = 0
    getNonBlank()

    if charClass == LETTER:
        addChar()
        getChar()
        while (charClass == LETTER | charClass == DIGIT):
            addChar()
            getChar()
        nextToken = IDENT

    elif charClass == DIGIT:
        addChar()
        getChar()
        while (charClass == DIGIT):
            addChar()
            getChar()
        
        nextToken = INT_LIT
        
    elif charClass == UNKNOWN:
        lookup(nextChar)
        getChar()

    elif charClass == LETTER:
        nextToken = EOF
        lexeme[0] = 'E'
        lexeme[1] = 'O'
        lexeme[2] = 'F'
        lexeme[3] = 0

    #fim do switch

    print("Next token is: %d, Next lexeme is %s\n",nextToken, lexeme)
    return nextToken


def main():
    in_fp = open('front.in', 'r')

    if(in_fp == None):
        print("ERROR - cannot open front.in \n")
    else:
        getChar()

        while (True):
            lex()

            if not (nextToken != EOF):
                break



