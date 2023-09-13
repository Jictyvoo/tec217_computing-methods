> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Deve-se fazer um comparativo dos diferentes métodos estudados em classe em cada um dos problemas abaixo. O resultado de cada questão deve estar uma tabela que demonstre a evolução dos métodos:
- métodos de quebra: 
	- iteração
	- limite inferior e superior
	- xr - raiz aproximada
	- f(xr)
	- erro aproximado e verdadeiro.
- métodos abertos: 
	- iteração
	- xr - raiz aproximada
	- f(xr)
	- erro aproximado e verdadeiro.

## Questão 01

	No estudo de crescimento populacional, engenheiros de transporte podem achar necessário determinar separadamente a tendência de crescimento da população de uma cidade e do subúrbio adjacente. A população da área urbana está diminuindo com o tempo de acordo com a função 1 [Pu(t)]. Enquanto a população suburbana cresce de acordo com o modelo descrito na função 2 [Ps(t)]. Determine o tempo e os valores de Pu(t) e Ps(t) quando os subúrbios forem 20% maiores que a cidade.

$$
\begin{align}
P_u(t) = P_{u,max} \cdot e^{-k_{u} \cdot t} + P_{u,min} \\
P_s(t) = \frac{P_{s,max}}{1+[\frac{P_{s,max}}{P_0} - 1] \cdot e^{-k_{s} \cdot t}} \\
\\
\\
&\text{Constantes:} \\
&P_{u,max} = 80000\text{pessoas} \\
&k_u = 0.05\text{/ano} \\
&P_{u,min} = 110000\text{pessoas} \\
&P_{s,max} = 320000\text{pessoas} \\
&P_0 = 10000\text{pessoas} \\
&k_s = 0.09\text{/ano}
\end{align}
$$
### Análise Inicial
Primeiro faz-se necessário configurar as equações para que a situação desejada seja aplicada:
$$
P_s(t) = 1.2 \cdot P_u(t)
$$

Agora, substituímos as expressões para Pu(t) e Ps(t) nas equações:
$$
\frac{P_{s,max}}{1+\left[\frac{P_{s,max}}{P_0} - 1\right] \cdot e^{-k_{s} \cdot t}} = 1.2 \cdot \left(P_{u,max} \cdot e^{-k_{u} \cdot t} + P_{u,min}\right)
$$

Após isso, é realizada a substituição das constantes:

$$
\begin{align}
\frac{320000}{1+\left[\frac{320000}{10000} - 1\right] \cdot e^{-0.09 \cdot t}} = 1.2 \cdot \left(80000 \cdot e^{-0.05 \cdot t} + 110000\right) \\
\\

\frac{320000}{1+31 \cdot e^{-0.09 \cdot t}} - (96000 \cdot e^{-0.05 \cdot t}) + 132000 = 0 \\
\\

f(t) = 0 \\
f(t) = \frac{320}{1 + (31 \cdot e^{-0.09t})} - 96 \cdot e^{-0.05t} - 132
\end{align}
$$

Tendo a função definida, foi realizada a análise gráfica dessa mesma função, de modo a verificar posteriormente os resultados.

![[01-grafical_analysis.png]]
Baseado no gráfico acima, é possível verificar que a função possui apenas uma raiz, possibilitando assim que possamos iniciar os métodos com valores entre 0 e 100.

Vale ressaltar também, que para cada um dos métodos a serem aplicados, foi utilizada a tolerância como sendo **1e-8**.
### Método da Bisseção
Para a aplicação do método da Falsa Posição, foram utilizados como valores de entrada **0** e **100**.
#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
Pumax = 80000;
ku = 0.05;
Pumin = 110000;
Psmax = 320000;
P0 = 10000;
ks = 0.09;

% Define the function
fX = @(x) (Psmax / (1 + (((Psmax / P0) - 1) * e^(-ks * x)))) - (1.2 * ((Pumax * e^(-ku * x)) + Pumin));

% Define function parameters
a = 0;
b = 100;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = bisectionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[bisectionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[bisectionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Bisseção.

| **Interaction** |   **a**   |    **b**   | **Found-X** |    **f(x)**   | **Relative-Error** |
|:---------------:|:---------:|:----------:|:-----------:|:-------------:|:------------------:|
|        1        |  0.000000 | 100.000000 |  50.000000  |  98147.974716 |      1.000000      |
|        2        |  0.000000 |  50.000000 |  25.000000  | -84516.926498 |      25.000000     |
|        3        | 25.000000 |  50.000000 |  37.500000  |  8560.314651  |      12.500000     |
|        4        | 25.000000 |  37.500000 |  31.250000  | -40300.841776 |      6.250000      |
|        5        | 31.250000 |  37.500000 |  34.375000  | -16171.256882 |      3.125000      |
|        6        | 34.375000 |  37.500000 |  35.937500  |  -3834.523586 |      1.562500      |
|        7        | 35.937500 |  37.500000 |  36.718750  |  2361.873142  |      0.781250      |
|        8        | 35.937500 |  36.718750 |  36.328125  |  -737.375631  |      0.390625      |
|        9        | 36.328125 |  36.718750 |  36.523438  |   812.084759  |      0.195312      |
|        10       | 36.328125 |  36.523438 |  36.425781  |   37.301187   |      0.097656      |
|        11       | 36.328125 |  36.425781 |  36.376953  |  -350.052109  |      0.048828      |
|        12       | 36.376953 |  36.425781 |  36.401367  |  -156.378990  |      0.024414      |
|        13       | 36.401367 |  36.425781 |  36.413574  |   -59.539760  |      0.012207      |
|        14       | 36.413574 |  36.425781 |  36.419678  |   -11.119498  |      0.006104      |
|        15       | 36.419678 |  36.425781 |  36.422729  |   13.090792   |      0.003052      |
|        16       | 36.419678 |  36.422729 |  36.421204  |    0.985634   |      0.001526      |
|        17       | 36.419678 |  36.421204 |  36.420441  |   -5.066935   |      0.000763      |
|        18       | 36.420441 |  36.421204 |  36.420822  |   -2.040652   |      0.000381      |
|        19       | 36.420822 |  36.421204 |  36.421013  |   -0.527509   |      0.000191      |
|        20       | 36.421013 |  36.421204 |  36.421108  |    0.229062   |      0.000095      |
|        21       | 36.421013 |  36.421108 |  36.421061  |   -0.149223   |      0.000048      |
|        22       | 36.421061 |  36.421108 |  36.421084  |    0.039919   |      0.000024      |
|        23       | 36.421061 |  36.421084 |  36.421072  |   -0.054652   |      0.000012      |
|        24       | 36.421072 |  36.421084 |  36.421078  |   -0.007366   |      0.000006      |
|        25       | 36.421078 |  36.421084 |  36.421081  |    0.016277   |      0.000003      |
|        26       | 36.421078 |  36.421081 |  36.421080  |    0.004455   |      0.000001      |
|        27       | 36.421078 |  36.421080 |  36.421079  |   -0.001456   |      0.000001      |
|        28       | 36.421079 |  36.421080 |  36.421080  |    0.001500   |      0.000000      |
|        29       | 36.421079 |  36.421080 |  36.421079  |    0.000022   |      0.000000      |
|        30       | 36.421079 |  36.421079 |  36.421079  |   -0.000717   |      0.000000      |
|        31       | 36.421079 |  36.421079 |  36.421079  |   -0.000347   |      0.000000      |
|        32       | 36.421079 |  36.421079 |  36.421079  |   -0.000163   |      0.000000      |
|        33       | 36.421079 |  36.421079 |  36.421079  |   -0.000070   |      0.000000      |
|        34       | 36.421079 |  36.421079 |  36.421079  |   -0.000024   |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[01.convergencia_bissecao.png]]


### Método da Falsa Posição
Para a aplicação do método da Falsa Posição, foram utilizados como valores de entrada **0** e **100**.
#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
Pumax = 80000;
ku = 0.05;
Pumin = 110000;
Psmax = 320000;
P0 = 10000;
ks = 0.09;

% Define the function
fX = @(x) (Psmax / (1 + (((Psmax / P0) - 1) * e^(-ks * x)))) - (1.2 * ((Pumax * e^(-ku * x)) + Pumin));

% Define function parameters
a = 0;
b = 100;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = falsePositionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[falsePositionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[falsePositionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Falsa Posição.

| **Interaction** |   **a**   |    **b**   | **Found-X** |    **f(x)**   | **Relative-Error** |
|:---------------:|:---------:|:----------:|:-----------:|:-------------:|:------------------:|
|        1        |  0.000000 | 100.000000 |  53.942558  | 119280.286853 |      1.000000      |
|        2        |  0.000000 |  53.942558 |  34.865594  | -12309.945011 |      19.076964     |
|        3        | 34.865594 |  53.942558 |  36.650198  |  1817.890994  |      1.784603      |
|        4        | 34.865594 |  36.650198 |  36.420565  |   -4.081795   |      0.229633      |
|        5        | 36.420565 |  36.650198 |  36.421079  |   -0.000548   |      0.000514      |
|        6        | 36.421079 |  36.650198 |  36.421079  |   -0.000000   |      0.000000      |
|        7        | 36.421079 |  36.650198 |  36.421079  |   -0.000000   |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[01.convergencia_falsa_posicao.png]]

### Método da Secante
Para a aplicação do método da Secante, foram utilizados como valores de entrada **0** e **100**.
#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
Pumax = 80000;
ku = 0.05;
Pumin = 110000;
Psmax = 320000;
P0 = 10000;
ks = 0.09;

% Define the function
fX = @(x) (Psmax / (1 + (((Psmax / P0) - 1) * e^(-ks * x)))) - (1.2 * ((Pumax * e^(-ku * x)) + Pumin));

% Define function parameters
a = 0;
b = 100;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = secantMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[secantMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[secantMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Secante.

| **Interaction** |   **a**   |    **b**   | **Found-X** |    **f(x)**   | **Relative-Error** |
|:---------------:|:---------:|:----------:|:-----------:|:-------------:|:------------------:|
|        1        |  0.000000 | 100.000000 |  53.942558  | 119280.286853 |      1.000000      |
|        2        | 53.942558 |  0.000000  |  34.865594  | -12309.945011 |      0.547157      |
|        3        | 34.865594 |  53.942558 |  36.650198  |  1817.890994  |      0.048693      |
|        4        | 36.650198 |  34.865594 |  36.420565  |   -4.081795   |      0.006305      |
|        5        | 36.420565 |  36.650198 |  36.421079  |   -0.000548   |      0.000014      |
|        6        | 36.421079 |  36.420565 |  36.421079  |    0.000000   |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[01.convergencia_secante.png]]

### Método da Iteração Linear

$$
\begin{align}
& f(t) = \frac{320}{1 + (31 \cdot e^{-0.09t})} - 96 \cdot e^{-0.05t} - 132 \\
\\
\text{Achando o g(t):} \\
& 9.6e^{-0.05t} + 13.2 = \frac{32}{1 + 31e^{-0.09t}} \\
& 1 + 31e^{-0.09t} = \frac{32}{9.6e^{-0.05t} + 13.2} \\
& e^{-0.09t} = \frac{1}{31} \cdot \left(\frac{32}{9.6e^{-0.05t} + 13.2} - 1\right) \\
\\
\\

\text{Após a reorganização, isola o t:} \\
& -0.09t = \ln\left(\frac{1}{31} \left(\frac{32}{9.6e^{-0.05t} + 13.2} - 1\right)\right) \\
& g(t) = \frac{\ln\left(\frac{1}{31} \left(\frac{32}{9.6e^{-0.05t} + 13.2} - 1\right)\right)}{-0.09}
\end{align}
$$

#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
Pumax = 80000;
ku = 0.05;
Pumin = 110000;
Psmax = 320000;
P0 = 10000;
ks = 0.09;

% Define the function
fX = @(x) (32 / (1 + (31 * exp(-ks * x)))) - 13.2 - (9.6 * exp(-ku * x));
gX = @(x) log(((32 / ((9.6 * exp(-0.05 * x)) + 13.2)) - 1) / 31) / -0.09;

% Define function parameters
x0 = 0;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[linearIterationMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[linearIterationMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```
#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Iteração Linear.

| **Interaction** |   **a**  | **Found-X** |   **f(x)**   | **Relative-Error** |
|:---------------:|:--------:|:-----------:|:------------:|:------------------:|
|        1        | 0.000000 |  48.239381  | 87394.692534 |      1.000000      |
|        2        | 0.000000 |  35.448395  | -7707.139467 |      0.360834      |
|        3        | 0.000000 |  36.529246  |  858.168144  |      0.029589      |
|        4        | 0.000000 |  36.409368  |  -92.909412  |      0.003292      |
|        5        | 0.000000 |  36.422351  |   10.089431  |      0.000356      |
|        6        | 0.000000 |  36.420941  |   -1.095293  |      0.000039      |
|        7        | 0.000000 |  36.421094  |   0.118908   |      0.000004      |
|        8        | 0.000000 |  36.421078  |   -0.012909  |      0.000000      |
|        9        | 0.000000 |  36.421080  |   0.001401   |      0.000000      |
|        10       | 0.000000 |  36.421079  |   -0.000152  |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[01.convergencia_linear_iteraction.png]]

### Método de Newton Raphson
Devido a problemas com os valores presentes na função, fez-se necessário realizar a redução dos mesmos, pois estavam causando overflow e impedindo que as operações de cálculo fossem efetuadas corretamente. Portanto, utilizou-se as seguintes equações:
$$
\begin{align}
& f(t) = \frac{32}{1 + (31 \cdot e^{-0.09t})} - 9.6 \cdot e^{-0.05t} - 13.2 \\
& f'(t) = 0.48 e^{-0.05 \cdot t} + \frac{89.28 e^{0.09 \cdot x}}{(e^{0.09 \cdot x}+31)^2}
\end{align}
$$

#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
Pumax = 80000;
ku = 0.05;
Pumin = 110000;
Psmax = 320000;
P0 = 10000;
ks = 0.09;

% Define the function
fX = @(x) (32 / (1 + (31 * exp(-ks * x)))) - 13.2 - (9.6 * exp(-ku * x));
f1X = @(x) (0.48 * exp(-0.05 * x)) + ((89.28 * exp(0.09 * x)) / (exp(0.09 * x) + 31)^2);

% Define function parameters
x0 = 0;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[newtonRaphsonMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[newtonRaphsonMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```
#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método de Newton Raphson.

| **Interaction** |   **a**  | **Found-X** |  **f(x)** | **Relative-Error** |
|:---------------:|:--------:|:-----------:|:---------:|:------------------:|
|        1        | 0.000000 |  38.435262  |  1.596533 |      1.000000      |
|        2        | 0.000000 |  36.414675  | -0.005081 |      0.055488      |
|        3        | 0.000000 |  36.421079  |  0.000000 |      0.000176      |
|        4        | 0.000000 |  36.421079  | -0.000000 |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[01.convergencia_newton_raphson.png]]



---

## Questão 02

	Uma carga total Q está uniformemente distribuída ao redor de um condutor circular de raio a. Uma carga q está localizada a uma distância x do centro do anel. A força exercida pelo anel na carga é dada por F. Encontre a distância x onde a força é de 1,25 N de acordo com as constantes passadas.

$$
\begin{align}
F = \frac{1}{4 \cdot \pi \cdot e_0} \cdot \frac{q \cdot Q \cdot x}{(x^2 + a^2)^{3/2}} \\
\\
\\
&\text{Constantes:} \\
&e_{0} = 8.9 \cdot 10^{-12} C^2/Nm^2 \\
&F = 1.25N \\
&a = 0.85m \\
&q = Q = 2 \cdot 10^{-5}C
\end{align}
$$
### Análise Inicial
Primeiro faz-se necessário configurar as equações para que a situação desejada seja aplicada:

$$
f(x) = \left(\frac{1}{4 \cdot \pi \cdot e_0} \cdot \frac{q \cdot Q \cdot x}{(x^2 + a^2)^{3/2}}\right) - F
$$

Após isso, é realizada a substituição das constantes:

$$
\begin{align}
f(x) = \left(\frac{1}{4 \cdot \pi \cdot (8.9 \cdot 10^{-12})} \cdot \frac{(2 \cdot 10^{-5})^2 \cdot x}{(x^2 + 0.85^2)^{3/2}}\right) - 1.25 \\
\end{align}
$$

Tendo a função definida, foi realizada a análise gráfica dessa mesma função, de modo a verificar posteriormente os resultados.

![[02-grafical_analysis.png]]
Baseado no gráfico acima, é possível verificar que a função possui apenas uma raiz, possibilitando assim que possamos iniciar os métodos com valores entre 0 e 0.5.

Vale ressaltar também, que para cada um dos métodos a serem aplicados, foi utilizada a tolerância como sendo **1e-8**.
### Método da Bisseção
Para a aplicação do método da Falsa Posição, foram utilizados como valores de entrada **0** e **0.5**.
#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
e0 = 8.9e-12;
F = 1.25;
r = 0.85;
q = 2e-5;
Q = q;

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;

% Define function parameters
a = 0;
b = 0.5;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = bisectionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[bisectionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[bisectionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Bisseção.

| **Interaction** |   **a**  |   **b**  | **Found-X** |  **f(x)** | **Relative-Error** |
|:---------------:|:--------:|:--------:|:-----------:|:---------:|:------------------:|
|        1        | 0.000000 | 0.500000 |   0.250000  |  0.035570 |      1.000000      |
|        2        | 0.000000 | 0.250000 |   0.125000  | -0.545022 |      0.125000      |
|        3        | 0.125000 | 0.250000 |   0.187500  | -0.233159 |      0.062500      |
|        4        | 0.187500 | 0.250000 |   0.218750  | -0.092890 |      0.031250      |
|        5        | 0.218750 | 0.250000 |   0.234375  | -0.027136 |      0.015625      |
|        6        | 0.234375 | 0.250000 |   0.242188  |  0.004603 |      0.007812      |
|        7        | 0.234375 | 0.242188 |   0.238281  | -0.011171 |      0.003906      |
|        8        | 0.238281 | 0.242188 |   0.240234  | -0.003260 |      0.001953      |
|        9        | 0.240234 | 0.242188 |   0.241211  |  0.000677 |      0.000977      |
|        10       | 0.240234 | 0.241211 |   0.240723  | -0.001290 |      0.000488      |
|        11       | 0.240723 | 0.241211 |   0.240967  | -0.000306 |      0.000244      |
|        12       | 0.240967 | 0.241211 |   0.241089  |  0.000186 |      0.000122      |
|        13       | 0.240967 | 0.241089 |   0.241028  | -0.000060 |      0.000061      |
|        14       | 0.241028 | 0.241089 |   0.241058  |  0.000063 |      0.000031      |
|        15       | 0.241028 | 0.241058 |   0.241043  |  0.000001 |      0.000015      |
|        16       | 0.241028 | 0.241043 |   0.241035  | -0.000029 |      0.000008      |
|        17       | 0.241035 | 0.241043 |   0.241039  | -0.000014 |      0.000004      |
|        18       | 0.241039 | 0.241043 |   0.241041  | -0.000006 |      0.000002      |
|        19       | 0.241041 | 0.241043 |   0.241042  | -0.000002 |      0.000001      |
|        20       | 0.241042 | 0.241043 |   0.241043  | -0.000000 |      0.000000      |
|        21       | 0.241043 | 0.241043 |   0.241043  |  0.000000 |      0.000000      |
|        22       | 0.241043 | 0.241043 |   0.241043  | -0.000000 |      0.000000      |
|        23       | 0.241043 | 0.241043 |   0.241043  |  0.000000 |      0.000000      |
|        24       | 0.241043 | 0.241043 |   0.241043  |  0.000000 |      0.000000      |
|        25       | 0.241043 | 0.241043 |   0.241043  |  0.000000 |      0.000000      |
|        26       | 0.241043 | 0.241043 |   0.241043  |  0.000000 |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[02.convergencia_bissecao.png]]


### Método da Falsa Posição
Para a aplicação do método da Falsa Posição, foram utilizados como valores de entrada **0** e **0.5**.
#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
e0 = 8.9e-12;
F = 1.25;
r = 0.85;
q = 2e-5;
Q = q;

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;

% Define function parameters
a = 0;
b = 0.5;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = falsePositionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[falsePositionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[falsePositionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Falsa Posição.

| **Interaction** |   **a**  |   **b**  | **Found-X** | **f(x)** | **Relative-Error** |
|:---------------:|:--------:|:--------:|:-----------:|:--------:|:------------------:|
|        1        | 0.000000 | 0.500000 |   0.335185  | 0.321566 |      1.000000      |
|        2        | 0.000000 | 0.335185 |   0.266601  | 0.098773 |      0.068584      |
|        3        | 0.000000 | 0.266601 |   0.247077  | 0.024075 |      0.019524      |
|        4        | 0.000000 | 0.247077 |   0.242408  | 0.005489 |      0.004669      |
|        5        | 0.000000 | 0.242408 |   0.241349  | 0.001231 |      0.001060      |
|        6        | 0.000000 | 0.241349 |   0.241111  | 0.000275 |      0.000238      |
|        7        | 0.000000 | 0.241111 |   0.241058  | 0.000061 |      0.000053      |
|        8        | 0.000000 | 0.241058 |   0.241046  | 0.000014 |      0.000012      |
|        9        | 0.000000 | 0.241046 |   0.241043  | 0.000003 |      0.000003      |
|        10       | 0.000000 | 0.241043 |   0.241043  | 0.000001 |      0.000001      |
|        11       | 0.000000 | 0.241043 |   0.241043  | 0.000000 |      0.000000      |
|        12       | 0.000000 | 0.241043 |   0.241043  | 0.000000 |      0.000000      |
|        13       | 0.000000 | 0.241043 |   0.241043  | 0.000000 |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[02.convergencia_falsa_posicao.png]]

### Método da Secante
Para a aplicação do método da Secante, foram utilizados como valores de entrada **0** e **0.5**.
#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
e0 = 8.9e-12;
F = 1.25;
r = 0.85;
q = 2e-5;
Q = q;

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;

% Define function parameters
a = 0;
b = 0.5;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = secantMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[secantMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[secantMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Secante.

| **Interaction** |   **a**  |   **b**  | **Found-X** |  **f(x)** | **Relative-Error** |
|:---------------:|:--------:|:--------:|:-----------:|:---------:|:------------------:|
|        1        | 0.000000 | 0.500000 |   0.335185  |  0.321566 |      1.000000      |
|        2        | 0.335185 | 0.000000 |   0.266601  |  0.098773 |      0.257253      |
|        3        | 0.266601 | 0.335185 |   0.236195  | -0.019674 |      0.128733      |
|        4        | 0.236195 | 0.266601 |   0.241245  |  0.000816 |      0.020935      |
|        5        | 0.241245 | 0.236195 |   0.241044  |  0.000006 |      0.000834      |
|        6        | 0.241044 | 0.241245 |   0.241043  | -0.000000 |      0.000006      |
|        7        | 0.241043 | 0.241044 |   0.241043  |  0.000000 |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[02.convergencia_secante.png]]

### Método da Iteração Linear

$$
\begin{align}
& f(x) = \left(\frac{1}{4 \cdot \pi \cdot (8.9 \cdot 10^{-12})} \cdot \frac{(2 \cdot 10^{-5})^2 \cdot x}{(x^2 + 0.85^2)^{3/2}}\right) - 1.25 \\
\\
\text{Achando o g(x):} \\
& 1.25 \cdot [(4 \cdot \pi \cdot (8.9 \cdot 10^{-12})) \cdot (x^2 + 0.85^2)^{3/2}]  = (2 \cdot 10^{-5})^2 \cdot x \\
\\
\\

\text{Após a reorganização, isola o x:} \\
& g(x) = \frac{1.25 \cdot [(4 \cdot \pi \cdot (8.9 \cdot 10^{-12})) \cdot (x^2 + 0.85^2)^{3/2}]}{(2 \cdot 10^{-5})^2}
\end{align}
$$

#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
e0 = 8.9e-12;
F = 1.25;
r = 0.85;
q = 2e-5;
Q = q;

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;
gX = @(x) (F * ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) / (q * Q);

% Define function parameters
x0 = 0;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[linearIterationMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[linearIterationMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```
#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Iteração Linear.

| **Interaction** |   **a**  | **Found-X** |  **f(x)** | **Relative-Error** |
|:---------------:|:--------:|:-----------:|:---------:|:------------------:|
|        1        | 0.000000 |   0.214638  | -0.110690 |      1.000000      |
|        2        | 0.000000 |   0.235491  | -0.022555 |      0.088552      |
|        3        | 0.000000 |   0.239818  | -0.004941 |      0.018044      |
|        4        | 0.000000 |   0.240770  | -0.001099 |      0.003953      |
|        5        | 0.000000 |   0.240982  | -0.000245 |      0.000879      |
|        6        | 0.000000 |   0.241029  | -0.000055 |      0.000196      |
|        7        | 0.000000 |   0.241040  | -0.000012 |      0.000044      |
|        8        | 0.000000 |   0.241042  | -0.000003 |      0.000010      |
|        9        | 0.000000 |   0.241043  | -0.000001 |      0.000002      |
|        10       | 0.000000 |   0.241043  | -0.000000 |      0.000000      |
|        11       | 0.000000 |   0.241043  | -0.000000 |      0.000000      |
|        12       | 0.000000 |   0.241043  | -0.000000 |      0.000000      |
|        13       | 0.000000 |   0.241043  | -0.000000 |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[02.convergencia_iteracao_linear.png]]

### Método de Newton Raphson
$$
\begin{align}
& f(x) = \left(\frac{1}{4 \cdot \pi \cdot (8.9 \cdot 10^{-12})} \cdot \frac{(2 \cdot 10^{-5})^2 \cdot x}{(x^2 + 0.85^2)^{3/2}}\right) - 1.25 \\
& f'(t) = \frac{3.57652}{(x^2 + 0.7225)^{3/2}} - \frac{10.7295 \cdot x^2}{(x^2+0.7225)^{5/2}}
\end{align}
$$

#### Chamada ao algoritmo:

```octave
% Define all constants that will be used
e0 = 8.9e-12;
F = 1.25;
r = 0.85;
q = 2e-5;
Q = q;

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;
f1X = @(x) (3.57652 / sqrt((x^2 + 0.7255)^3)) - ((10.7295 * x^2) / sqrt((x^2 + 0.7255)^2));

% Define function parameters
x0 = 0;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[newtonRaphsonMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[newtonRaphsonMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```
#### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método de Newton Raphson.

| **Interaction** |   **a**  | **Found-X** |  **f(x)** | **Relative-Error** |
|:---------------:|:--------:|:-----------:|:---------:|:------------------:|
|        1        | 0.000000 |   0.215976  | -0.104875 |      1.000000      |
|        2        | 0.000000 |   0.238661  | -0.009628 |      0.095052      |
|        3        | 0.000000 |   0.240856  | -0.000751 |      0.009113      |
|        4        | 0.000000 |   0.241028  | -0.000057 |      0.000714      |
|        5        | 0.000000 |   0.241042  | -0.000004 |      0.000055      |
|        6        | 0.000000 |   0.241043  | -0.000000 |      0.000004      |
|        7        | 0.000000 |   0.241043  | -0.000000 |      0.000000      |
|        8        | 0.000000 |   0.241043  | -0.000000 |      0.000000      |
|        9        | 0.000000 |   0.241043  | -0.000000 |      0.000000      |

##### Gráfico de Convergência
Com base na tabela gerada anteriormente, é possível criar o gráfico de convergência da saída do método
![[02.convergencia_newton_raphson.png]]

---

## Conclusão

Os métodos da bisseção, da falsa posição e da secante mostraram-se eficazes e capazes de convergir para soluções com precisão razoável. Além disso, eles não exigiram um esforço inicial muito maior em comparação com os outros dois métodos.

No entanto, é importante observar que o método da bisseção possui uma desvantagem notável, que é a oscilação inicial dos valores da função durante as primeiras iterações, como apresentado nos gráficos após a execução do método.

Essa oscilação inicial ocorre porque o método da bisseção divide o intervalo pela metade em cada iteração, e o ponto médio do intervalo é escolhido como a próxima estimativa. No início, essa estimativa pode estar longe da raiz real, resultando em valores da função que alternam de um lado para o outro da raiz, produzindo uma oscilação na convergência.

O método da iteração linear, embora tenha demonstrado ser uma abordagem razoável em termos de iterações, acabou sendo mais lento do que os métodos da secante e da falsa posição, que são mais fáceis de aplicar.

Por outro lado, o método de Newton-Raphson se destacou em termos de convergência rápida, requerendo menos iterações para encontrar uma solução. No entanto, vale ressaltar que ele exige o cálculo da derivada da função, o que pode ser complicado em alguns casos.

Em suma, todos os métodos foram capazes de encontrar uma raiz aproximada para a função dada, e os métodos de Newton-Raphson e Secante se destacaram em termos de quantidade de iterações necessárias. A escolha do método ideal depende da natureza da função a qual o método será aplicado.

Para funções mais complexas, como a da questão 1, os métodos da Falsa Posição e da Secante são recomendados devido à facilidade de aplicação e convergência. Enquanto os métodos de Newton e da Iteração Linear são mais adequados para casos em que os métodos da Falsa Posição e da Secante não convergem facilmente.
