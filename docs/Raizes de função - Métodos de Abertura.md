> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Todos os algoritmos devem ser desenvolvidos e aplicados na função abaixo:
$$
\begin{align}
f(x) = x^2 -3x + e^x - 2
\end{align}
$$
Essa função possui duas raízes, uma positiva e uma negativa. Deve ser encontrada a menor raíz em módulo. Para isso será utilizada a tolerância de 1e-4.
- Fazer manualmente duas iterações passo a passo de cada método.
- Implementar os algoritmos.  
- Aplicar os algoritmos para obter a solução numérica da questão.

## Análise Gráfica inicial
![[02.first_func.png]]
Baseado no gráfico abaixo, foi optado a busca da raiz à esquerda, utilizando valores entre -1 e 0.
## Método da Iteração Linear

$$
\begin{align}
& f(x) = x^2 - 3x + e^x - 2 \\
& g(x) = \frac{x^2 - 2 + e^x}{3}
\end{align}
$$
### Iterações Manuais:

#### **Iteração 1:**
$$
\begin{align}
x_0 = -1 \\
x_1 = g(x_0) = \frac{(-1)^2 - 2 + e^{-1}}{3} \approx -0.210707 \\
\end{align}
$$
Agora, é calculado o valor obtido pela raiz para verificar a convergência:

$$
f(x_1) = (-0.210707)^2 - 3(-0.210707) + e^{-0.210707} - 2 \approx -0.513471
$$
Como é a primeira iteração, o erro continua em 1

#### **Iteração 2:**
$$
\begin{align}
x_1 = -0.210707 \\
x_2 = g(x_1) = \frac{(-0.210707)^2 - 2 + e^{-0.210707}}{3} \approx -0.381864 \\
\end{align}
$$

Agora, é calculado o valor obtido pela raiz para verificar a convergência:
$$
f(x_2) = (-0.381864)^2 - 3(-0.381864) + e^{-0.381864} - 2 \approx -0.026001
$$
$$
\text{Erro relativo} = \left|\frac{x_2 - x_1}{x_2}\right| \approx \left|\frac{-0.381864 - (-0.210707)}{-0.381864}\right| \approx 0.448215
$$

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% linearIterationMethod.m
function [result, totalIteration, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    absoluteError = 1;
    totalIteration = 0;
    oldResult = 0;
    result = x0;

    while (abs(absoluteError) >= epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        oldResult = result;
        result = gX(result);

        if result != 0 && totalIteration > 1
            absoluteError = abs(oldResult - result) / result;
        end

        rootResult = fX(result);
        if rootResult == 0
            break
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
```

```octave
% Define the function
fX = @(x)((x^2) - (3 * x) + (e^x) -2);
gX = @(x)((x^2) - 2 + (e^x)) / 3;

% Define function parameters
x0 = -1;
epsilon = 1e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```


### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Iteração Linear.

| Interaction |     a     |  Found-X  |    f(x)   | Relative-Error |
|:-----------:|:---------:|:---------:|:---------:|:--------------:|
|      1      | -1.000000 | -0.210707 | -0.513471 |    1.000000    |
|      2      | -1.000000 | -0.381864 | -0.026001 |    -0.448215   |
|      3      | -1.000000 | -0.390531 |  0.000804 |    -0.022193   |
|      4      | -1.000000 | -0.390263 | -0.000028 |    -0.000687   |
|      5      | -1.000000 | -0.390272 |  0.000001 |    -0.000024   |


## Método de Newton Raphson
$$
\begin{align}
& f(x) = x^2 - 3x + e^x - 2 \\
& f'(x) = 2x + e^x - 3
\end{align}
$$
### Iterações Manuais:

#### **Iteração 1:**
$$
\begin{align}
x_0 = -1 \\
f(x_0) = (-1)^2 - 3(-1) + e^{-1} - 2 \approx 2.367879 \\
f'(x_0) = 2(-1) + e^{-1} - 3 \approx -4.632121 \\
\\
& \text{Calcula-se o valor de }x_1 \\
x_1 = x_0 - \frac{f(x_0)}{f'(x_0)} \approx -1 - \frac{0.318731}{-5.71828} \approx -0.488813
\end{align}
$$

Agora, é calculado o valor obtido pela raiz para verificar a convergência:

$$
f(x_1) = (-0.488813)^2 - 3(-0.488813) + e^{-0.488813} - 2 \approx 0.318731
$$
Como é a primeira iteração, o erro continua em 1

#### **Iteração 2:**
$$
\begin{align}
x_1 = -0.488813 \\
f(x_1) = (-0.488813)^2 - 3(-0.488813) + e^{-0.488813} - 2 \approx 0.318731 \\
f'(x_1) = 2(-0.488813) + e^{-0.488813} - 3 \approx -3.364272 \\
\\
& \text{Calcula-se o valor de }x_2 \\
x_2 = x_1 - \frac{f(x_1)}{f'(x_1)} \approx -0.488813 - \frac{0.318731}{-3.364272} \approx -0.394073
\end{align}
$$

Agora, é calculado o valor obtido pela raiz para verificar a convergência:
$$
f(x_2) = (-0.394073)^2 - 3(-0.394073) + e^{-0.394073} - 2 \approx 0.011817
$$
$$
\text{Erro relativo} = \left|\frac{x_2 - x_1}{x_2}\right| \approx \left|\frac{-0.394073 - (-0.488813)}{-0.394073}\right| \approx 0.240413
$$

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% newtonRaphsonMethod.m
function [result, totalIteration, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    relativeError = 1;
    totalIteration = 0;
    oldResult = 0;
    result = x0;

    while (abs(relativeError) >= epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        oldResult = result;
        result = result - (fX(oldResult) / f1X(oldResult));

        if result != 0 && totalIteration > 1
            relativeError = abs(oldResult - result) / result;
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
```

```octave
% Define the function
fX = @(x)((x^2) - (3 * x) + (e^x) -2);
f1X = @(x)((2 * x) + (e^x) - 3);

% Define function parameters
x0 = -1;
epsilon = 1e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```


### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método de Newton Raphson.

| Interaction |     x0    |  Found-X  |   f(x)   | Relative-Error |
|:-----------:|:---------:|:---------:|:--------:|:--------------:|
|      1      | -1.000000 | -0.488813 | 0.318731 |    1.000000    |
|      2      | -1.000000 | -0.394073 | 0.011817 |    -0.240413   |
|      3      | -1.000000 | -0.390278 | 0.000019 |    -0.009724   |
|      4      | -1.000000 | -0.390272 | 0.000000 |    -0.000016   |


## Método da Secante
$$
\begin{align}
& f(x) = x^2 - 3x + e^x - 2
\end{align}
$$
### Iterações Manuais:

#### **Iteração 1:**
$$
\begin{align}
x_0 = 0; x_1 = -1 \\
f(x_0) = (0)^2 - 3(0) + e^{0} - 2 \approx -1 \\
f(x_1) = (-1)^2 - 3(-1) + e^{-1} - 2 \approx 2.367879 \\
\\
& \text{Calcula-se o valor de }x_2 \\
x_2 = \frac{((x_0 \cdot f(x_1)) - (x_1 \cdot f(x_0))}{f(x_1) - f(x_0)} \approx -0.296923
\end{align}
$$

Agora, é calculado o valor obtido pela raiz para verificar a convergência:

$$
f(x_2) = (-0.296923)^2 - 3(-0.296923) + e^{-0.296923} - 2 \approx -0.277967
$$
Como é a primeira iteração, o erro continua em 1

#### **Iteração 2:**
$$
\begin{align}
x_0 = -0.296923; x_1 = 0 \\
f(x_0) = (-0.296923)^2 - 3(-0.296923) + e^{-0.296923} - 2 \approx -0.277967 \\
f(x_1) = (0)^2 - 3(0) + e^{0} - 2 \approx -1 \\
\\
& \text{Calcula-se o valor de }x_2 \\
x_2 = \frac{((x_0 \cdot f(x_1)) - (x_1 \cdot f(x_0))}{f(x_1) - f(x_0)} \approx -0.411232
\end{align}
$$

Agora, é calculado o valor obtido pela raiz para verificar a convergência:
$$
f(x_2) = (-0.411232)^2 - 3(-0.411232) + e^{-0.411232} - 2 \approx 0.065640
$$
$$
\text{Erro relativo} = \left|\frac{x_2 - x_1}{x_2}\right| \approx \left|\frac{-0.296923 - (-0.411232)}{-0.411232}\right| \approx 0.277967
$$

### Algoritmo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% secantMethod.m
function [result, totalIteration, err] = secantMethod(x0, x1, fX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    relativeError = 1;
    totalIteration = 0;

    while (abs(relativeError) >= epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        fPrevious = fX(x0);
        fCurrent = fX(x1);

        result = ((x0 * fCurrent) - (x1 * fPrevious)) / (fCurrent - fPrevious);
        if result != 0 && totalIteration > 1
            relativeError = abs(x0 - result) / result;
        end
        x1 = x0;
        x0 = result;
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
```

```octave
% Define the function
fX = @(x)((x^2) - (3 * x) + (e^x) -2);

% Define function parameters
x0 = 0;
x1 = -1;
epsilon = 1e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = secantMethod(x0, x1, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```


### Resultados
Abaixo se encontra a tabela referente à saída do programa com a utilização do método de Newton Raphson.

| Interaction |     x0    |     x1    |  Found-X  |    f(x)   | Relative-Error |
|:-----------:|:---------:|:---------:|:---------:|:---------:|:--------------:|
|      1      |  0.000000 | -1.000000 | -0.296923 | -0.277967 |    1.000000    |
|      2      | -0.296923 |  0.000000 | -0.411232 |  0.065640 |    -0.277967   |
|      3      | -0.411232 | -0.296923 | -0.389395 | -0.002720 |    -0.056078   |
|      4      | -0.389395 | -0.411232 | -0.390264 | -0.000024 |    -0.002226   |
|      5      | -0.390264 | -0.389395 | -0.390272 |  0.000000 |    -0.000020   |


## O Problema Final
	O tamanho critico de um reator nuclear é determinado por uma equação de criticidade. Uma forma simples é dada por: tan(0.1x) = 9.2 e^-x A solução fisicamente significativa é a menor raiz positiva satisfazendo 3 < x < 4. Determine a menor raiz positiva. Defina α como 1/f'(xi). Calcule a estimativa da derivada através de diferença finita nos pontos (x=3; f(3)) e (x=4;f(4))

Para realizar o cálculo corretamente, foi separada a função f(x) e descoberta sua derivada, de modo a ser possível encontrar sua raiz a partir do valor inicial 3.5, que se encontra entre 3 e 4 conforme solicitado pela questão.
$$
\begin{align}
f(x) = \tan(0.1x) - 9.2 \cdot e^{-x} \\
f'(x) = 9.2 \cdot e^{-x} + 0.1\left(\frac{1}{\cos^2(0.1x)}\right) \\
\end{align}
$$

Dessa forma, ao se aplicar o algoritmo de Newton-Raphson, com uma tolerância igual a **1e-8**, foi obtido o valor final da raiz como `3.292924614900908`

| Interaction |     a    |  Found-X |    f(x)   | Relative-Error |
|:-----------:|:--------:|:--------:|:---------:|:--------------:|
|      1      | 3.500000 | 3.277030 | -0.007249 |    1.000000    |
|      2      | 3.500000 | 3.292832 | -0.000042 |    0.004799    |
|      3      | 3.500000 | 3.292925 | -0.000000 |    0.000028    |
|      4      | 3.500000 | 3.292925 |  0.000000 |    0.000000    |

Com o objetivo de também ser realizado o cálculo do alpha para aplicar o algoritmo do ponto fixo, utilizou-se a função:
$$
g(x) = x - \left(\frac{1}{f'(x)} \cdot f(x) \right)
$$
Obtendo por fim o mesmo resultado que o método de Newton-Raphson.

### Cálculo da Estimativa da Derivada

Para realizar o cálculo da derivada de f(x) em dois pontos, x = 3 e x = 4, foi utilizado o método da diferença finita com um tamanho de passo **h = 1e-3**. Os cálculos foram feitos da seguinte maneira:

1. Para x = 3:
	- O valor real da derivada é de aproximadamente 0.5676099.
$$
\begin{align}
f(3)\text{ foi calculado como }\tan(0.1 \cdot 3) - 9.2 \cdot e^{-3} \approx -0.148704779 \\
f'(3)\text{ foi estimado usando a diferença finita como }\frac{f(3 + 0.001) - f(3)}{0.001} \approx 0.567384366 \\
\end{align}
$$

2. Para x = 4:
	- O valor real da derivada é de aproximadamente 0.28630006.
$$
\begin{align}
f(4)\text{ foi calculado como }\tan(0.1 \cdot 4) - 9.2 \cdot e^{-4} \approx 0.254289341 \\
f'(4)\text{ foi estimado usando a diferença finita como }\frac{f(4 + 0.001) - f(4)}{0.001} \approx 0.286300049 \\
\end{align}
$$

A diferença entre os valores estimados e os valores obtidos é bastante pequena em ambos os casos. Devido à essa diferença, foi-se aos poucos diminuindo o valor do tamanho de passo, de modo a perceber que o valor da derivada se aproximava cada vez mais do valor esperado.

Assim sendo, torna-se possível de adquirir o valor da derivada corretamente sem necessariamente realizar a derivação da função antes de enviar à entrada do método.
