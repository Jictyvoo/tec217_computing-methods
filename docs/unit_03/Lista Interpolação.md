> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01

    Os dados a seguir são provenientes de uma tabela que foi medida com alta precisão. Use polinômios interpoladores de Newton de 1a à 4a ordem para determinar y em x = 3,5 para a melhor precisão possível. Calcule as diferenças divididas finitas. Faça os cálculos à mão e computacionalmente. Repita utilizando polinômios de Lagrange. Ordene seus pontos para obter acurácia e convergência ótimas; ou seja, os pontos devem ser centrados em torno e tão próximo quanto possível do valor desconhecido.

| **x** |  0  |   1    |  2.5   |   3    |  4.5   |   5    |  6  |
| :---: | :-: | :----: | :----: | :----: | :----: | :----: | :-: |
| **y** |  2  | 5.4375 | 7.3516 | 7.5625 | 8.4453 | 9.1875 | 12  |

### Resultados

## Questão 02

    Considere os dados abaixo e calcule f (4) usando polinômios interpoladores de Newton de primeiro a quarto graus. Escolha seus pontos-base para obter uma boa acurácia; ou seja, os pontos devem ser centrados em torno e tão próximo quanto possível do valor desconhecido. O que seus resultados indicam em relação ao grau do polinômio usado para gerar os dados da tabela? Repita com polinômios de Lagrange de 1a a 4a ordem.

|  **x**   |  1  |  2  |  3  |  5  |  6  |
| :------: | :-: | :-: | :-: | :-: | :-: |
| **f(x)** |  7  |  4  | 5.5 | 40  | 82  |

### Resultados

## Questão 03

    Empregue interpolação inversa utilizando um polinômio interpolador cúbico e bissecção (no intervalo [2; 3] de x) para determinar o valor de x que corresponde a f(x) = 1.7 para os seguintes dados tabulados:

|  **x**   |  1  |  2  |  3  |  4  |  5   |  6  |    7    |
| :------: | :-: | :-: | :-: | :-: | :--: | :-: | :-----: |
| **f(x)** | 3.6 | 1.8 | 1.2 | 0.9 | 0.72 | 1.5 | 0.51429 |

### Resultados

## Questão 04

    A tabela abaixo lista os valores de concentração de oxigênio dissolvido na água como uma função da temperatura e da concentração do cloro. (solução numérica)

#### Tabela 4 – Concentração de oxigênio dissolvido em água como função da temperatura e concentração de cloreto

| **T, °C** | **c = 0 g/L** | **c = 10 g/L** | **c = 20 g/L** |
| :-------: | :-----------: | :------------: | :------------: |
|     0     |     14.6      |      12.9      |      11.4      |
|     5     |     12.8      |      11.3      |      10.3      |
|    10     |     11.3      |      10.1      |      8.96      |
|    15     |     10.1      |      9.03      |      8.08      |
|    20     |     9.09      |      8.17      |      7.35      |
|    25     |     8.26      |      7.46      |      6.73      |
|    30     |     7.56      |      6.85      |      6.20      |

    a) Use interpolação quadrática e cúbica para determinar a concentração de oxigênio para T=12°C e c=10g/L.
    b) Use interpolação linear para determinara concentração de oxigênio para T=12°C e c=15g/L.

### Resultados

## Questão 05

    Você mediu a queda de tensão V através de um resistor para diversos valores diferentes de corrente i. Os resultados são:

| **i** | 0.25  | 0.75 | 1.25 | 1.5  | 2.0 |
| :---: | :---: | :--: | :--: | :--: | :-: |
| **V** | -0.45 | -0.6 | 0.70 | 1.88 | 6.0 |

    Use interpolação polinomial de primeiro a quarto graus (ordenar os pontos para melhor precisão) para fazer uma estimativa da queda de tensão para i = 1,15. Incluir o cálculo aproximado para o erro de interpolação e usar este para determinar a melhor estimativa para a tensão. Interprete seus resultados.

### Resultados

## Questão 06

    Empregue interpolação inversa para determinar o valor de x que corresponde a f (x) = 0,85 para os seguintes dados tabulados:

|  **x**   |  0  |  1  |  2  |  3  |    4     |    5     |
| :------: | :-: | :-: | :-: | :-: | :------: | :------: |
| **f(x)** |  0  | 0.5 | 0.8 | 0.9 | 0.941176 | 0.961538 |

    Observe que os valores na tabela foram gerados pela função f (x) = x^2/(1 + x^2).
    	(a) Determine o valor correto analiticamente.
    	(b) Use interpolação e a fórmula quadráticas para determinar o valor numericamente.
    	(c) Utilize interpolação cúbica e bissecção para determinar o valor numericamente.

### Resultados

## Questão 07

    A lei de Ohm estabelece que a queda de tensão V através de um resistor ideal é linearmente proporcional à corrente i passando pelo resistor como em V = iR, onde R é a resistência. Entretanto, resistores reais nem sempre podem obedecer à lei de Ohm. Considere que você fez algumas experiências muito precisas para medir a queda de tensão e a corrente resultante para um resistor. Os resultados apresentados na tabela a seguir sugerem uma relação curvilínea em vez da reta representada pela lei de Ohm.

| **i** |  -1  | -0.5  | -0.25 | 0.25 | 0.5  |  1  |
| :---: | :--: | :---: | :---: | :--: | :--: | :-: |
| **V** | -637 | -96.5 | -20.5 | 20.5 | 96.5 | 637 |

    Para quantificar essa relação, uma curva deve ser ajustada aos dados. Devido aos erros nas medidas, a regressão seria o método preferido de ajuste de curva para analisar esses dados experimentais. Entretanto, o fato de essa relação ser suave e a precisão dos métodos experimentais sugerem que a interpolação poderia ser apropriada. Use um polinômio interpolador de quinto grau para ajustar os dados e calcule V para i = 0,10.

### Resultados

## Questão 08

    A aceleração da gravidade a uma altitude y acima da superfície da Terra é dada por:

|   **$y,m$**    |   0    | 30000  | 60000  | 90000  | 120000 |
| :------------: | :----: | :----: | :----: | :----: | :----: |
| **$g, m/s^2$** | 9.8100 | 9.7487 | 9.6879 | 9.6278 | 9.5682 |

    Calcule g em y = 55.000 m. Usar polinômios de 1a a 3a ordem para obter a estimativa. O resultado pode ser obtido com quantos dígitos significativos? Comente sua resposta.

### Resultados
