> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Para as implementações dos métodos solicitados logo abaixo, deve-se obrigatoriamente realizar ao menos duas iterações manuais.

## Achando valores para C e d

### Inicialização

Primeiramente faz-se necessária a inicialização com os valores de `C` e `d` como os seguintes:

$$
\begin{align}

d=
\begin{bmatrix}
0 & 0 & 0
\end{bmatrix}
\\
\\
C=
\begin{bmatrix}
0 & 0 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
\end{align}
$$

### Achando C

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -\frac{A[1,2]}{A[1,1]} & -\frac{A[1,3]}{A[1,1]} \\
-\frac{A[2,1]}{A[2,2]} & 0 & -\frac{A[2,3]}{A[2,2]} \\
-\frac{A[3,1]}{A[3,3]} & -\frac{A[3,2]}{A[3,3]} & 0
\end{bmatrix}
$$

- É realizada a iteração linha a linha coluna por coluna, caso a linha e coluna tenham o memso índice, o valor permanece como zero.
- Calculamos os valores de C\[linha, coluna] da matriz C usando a fórmula:

$$C[linha, coluna] = -\frac{A[linha, coluna]}{A[linha, linha]}$$

Por exemplo, calculamos $C[1,2]$ como $-\frac{A[1,2]}{A[1,1]}$, assumindo que o valor de $A[1,2]=54$ e $A[1,1]=2$, teremos como resultado: $-\frac{54}{2} = 27$.

Assim sendo, esses passos se repetem até termos os valores de C e d completos.

### Achando d

Assumindo que o conjunto de resultados é representado como o vetor $b$, o vetor $d$ pode ser representado da seguinte forma:

$$
d=
\begin{bmatrix}
\frac{b[1]}{A[1,1]} & \frac{b[2]}{A[2,2]} & \frac{b[3]}{A[3,3]} \\
\end{bmatrix}
$$

## Questão 01 - Método de Jacobi

    Resolva o sistema usando o método de Jacobi. Com erro relativo menor que a tolerância e = 0,05. Usar x0 = (0,0,0). Incluir critérios de convergência: diagonal dominante.

$$
\begin{align}
\epsilon = 0.05 \\
x_0 = [0,0,0] \\
\\
5x_1 + 1x_2 + 1x_3 = 5 \\
3x_1 + 4x_2 + 1x_3 = 6 \\
3x_1 + 3x_2 + 6x_3 = 0 \\
\end{align}
$$

### Resolução através do método

#### Convergência: Diagonal Dominante

Realiza-se a iteração das linhas da matriz A e, para cada linha, é calculada a soma dos valores absolutos dos elementos fora da diagonal principal.

$$
\begin{align}
\text{Para a linha L1:} \\
&sum^{L1} = abs(1) + abs(1) \\
\\
\\
\text{Para a linha L2:} \\
&sum^{L2} = abs(3) + abs(1) \\
\\
\\
\text{Para a linha L3:} \\
&sum^{L3} = abs(3) + abs(3)
\end{align}
$$

Após realizar o cálculo da soma ignorando o elemento da diagonal principal, é necessário realizar a comparação entre esse valor da diagonal principal e o resultado da soma.

$$
\begin{align}
\text{L1: } 5 >= 2\text{ (Verdadeiro)} \\
\text{L2: } 4 >= 4\text{ (Verdadeiro)} \\
\text{L3: } 6 >= 6\text{ (Verdadeiro)}
\end{align}
$$

Logo, é possível assegurar que a matriz é diagonal dominante.

#### Iterações manuais

##### Passo 1

$$
x^{(1)} = (C \cdot x^{(0)}) + d
$$

$$
x^{(1)} = \left(\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & -0.25 \\
-0.5 & -0.5 & 0
\end{bmatrix} \cdot \begin{bmatrix}
0 \\
0 \\
0
\end{bmatrix}\right) + \begin{bmatrix}
1 \\
1.5 \\
0
\end{bmatrix}
$$

Portanto, $x^{(1)} = \begin{bmatrix} 1 \\ 1.5 \\ 0 \end{bmatrix}$.

##### Passo 2

$$
x^{(2)} = (C \cdot x^{(1)}) + d
$$

$$
x^{(2)} = \left(\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & -0.25 \\
-0.5 & -0.5 & 0
\end{bmatrix} \cdot \begin{bmatrix}
1 \\
1.5 \\
0
\end{bmatrix}\right) + \begin{bmatrix}
1 \\
1.5 \\
0
\end{bmatrix}
$$

Portanto, $x^{(2)} = \begin{bmatrix} 0.7 \\ 0.75 \\ -1.25 \end{bmatrix}$.

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% jacobiMethod.m
function [foundRoots, currentIteration, err] = jacobiMethod(inputMatrix, initialGuess, maxIterations, epsilon)
    err = ""; % Initialize err to an empty string
    currentIteration = 0;

    if isDiagonallyDominant(inputMatrix) == false
        foundRoots = NaN
        err = "the given matrix is not diagonally dominant";
        return
    endif

    matrixSize = size(inputMatrix, 1);

    workingMatrix = inputMatrix(:, 1:matrixSize);
    results = inputMatrix(:, matrixSize + 1);

    C = zeros(matrixSize, matrixSize);
    d = zeros(matrixSize, 1);

    for eqIndex = 1:matrixSize
        for index = 1:matrixSize
            setValue = 0;
            if eqIndex == index
                divisor = results(eqIndex);
                dividend = workingMatrix(eqIndex, eqIndex);
                d(eqIndex) = divisor / dividend;
            else
                setValue = -workingMatrix(eqIndex, index) / workingMatrix(eqIndex, eqIndex);
            endif
            C(eqIndex, index) = setValue;
        endfor
    endfor

    foundRoots = zeros(matrixSize, 1);
    foundRoots(1) = initialGuess;
    while currentIteration < maxIterations
        currentIteration += 1;
        [foundRoots, maxError, success] = calculateRoot(C, d, matrixSize, epsilon, foundRoots);
        if success
            return;
        endif
    endwhile
    err = printf("failed to converge in %d iterations", currentIteration);
endfunction

function isDiagonalDominant = isDiagonallyDominant(A)
    % Get the dimensions of the matrix
    matrixSize = size(A, 1);

    for index = 1:matrixSize
        sum = 0;
        for subIndex = 1:matrixSize
            if index != subIndex
                % Add the absolute value to the accumulator
                absoluteAij = abs(A(index, subIndex));
                sum += absoluteAij;
            endif
        endfor
        % Check if the value on the main diagonal is less than sum
        if abs(A(index, index)) < sum
            isDiagonalDominant = false;
            return;
        endif
    endfor

    % The matrix is diagonally dominant
    isDiagonalDominant = true;
endfunction

function [foundRoots, maximumXError, success] = calculateRoot(C, d, matrixSize, epsilon, xPrev)
    foundRoots = C * xPrev + d
    maximumXError = max(abs((foundRoots - xPrev) ./ foundRoots)) * 100;
    success = maximumXError < epsilon;
endfunction
```

#### Chamada ao algoritmo

```octave
inputMatrix = [
            5, 1, 1, 5;
            3, 4, 1, 6;
            3, 3, 6, 0
            ];
initialGuess = 0;
maxIterations = 1000;
epsilon = 0.05;

% Run the Jacobi method.
[foundRoots, totalIterations, err] = jacobiMethod(inputMatrix, initialGuess, maxIterations, epsilon);

% Report error if there is one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: Jacobi\n');
fprintf('Total Iterations: %d\n', totalIterations);
fprintf('Roots:\n'), disp(foundRoots);
```

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
5 & 1 & 1 \\
3 & 4 & 1 \\
3 & 3 & 6
\end{bmatrix}

\begin{bmatrix}
5 \\
6 \\
0
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- L\[1,1] = 0.0000

#### Passo 2

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- L\[1,2] = -0.2000

#### Passo 3

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- L\[1,3] = -0.2000

#### Passo 4

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- L\[2,1] = -0.7500

#### Passo 5

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- L\[2,2] = 0.0000

#### Passo 6

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & -0.25 \\
0 & 0 & 0
\end{bmatrix}
$$

**Multiplier:** -0.2500
**Operation:**

- L\[2,3] = -0.2500

#### Passo 7

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & -0.25 \\
-0.5 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- L\[3,1] = -0.5000

#### Passo 8

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & -0.25 \\
-0.5 & -0.5 & 0
\end{bmatrix}
$$

**Operation:**

- L\[3,2] = -0.5000

#### Passo 9

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.2 & -0.2 \\
-0.75 & 0 & -0.25 \\
-0.5 & -0.5 & 0
\end{bmatrix}
$$

**Operation:**

- L\[3,3] = 0.0000

#### Valor do vetor `d`

Enquanto eram realizadas operações para montar a matriz C, também foi realizada a montagem do vetor d.

| **Step** | **d**      | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ---------- | :-------------: | :---------: | :------------: |
|    1     | \[1,0,0]   |       5.0       |     5.0     |      1.0       |
|    2     | \[1,1.5,0] |       6.0       |     4.0     |      1.5       |
|    3     | \[1,1.5,0] |       0.0       |     6.0     |      0.0       |

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:
**Raízes Encontradas:** 1.0001,1.0002,-0.9998

De modo a facilitar a visualização, os valores de `a, b, c` correspondem ao valor anterior do X.

| **Interaction** |  **a**   |  **b**   |   **c**   | **Found-X**                                                  | **Relative-Error** |
| :-------------: | :------: | :------: | :-------: | ------------------------------------------------------------ | :----------------: |
|        1        | 0.000000 | 0.000000 | 0.000000  | \[1,1.5,0]                                                   |     100.000000     |
|        2        | 1.000000 | 1.500000 | 0.000000  | \[0.7,0.75,-1.25]                                            |     100.000000     |
|        3        | 0.700000 | 0.750000 | -1.250000 | \[1.1,1.2875,-0.725]                                         |     72.413793      |
|        4        | 1.100000 | 1.287500 | -0.725000 | \[0.8875,0.85625,-1.19375]                                   |     50.364964      |
|        5        | 0.887500 | 0.856250 | -1.193750 | \[1.0675,1.1328125,-0.871875]                                |     36.917563      |
|        6        | 1.067500 | 1.132812 | -0.871875 | \[0.9478125,0.9173437500000001,-1.10015625]                  |     23.488332      |
|        7        | 0.947812 | 0.917344 | -1.100156 | \[1.0365625,1.0641796875,-0.932578125]                       |     17.969339      |
|        8        | 1.036563 | 1.064180 | -0.932578 | \[0.9736796875,0.95572265625,-1.05037109375]                 |     11.348170      |
|        9        | 0.973680 | 0.955723 | -1.050371 | \[1.0189296875,1.0323330078125,-0.964701171875]              |      8.880462      |
|       10        | 1.018930 | 1.032333 | -0.964701 | \[0.9864736328125,0.97697802734375,-1.02563134765625]        |      5.940748      |
|       11        | 0.986474 | 0.976978 | -1.025631 | \[1.0097306640625,1.0165526123046875,-0.9817258300781251]    |      4.472279      |
|       12        | 1.009731 | 1.016553 | -0.981726 | \[0.9930346435546875,0.9881334594726563,-1.0131416381835936] |      3.100831      |
|       13        | 0.993035 | 0.988133 | -1.013142 | \[1.0050016357421874,1.0085094268798827,-0.9905840515136719] |      2.277201      |
|       14        | 1.005002 | 1.008509 | -0.990584 | \[0.9964149249267579,0.9938947860717775,-1.006755531311035]  |      1.606297      |
|       15        | 0.996415 | 0.993895 | -1.006756 | \[1.0025721490478514,1.0043776891326903,-0.9951548554992677] |      1.165716      |
|       16        | 1.002572 | 1.004378 | -0.995155 | \[0.9981554332733155,0.9968596020889283,-1.0034749190902708] |      0.829125      |
|       17        | 0.998155 | 0.996860 | -1.003475 | \[1.0013230634002686,1.0022521548175811,-0.9975075176811219] |      0.598231      |
|       18        | 1.001323 | 1.002252 | -0.997508 | \[0.9990510725727082,0.998384581870079,-1.0017876091089248]  |      0.427245      |
|       19        | 0.999051 | 0.998385 | -1.001788 | \[1.0006806054477693,1.0011585978477,-0.9987178272213936]    |      0.307372      |
|       20        | 1.000681 | 1.001159 | -0.998718 | \[0.9995118458747387,0.9991690027195215,-1.0009196016477346] |      0.219975      |
|       21        | 0.999512 | 0.999169 | -1.000920 | \[1.0003501197856426,1.0005960160058796,-0.9993404242971301] |      0.158022      |
|       22        | 1.000350 | 1.000596 | -0.999340 | \[0.99974888165825,0.9995725162350506,-1.000473067895761]    |      0.113211      |
|       23        | 0.999749 | 0.999573 | -1.000473 | \[1.0001801103321422,1.0003066057302528,-0.9996606989466503] |      0.081264      |
|       24        | 1.000180 | 1.000307 | -0.999661 | \[0.9998708186432795,0.9997800919875559,-1.0002433580311973] |      0.058252      |
|       25        | 0.999871 | 0.999780 | -1.000243 | \[1.0000926532087282,1.0001577255253395,-0.9998254553154178] |      0.041798      |

---

## Questão 02 - Método de Gauss-Seidel

    Resolva o sistema pelo método de Gauss-Seidel. Mostre se o critério de Sassenfeld é satisfeito como avaliação de convergência.

$$
\begin{align}
\epsilon = 1 \cdot 10^{-5} \\
x_0 = [0,0,0] \\
\\
1x_1 + 0.5x_2 - 0.1x_3 = 0.2 \\
0.2x_1 + 1x_2 - 0.2x_3 = -2 \\
-0.1x_1 - 0.2x_2 + 1x_3 = 1 \\
\end{align}
$$

### Resolução através do método

#### Convergência: Sassenfeld

Primeiro é realizado o calculo dos valores dos coeficientes de Sassenfeld, para posteriormente verificar se eles são menores que 1 para garantir a convergência do método iterativo.

Vamos aplicar o critério de convergência de Sassenfeld às equações fornecidas:

$$
\begin{align}
1x_1 + 0.5x_2 - 0.1x_3 = 0.2 \quad &\Rightarrow \quad S_1 = |0.5| + |-0.1| = 0.6 \\
0.2x_1 + 1x_2 - 0.2x_3 = -2 \quad &\Rightarrow \quad S_2 = |0.2 \cdot 0.6| + |-0.2| = 0.32 \\
-0.1x_1 - 0.2x_2 + 1x_3 = 1 \quad &\Rightarrow \quad S_3 = |-0.1 \cdot 0.6| + |-0.2 \cdot 0.32| = 0.124 \\
\end{align}
$$

Agora, é realizada a verificação se a divisão dos valores de $\frac{S_i}{A[i,i]}$ é inferior a 1, dessa forma satisfazendo o critério.

- Para a primeira equação $\frac{0.6}{1} = 0.6$
- Para a segunda equação $\frac{0.32}{1} = 0.32$
- Para a terceira equação $\frac{0.124}{1} = 0.124$

Logo, é possível assegurar que a matriz satisfaz o critério.

#### Iterações manuais

Para iniciar o processo, assume-se o valor inicial de $a,b,c$ como 0, e posteriormente os valores de X encontrados serão usados para atualizar $a,b,c$.

##### Passo 1

1. Para $x_1$ :

$$
\begin{align}
a = (-1 \cdot a -0.5 \cdot b + 0.1 \cdot c) + (-1 \cdot x_1 -0.5 \cdot x_2 + 0.1 \cdot x_3) \\
a = 0 \\
\\
a = \frac{a}{1} + \frac{0.2}{1} \\
a = 0.2
\end{align}
$$

2. Para $x_2$ :

$$
\begin{align}
b = (-0.2 \cdot a -1 \cdot b + 0.2 \cdot c) + (-0.2 \cdot x_1 -1 \cdot x_2 + 0.2 \cdot x_3) \\
b = -0.04 \\
\\
b = \frac{b}{1} + \frac{-2}{1} \\
b = -2.04
\end{align}
$$

3. Para $x_3$ :

$$
\begin{align}
c = (0.1 \cdot a +0.2 \cdot b -1 \cdot c) + (0.1 \cdot x_1 +0.2 \cdot x_2 - 1 \cdot x_3) \\
c = -0.388 \\
\\
c = \frac{b}{1} + \frac{1}{1} \\
c = 0.612
\end{align}
$$

Dessa forma, ao fim do passo 1, temos $x_1=0.2$, $x_2 = -2.04$, $x_3 = 0.612$

##### Passo 2

1. Para $x_1$ :

$$
\begin{align}
a = (-1 \cdot a -0.5 \cdot b + 0.1 \cdot c) + (-1 \cdot x_1 -0.5 \cdot x_2 + 0.1 \cdot x_3) \\
a = 1.0812 \\
\\
a = \frac{a}{1} + \frac{0.2}{1} \\
a = 1.2812
\end{align}
$$

2. Para $x_2$ :

$$
\begin{align}
b = (-0.2 \cdot a -1 \cdot b + 0.2 \cdot c) + (-0.2 \cdot x_1 -1 \cdot x_2 + 0.2 \cdot x_3) \\
b = −0.13384 \\
\\
b = \frac{b}{1} + \frac{-2}{1} \\
b = -2.13384
\end{align}
$$

3. Para $x_3$ :

$$
\begin{align}
c = (0.1 \cdot a +0.2 \cdot b -1 \cdot c) + (0.1 \cdot x_1 +0.2 \cdot x_2 - 1 \cdot x_3) \\
c = -0.298648 \\
\\
c = \frac{b}{1} + \frac{1}{1} \\
c = 0.701352
\end{align}
$$

Dessa forma, ao fim do passo 1, temos $x_1 = 1.2812$, $x_2 = -2.13384$, $x_3 = 0.701352$

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% gaussSeidelMethod.m
function [foundRoots, currentIteration, err] = gaussSeidelMethod(inputMatrix, X0, maxIterations, epsilon)
    err = ""; % Initialize err to an empty string
    currentIteration = 0;

    % Checking if the given matrix is a square one
    [matrixSize, eqLength] = size(inputMatrix);
    if matrixSize != eqLength -1
        err = "the given matrix isn't a square one";
        foundRoots = NaN;
        return;
    endif

    % Checking Sassenfeld criteria
    if sassenfeld(inputMatrix) == false
        foundRoots = NaN
        err = "the given matrix does not respect the Sassenfeld criteria";
        return;
    endif

    workingMatrix = inputMatrix(:, 1:matrixSize);
    results = inputMatrix(:, matrixSize + 1);

    [C, d] = mountCAndD(workingMatrix, results);

    foundRoots = X0;
    while currentIteration < maxIterations
        currentIteration += 1;
        oldSolution = foundRoots;
        foundRoots = calculateRoots(matrixSize, workingMatrix, foundRoots, oldSolution, results);

        maximumXError = max(abs((foundRoots - oldSolution) ./ foundRoots) * 100);
        if maximumXError < epsilon
            err = "";
            return;
        endif
    endwhile

    err = printf("failed to converge in %d iterations", currentIteration);
endfunction

function ssnfd = sassenfeld(A)
    matrixSize = size(A, 1);
    betas = zeros(matrixSize, 1);

    for index = 1:matrixSize
        tempValue = 0;
        for loopCount = 1:index - 1
            tempValue += abs(A(index, loopCount)) * betas(j);
        endfor
        for subIndex = index + 1:matrixSize
            tempValue += abs(A(index, subIndex));
        endfor

        betas(index) = tempValue / abs(A(index, index));

        if betas(index) >= 1
            ssnfd = false;
            return;
        endif
    endfor

    ssnfd = true;
endfunction

function [C, d] = mountCAndD(A, results)
    matrixSize = size(A, 1);
    C = zeros(matrixSize, matrixSize);
    d = zeros(matrixSize, 1);

    for eqIndex = 1:matrixSize
        for index = 1:matrixSize
            setValue = 0;
            if eqIndex != index
                setValue = -A(eqIndex, index) / A(eqIndex, eqIndex);
            endif
        endfor
        divisor = results(eqIndex);
        dividend = A(eqIndex, eqIndex);
        d(eqIndex) = divisor / dividend;
    endfor
endfunction

function foundRoots = calculateRoots(matrixSize, A, foundRoots, oldRoots, results)
    for index = 1:matrixSize
        firstSum = sum(A(index, 1:index - 1) .* foundRoots(1:index - 1));
        secondSum = sum(A(index, index + 1:matrixSize) .* oldRoots(index + 1:matrixSize));
        foundRoots(index) = (results(index) - firstSum - secondSum) / A(index, index);
    endfor
endfunction
```

#### Chamada ao algoritmo

```octave
inputMatrix = [
            1, 0.5, -0.1, 0.2;
            0.2, 1, -0.2, -2;
            -0.1, -0.2, 1, 1
            ];
X0 = [0, 0, 0];
maxIterations = 1000;
epsilon = 1e-5;

% Run the Gauss-Seidel method.
[foundRoots, totalIterations, err] = gaussSeidelMethod(inputMatrix, X0, maxIterations, epsilon);

% Report error if there is one.
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: Gauss-Seidel\n');
fprintf('Total Iterations: %d\n', totalIterations);
fprintf('Roots:\n'), disp(foundRoots);
```

### Resultados

Durante a execução do algoritmo, foram coletadas informações referentes aos valores utilizados e a forma como o mesmo realizou a manipulação da matriz.

$$
\begin{bmatrix}
1 & 0.5 & -0.1 \\
0.2 & 1 & -0.2 \\
-0.1 & -0.2 & 1
\end{bmatrix}

\begin{bmatrix}
0.2 \\
-2 \\
1
\end{bmatrix}
$$

#### Passo 1

**Matrix:**

$$
C=
\begin{bmatrix}
0 & 0 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[1,1] = 0.0000$

#### Passo 2

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[1,2] = -0.5000$

#### Passo 3

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
0 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[1,3] = 0.1000$

#### Passo 4

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
-0.2000 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[2,1] = -0.2000$

#### Passo 5

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
-0.2000 & 0 & 0 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[2,2] = 0.0000$

#### Passo 6

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
-0.2000 & 0 & 0.2000 \\
0 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[2,3] = 0.2000$

#### Passo 7

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
-0.2000 & 0 & 0.2000 \\
0.1000 & 0 & 0
\end{bmatrix}
$$

**Operation:**

- $L[3,1] = 0.1000$

#### Passo 8

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
-0.2000 & 0 & 0.2000 \\
0.1000 & 0.2000 & 0
\end{bmatrix}
$$

**Operation:**

- $L[3,2] = 0.2000$

#### Passo 9

**Matrix:**

$$
C=
\begin{bmatrix}
0 & -0.5 & 0.1000 \\
-0.2000 & 0 & 0.2000 \\
0.1000 & 0.2000 & 0
\end{bmatrix}
$$

**Operation:**

- $L[3,3] = 0.0000$

#### Valor do vetor `d`

Enquanto eram realizadas operações para montar a matriz C, também foi realizada a montagem do vetor d.

| **Step** | **d**       | **DividendSum** | **Divisor** | **Added Root** |
| :------: | ----------- | :-------------: | :---------: | :------------: |
|    1     | \[0.2,0,0]  |       1.0       |     0.2     |      0.2       |
|    2     | \[0.2,-2,0] |       1.0       |    -2.0     |      -2.0      |
|    3     | \[0.2,-2,1] |       1.0       |     1.0     |      1.0       |

#### Raízes

Posteriormente foi realizada a execução do resto do método, de modo a encontrar os valores das raízes. As iterações podem ser visualizadas nas tabelas abaixo:
**Raízes Encontradas:** 1.3333,-2.1249,0.7083

De modo a facilitar a visualização, os valores de `a, b, c` correspondem ao valor anterior do X.

| **Interaction** |  **a**   |   **b**   |  **c**   | **Found-X**                                                  | **Relative-Error** |
| :-------------: | :------: | :-------: | :------: | ------------------------------------------------------------ | :----------------: |
|        1        | 0.000000 | 0.000000  | 0.000000 | \[0.2,-2.04,0.612]                                           |     100.000000     |
|        2        | 0.200000 | -2.040000 | 0.612000 | \[1.2812,-2.13384,0.701352]                                  |     84.389635      |
|        3        | 1.281200 | -2.133840 | 0.701352 | \[1.3370552,-2.12714064,0.708277392]                         |      4.177479      |
|        4        | 1.337055 | -2.127141 | 0.708277 | \[1.3343980592,-2.12522413344,0.708394979232]                |      0.199127      |
|        5        | 1.334398 | -2.125224 | 0.708395 | \[1.3334515646432001,-2.1250113170822402,0.7083428930478719] |      0.070981      |
|        6        | 1.333452 | -2.125011 | 0.708343 | \[1.3333399478459074,-2.124999410959607,0.7083341125926693]  |      0.008371      |
|        7        | 1.333340 | -2.124999 | 0.708334 | \[1.3333331167390705,-2.12499980082928,0.708333351508051]    |      0.000512      |
|        8        | 1.333333 | -2.125000 | 0.708333 | \[1.333333235565445,-2.1249999768114787,0.7083333281942488]  |      0.000009      |

---

## Questão 03 - Gauss-Seidel SOR
	3.Elabore mudanças no algoritmo para implementar o método de SOR. Para o sistema a seguir encontre o melhor relaxamento utilizando um critério de parada de es = 10% .


$$
\begin{align}
\epsilon = 1 \cdot 10^{-1} \\
x_0 = [0,0] \\
\omega = 1 \\
\\
-3x_1 + 12x_2 = 9 \\
10x_1 - 2x_2 = 8 \\
\end{align}
$$

#### Iterações manuais

Para iniciar o processo, assume-se o valor inicial de $a,b$ como 0, e posteriormente os valores de X encontrados serão usados para atualizar $a,b$.

##### Passo 1

1. Para $x_1$ :

$$
\begin{align}
a = (3 \cdot a -12 \cdot b) + (3 \cdot x_1 -12 \cdot x_2) \\
a = 0 \\
\\
a = \frac{a}{-3} + \frac{9}{-3} \\
a = -3 \\
\\
\\
\text{Aplicando o fator } \omega: a = ((1 - \omega) \cdot 0) + (\omega \cdot a) \\
a = -3
\end{align}
$$

2. Para $x_2$ :

$$
\begin{align}
b = (-10 \cdot a +2 \cdot b) + (-10 \cdot x_1 -2 \cdot x_2) \\
b = 30 \\
\\
b = \frac{b}{-2} + \frac{8}{-2} \\
b = -19 \\
\\
\\
\text{Aplicando o fator } \omega: b = ((1 - \omega) \cdot 0) + (\omega \cdot b) \\
b = -19
\end{align}
$$

Dessa forma, ao fim do passo 1, temos $x_1=-3$, $x_2 = -19$

##### Passo 2

1. Para $x_1$ :

$$
\begin{align}
a = (3 \cdot a -12 \cdot b) + (3 \cdot x_1 -12 \cdot x_2) \\
a = 228 \\
\\
a = \frac{a}{-3} + \frac{9}{-3} \\
a = -79 \\
\\
\\
\text{Aplicando o fator } \omega: a = ((1 - \omega) \cdot 0) + (\omega \cdot a) \\
a = -79
\end{align}
$$

2. Para $x_2$ :

$$
\begin{align}
b = (-10 \cdot a +2 \cdot b) + (-10 \cdot x_1 -2 \cdot x_2) \\
b = 790 \\
\\
b = \frac{b}{-2} + \frac{8}{-2} \\
b = -399 \\
\\
\\
\text{Aplicando o fator } \omega: b = ((1 - \omega) \cdot 0) + (\omega \cdot b) \\
b = -399
\end{align}
$$

Dessa forma, ao fim do passo 1, temos $x_1 = -79$, $x_2 = -399$

### Resultados
Durante a execução do algoritmo, notou-se que não foi possível encontrar os valores das raízes do sistema proposto ao utilizar o método de **Gauss-Seidel**. Ao realizar as verificações de convergência utilizadas anteriormente, notou-se que esses sistema não converge utilizando nenhuma delas.
