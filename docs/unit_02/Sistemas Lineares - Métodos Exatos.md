> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Deve ser realizada a implementação e execução do método de Eliminação de Gauss, de forma ingênua e com pivoteamento parcial.

## Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% gaussEliminationMethod.m
function [det, foundRoots, err] = gaussEliminationMethod(inputMatrix, usePivot)
    err = ""; % Initialize err to an empty string

    % Checking if the given matrix is a square one
    [totalEquations, eqLength] = size(inputMatrix);
    if totalEquations != eqLength -1
        err = "the given matrix isn't a square one";
        det = NaN;
        foundRoots = NaN;
        return;
    end

    % If pivot option is enabled, do the pivot
    if usePivot
        inputMatrix = pivot(inputMatrix, 1, 1);
    end
    for eqIndex = 1:totalEquations
        if eqIndex == totalEquations
            break;
        end

        % Eliminate the next equations
        for index = eqIndex + 1:totalEquations
            divisionFactor = inputMatrix(index, eqIndex) / inputMatrix(eqIndex, eqIndex);
            inputMatrix(index, :) -= divisionFactor * inputMatrix(eqIndex, :);
        end
    end

    % Calculate determinant
    det = 1;
    for index = 1:totalEquations
        det *= inputMatrix(index, index);
    end

    % Calculate roots from bottom-up
    foundRoots = zeros(1, totalEquations);
    for eqIndex = totalEquations:-1:1
        sum = inputMatrix(eqIndex, totalEquations + 1);
        for i = eqIndex + 1:totalEquations
            sum -= inputMatrix(eqIndex, i) * foundRoots(i);
        end

        foundRoots(eqIndex) = sum / inputMatrix(eqIndex, eqIndex);
    end
end

function [matrix] = pivot(matrix, currentIndex, lookupColumn)
    % Find the index of the maxValue value in column
    maxValue = abs(matrix(currentIndex, lookupColumn));
    swapIndex = currentIndex;
    [totalEquations, eqLength] = size(matrix)

    for index = currentIndex + 1:totalEquations
        absValue = abs(matrix(index, lookupColumn));
        if absValue > maxValue
            swapIndex = index;
            maxValue = absValue;
        end
    end

    % Swap the current row and the maxValue row if necessary
    if swapIndex ~= currentIndex
        temp = matrix(swapIndex, :);
        matrix(swapIndex, :) = matrix(currentIndex, :);
        matrix(currentIndex, :) = temp;
    end
end
```

## Questão 01

    Resolva MANUALMENTE o sistema por Eliminação de Gauss ingênua. Mostre todos os passos dos cálculos. Substitua os resultados nas equações originais para verificar as respostas. Implemente o algoritmo de Eliminação de Gauss ingênua. Aplique o script ao sistema. A saída numérica deve mostrar todos os passos.

$$
\begin{align}
10x_1 + 2x_2 − 1x_3 = 27 \\
−3x_1 − 5x_2 + 2x_3 = −61,5 \\
1x_1 + 1x_2 + 6x_3 = −21,5 \\
\end{align}
$$

Inicialmente realizou-se a adaptação do sistema de equações, para um sistema de matriz, seguindo o modelo \[A | b\] . A matriz aumentada \[A | b\] é representada como:

$$
\begin{matrix}
10 & 2 & -1 \\
-3 & -5 & 2 \\
1 & 1 & 6
\end{matrix}
\begin{matrix}
| \\
| \\
|
\end{matrix}
\begin{matrix}
27 \\
-61.5 \\
-21.5
\end{matrix}
$$

Em seguida, faz-se necessário realizar a transformação dessa matriz em uma matriz triangular superior. De modo a seguir manualmente o que é proposto pelo algoritmo, deve-se utilizar sempre a primeira linha para obter o primeiro pivô.

Para a segunda linha, de modo a eliminar a primeira coluna, assume-se o multiplicador atribuído à primeira linha como sendo `-3/10`

$$
L2 = L2 - (-3/10) \cdot L1
$$

$$
\begin{matrix}
10  &  2    & -1   \\
0   & -\frac{22}{5} & \frac{17}{10} \\
1   &  1    &  6
\end{matrix}

\begin{matrix}
| \\
| \\
|
\end{matrix}

\begin{matrix}
27 \\
-\frac{267}{5}\\
-21.5
\end{matrix}
$$

Em seguida, assume-se um novo valor para o multiplicador, para eliminar a primeira coluna da terceira linha. O novo valor do multiplicador se torna `1/10`

$$
L3 = L3 - (1/10) \cdot L1
$$

$$
\begin{matrix}
10  &  2    & -1    \\
0   & -\frac{22}{5} & \frac{17}{10}  \\
0   &  \frac{4}{5}    &  \frac{61}{10}
\end{matrix}

\begin{matrix}
| \\
| \\
|
\end{matrix}

\begin{matrix}
27 \\
-\frac{267}{5}\\
-\frac{121}{5}
\end{matrix}
$$

Por fim, é realizada a eliminação da segunda coluna da terceira linha, assumindo assim o novo valor do multiplicador como `-2/11`. Vale lembrar que nessa etapa é usada a segunda linha para a subtração, ao invés da primeira linha, como previamente.

$$
L3 = L3 - (-2/11) \cdot L2
$$

$$
\begin{matrix}
10  &  2    & -1   \\
0   & -\frac{22}{5} & \frac{17}{10}  \\
0   &  0    &  \frac{141}{22}
\end{matrix}

\begin{matrix}
| \\
| \\
|
\end{matrix}

\begin{matrix}
27 \\
-\frac{267}{5}\\
-\frac{373}{11}
\end{matrix}
$$

Agora temos uma matriz triangular superior. Podemos trabalhar de baixo para cima para encontrar as raízes do sistema

$$
\begin{align}
x_3 \text{ a partir da equação 3 do sistema:} \\
\\
&\frac{141}{22}x_3 = -\frac{373}{11} \\
&x_3 = -\frac{746}{141} \\
\\
x_2 \text{ a partir da equação 2 do sistema:} \\
\\
&-\frac{22}{5}x_2 = -\frac{267}{5} - \frac{17}{10}x_3 \\
&x_2 = -\frac{267}{5} - \left(\frac{17}{10} \cdot -\frac{746}{141}\right) = -\frac{31306}{705} \\
&x_2 = \frac{1423}{141} \\
\\
x_1 \text{ a partir da equação 1 do sistema:} \\
\\
&10x_1 = 27 - 2x_2 + x_3 \\
&x_1 = 27 - \left(2 \cdot \frac{1423}{141} \right) + \left(-\frac{746}{141}\right) = \frac{215}{141} \\
&x_1 = \frac{43}{282}
\end{align}
$$

Valores das raízes:

$$
 X = \{0.15248, 10.09219, -5.29078\}
$$

### Chamada ao algoritmo

```octave
inputMatrix = [
            10, 2, -1, 27;
            -3, -5, 2, -61.5;
            1, 1, 6, -21.5
            ];

[det, foundRoots, err] = gaussEliminationMethod(inputMatrix, false);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussElimination\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:'), disp(foundRoots);
```

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a triangularização da matriz. Essas informações podem ser visualizadas nas tabelas a seguir

| **Step** |                                    **Matrix**                                    | **Multiplier** |       **Operation**       |
| :------: | :------------------------------------------------------------------------------: | :------------: | :-----------------------: |
|    1     |               \[\[10 2 -1 27] \[0 -4.4 1.7 -53.4] \[1 1 6 -21.5]]                |   -0.300000    | L2 = L2 - (-0.3000 \* L1) |
|    2     |             \[\[10 2 -1 27] \[0 -4.4 1.7 -53.4] \[0 0.8 6.1 -24.2]]              |    0.100000    | L3 = L3 - (0.1000 \* L1)  |
|    3     | \[\[10 2 -1 27] \[0 -4.4 1.7 -53.4] \[0 0 6.409090909090908 -33.90909090909091]] |   -0.181818    | L3 = L3 - (-0.1818 \* L2) |

Da mesma forma, foram coletados os dados referentes aos passos de obtenção dos valores das raízes. Esses passos podem ser vistos na tabela abaixo:

| **Step** |                          **Roots**                           | **DivisorSum** | **Dividend** | **Added Root** |
| :------: | :----------------------------------------------------------: | :------------: | :----------: | :------------: |
|    1     |                  \[0 0 -5.290780141843972]                   |   -33.909091   |   6.409091   |   -5.290780    |
|    2     |          \[0 10.092198581560282 -5.290780141843972]          |   -44.405674   |  -4.400000   |   10.092199    |
|    3     | \[0.15248226950354632 10.092198581560282 -5.290780141843972] |    1.524823    |  10.000000   |    0.152482    |

**Valor do Determinante:** -282
**Raízes Encontradas:** 0.1525, 10.0922, -5.2908

## Questão 02

    Resolva o sistema por Eliminação de Gauss com Pivotamento parcial. Mostre todos os passos dos cálculos. Substitua os resultados nas equações originais para verificar as respostas. Implemente o algoritmo de Eliminação de Gauss ingênua. Aplique o script ao sistema.

$$
\begin{align}
2x_1 - 6x_2 − 1x_3 = -38 \\
−3x_1 − 1x_2 + 7x_3 = −34 \\
-8x_1 + 1x_2 - 2x_3 = −20 \\
\end{align}
$$

### Chamada ao algoritmo

```octave
inputMatrix = [
            2, -6, -1, -38;
            -3, -1, 7, -34;
            -8, 1, -2, -20
            ];

[det, foundRoots, err] = gaussEliminationMethod(inputMatrix, true);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussElimination\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:'), disp(foundRoots);
```

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a triangularização da matriz. Essas informações podem ser visualizadas nas tabelas a seguir

| **Step** |                                      **Matrix**                                      | **Multiplier** |       **Operation**       |
| :------: | :----------------------------------------------------------------------------------: | :------------: | :-----------------------: |
|    1     |               \[\[-8 1 -2 -20] \[0 -1.375 7.75 -26.5] \[2 -6 -1 -38]]                |    0.375000    | L2 = L2 - (0.3750 \* L1)  |
|    2     |             \[\[-8 1 -2 -20] \[0 -1.375 7.75 -26.5] \[0 -5.75 -1.5 -43]]             |   -0.250000    | L3 = L3 - (-0.2500 \* L1) |
|    3     | \[\[-8 1 -2 -20] \[0 -1.375 7.75 -26.5] \[0 0 -33.90909090909091 67.81818181818181]] |    4.181818    | L3 = L3 - (4.1818 \* L2)  |

Da mesma forma, foram coletados os dados referentes aos passos de obtenção dos valores das raízes. Esses passos podem ser vistos na tabela abaixo:

| **Step** | **Roots** | **DivisorSum** | **Dividend** | **Added Root** |
| :------: | :-------: | :------------: | :----------: | :------------: |
|    1     | \[0 0 -2] |   67.818182    |  -33.909091  |   -2.000000    |
|    2     | \[0 8 -2] |   -11.000000   |  -1.375000   |    8.000000    |
|    3     | \[4 8 -2] |   -32.000000   |  -8.000000   |    4.000000    |

**Valor do Determinante:** -373
**Raízes Encontradas:** 4, 8, -2

#### Verificando resultados

Agora, vamos substituir as raízes encontradas nas equações originais para verificar se elas são satisfeitas:

Para os valores

$$
\begin{align}
x_1 = 4 \\
x_2 = 8 \\
x_3 = -2
\end{align}
$$

1. Primeira equação:

$$
\begin{align}
2(4) - 6(8) - 1(-2) = 8 - 48 + 2 = -38 \\
\text{A primeira equação é satisfeita.}
\end{align}
$$

2. Segunda equação:

$$
\begin{align}
-3(4) - 1(8) + 7(-2) = -12 - 8 - 14 = -34 \\
\text{A segunda equação é satisfeita.}
\end{align}
$$

3. Terceira equação:

$$
\begin{align}
-8(4) + 1(8) - 2(-2) = -32 + 8 + 4 = -20 \\
\text{A terceira equação é satisfeita.}
\end{align}
$$

Portanto, as raízes encontradas satisfazem todas as equações originais do sistema. Isso confirma que as respostas estão corretas.
