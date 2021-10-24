# Atividade 03

Erick Correia Silva

## Questão 01
## A)  A = A + (B + C) "derivação mais à esquerda"
< assign > -> < id > = < expr >

           -> A = < expr > 
           
           -> A = < expr > + < term > 
           
           -> A = < term > + < term > 
           
           -> A = < factor > + < term > 
           
           -> A = < id > + < term > 
           
           -> A = A + < term > 
           
           -> A = A + < factor > 
           
           -> A = A + ( < expr > ) 
           
           -> A = A + ( < expr > + < term > ) 
           
           -> A = A + ( < term > + < term > )
           
           -> A = A + ( < factor > + < term > ) 
           
           -> A = A + ( < id > + < term > ) 
           
           -> A = A + ( B + < term > ) 
           
           -> A = A + ( B + < factor > ) 
           
           -> A = A + ( B + < id > ) 
           
           -> A = A + ( B + C ) 
         
![Questão 1A Atividade 03 LIP](https://user-images.githubusercontent.com/39568346/138574089-35f3bd80-b799-4a49-b098-25369e105f33.png)

## B)  A = B * (C * (A + B)) "derivação mais à esquerda"
< assign > -> < id > = < expr >

           -> A = < expr >
           
           -> A = < term >
           
           -> A = < term > * < factor >
           
           -> A = < factor > * < factor >
           
           -> A = < id > * < factor >
           
           -> A = B * < factor >
           
           -> A = B * ( < expr > )
           
           -> A = B * ( < term > )
           
           -> A = B * ( < factor > * < factor > )
           
           -> A = B * ( < id > * < factor > )
           
           -> A = B * ( C * < factor > )
           
           -> A = B * ( C * ( < expr > ) )
           
           -> A = B * ( C * ( < expr > + < term > ) )
           
           -> A = B * ( C * ( < term > + < term > ) )
           
           -> A = B * ( C * ( < factor > + < term > ) )
           
           -> A = B * ( C * ( < id > + < term > ) )
           
           -> A = B * ( C * ( A + < term > ) )

           -> A = B * ( C * ( A + < factor > ) )
          
           -> A = B * ( C * ( A + < id > ) )
                     
           -> A = B * ( C * ( A + B ) )


![Quastão 1B Atividade 03](https://user-images.githubusercontent.com/39568346/138574218-706eb7b5-5db0-4209-b240-662166d39413.png)

## Questão 02
## Resposta) (a 23 (m x y)) "derivação mais à esquerda"
## obs: como a liguagem não possui um tabela de simbolos, o final desta questão deve ficar (identificador numero (  identificador identificador identificador ) 
      
< lexp >   -> < lista >

           -> (< lexp-seq >)
           
           -> (< lexp-seq > < lexp >)
           
           -> (< lexp-seq > < lista >)
           
           -> (< lexp-seq > (< lexp-seq >))
           
           -> (< lexp-seq > (< lexp-seq > < lexp >))
           
           -> (< lexp-seq > (< lexp-seq > < atomo >))
           
           -> (< lexp-seq > (< lexp-seq >  identificador ))
           
           -> (< lexp-seq > (< lexp-seq > < lexp > identificador ))
           
           -> (< lexp-seq > (< lexp-seq > < atomo > identificador ))
           
           -> (< lexp-seq > (< lexp-seq > identificador identificador ))
           
           -> (< lexp-seq > (< lexp > identificador identificador ))
           
           -> (< lexp-seq > (< atomo > identificador identificador ))
           
           -> (< lexp-seq > (identificador identificador identificador ))
           
           -> (< lexp-seq > < lexp > (identificador identificador identificador ))
           
           -> (< lexp-seq > < atomo > (identificador identificador identificador ))
           
           -> (< lexp-seq > numero (identificador identificador identificador ))
           
           -> (< lexp > umero (identificador > < identificador > < identificador >))
           
           -> (< atomo > numero (identificador identificador identificador ))
           
           -> (identificador  numero (identificador identificador identificador ))
           
![Questão 2 Ativiadade 03 LIP](https://user-images.githubusercontent.com/39568346/138575290-d910c8bc-5a4c-4686-b65e-b5336a743a5a.png)
           
           
## Questão 03
## Resposta)
 Switch em C exemplo
 
 switch (variável){
 
   case constante1:
   
     Instruções;
     
   break;

   case constante2:
   
     Instruções;
     
   break;

   default
   
     Instruções;
}

## BNF Switch da linguagem C

## obs: <stmt> esta relacionando a expressões aritmeticas, ou a expressões (instruções) em geral 


< switch_stmt > -> switch ( < controle > ) { < lista_casos > }

< controle > -> < variavel > | < stmt >

< lista_casos > -> case < constante > : < stmt > ; break ;  defalt < stmt > ; | case < constante > : < stmt > ; break ;
           

           





  


