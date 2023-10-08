> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Para as implementações dos métodos solicitados logo abaixo, deve-se obrigatoriamente realizar ao menos duas iterações manuais.

## Questão 01

    Das equações a seguir resolva por eliminação de Gauss com pivoteamento parcial, use os elementos da diagonal para calcular o determinante. Substitua seus resultados nas equações originais para verificar suas respostas. Mostre os passos dos cálculos.

$$
\begin{align}
2x_1 - 6x_2 - 1x_3 = -38 \\
-3x_1 - 1x_2 + 7x_3 = -34 \\
-8x_1 + 1x_2 - 2x_3 = -20 \\
\end{align}
$$

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
2 & -6 & -1 \\
-3 & -1 & 7 \\
-8 & 1 & -2
\end{bmatrix}

\begin{bmatrix}
-38 \\
-34 \\
-20
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
\begin{bmatrix}
2 & -6 & -1 & -38 \\
0 & -10 & 5.5 & -91 \\
-8 & 1 & -2 & -20
\end{bmatrix}
$$

**Multiplier:** -1.5000
**Operation:**

- $L2 = L2 - (-1.5000 * L1)$

#### Passo 2

**Matrix:**

$$
\begin{bmatrix}
2 & -6 & -1 & -38 \\
0 & -10 & 5.5 & -91 \\
0 & -23 & -6 & -172
\end{bmatrix}
$$

**Multiplier:** -4.0000
**Operation:**

- $L3 = L3 - (-4.0000 * L1)$

#### Passo 3

**Matrix:**

$$
\begin{bmatrix}
2 & -6 & -1 & -38 \\
0 & -10 & 5.5 & -91 \\
0 & 0 & -18.65 & 37.30
\end{bmatrix}
$$

**Multiplier:** 2.3000
**Operation:**

- $L3 = L3 - (2.3000 * L2)$

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:

| **Step** |    **Roots**    | **DividendSum** | **Divisor** | **Added Root** |
| :------: | :-------------: | :-------------: | :---------: | :------------: |
|    1     | \[0,0,-1.99999] |    37.300000    | -18.650000  |   -2.000000    |
|    2     | \[0,8,-1.99999] |   -80.000000    | -10.000000  |    8.000000    |
|    3     | \[4,8,-1.99999] |    8.000000     |  2.000000   |    4.000000    |

**Valor do Determinante:** $-282$.
**Raízes Encontradas:** $4$, $8$, $-1.9999$.

---

### Verificando os resultados

Agora, vamos substituir as raízes encontradas nas equações originais para verificar se elas são satisfeitas:

Para os valores

$$
\begin{align}
x_1 = 4 \\
x_2 = 8 \\
x_3 = -1.9999
\end{align}
$$

1. Primeira equação:

$$
\begin{align}
2(4) - 6(8) - 1(-1.9999) = 8 - 48 + 1.9999 \approx -38 \\
\text{A primeira equação é aproximadamente satisfeita.}
\end{align}
$$

2. Segunda equação:

$$
\begin{align}
-3(4) - 1(8) + 7(-1.9999) \approx -12 - 8 - 13.9993 \approx -34 \\
\text{A segunda equação é aproximadamente satisfeita.}
\end{align}
$$

3. Terceira equação:

$$
\begin{align}
-8(4) + 1(8) - 2(-1.9999) \approx -32 + 8 + 3.9998 \approx -20 \\
\text{A terceira equação é aproximadamente satisfeita.}
\end{align}
$$

Portanto, as raízes encontradas aproximadamente satisfazem todas as equações originais do sistema. Isso indica que as respostas estão próximas de serem corretas.

## Questão 02

    O sistema mostra 3 tanques interligados por tubos. Como indicado, a taxa de transferência de produtos através de cada tubo é igual a vazão Q multiplicada pela concentração do reator do qual o fluxo se origina. Se o sistema estiver em estado estacionário, a transferência para dentro de cada reator vai balancear a transferência para fora. Deduza e resolva as 3 equações lineares para suas concentrações por Gauss Jordan. Mostre os passos dos cálculos.

$$
\begin{align}
\text{Constantes: } \\
Q33 = 120 \\
Q13 = 40 \\
Q12 = 90 \\
Q23 = 60 \\
Q21 = 30 \\
\\
(Q_{13} + Q_{12})x_1 - Q21x_2 + 0x_3 = 200 \\
Q_{12}x_1 - (Q_{21}+Q_{23})x_2 + 0x_3 = 0 \\
Q_{13}x_1 + Q_{23}x_2 - Q_{33}x_3 = -500 \\
\end{align}
$$

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
130 & -30 & 0 \\
90 & -90 & 0 \\
40 & 60 & -120
\end{bmatrix}

\begin{bmatrix}
200 \\
0 \\
-500
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
C=
\begin{bmatrix}
130 & -30 & 0 & 1 & 0 & 0 & 200 \\
90 & -90 & 0 & 0 & 1 & 0 & 0 \\
40 & 60 & -120 & 0 & 0 & 1 & -500
\end{bmatrix}
$$

**Operation:**

- $L1 \iff L1$

#### Passo 2

**Matrix:**

$$
C=
\begin{bmatrix}
1 & -0.2308 & 0 & 0.0077 & 0 & 0 & 1.5385 \\
90 & -90 & 0 & 0 & 1 & 0 & 0 \\
40 & 60 & -120 & 0 & 0 & 1 & -500
\end{bmatrix}
$$

**Multiplier:** 130.0000

**Operation:**

- $L1 = L1 / 130.0000$

#### Passo 3

**Matrix:**

$$
C=
\begin{bmatrix}
1 & -0.2308 & 0 & 0.0077 & 0 & 0 & 1.5385 \\
0 & -69.2308 & 0 & -0.6923 & 1 & 0 & -138.4615 \\
40 & 60 & -120 & 0 & 0 & 1 & -500
\end{bmatrix}
$$

**Multiplier:** 90.0000

**Operation:**

- $L2 = L2 - (90.0000 * L1)$

#### Passo 4

**Matrix:**

$$
C=
\begin{bmatrix}
1 & -0.2308 & 0 & 0.0077 & 0 & 0 & 1.5385 \\
0 & -69.2308 & 0 & -0.6923 & 1 & 0 & -138.4615 \\
0 & 69.2308 & -120 & -0.3077 & 0 & 1 & -561.5385
\end{bmatrix}
$$

**Multiplier:** 40.0000

**Operation:**

- $L3 = L3 - (40.0000 * L1)$

#### Passo 5

**Matrix:**

$$
C=
\begin{bmatrix}
1 & -0.2308 & 0 & 0.0077 & 0 & 0 & 1.5385 \\
0 & -69.2308 & 0 & -0.6923 & 1 & 0 & -138.4615 \\
0 & 69.2308 & -120 & -0.3077 & 0 & 1 & -561.5385
\end{bmatrix}
$$

**Operation:**

- $L2 \iff L2$

#### Passo 6

**Matrix:**

$$
C=
\begin{bmatrix}
1 & -0.2308 & 0 & 0.0077 & 0 & 0 & 1.5385 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 69.2308 & -120 & -0.3077 & 0 & 1 & -561.5385
\end{bmatrix}
$$

**Multiplier:** -69.2308

**Operation:**

- $L2 = L2 / -69.2308$

#### Passo 7

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.0077 & -0.0033 & 0 & 2 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 69.2308 & -120 & -0.3077 & 0 & 1 & -561.5385
\end{bmatrix}
$$

**Multiplier:** -0.2308

**Operation:**

- $L1 = L1 - (-0.2308 * L2)$

#### Passo 8

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.0077 & -0.0033 & 0 & 2 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 0 & -120 & -1 & 1 & 1 & -700
\end{bmatrix}
$$

**Multiplier:** 69.2308

**Operation:**

- $L3 = L3 - (69.2308 * L2)$

#### Passo 9

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.0077 & -0.0033 & 0 & 2 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 0 & -120 & -1 & 1 & 1 & -700
\end{bmatrix}
$$

**Operation:**

- $L3 \iff L3$

#### Passo 10

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.0077 & -0.0033 & 0 & 2 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 0 & 1 & 0.0083 & -0.0083 & -0.0083 & 5.8333
\end{bmatrix}
$$

**Multiplier:** -120.0000

**Operation:**

- $L3 = L3 / -120.0000$

#### Passo 11

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.0077 & -0.0033 & 0 & 2 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 0 & 1 & 0.0083 & -0.0083 & -0.0083 & 5.8333
\end{bmatrix}
$$

**Operation:**

- $L1 = L1 - (0.0000 * L3)$

#### Passo 12

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.0077 & -0.0033 & 0 & 2 \\
0 & 1 & 0 & 0.0144 & -0.01 & 0 & 2 \\
0 & 0 & 1 & 0.0083 & -0.0083 & -0.0083 & 5.8333
\end{bmatrix}
$$

**Operation:**

- $L2 = L2 - (-0.0000 * L3)$

#### Determinante e Raízes

Como transformamos a matriz em uma matriz identidade, os valores das raízes são obtidos diretamente.

**Valor do Determinante:** 1080000
**Raízes Encontradas:** 2, 2, 5.83

##### Matriz Inversa

| 0.0100 | -0.0033 | 0.0000  |
| :----: | :-----: | :-----: |
| 0.0100 | -0.0144 | -0.0000 |
| 0.0083 | -0.0083 | -0.0083 |

---

## Questão 03

    Um engenheiro civil envolvido em uma obra precisa de 4.800, 5.800 e 5.700 m3 de areia, cascalho fino e cascalho grosso, respectivamente, para um projeto de construção. Existem 3 minas de onde esses materiais podem ser obtidos. A composição dessas minas está na tabela abaixo. Quantos m3 devem ser minerados de cada mina para atender as necessidades do engenheiro? Use eliminação de Gauss sem pivoteamento.

$$
\begin{align}
0.52x_1 + 0.2x_2 + 0.25x_3 = 4800 \\
0.3x_1 + 0.5x_2 + 0.2x_3 = 5800 \\
0.18x_1 + 0.3x_2 + 0.55x_3 = 5700 \\
\end{align}
$$

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
2 & -6 & -1 \\
-3 & -1 & 7 \\
-8 & 1 & -2
\end{bmatrix}

\begin{bmatrix}
-38 \\
-34 \\
-20
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
\begin{bmatrix}
0.52 & 0.2 & 0.25 & 4800 \\
0 & 0.38461538461538464 & 0.05576923076923079 & 3030.769230769231 \\
0.18 & 0.3 & 0.55 & 5700
\end{bmatrix}
$$

**Multiplier:** 0.576923
**Operation:**

- $L2 = L2 - (0.5769 * L1)$

#### Passo 2

**Matrix:**

$$
\begin{bmatrix}
0.52 & 0.2 & 0.25 & 4800 \\
0 & 0.38461538461538464 & 0.05576923076923079 & 3030.769230769231 \\
0 & 0.23076923076923075 & 0.4634615384615385 & 4038.4615384615386
\end{bmatrix}
$$

**Multiplier:** 0.346154
**Operation:**

- $L3 = L3 - (0.3462 * L1)$

#### Passo 3

**Matrix:**

$$
\begin{bmatrix}
0.52 & 0.2 & 0.25 & 4800 \\
0 & 0.38461538461538464 & 0.05576923076923079 & 3030.769230769231 \\
0 & -2.7755575615628914e-17 & 0.43 & 2220
\end{bmatrix}
$$

**Multiplier:** 0.600000
**Operation:**

- $L3 = L3 - (0.6000 * L2)$

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:

| **Step** | **Roots**                                                 | **DividendSum** | **Divisor** | **Added Root** |
| :------: | --------------------------------------------------------- | :-------------: | :---------: | :------------: |
|    1     | \[0,0,5162.790697674419]                                  |   2220.000000   |  0.430000   |  5162.790698   |
|    2     | \[0,7131.395348837208,5162.790697674419]                  |   2742.844365   |  0.384615   |  7131.395349   |
|    3     | \[4005.8139534883717,7131.395348837208,5162.790697674419] |   2083.023256   |  0.520000   |  4005.813953   |

**Valor do Determinante:** $0.086$.
**Raízes Encontradas:** $4005.8139$, $7131.3953$, $5162.7907$.

---

## Questão 04

    a)Use a decomposição LU para resolver o seguinte sistema de equações. Mostre os passos.
    b) resolva o sistema para um vetor do lado direito alternativo aproveitando as matrizes L e U obtidas em (a):

$$
\begin{align}
7x_1 + 2x_2 - 3x_3 = -12 \\
2x_1 + 5x_2 - 3x_3 = -20 \\
1x_1 - 1x_2 - 6x_3 = -26
\end{align}
$$

### Resultados

Como o objetivo era realizar a aplicação do método de decomposição LU, foram coletados os dados referente à obtenção das matrizes L e U. Esses dados podem ser visualizados a seguir.

#### Matriz L

Inicialmente é utilizado o valor da matriz L como sendo a matriz identidade. Nesse caso, como sendo I 3x3.

$$
L=
\begin{bmatrix}
1 & 0 & 0 \\
0 & 1 & 0 \\
0 & 0 & 1
\end{bmatrix}
$$

| **Step** | **Matrix**                                                                          | **Multiplier** |   **Operation**   |
| :------: | ----------------------------------------------------------------------------------- | :------------: | :---------------: |
|    1     | \[\[1,0,0],\[0.2857142857142857,1,0],\[0,0,1]]                                      |    0.285714    | L\[1,2] = 0.2857  |
|    2     | \[\[1,0,0],\[0.2857142857142857,1,0],\[0.14285714285714285,0,1]]                    |    0.142857    | L\[1,3] = 0.1429  |
|    3     | \[\[1,0,0],\[0.2857142857142857,1,0],\[0.14285714285714285,-0.29032258064516125,1]] |   -0.290323    | L\[2,3] = -0.2903 |

#### Matriz U

Para o calculo da matriz U, assume-se como valor inicial da mesma o mesmo da matriz de entrada.

$$
U=
\begin{bmatrix}
7 & 2 & -3 \\
2 & 5 & -3 \\
1 & -1 & -6
\end{bmatrix}
$$

| **Step** | **Matrix**                                                                                         | **Multiplier** |             **Operation** |
| :------: | -------------------------------------------------------------------------------------------------- | :------------: | ------------------------: |
|    1     | \[\[7,2,-3],\[0,4.428571428571429,-2.142857142857143],\[1,-1,-6]]                                  |    0.285714    |  L1 = L1 - (0.2857 \* L2) |
|    2     | \[\[7,2,-3],\[0,4.428571428571429,-2.142857142857143],\[0,-1.2857142857142856,-5.571428571428571]] |    0.142857    |  L1 = L1 - (0.1429 \* L3) |
|    3     | \[\[7,2,-3],\[0,4.428571428571429,-2.142857142857143],\[0,0,-6.193548387096774]]                   |   -0.290323    | L2 = L2 - (-0.2903 \* L3) |

#### Cálculo das Raízes

##### B1

De modo a realizar a obtenção das raízes, deve-se inicialmente calcular o vetor y, através da resolução do triângulo inferior da matriz L. Para realização desse cálculo, utiliza-se o valor indicado:

$$
b_1 =
\begin{bmatrix}
-12 \\
-20 \\
-26
\end{bmatrix}
$$

| **Step** | **Y**                                          | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ---------------------------------------------- | :-------------: | :---------: | :------------: |
|    1     | \[-12,0,0]                                     |   -12.000000    |  1.000000   |   -12.000000   |
|    2     | \[-12,-16.571428571428573,0]                   |   -16.571429    |  1.000000   |   -16.571429   |
|    3     | \[-12,-16.571428571428573,-29.096774193548384] |   -29.096774    |  1.000000   |   -29.096774   |

E posteriormente deve-se utilizar o vetor resultante y para calcular o vetor de raízes x, através da resolução do triângulo superior da matriz U.

| **Step** | **Roots**                                                   | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ----------------------------------------------------------- | :-------------: | :---------: | :------------: |
|    1     | \[0,0,4.697916666666666]                                    |   -29.096774    |  -6.193548  |    4.697917    |
|    2     | \[0,-1.4687500000000004,4.697916666666666]                  |    -6.504464    |  4.428571   |   -1.468750    |
|    3     | \[0.7187499999999998,-1.4687500000000004,4.697916666666666] |    5.031250     |  7.000000   |    0.718750    |

**Valor do Determinante:** $-192$
**Raízes Encontradas:** $0.7188$, $-1.4688$, $4.6979$

##### B2

De modo a realizar a obtenção das raízes, deve-se inicialmente calcular o vetor y, através da resolução do triângulo inferior da matriz L. Para realização desse cálculo, utiliza-se o valor indicado:

$$
b_2 =
\begin{bmatrix}
12 \\
18 \\
-6
\end{bmatrix}
$$

| **Step** | **Y**                                       | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ------------------------------------------- | :-------------: | :---------: | :------------: |
|    1     | \[12,0,0]                                   |    12.000000    |  1.000000   |    0.000000    |
|    2     | \[12,14.571428571428571,0]                  |    14.571429    |  1.000000   |   14.571429    |
|    3     | \[12,14.571428571428571,-3.483870967741936] |    -3.483871    |  1.000000   |   12.000000    |

E posteriormente deve-se utilizar o vetor resultante y para calcular o vetor de raízes x, através da resolução do triângulo superior da matriz U.

| **Step** | **Roots**                                                   | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ----------------------------------------------------------- | :-------------: | :---------: | :------------: |
|    1     | \[0,0,0.5625000000000001]                                   |    -3.483871    |  -6.193548  |    0.562500    |
|    2     | \[0,3.5624999999999996,0.5625000000000001]                  |    15.776786    |  4.428571   |    3.562500    |
|    3     | \[0.9375000000000002,3.5624999999999996,0.5625000000000001] |    6.562500     |  7.000000   |    0.937500    |

**Valor do Determinante:** $-192$
**Raízes Encontradas:** $0.9375$, $3.5625$, $0.5625$

## Questão 05

    Determine a matriz inversa para o sistema a seguir. Verifique seu resultado através de [A][A]-1. Não use estratégia de pivoteamento.

$$
\begin{align}
10x_1 + 2x_2 - 1x_3 = 27 \\
-3x_1 - 6x_2 + 2x_3 = -61.5 \\
1x_1 + 1x_2 + 5x_3 = -21.5 \\
\end{align}
$$

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
10 & 2 & -1 \\
-3 & -6 & 2 \\
1 & 1 & 5
\end{bmatrix}

\begin{bmatrix}
27 \\
-61.5 \\
-21.5
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0.2 & -0.1 & 0.1 & 0 & 0 & 2.7 \\
-3 & -6 & 2 & 0 & 1 & 0 & -61.5 \\
1 & 1 & 5 & 0 & 0 & 1 & -21.5
\end{bmatrix}
$$

**Multiplier:** 10.0000

**Operation:**

- $L1 = L1 / 10.0000$

#### Passo 2

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0.2 & -0.1 & 0.1 & 0 & 0 & 2.7 \\
0 & -5.4 & 1.7 & 0.3 & 1 & 0 & -53.4 \\
1 & 1 & 5 & 0 & 0 & 1 & -21.5
\end{bmatrix}
$$

**Multiplier:** -3.0000

**Operation:**

- $L2 = L2 - (-3.0000 * L1)$

#### Passo 3

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0.2 & -0.1 & 0.1 & 0 & 0 & 2.7 \\
0 & -5.4 & 1.7 & 0.3 & 1 & 0 & -53.4 \\
0 & 0.8 & 5.1 & -0.1 & 0 & 1 & -24.2
\end{bmatrix}
$$

**Multiplier:** 1.0000

**Operation:**

- $L3 = L3 - (1.0000 * L1)$

#### Passo 4

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0.2 & -0.1 & 0.1 & 0 & 0 & 2.7 \\
0 & 1 & -0.3148 & -0.0556 & -0.1852 & 0 & 9.8889 \\
0 & 0.8 & 5.1 & -0.1 & 0 & 1 & -24.2
\end{bmatrix}
$$

**Multiplier:** -5.4000

**Operation:**

- $L2 = L2 / -5.4000$

#### Passo 5

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & -0.0370 & 0.1111 & 0.0370 & 0 & 0.7222 \\
0 & 1 & -0.3148 & -0.0556 & -0.1852 & 0 & 9.8889 \\
0 & 0.8 & 5.1 & -0.1 & 0 & 1 & -24.2
\end{bmatrix}
$$

**Multiplier:** 0.2000

**Operation:**

- $L1 = L1 - (0.2000 * L2)$

#### Passo 6

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & -0.0370 & 0.1111 & 0.0370 & 0 & 0.7222 \\
0 & 1 & -0.3148 & -0.0556 & -0.1852 & 0 & 9.8889 \\
0 & 0 & 5.3519 & -0.0556 & 0.1481 & 1 & -32.1111
\end{bmatrix}
$$

**Multiplier:** 0.8000

**Operation:**

- $L3 = L3 - (0.8000 * L2)$

#### Passo 7

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & -0.0370 & 0.1111 & 0.0370 & 0 & 0.7222 \\
0 & 1 & -0.3148 & -0.0556 & -0.1852 & 0 & 9.8889 \\
0 & 0 & 1 & -0.0104 & 0.0277 & 0.1869 & -6
\end{bmatrix}
$$

**Multiplier:** 5.3519

**Operation:**

- $L3 = L3 / 5.3519$

#### Passo 8

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.1107 & 0.0381 & 0.0069 & 0.5 \\
0 & 1 & -0.3148 & -0.0556 & -0.1852 & 0 & 9.8889 \\
0 & 0 & 1 & -0.0104 & 0.0277 & 0.1869 & -6
\end{bmatrix}
$$

**Multiplier:** -0.0370

**Operation:**

- $L1 = L1 - (-0.0370 * L3)$

#### Passo 9

**Matrix:**

$$
C=
\begin{bmatrix}
1 & 0 & 0 & 0.1107 & 0.0381 & 0.0069 & 0.5 \\
0 & 1 & 0 & -0.0588 & -0.1765 & 0.0588 & 8 \\
0 & 0 & 1 & -0.0104 & 0.0277 & 0.1869 & -6
\end{bmatrix}
$$

**Multiplier:** -0.3148

**Operation:**

- $L2 = L2 - (-0.3148 * L3)$

#### Determinante e Raízes

Como transformamos a matriz em uma matriz identidade, os valores das raízes são obtidos diretamente.

**Valor do Determinante:** -288.99
**Raízes Encontradas:** 0.5, 8, -6

##### Matriz Inversa

| 0.1107  | 0.0381  | 0.0069 |
| :-----: | :-----: | :----: |
| -0.0588 | -0.1765 | 0.0588 |
| -0.0104 | 0.0277  | 0.1869 |

---

## Questão 06

    Determine as normas ||A||l e ||A||inf para a matriz A. Antes de determinar as normas, normalize a matriz fazendo o maior elemento de cada linha igual a 1.

$$
A =
\begin{bmatrix}
8 & 2 & -10 \\
-9 & 1 & 3 \\
15 & -1 & 6 \\
\end{bmatrix}
$$

### Resultados

#### Normalização

##### Passo 1: Normalização da linha L1

- O maior elemento da primeira linha é 10.
- $L1 = L1 / 10$

$$
\begin{bmatrix}
\frac{4}{5} & \frac{1}{5} & -1 \\
-9 & 1 & 3 \\
15 & -1 & 6 \\
\end{bmatrix}
$$

##### Passo 2: Normalização da linha L2

- O maior elemento da segunda linha é 3.
- $L2 = L2 / 9$

$$
\begin{bmatrix}
\frac{4}{5} & \frac{1}{5} & -1 \\
-1 & \frac{1}{9} & \frac{1}{3} \\
15 & -1 & 6 \\
\end{bmatrix}
$$

##### Passo 3: Normalização da linha L3

- O maior elemento da terceira linha é 15.
- $L3 = L3 / 15$

$$
\begin{bmatrix}
\frac{4}{5} & \frac{1}{5} & -1 \\
-1 & \frac{1}{9} & \frac{1}{3} \\
1 & -\frac{1}{15} & \frac{2}{5} \\
\end{bmatrix}
$$

#### Norma $\|A\|_1$

A norma $\|A\|_1$ de uma matriz é o máximo da soma dos valores absolutos dos elementos da matriz.

- Soma dos valores absolutos da primeira **coluna**: $\left|\frac{4}{5}\right| + \left|-1\right| + \left|1\right| = \frac{14}{5}$
- Soma dos valores absolutos da segunda **coluna**: $\left|\frac{1}{5}\right| + \left|\frac{1}{9}\right| + \left|-\frac{1}{15}\right| = \frac{17}{45}$
- Soma dos valores absolutos da terceira **coluna**: $\left|-1\right| + \left|-\frac{1}{3}\right| + \left|\frac{2}{5}\right| = \frac{26}{15}$

O máximo desses valores é $\frac{14}{5}$

#### Norma $\|A\|_\infty$

A norma $\|A\|_\infty$ de uma matriz é o máximo dos valores absolutos das somas das linhas da matriz.

- Soma dos valores absolutos da primeira linha: $\left|\frac{4}{5}\right| + \left|\frac{1}{5}\right| + \left|-1\right| = 2$
- Soma dos valores absolutos da segunda linha: $\left|-1\right| + \left|\frac{1}{9}\right| + \left|\frac{1}{3}\right| = \frac{13}{9}$
- Soma dos valores absolutos da terceira linha: $\left|1\right| + \left|-\frac{1}{15}\right| + \left|\frac{2}{5}\right| = \frac{22}{15}$

O máximo desses valores é $2$.

## Questão 08

    Use o método de Gauss-Seidel para resolver o sistema a seguir, até que o erro relativo percentual caia abaixo de es = 5%.

$$
\begin{align}
\epsilon = 0.05 \\
x_0 = [0,0,0] \\
\\
0.8x_1 - 0.4x_2 + 0x_3 = 41 \\
-0.4x_1 + 0.8x_2 - 0.4x_3 = 25 \\
0x_1 - 0.4x_2 + 0.8x_3 = 105 \\
\end{align}
$$

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
0.8 & -0.4 & 0 \\
-0.4 & 0.8 & -0.4 \\
0 & -0.4 & 0.8
\end{bmatrix}

\begin{bmatrix}
41 \\
25 \\
105
\end{bmatrix}
$$

Aqui estão os passos organizados com base nas informações fornecidas:

#### Valor da matriz C

##### Passo 1

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.000000

**Operation:**

- $L[1,1] = 0.0000$

##### Passo 2

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.500000

**Operation:**

- $L[2,1] = 0.5000$

##### Passo 3

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.000000

**Operation:**

- $L[3,1] = -0.0000$

##### Passo 4

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0.5 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.500000

**Operation:**

- $L[1,2] = 0.5000$

##### Passo 5

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0.5 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.000000

**Operation:**

- $L[2,2] = 0.0000$

##### Passo 6

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0.5 & 0 & 0.5 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.500000

**Operation:**

- $L[3,2] = 0.5000$

##### Passo 7

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0.5 & 0 & 0.5 \\
-0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.000000

**Operation:**

- $L[1,3] = -0.0000$

##### Passo 8

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0.5 & 0 & 0.5 \\
-0 & 0.5 & 0
\end{bmatrix}
$$

**Multiplier:** 0.500000

**Operation:**

- $L[2,3] = 0.5000$

##### Passo 9

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.5 & -0 \\
0.5 & 0 & 0.5 \\
-0 & 0.5 & 0
\end{bmatrix}
$$

**Multiplier:** 0.000000

**Operation:**

- $L[3,3] = 0.0000$

#### Valor do vetor `d`

Enquanto eram realizadas operações para montar a matriz C, também foi realizada a montagem do vetor d.

| **Step** | **d**                 | **DividendSum** | **Divisor** | **Added Root** |
| :------: | --------------------- | :-------------: | :---------: | :------------: |
|    1     | \[51.25,0,0]          |    41.000000    |  0.800000   |    0.000000    |
|    2     | \[51.25,31.25,0]      |    25.000000    |  0.800000   |   31.250000    |
|    3     | \[51.25,31.25,131.25] |   105.000000    |  0.800000   |   51.250000    |

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:
**Raízes Encontradas:** $173.7041$, $244.9541$, $253.7270$

De modo a facilitar a visualização, os valores de `a, b, c` correspondem ao valor anterior do X.

| **Interaction** |   **a**    |   **b**    |   **c**    | **Found-X**                                                 | **f(x)** | **Relative-Error** |
| :-------------: | :--------: | :--------: | :--------: | ----------------------------------------------------------- | :------: | :----------------: |
|        1        |  0.000000  |  0.000000  |  0.000000  | \[51.25,56.875,159.6875]                                    | 0.000000 |     100.000000     |
|        2        | 51.250000  | 56.875000  | 159.687500 | \[79.6875,150.9375,206.71875]                               | 0.000000 |     62.318841      |
|        3        | 79.687500  | 150.937500 | 206.718750 | \[126.71875,197.96875,230.234375]                           | 0.000000 |     37.114673      |
|        4        | 126.718750 | 197.968750 | 230.234375 | \[150.234375,221.484375,241.9921875]                        | 0.000000 |     15.652626      |
|        5        | 150.234375 | 221.484375 | 241.992188 | \[161.9921875,233.2421875,247.87109375]                     | 0.000000 |      7.258259      |
|        6        | 161.992188 | 233.242188 | 247.871094 | \[167.87109375,239.12109375,250.810546875]                  | 0.000000 |      3.502036      |
|        7        | 167.871094 | 239.121094 | 250.810547 | \[170.810546875,242.060546875,252.2802734375]               | 0.000000 |      1.720885      |
|        8        | 170.810547 | 242.060547 | 252.280273 | \[172.2802734375,243.5302734375,253.01513671875]            | 0.000000 |      0.853102      |
|        9        | 172.280273 | 243.530273 | 253.015137 | \[173.01513671875,244.26513671875,253.382568359375]         | 0.000000 |      0.424739      |
|       10        | 173.015137 | 244.265137 | 253.382568 | \[173.382568359375,244.632568359375,253.5662841796875]      | 0.000000 |      0.211920      |
|       11        | 173.382568 | 244.632568 | 253.566284 | \[173.5662841796875,244.8162841796875,253.65814208984375]   | 0.000000 |      0.105848      |
|       12        | 173.566284 | 244.816284 | 253.658142 | \[173.65814208984375,244.90814208984375,253.70407104492188] | 0.000000 |      0.052896      |
|       13        | 173.658142 | 244.908142 | 253.704071 | \[173.70407104492188,244.95407104492188,253.72703552246094] | 0.000000 |      0.026441      |

### Gráfico de Convergência

![[assets/unit02_challenge_quest08_convgraph.png]]

## Questão 09

    Use o método de Jacob para resolver o seguinte sistema, até que o erro relativo percentual caia abaixo de es = 5%.

$$
\begin{align}
\epsilon = 0.05 \\
x_0 = [0,0,0] \\
\\
10x_1 + 2x_2 - 1x_3 = 27 \\
-3x_1 - 6x_2 + 2x_3 = -61.5 \\
1x_1 + 1x_2 + 5x_3 = -21.5 \\
\end{align}
$$

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
10 & 2 & -1 \\
-3 & -6 & 2 \\
1 & 1 & 5
\end{bmatrix}

\begin{bmatrix}
27 \\
-61.5 \\
-21.5
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
\begin{bmatrix}
0 & 0 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.000000

**Operation:**

- $L[1,1] = 0.0000$

#### Passo 2

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.200000

**Operation:**

- $L[2,1] = -0.2000$

#### Passo 3

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.100000

**Operation:**

- $L[3,1] = 0.1000$

#### Passo 4

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
-0.5 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.500000

**Operation:**

- $L[1,2] = -0.5000$

#### Passo 5

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
-0.5 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.000000

**Operation:**

- $L[2,2] = 0.0000$

#### Passo 6

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
-0.5 & 0 & 0.333333 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.333333

**Operation:**

- $L[3,2] = 0.3333$

#### Passo 7

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
-0.5 & 0 & 0.333333 \\
-0.2 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.200000

**Operation:**

- $L[1,3] = -0.2000$

#### Passo 8

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
-0.5 & 0 & 0.333333 \\
-0.2 & -0.2 & 0
\end{bmatrix}
$$

**Multiplier:** -0.200000

**Operation:**

- $L[2,3] = -0.2000$

#### Passo 9

**Matrix:**

$$
\begin{bmatrix}
0 & -0.2 & 0.1 \\
-0.5 & 0 & 0.333333 \\
-0.2 & -0.2 & 0
\end{bmatrix}
$$

**Multiplier:** 0.000000

**Operation:**

- $L[3,3] = 0.0000$

#### Valor do vetor `d`

Enquanto eram realizadas operações para montar a matriz C, também foi realizada a montagem do vetor d.

| **Step** | **d**             | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ----------------- | :-------------: | :---------: | :------------: |
|    1     | \[2.7,0,0]        |    27.000000    |  10.000000  |    0.000000    |
|    2     | \[2.7,10.25,0]    |   -61.500000    |  -6.000000  |   10.250000    |
|    3     | \[2.7,10.25,-4.3] |   -21.500000    |  5.000000   |    2.700000    |

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:
**Raízes Encontradas:** $0.5000$, $7.9999$, $-5.9999$.

De modo a facilitar a visualização, os valores de `a, b, c` correspondem ao valor anterior do X.

| **Interaction** |  **a**   |   **b**   |   **c**   | **Found-X**                                                  | **Relative-Error** |
| :-------------: | :------: | :-------: | :-------: | ------------------------------------------------------------ | :----------------: |
|        1        | 0.000000 | 0.000000  | 0.000000  | \[2.7,10.25,-4.3]                                            |     100.000000     |
|        2        | 2.700000 | 10.250000 | -4.300000 | \[0.21999999999999975,7.466666666666667,-6.890000000000001]  |    1127.272727     |
|        3        | 0.220000 | 7.466667  | -6.890000 | \[0.5176666666666669,7.843333333333334,-5.8373333333333335]  |     57.501610      |
|        4        | 0.517667 | 7.843333  | -5.837333 | \[0.5476000000000001,8.045388888888889,-5.9722]              |      5.466277      |
|        5        | 0.547600 | 8.045389  | -5.972200 | \[0.49370222222222226,7.9854666666666665,-6.018597777777778] |     10.917062      |
|        6        | 0.493702 | 7.985467  | -6.018598 | \[0.5010468888888888,7.99694962962963,-5.9958337777777775]   |      1.465864      |
|        7        | 0.501047 | 7.996950  | -5.995834 | \[0.5010266962962961,8.000865296296297,-5.999599303703704]   |      0.062763      |
|        8        | 0.501027 | 8.000865  | -5.999599 | \[0.4998670103703704,7.999620217283951,-6.000378398518519]   |      0.231999      |
|        9        | 0.499867 | 7.999620  | -6.000378 | \[0.5000381166913579,7.999940361975309,-5.9998974455308645]  |      0.034219      |

---

### Gráfico de Convergência

![[assets/unit02_challenge_quest09_convgraph.png]]

## Questão 10

    Use o método Gauss-Seidel com relaxamento (w = 1.2) para resolver o sistema a seguir com uma tolerância de es = 5%. Se necessário rearranje as equações para garantir convergência.

$$
\begin{align}
\epsilon = 0.05 \\
x_0 = [0,0,0] \\
\omega = 1.2 \\
\\
2x_1 - 6x_2 -1x_3 = -38 \\
-3x_1 - 1x_2 + 7x_3 = -34 \\
-8x_1 + 1x_2 - 2x_3 = -20
\end{align}
$$

### Resultados

Para executar o método com sucesso, fez-se necessário por a ordem das linhas da matriz como $L3$, $L2$, $L1$.

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
-8 & 1 & -2 \\
2 & -6 & -1 \\
-3 & -1 & 7
\end{bmatrix}

\begin{bmatrix}
-20 \\
-38 \\
-34
\end{bmatrix}
$$

Aqui estão os passos organizados de acordo com o modelo fornecido:

#### Achando C

##### Passo 1

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.0000

**Operation:**

- $L[1,1] = 0.0000$

##### Passo 2

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.1250

**Operation:**

- $L[2,1] = 0.1250$

##### Passo 3

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.2500

**Operation:**

- $L[3,1] = -0.2500$

##### Passo 4

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0.3333 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.3333

**Operation:**

- $L[1,2] = 0.3333$

##### Passo 5

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0.3333 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.0000

**Operation:**

- $L[2,2] = 0.0000$

##### Passo 6

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0.3333 & 0 & -0.1667 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.1667

**Operation:**

- $L[3,2] = -0.1667$

##### Passo 7

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0.3333 & 0 & -0.1667 \\
0.4286 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** 0.4286

**Operation:**

- $L[1,3] = 0.4286$

##### Passo 8

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0.3333 & 0 & -0.1667 \\
0.4286 & 0.1429 & 0
\end{bmatrix}
$$

**Multiplier:** 0.1429

**Operation:**

- $L[2,3] = 0.1429$

##### Passo 9

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0.125 & -0.25 \\
0.3333 & 0 & -0.1667 \\
0.4286 & 0.1429 & 0
\end{bmatrix}
$$

**Operation:**

- $L[3,3] = 0.0000$

#### Achando d

| **Step** | **d**                                       | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ------------------------------------------- | :-------------: | :---------: | :------------: |
|    1     | \[2.5,0,0]                                  |   -20.000000    |  -8.000000  |      2.50      |
|    2     | \[2.5,6.333333333333333,0]                  |   -38.000000    |  -6.000000  |      6.33      |
|    3     | \[2.5,6.333333333333333,-4.857142857142857] |   -34.000000    |  7.000000   |     -4.86      |

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:
**Raízes Encontradas:** $3.99$, $7.99$, $-2.00$

De modo a facilitar a visualização, os valores de `a, b, c` correspondem ao valor anterior do X.

| **Interaction** |  **a**   |  **b**   |   **c**   |                         **Found-X**                         | **Relative-Error** |
| :-------------: | :------: | :------: | :-------: | :---------------------------------------------------------: | :----------------: |
|        1        | 0.000000 | 0.000000 | 0.000000  |          \[3,8.799999999999999,-2.777142857142857]          |     100.000000     |
|        2        | 3.000000 | 8.800000 | -2.777143 | \[4.5531428571428565,8.216685714285715,-1.522951836734693]  |     82.352638      |
|        3        | 4.553143 | 8.216686 | -1.522952 | \[3.778759836734694,7.772757159183674,-2.2481462036151605]  |     32.257438      |
|        4        | 3.778760 | 7.772757 | -2.248146 |  \[4.084605467615161,8.12891999593236,-1.8847588052007658]  |     19.280313      |
|        5        | 4.084605 | 8.128920 | -1.884759 | \[3.9678445474270516,7.938305580824502,-2.050161514998877]  |      8.067789      |
|        6        | 3.967845 | 7.938306 | -2.050162 |  \[4.012225382137928,8.027261339690046,-1.97900698509671]   |      3.595466      |
|        7        | 4.012225 | 8.027261 | -1.979007 |  \[3.995346220054934,7.988487617103306,-2.008565526877553]  |      1.471624      |
|        8        | 3.995346 | 7.988488 | -2.008566 | \[4.001773556617775,8.004725004601958,-1.9965647790035832]  |      0.601070      |
|        9        | 4.001774 | 8.004725 | -1.996565 | \[3.9993234730678133,7.9980973441074505,-2.001361141917416] |      0.239655      |
|       10        | 3.999323 | 7.998097 | -2.001361 |  \[4.00025824957778,8.000756059393105,-1.9994653473662682]  |      0.094815      |
|       11        | 4.000258 | 8.000756 | -1.999465 |  \[3.99990136320329,7.999702402875948,-2.0002086746720344]  |      0.037162      |

---

### Gráfico de Convergência

![[assets/unit02_challenge_quest10_convgraph.png]]
