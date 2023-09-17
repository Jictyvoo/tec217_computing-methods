> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01

Encontre a primeira raiz não trivial da função a seguir usando um método gráfico e o método da bissecção com intervalo inicial de 0,5 a 1 (obs.: x está em radianos).
Realize duas iterações à mão, em seguida faça os cálculos através do algoritmo até que Ea seja menor que Es = 2%.

$$
\begin{align}
f(x) = sen(x) - x^2
\end{align}
$$

### Realizando o método gráfico

![[01.quest_01.png]]

### Iterações Manuais:

#### **Iteração 1:**

- Intervalo: \(a = 0.5\), \(b = 1.0\)
- Encontrar o ponto médio: \(x = (a + b) / 2 = 0.75\)
- Avaliar a função: \(f(0.75) = \sin(0.75) - 0.75^2 \approx 0.119139\)
- Absolute Error: \(|0.75 - 0.5| = 0.25\)

#### **Iteração 2:**

- Intervalo: \(a = 0.75\), \(b = 1.0\)
- Encontrar o ponto médio: \(x = (a + b) / 2 = 0.875\)
- Avaliar a função: \(f(0.875) = \sin(0.875) - 0.875^2 \approx 0.001919\)
- Absolute Error: \(|0.875 - 0.75| = 0.125\)

### Algoritmo de Bisseção:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% bisectionMethod.m
function [result, totalIteration, err] = bisectionMethod(a, b, fX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    if fX(a) * fX(b) >= 0
        err = "you have not assumed right a and b\n";
        result = NaN;
        totalIteration = 0;
        return
    end

    absoluteError = 1;
    totalIteration = 0;
    result = 0;

    while(absoluteError > epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        oldResult = result;
        result = (a+b)/2;
        if totalIteration > 1
            absoluteError = abs(oldResult - result);
        end
        fR = fX(result);
        if fR == 0
            return
        end
        if fX(a) * fR < 0
            b = result;
        else
            a = result;
        end
    end
    if totalIteration >= maxIterations
        err  ="Failed to find root after all iterations.";
    end
endfunction
```

```octave
% Define the function
fX = @(x)(sin(x) - x^2);

% Define function parameters
a = 0.5;
b = 1;
epsilon = 2e-2;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = bisectionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

### Resultados

Abaixo se encontra a tabela referente à saída do programa com a utilização do método da bisseção.

| Interaction |    a     |    b     | Found-X  |   f(x)    | Absolute-Error |
| :---------: | :------: | :------: | :------: | :-------: | :------------: |
|      1      | 0.500000 | 1.000000 | 0.750000 | 0.119139  |    1.000000    |
|      2      | 0.750000 | 1.000000 | 0.875000 | 0.001919  |    0.125000    |
|      3      | 0.875000 | 1.000000 | 0.937500 | -0.072825 |    0.062500    |
|      4      | 0.875000 | 0.937500 | 0.906250 | -0.034092 |    0.031250    |
|      5      | 0.875000 | 0.906250 | 0.890625 | -0.015748 |    0.015625    |

## Questão 02

Aplique o método da falsa posição na função abaixo no intervalo \[0,1\] considerando Es = 5\*10^-4

$$
\begin{align}
f(x) = x^3 - 9x + 3
\end{align}
$$

### Realizando o método gráfico

![[01.quest_02.png]]

### Iterações Manuais:

Para realizar os calculos, será utilizada a seguinte fórmula:

$$
r = b - \frac{f(b) \cdot (a - b)}{f(a) - f(b)}
$$

Pontos Iniciais: 0 e 1

#### **Iteração 1:**

$$
\begin{align}
&a = 0; b = 1 \\
\\
f(a) = f(0) = 0^3 - 9 \cdot 0 + 3 = 3 \\
f(b) = f(1) = 1^3 - 9 \cdot 1 + 3 = -5 \\
\end{align}
$$

Agora, calcula-se o próximo ponto:

$$
x_{\text{next}} = 1 - \frac{-5 \cdot (1 - 0)}{-5 - 3} = 1 - \frac{5}{8} = \frac{3}{8} = 0.375
$$

---

#### **Iteração 2:**

$$
\begin{align}
&a = 0; b = 0.375 \\
\\
f(a) = f(0) = 0^3 - 9 \cdot 0 + 3 = 3 \\
f(b) = f(0.375) = 0.375^3 - 9 \cdot 0.375 + 3 = -0.322265625 \\
\end{align}
$$

Agora, calcula-se o próximo ponto:

$$
x_{\text{next}} = 0.375 - \frac{-0.322265625 \cdot (0.375 - 0)}{-0.322265625 - 3} \approx \frac{64}{189} \approx 0.338624
$$

### Algoritmo de Falso Positivo:

A implementação do algoritmo solicitado se encontra abaixo:

```octave
%% falsePositiveMethod.m
function [result, totalIteration, err] = falsePositiveMethod(a, b, fX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    if fX(a) * fX(b) >= 0
        err = "there's no root on given interval\n";
        result = NaN;
        totalIteration = 0;
        return
    end

    absoluteError = 1;
    totalIteration = 0;
    result = 0;

    while abs(absoluteError) > epsilon && totalIteration < maxIterations
        totalIteration += 1;
        fResultA = fX(a);
        fResultB = fX(b);
        oldResult = result;

        if fResultA == fResultB
            err = "division by zero error";
            return;
        end
        result = b - (fResultB * (a - b)) / (fResultA - fResultB);
        if totalIteration > 1
            absoluteError = abs(oldResult - result);
        end

        rootResult = fX(result);

        if rootResult == 0
            return;
        end

        if fResultA * rootResult < 0
            b = result;
        else
            a = result;
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end
endfunction
```

```octave
% Define the function
fX = @(x)(x^3 - 9*x + 3);

% Define function parameters
a = 0;
b = 1;
epsilon = 5e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = falsePositiveMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
```

### Resultados

Abaixo se encontra a tabela referente à saída do programa com a utilização do método da bisseção.

| Interaction |    a     |    b     | Found-X  |   f(x)    | Absolute-Error |
| :---------: | :------: | :------: | :------: | :-------: | :------------: |
|      1      | 0.000000 | 1.000000 | 0.375000 | -0.322266 |    1.000000    |
|      2      | 0.000000 | 0.375000 | 0.338624 | -0.008790 |    0.036376    |
|      3      | 0.000000 | 0.338624 | 0.337635 | -0.000226 |    0.000989    |
|      4      | 0.000000 | 0.337635 | 0.337610 | -0.000006 |    0.000025    |
