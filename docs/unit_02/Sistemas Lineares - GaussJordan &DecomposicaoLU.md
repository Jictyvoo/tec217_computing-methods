> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01

    Resolva o sistema linear por Eliminação de Gauss-Jordan. Mostre todos os passos dos cálculos. Encontre a matriz inversa.

$$
\begin{align}
-0.04x_1 + 0.04x_2 − 0.12x_3 = 3 \\
−0.56x_1 − 1.56x_2 + 0.32x_3 = 1 \\
-0.24x_1 + 1.24x_2 + -0.28x_3 = 0 \\
\end{align}
$$

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% gaussJordanMethod.m
function [det, invMatrix, roots, err] = gaussJordanMethod(inputMatrix)
    err = ""; % Initialize err to an empty string

    % Checking if the given matrix is a square one
    [totalEquations, eqLength] = size(inputMatrix);
    if totalEquations != eqLength -1
        err = "the given matrix isn't a square one";
        det = NaN;
        roots = NaN;
        invMatrix = NaN;
        return;
    end

    % Generate the augmentedMatrix by insert the identity on it
    augmentedMatrix = [inputMatrix, eye(totalEquations)];

    det = 1
    for eqIndex = 1:totalEquations
        % Swap equation by the one in pivot element and use the value to normalize it next
        [augmentedMatrix, pivotElement] = pivot(augmentedMatrix, eqIndex, eqIndex);
        det *= pivotElement

        % Normalization
        augmentedMatrix(eqIndex, :) = augmentedMatrix(eqIndex, :) / augmentedMatrix(eqIndex, eqIndex);

        % Elimination
        for index = setdiff(1:totalEquations, eqIndex)
            augmentedMatrix(index, :) = augmentedMatrix(index, :) - augmentedMatrix(eqIndex, :) * augmentedMatrix(index, eqIndex);
        end
    end

    invMatrix = augmentedMatrix(:, totalEquations + 2:(totalEquations * 2) + 1);
    roots = augmentedMatrix(:, totalEquations + 1);
end

function [matrix, maxValue] = pivot(matrix, currentIndex, lookupColumn)
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

#### Chamada ao algoritmo

```octave
inputMatrix = [
            -0.04, 0.04, 0.12, 3;
            0.56, -1.56, 0.32, 1;
            -0.24, 1.24, -0.28, 0
            ];

[det, invMatrix, foundRoots, err] = gaussJordanMethod(inputMatrix);

% Report error if it has one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussJordan\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:\n'), disp(foundRoots);
fprintf('Inverse:\n'), disp(invMatrix);
```

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a triangularização da matriz.

Inicialmente foi realizada a inserção da matriz identidade na lista de operações, para dessa forma achar o valor da matriz inversa.

$$
\begin{bmatrix}
-0.04 & 0.04 & 0.12 \\
0.56 & -1.56 & 0.32 \\
-0.24 & 1.24 & -0.28
\end{bmatrix}

\begin{bmatrix}
1 & 0 & 0 \\
0 & 1 & 0 \\
0 & 0 & 1
\end{bmatrix}

\begin{bmatrix}
3 \\
1 \\
0
\end{bmatrix}
$$

Aqui estão as matrizes geradas passo a passo de acordo com a tabela:

Ao longo dos passos, também foi realizada a permutação entre linhas, para que os cálculos não tivessem erros de ponto flutuante.

#### Passo 1

**Multiplier:** 0.5600
**Operation:**

- L1 = L1 / 0.5600
- L1 <=> L2

$$
\begin{bmatrix}
1 & -2.7857 & 0.5714 \\
-0.04 & 0.04 & 0.12 \\
-0.24 & 1.24 & -0.28
\end{bmatrix}

\begin{bmatrix}
0 & 1.7857 & 0 \\
1 & 0 & 0 \\
0 & 0 & 1
\end{bmatrix}

\begin{bmatrix}
1.7857 \\
3 \\
0
\end{bmatrix}
$$

#### Passo 2

**Multiplier:** -0.0400
**Operation:**

- L2 = L2 - (-0.0400 \* L1)

$$
\begin{bmatrix}
1 & -2.7857 & 0.5714 \\
0 & -0.0714 & 0.1429 \\
-0.24 & 1.24 & -0.28
\end{bmatrix}

\begin{bmatrix}
0 & 1.7857 & 0 \\
1 & 0.0714 & 0 \\
0 & 0 & 1
\end{bmatrix}

\begin{bmatrix}
1.7857 \\
3.0714 \\
0
\end{bmatrix}
$$

#### Passo 3

**Multiplier:** -0.2400
**Operation:**

- L3 = L3 - (-0.2400 \* L1)

$$
\begin{bmatrix}
1 & -2.7857 & 0.5714 \\
0 & -0.0714 & 0.1429 \\
0 & 0.5714 & -0.1429
\end{bmatrix}

\begin{bmatrix}
0 & 1.7857 & 0 \\
1 & 0.0714 & 0 \\
0 & 0.4286 & 1
\end{bmatrix}

\begin{bmatrix}
1.7857 \\
3.0714 \\
0.4286
\end{bmatrix}
$$

#### Passo 4

**Multiplier:** 0.5714
**Operation:**

- L2 <=> L3
- L2 = L2 / 0.5714

$$
\begin{bmatrix}
1 & -2.7857 & 0.5714 \\
0 & 1 & -0.2500 \\
0 & -0.0714 & 0.1429
\end{bmatrix}

\begin{bmatrix}
0 & 1.7857 & 0 \\
0 & 0.7500 & 1.7500 \\
1 & 0.0714 & 0
\end{bmatrix}

\begin{bmatrix}
1.7857 \\
0.7500 \\
3.0714
\end{bmatrix}
$$

#### Passo 5

**Multiplier:** -2.7857
**Operation:**

- L1 = L1 - (-2.7857 \* L2)

$$
\begin{bmatrix}
1 & 0 & -0.1250 \\
0 & 1 & -0.2500 \\
0 & -0.0714 & 0.1429
\end{bmatrix}

\begin{bmatrix}
0 & 3.8750 & 4.8750 \\
0 & 0.7500 & 1.7500 \\
1 & 0.0714 & 0
\end{bmatrix}

\begin{bmatrix}
3.8750 \\
0.7500 \\
3.0714
\end{bmatrix}
$$

#### Passo 6

**Multiplier:** -0.0714
**Operation:**

- L3 = L3 - (-0.0714 \* L2)

$$
\begin{bmatrix}
1 & 0 & -0.1250 \\
0 & 1 & -0.2500 \\
0 & 0 & 0.1250
\end{bmatrix}

\begin{bmatrix}
0 & 3.8750 & 4.8750 \\
0 & 0.7500 & 1.7500 \\
1 & 0.1250 & 0.1250
\end{bmatrix}

\begin{bmatrix}
3.8750 \\
0.7500 \\
3.1250
\end{bmatrix}
$$

#### Passo 7

**Multiplier:** 0.1250
**Operation:**

- L3 = L3 / 0.1250

$$
\begin{bmatrix}
1 & 0 & -0.1250 \\
0 & 1 & -0.2500 \\
0 & 0 & 1
\end{bmatrix}

\begin{bmatrix}
0 & 3.8750 & 4.8750 \\
0 & 0.7500 & 1.7500 \\
8.0000 & 1 & 0.9999
\end{bmatrix}

\begin{bmatrix}
3.8750 \\
0.7500 \\
25.0000
\end{bmatrix}
$$

#### Passo 8

**Multiplier:** -0.1250
**Operation:**

- L1 = L1 - (-0.1250 \* L3)

$$
\begin{bmatrix}
1 & 0 & 0 \\
0 & 1 & -0.2500 \\
0 & 0 & 1
\end{bmatrix}

\begin{bmatrix}
1.0000 & 3.9999 & 4.9999 \\
0.0000 & 0.7499 & 1.7499 \\
8.0000 & 1.0000 & 0.9999
\end{bmatrix}

\begin{bmatrix}
7.0000 \\
0.7499 \\
25.0000
\end{bmatrix}
$$

#### Passo 9

**Multiplier:** -0.2500
**Operation:**

- L2 = L2 - (-0.2500 \* L3)

$$
\begin{bmatrix}
1 & 0 & 0 \\
0 & 1 & 0 \\
0 & 0 & 1
\end{bmatrix}

\begin{bmatrix}
1.0000 & 4.0000 & 5.0000 \\
2.0000 & 1.0000 & 1.9999 \\
8.0000 & 1.0000 & 0.9999
\end{bmatrix}

\begin{bmatrix}
7.0000 \\
7.0000 \\
25.0000
\end{bmatrix}
$$

#### Determinante e Raízes

Como transformamos a matriz em uma matriz identidade, os valores das raízes são obtidos diretamente.

**Valor do Determinante:** 0.04
**Raízes Encontradas:** 7, 7, 25

##### Matriz Inversa

| 1   | 4   | 5   |
| --- | --- | --- |
| 2   | 1   | 2   |
| 8   | 1   | 1   |

---

## Questão 02

    Resolva o sistema linear por Decomposição LU. Mostre todos os passos dos cálculos (As matrizes L e U, vetor intermediário y, e vetor solução x). Aplicar as matrizes obtidas na decomposição LU para o cálculo de x usando b2 como vetor independente.

$$
\begin{align}
2x_1 + 1x_2 - 3x_3 = 2 \\
−1x_1 + 3x_2 + 2x_3 = 0 \\
3x_1 + 1x_2 - 3x_3 = 1 \\

\\
b_2 =
\begin{bmatrix}
3 \\
1 \\
2
\end{bmatrix}

\end{align}
$$

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% luDecompositionMethod.m
function [LMatrix, UMatrix, det, roots, err] = luDecompositionMethod(inputMatrix)
    err = ""; % Initialize err to an empty string

    % Checking if the given matrix is a square one
    [matrixSize, eqLength] = size(inputMatrix);
    if matrixSize != eqLength -1
        err = "the given matrix isn't a square one";
        det = NaN;
        roots = NaN;
        LMatrix = NaN;
        UMatrix = NaN;
        return;
    end

    det = 1;
    % Initialize the L and U matrices
    LMatrix = eye(matrixSize);

    % Split the input matrix to obtain uMatrix and results separatedly
    UMatrix = inputMatrix(:, 1:matrixSize);
    results = inputMatrix(:, matrixSize + 1);

    for eqIndex = 1:matrixSize
        pivot = UMatrix(eqIndex, eqIndex);

        for rowIndex = eqIndex + 1:matrixSize
            factor = UMatrix(rowIndex, eqIndex) / pivot;
            UMatrix(rowIndex, :) -= factor * UMatrix(eqIndex, :);
            LMatrix(rowIndex, eqIndex) = factor; % store the multiplier in L matrix
        end

        det *= UMatrix(eqIndex, eqIndex);
    end

    % Calculate roots by forward and backward substitution
    y = forwardSubstitution(LMatrix, matrixSize, results);
    roots = backwardSubstitution(UMatrix, matrixSize, y);

end

% Solve Inferior triangle
function [x] = forwardSubstitution(L, matrixSize, b)
    x = zeros(matrixSize, 1);
    x(1) = b(1) / L(1, 1);

    for index = 2:matrixSize
        x(index) = (b(index) - L(index, 1:index - 1) * x(1:index - 1)) / L(index, index);
    end
end

% Solve superior triangle
function [x] = backwardSubstitution(U, matrixSize, b)
    x = zeros(matrixSize, 1);
    x(matrixSize) = b(matrixSize) / U(matrixSize, matrixSize);

    for index = matrixSize - 1:-1:1
        x(index) = (b(index) - U(index, index + 1:matrixSize) * x(index + 1:matrixSize)) / U(index, index);
    end
end
```

#### Chamada ao algoritmo

```octave
inputMatrix = [
            2, 1, -1, 3;
            -1, 3, 2, 1;
            3, 1, -3, 2;
            ];

[LMatrix, UMatrix, det, foundRoots, err] = luDecompositionMethod(inputMatrix);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: LU Decomposition\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:\n'), disp(foundRoots);

% Report L and U Matrices
disp("LMatrix: "); disp(LMatrix);
disp("UMatrix: "); disp(UMatrix);
```

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

| **Step** | **Matrix**                                            | **Multiplier** |     **Operation** |
| :------: | :---------------------------------------------------- | :------------: | ----------------: |
|    1     | \[\[1,0,0],\[-0.5,1,0],\[0,0,1]]                      |   -0.500000    | L\[1,2] = -0.5000 |
|    2     | \[\[1,0,0],\[-0.5,1,0],\[1.5,0,1]]                    |    1.500000    |  L\[1,3] = 1.5000 |
|    3     | \[\[1,0,0],\[-0.5,1,0],\[1.5,-0.14285714285714285,1]] |   -0.142857    | L\[2,3] = -0.1429 |

#### Matriz U

Para o calculo da matriz U, assume-se como valor inicial da mesma o mesmo da matriz de entrada.

$$
U=
\begin{bmatrix}
2 & 1 & -3 \\
-1 & 3 & 2 \\
3 & 1 & -3
\end{bmatrix}
$$

| **Step** | **Matrix**                                           | **Multiplier** |             **Operation** |
| :------: | ---------------------------------------------------- | :------------: | ------------------------: |
|    1     | \[\[2,1,-1],\[0,3.5,1.5],\[3,1,-3]]                  |   -0.500000    | L1 = L1 - (-0.5000 \* L2) |
|    2     | \[\[2,1,-1],\[0,3.5,1.5],\[0,-0.5,-1.5]]             |    1.500000    |  L1 = L1 - (1.5000 \* L3) |
|    3     | \[\[2,1,-1],\[0,3.5,1.5],\[0,0,-1.2857142857142858]] |   -0.142857    | L2 = L2 - (-0.1429 \* L3) |

#### Cálculo das Raízes

De modo a realizar a obtenção das raízes, deve-se inicialmente calcular o vetor y, através da resolução do triângulo inferior da matriz L. Para realização desse cálculo, utiliza-se o valor indicado:

$$
b_2 =
\begin{bmatrix}
3 \\
1 \\
2
\end{bmatrix}
$$

| **Step** |            **Y**            | **DividendSum** | **Divisor** | **Added Root** |
| :------: | :-------------------------: | :-------------: | :---------: | :------------: |
|    1     |          \[3,0,0]           |    3.000000     |  1.000000   |    0.000000    |
|    2     |         \[3,2.5,0]          |    2.500000     |  1.000000   |    2.500000    |
|    3     | \[3,2.5,-2.142857142857143] |    -2.142857    |  1.000000   |    3.000000    |

E posteriormente deve-se utilizar o vetor resultante y para calcular o vetor de raízes x, através da resolução do triângulo superior da matriz U.

| **Step** |                 **Roots**                 | **DividendSum** | **Divisor** | **Added Root** |
| :------: | :---------------------------------------: | :-------------: | :---------: | :------------: |
|    1     |         \[0,0,1.6666666666666665]         |    -2.142857    |  -1.285714  |    1.666667    |
|    2     |         \[0,0,1.6666666666666665]         |    0.000000     |  3.500000   |    0.000000    |
|    3     | \[2.333333333333333,0,1.6666666666666665] |    4.666667     |  2.000000   |    2.333333    |

**Valor do Determinante:** -9
**Raízes Encontradas:** 2.3333, 0, 1.6667
