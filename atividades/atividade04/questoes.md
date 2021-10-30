# Atividade 04

Erick Correia Silva

## Questão 01


R) Usando a sentença a + b + c ,podemo encontra duas derivações diferente (ou seja para a mesmam sentença temos duas arvores de análise)

Primeira forma: (derivação a esquerda)

 < S > -> < A >
        
       -> < A > + < A >
       
       -> < A > + < A > + < A >
       
       -> < id > + < A > + < A >
       
       -> a + < A > + < A >
       
       -> a + < id > + < A >
       
       -> a + b + < A >
       
       -> a + b + < id >
       
       -> a + b + c
       
Segunda forma: (derivação a esquerda)

 < S > -> < A >
        
       -> < A > + < A >
       
       -> < id > + < A >
       
       -> a + < A >
       
       -> a + < A > +  < A >
       
       -> a + < id > + < A >
       
       -> a + b + < A >
       
       -> a + b + < id >
       
       -> a + b + c
       
Arvores de derivação: 

![Questão 01 ativiadade 04 ](https://user-images.githubusercontent.com/39568346/139163505-4afa0f89-57e4-43c3-afa0-eb24f08a5a18.png)

## Questão 02

R) Transformando a linguagem em EBNF temos:

< program > -> begin  < stmt_list >  end

< stmt_list > -> < stmt >  | <stmt> ; <stmt_list>  

< stmt > -> < var > = < expression >
 
< expression > -> < var > {  var ( + | - ) < var > }
 
< var > -> A | B | C 


## Questão 03

R) Transformando a linguagem em EBNF temos:

< assign > -> < id > = < expr > 
 
< id > -> A|B|C
 
< expr > -> < id > { < expr> ( + | * ) < expr> }


## Questão 04
 
obs: Formato da demonstração da gramática de atributos 
 
 ** Regra Sintática
 
 ** Regra Semântica
 
 R)
 
Regra Sintática: < assign > -> < var > = < expr > 

Regra Semântica:  ( < expr >.expected_type <- < var >.actual_type )
 
Regra Sintática: < expr > -> < var >[2] + < var >[3]
 
Regra Semântica: < var >[2].env <- < expr >.env
      
                 < var >[3].env <- < expr >.env
                 
                 < expr >.actual_type <- < var >[2].actual_type
                  
                 Predicado: < var >[2] .actual_type = < var >[3].actual_type
                  
Regra Sintática: < expr > -> < var >
                  
Regra Semântica: < expr >.actual_type <- < var >.actual_type
                  
Predicado: < expr >.actual_type == < expr >.expected_type
                  
Regra sintática: < var > -> A | B | C 
                  
Regra semântica: < var >.actual_type <- look-up( < var >.string )

                  
                  
                  

 
 


