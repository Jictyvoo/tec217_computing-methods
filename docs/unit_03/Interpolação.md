> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Deve-se realizar a implementação de 2 métodos de interpolação, **Newton** e **Lagrange**.

## Questão 01

    Considere a função f(x) = e^x + sen(x) tabelada como segue:

|  **$x_i$**   |  0  | 0.5  | 1.0  |
| :----------: | :-: | :--: | :--: |
| **$f(x_i)$** | 1.0 | 2.12 | 3.55 |

    Determine o polinômio interpolador de grau 2 **$P_{2}(x)$** analiticamente usando a fórmula de Newton, e avalie f(0.7). Trace o gráfico do polinômio obtido. Implemente o algoritmo de Newton e compare os resultados.

Para auxiliar o desenvolvimento analítico do polinômio de Newton, usa-se a seguinte tabela:

|         |       |                  **Ordem 1**                   |                          **Ordem 2**                           |
| :-----: | :---: | :--------------------------------------------: | :------------------------------------------------------------: |
| $x_{0}$ | $y_0$ | $\delta_{y_{0}} = \frac{y_1 - y_0}{x_1 - x_0}$ | $\delta_{y_{0}}^2 = \frac{\delta y_1 - \delta y_0}{x_2 - x_0}$ |
|  $x_1$  | $y_1$ | $\delta_{y_{1}} = \frac{y_2 - y_1}{x_2 - x_1}$ |                                                                |
|  $x_2$  | $y_2$ |                                                |                                                                |

### Analiticamente

|       |        |                  **Ordem 1**                   |                          **Ordem 2**                           |
| :---: | :----: | :--------------------------------------------: | :------------------------------------------------------------: |
|  $0$  |  $1$   |  $\delta_{y_{0}} = \frac{2.12 - 1}{0.5 - 0}$   | $\delta_{y_{0}}^2 = \frac{\delta y_1 - \delta y_0}{x_2 - x_0}$ |
| $0.5$ | $2.12$ | $\delta_{y_{1}} = \frac{3.55 - 2.12}{1 - 0.5}$ |                                                                |
|  $1$  | $3.55$ |                                                |                                                                |

Após o calculo da primeira ordem, temos:

$$
\begin{align}
\delta_{y_{0}} = \frac{2.12 - 1}{0.5 - 0} = 2.24 \\
\delta_{y_{1}} = \frac{3.55 - 2.12}{1 - 0.5} = 2.86 \\
\end{align}
$$

Ao continuar agora o cálculo para ordem 2, temos:

$$
\begin{align}
\delta_{y_{0}}^2 = \frac{\delta y_1 - \delta y_0}{x_2 - x_0} \\
\\
\delta_{y_{0}}^2 = \frac{2.86 - 2.24}{1 - 0}
&\therefore \delta_{y_{0}}^2 = 0.62
\end{align}
$$

Assim, no final a tabela ficará da seguinte forma:

| **x** | **y**  |       **Ordem 1**       |        **Ordem 2**        |
| :---: | :----: | :---------------------: | :-----------------------: |
|  $0$  |  $1$   | $\delta_{y_{0}} = 2.24$ | $\delta_{y_{0}}^2 = 0.62$ |
| $0.5$ | $2.12$ | $\delta_{y_{1}} = 2.86$ |                           |
|  $1$  | $3.55$ |                         |                           |

Após a montagem da tabela, é possível montar o polinômio de Newton:

$$
\begin{align}
P_2(x) = 1 + (x-0) \cdot 2.24 + (x-0) \cdot (x-0.5) \cdot 0.62 \\
P_2(x) = 1 + 2.24x + x^2 \cdot 0.62 - 0.31x \\
\\
\therefore P_2(x) = 0.62x^2 + 1.93x + 1
\end{align}
$$

Tendo o polinômio de Newton montado, é possível avaliar o valor do mesmo para $x=0.7$:

$$
\begin{align}
P(0.7) = 0.62 \cdot 0.7^2 + 1.93 \cdot 0.7 + 1 \\
\therefore P(0.7) = 2.6548
\end{align}
$$

### Algoritmo

```octave
function [yint, err] = newtonInterpolationMethod(x, y, xx)
    err = ""; % Initialize err to an empty string
    sample_size = length(x);

    % Check for consistent input
    if sample_size != length(y)
        err = "x and y must have the same length";
        yint = NaN;
        return;
    endif

    % Initialize b as a zero 2D matrix
    b = zeros(sample_size);

    % Set the first column of b as y
    b(:, 1) = y;

    % Calculate finite differences
    for ctrlIdx = 2:sample_size
        for index = 1:(sample_size - ctrlIdx + 1)
            b(index, ctrlIdx) = (b(index + 1, ctrlIdx - 1) - b(index, ctrlIdx - 1)) / (x(index + ctrlIdx - 1) - x(index));
        endfor
    endfor

    % Apply finite differences to interpolate
    xt = 1;
    yint = b(1, 1);
    for index = 1:(sample_size - 1)
        xt *= (xx - x(index));
        yint += (b(1, index + 1) * xt);
    endfor

endfunction
```

#### Chamada ao algoritmo

```octave
x = [0; 0.5; 1];
y = [1; 2.12; 3.55];

[foundY, err] = newtonInterpolationMethod(x, y, 0.7);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: NewtonInterpolationMethod\n');
fprintf('Result:'), disp(foundY);
```

### Resultados

Tanto os cálculos realizados analiticamente, quanto os valores obtidos através do algoritmo exposto, tiveram como resultado o valor de 2.6548 para $x=0.7$, provando assim que a implementação do algoritmo se mostrou correta.

![[unit03_quest01_compare_newtoninterpolation.png]]

Ao analisar o gráfico que compara o polinômio $P_2(x)$ com a função $f(x) = e^x + \sin(x)$, podemos observar que, em geral, as duas curvas são bastante próximas, especialmente para valores de $x$ em torno de 0 e positivos.

Porém como pode ser observado na imagem abaixo, para valores negativos de $x$, as curvas começam a se afastar, indicando que o polinômio interpolador de grau 2 $P_2(x)$ não reproduz perfeitamente o comportamento da função $f(x)$ nessa região específica.

![[unit03_quest01_compare_newtoninterpolation_geogebra.png]]

É importante notar que a função $f(x) = e^x + \sin(x)$ é uma função analítica suave, enquanto o polinômio $P_2(x)$ é uma aproximação polinomial.

## Questão 02

    Considere a função f(x) definida nos seguintes pontos, conforme a tabela:

|  **$x_i$**   |  0  | 0.5 | 1.0 |
| :----------: | :-: | :-: | :-: |
| **$f(x_i)$** | 1.3 | 2.5 | 0.9 |

    Determine o polinômio interpolador de grau 2 **$P_{2}(x)$** analiticamente usando a fórmula de Lagrange Li, trace a curva de cada um dos polinômios Li num mesmo gráfico. Acrescente ao gráfico a curva do polinômio $P_{2}(x)$. Implemente o algoritmo de Lagrange e compare os resultados.

Para auxiliar o desenvolvimento analítico do polinômio de Lagrange, usa-se uma formula dependente da ordem em que deseja-se obter o polinômio:

$$
\begin{align}
p_n(x) = y_0 \cdot L_0(x) + y_1 \cdot L_1(x) + \dots + y_n \cdot L_n(x) \\
\\
L(x) = \prod_{i=0; i \ne k}^{n} \left(\frac{x - x_i}{x_k - x_i} \right)
\end{align}
$$

Com o intuito de simplificar o desenvolvimento analítico, foi montada a seguinte tabela:

| **Ordem** |          **Simples**           |                  **Lagrange**                   |
| :-------: | :----------------------------: | :---------------------------------------------: |
|   $1_a$   |   $f_1(x) = a_{1} + a_{2}x$    |       $f_1(x) = L_1 f(x_1) + L_2 f(x_2)$        |
|   $2_a$   | $f_2(x) = a_1 + a_2x + a_3x^2$ | $f_2(x) = L_1 f(x_1) + L_2 f(x_2) + L_3 f(x_3)$ |

### Analiticamente

$$
\begin{align}
L_0(x) = \frac{(x - x_1)(x - x_2)}{(x_0 - x_1)(x_0 - x_2)} \\
\text{Substituindo: } L_0(x) = \frac{(x - 0.5)(x - 1)}{(0 - 0.5)(0 - 1)} \\
\therefore L_0(x) = \frac{x^2 - 1.5x + 0.5}{0.5} \\
\\
L_1(x) = \frac{(x - x_0)(x - x_2)}{(x_1 - x_0)(x_1 - x_2)} \\
\text{Substituindo: } L_1(x) = \frac{(x - 0)(x - 1)}{(0.5 - 0)(0.5 - 1)} \\
\therefore L_1(x) = \frac{x^2 - x}{-0.25} \\
\\
L_2(x) = \frac{(x - x_0)(x - x_1)}{(x_2 - x_0)(x_2 - x_1)} \\
\text{Substituindo: } L_2(x) = \frac{(x - 0)(x - 0.5)}{(1 - 0)(1 - 0.5)} \\
\therefore L_2(x) = \frac{x^2 - 0.5x}{0.5} \\
\end{align}
$$

Após a montagem dos valores de $L_i$, é possível montar o polinômio de Lagrange:

$$
\begin{align}
P_2(x) = y_0 \cdot L_0(x) + y_1 \cdot L_1(x) + y_2 \cdot L_2(x) \\
P_2(x) = 1.3 \cdot \frac{x^2 - 1.5x + 0.5}{0.5} + 2.5 \cdot \frac{x^2 - x}{-0.25} + 0.9 \cdot \frac{x^2 - 0.5x}{0.5} \\
\\
P_2(x) = 1.3 \cdot (2x^2 - 3x + 1) + 2.5 \cdot (-4x^2 + 4x) + 0.9 \cdot (2x^2 - x) \\
P_2(x) = 2.6x^2 - 3.9x + 1.3 - 10x^2 + 10x + 1.8x^2 - 0.9x \\
\therefore P_2(x) = -5.6x^2 + 5.2x + 1.3 \\
\end{align}
$$

Tendo o polinômio de Lagrange montado, é possível avaliar o valor do mesmo para $x=0.8$:

$$
\begin{align}
P_2(0.8) = -5.6 \cdot (0.8^2) + (5.2 \cdot 0.8) + 1.3 \\
\therefore P(0.8) = 1.876
\end{align}
$$

### Algoritmo

```octave
function [result, err] = lagrangeInterpolationMethod(x, y, xx)
    err = ""; % Initialize err to an empty string
    sampleSize = length(x);

    % Check for consistent input
    if sampleSize != length(y)
        err = "x and y must have the same length";
        result = NaN;
        return;
    endif

    result = 0;
    for ctrlIndex = 1:sampleSize
        term = y(ctrlIndex);

        % Start to perform the prod calculation
        for index = 1:sampleSize
            if ctrlIndex ~= index
                term *= (xx - x(index)) / (x(ctrlIndex) - x(index));
            end
        end
        result += term;
    end

endfunction
```

#### Chamada ao algoritmo

```octave
x = [0; 0.5; 1];
y = [1.3; 2.5; 0.9];

[foundY, err] = lagrangeInterpolationMethod(x, y, 0.8);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: lagrangeInterpolationMethod\n');
fprintf('Result:'), disp(foundY);
```

### Resultados

Tanto os cálculos realizados analiticamente, quanto os valores obtidos através do algoritmo exposto, tiveram como resultado o valor de 1.876 para $x=0.8$, provando assim que a implementação do algoritmo se mostrou correta.

![[unit03_quest02_compare.png]]

## Questão 03

    Encontre os coeficientes dos polinômios interpoladores das questões 1 e 2 através do sistema de equações simultâneas obtido com os pontos disponíveis (matriz de Vandermonde), e trace o gráfico dos polinômios encontrados.

A matriz de Vandermonde \(V\) é dada por:

$$
V = \begin{bmatrix}
0 & 0 & 1 \\
0.25 & 0.5 & 1 \\
1 & 1 & 1 \\
\end{bmatrix}
$$

A equação do sistema linear $Ax = b$ associada a essa matriz de Vandermonde pode ser expressa como:

$$
\begin{bmatrix}
0 & 0 & 1 \\
0.25 & 0.5 & 1 \\
1 & 1 & 1 \\
\end{bmatrix}
\begin{bmatrix}
c_1 \\
c_2 \\
c_3 \\
\end{bmatrix}
=
\begin{bmatrix}
f(x_1) \\
f(x_2) \\
f(x_3) \\
\end{bmatrix}
$$

Onde os valores de $f(x_i)$ são substituídos pelos valores $y_i$ das questões 1 e 2.

Executando os cálculos através de GaussJordan, têm-se:

### Para a questão 01

$$
\begin{bmatrix}
0 & 0 & 1 \\
0.25 & 0.5 & 1 \\
1 & 1 & 1 \\
\end{bmatrix}
\begin{bmatrix}
c_1 \\
c_2 \\
c_3 \\
\end{bmatrix}
=
\begin{bmatrix}
1 \\
2.12 \\
3.55 \\
\end{bmatrix}
$$

Os valores encontrados para as raízes foram:

$$
\begin{align}
\begin{bmatrix}
0.6199999999999992 \\
1.9300000000000006 \\
1.0000000000000000
\end{bmatrix}
\\
\\
\text{Polinômio: } \\
P_2(x) = 0.6199999999999992x^2 + 1.9300000000000006x + 1.0000000000000000
\end{align}
$$

O polinômio obtido a partir da matriz de vandermonde se assemelha o suficiente ao obtido manualmente e algoritmicamente na questão 01.

### Para a questão 02

$$
\begin{bmatrix}
0 & 0 & 1 \\
0.25 & 0.5 & 1 \\
1 & 1 & 1 \\
\end{bmatrix}
\begin{bmatrix}
c_1 \\
c_2 \\
c_3 \\
\end{bmatrix}
=
\begin{bmatrix}
1.3 \\
2.5 \\
0.9 \\
\end{bmatrix}
$$

Os valores encontrados para as raízes foram:

$$
\begin{align}
\begin{bmatrix}
-5.600000000000000 \\
5.199999999999999 \\
1.300000000000000
\end{bmatrix}
\\
\\
\text{Polinômio: } \\
P_2(x) = -5.600000000000000x^2 + 5.199999999999999x + 1.300000000000000
\end{align}
$$

O polinômio obtido a partir da matriz de vandermonde se assemelha o suficiente ao obtido manualmente e algoritmicamente na questão 02.

Abaixo é possível ter um comparativo entre os gráficos de ambos os polinômios:
![[unit03_quest03_polynouns_comparation.png]]
