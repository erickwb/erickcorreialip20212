# Atividade 07

# Correção: 1,0

## Erick Correia Silva 


Linguagem: 
1. E -> E + T
2. E -> T 
3. T -> T * F
4. T -> F
5. F -> ( E )
6. F -> id


Cadeia:
id * (id + id)


# Tabela 

Pilha   | Entrada           | Ação
------- | ----------------- | ----------
0       | id * (id + id)  $ |   Deslocar 5
0id5    |    * (id + id) $  |   Reduzir 6 (use GOTO[0, F])
0F3     |    * (id + id) $  |   Reduzir 4 (use GOTO[0, T])
0T2     |    * (id + id) $  |   Desloca 7 
0T2 * 7 |    (id + id) $    |   Desloca 4
0T2 * 7 ( 4 | id + id) $    |   Desloca 5
0T2 * 7 ( 4id5 | + id) $    |   Reduzir 6 (use GOTO[4, F])
0T2 * 7 ( 4F3  | + id ) $   |   Reduzir 4 (use GOTO[4, T])
0T2 * 7 ( 4T2  | + id) $    |   Reduzir 2 (use GOTO[4, E])
0T2 * 7 ( 4E8  | + id) $    |   Desloca 6
0T2 * 7 ( 4E8 + 6 |  id) $  |   Desloca 5
0T2 * 7 ( 4E8 + 6id5 |  ) $ |   Reduzir 6 (use GOTO[6, F])
0T2 * 7 ( 4E8 + 6F3  |  ) $ |   Reduzir 4 (use GOTO[4, F])
0T2 * 7 ( 4E8 + 6T9  |  ) $ |   Reduzir 1 (use GOTO[4, E])
0T2 * 7 ( 4E8        |  ) $ |   Desloca 11
0T2 * 7 ( 4E8 ) 11   |   $  |   Reduzir 5 (use GOTO[7, F])
0T2 * 7F10           |   $  |   Reduzir 3 (use GOTO[0, T])
0T2                  |   $  |   Reduzir 2 (use GOTO[0, E])
0E1                  |   $  |   Aceita










