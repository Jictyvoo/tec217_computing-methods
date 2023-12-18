> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Algoritmo

```octave
function [result, err] = minimumSquarePolynomialRegressionMethod(x, y, polyDegree)
    err = ""; % Initialize err to an empty string

    [sampleSize] = length(x);

    % Checking if the given inputs has the same size
    if sampleSize != length(y)
        err = "x and y must have the same length";
        result = NaN;
        return;
    end

    n = polyDegree + 1;
    a = zeros(n, n);
    b = zeros(n, 1);

    % Generating the linear system
    b(1) = sum(y);
    a(1, 1) = sampleSize;
    for index = 1:n
        k1 = index - 1;
        k2 = index;
        for subIndex = 2:n
            sumAuxiliary = sum(power(x, k2));
            a(index, subIndex) = sumAuxiliary;
            a(subIndex, index) = a(index, subIndex);
            k2 += 1;
        end

        sumAuxiliary = 0;
        if index ~= 1
            for indexIterator = 1:1:sampleSize
                sumAuxiliary += y(indexIterator) * x(indexIterator)^(k1);
            end
            b(index) = sumAuxiliary;
        end
    end

    % Start the linear system solution

    % Superior Triangular Matrix
    [a, b] = solveSuperiorTriangle(n, a, b);

    % Start the system solution calculation
    result = zeros(n, 1);
    result(n) = b(n) / a(n, n);
    for k = n - 1:-1:1
        s = 0;
        for j = k + 1:n
            s = s + a(k, j) * result(j);
        end
        result(k) = (b(k) - s) / a(k, k);
    end
end

function [a, b] = solveSuperiorTriangle(n, a, b)
    for k = 1:n - 1
        for i = k + 1:n
            m = a(i, k) / a(k, k);
            a(i, k) = 0;
            for j = k + 1:n
                a(i, j) -= m * a(k, j);
            end
            b(i) = b(i) - m * b(k);
        end
    end
    return
end

function result = power(x, y)
    result = x.^y;
end
```

## Questão 01

    Teste o Algoritmo aplicando ao conjunto de dados da tabela abaixo:

| **X** | **Y** |
| :---: | :---: |
|   0   |   0   |
|   1   |  20   |
|   2   |  60   |
|   3   |  68   |
|   4   |  77   |
|   5   |  100  |

    Ajuste regressões de 1a a 5a ordem. Faça o gráfico dos valores fornecidos por cada polinômio. Acrescente a esse gráfico os valores dos pontos da tabela. Comparando visualmente a curva dos valores de cada aproximação e os valores da tabela, você pode concluir qual aproximação foi melhor? Inclua o calculo do coeficiente de determinação para sustentar sua análise.

### Chamada ao algoritmo

```octave
x = [0, 1, 2, 3, 4, 5];
y = [0, 20, 60, 68, 77, 100];

for degree = 1:5
    [result, err] = minimumSquarePolynomialRegressionMethod(x, y, double(degree));
    if ~isempty(err)
        disp(err);
        return;
    end
    disp(['PolynomialRegression - ' num2str(degree) ' degree']);
    disp(result');
end
```

### Resultados

#### 1ª ordem

- Coeficientes: $5.67x + 19.4$
- $R^2$: 0.1064

#### 2ª ordem

- Coeficientes: $-0.82x^2 + 29.13x - 1.95$
- $R^2$: 0.8088

#### 3ª ordem

- Coeficientes: $-2.74x^3 + 37.88x^2 - 6.74x + 0.64$
- $R^2$: -51.5908

#### 4ª ordem

- Coeficientes: $-0.42x^4 - 0.81x^3 + 34.85x^2 - 12.90x + 1.35$
- $R^2$: -30.6123

#### 5ª ordem

- Coeficientes: $-3.83 \times 10^{-11}x^5 - 49.58x^4 + 118.71x^3 - 60.54x^2 + 12.29x - 0.87$
- $R^2$: -50727.2701

#### Gráfico comparativo

Ao realizar a execução do algoritmo, foi possível gerar um gráfico comparativo como o indicado abaixo:

![[unit03_exercise03_quest01.png]]

Avaliando visualmente os gráficos, parece que o polinômio de 2ª ordem se ajusta melhor aos dados da tabela. Ele segue a tendência geral dos pontos e parece capturar a forma da curva de maneira mais precisa em comparação com os outros polinômios.

O coeficiente de determinação ($R^2$) para cada regressão polinomial, fornece uma medida quantitativa da qualidade do ajuste aos dados. Os resultados indicam o seguinte:

- O polinômio de 2ª ordem (quadrático) apresenta o maior $R^2$ (0.8088), indicando um bom ajuste aos dados.
- O polinômio de 1ª ordem (linear) tem um $R^2$ relativamente baixo (0.1064), indicando um ajuste mais fraco.
- Os polinômios de graus mais elevados (3ª a 5ª ordem) apresentam $R^2$ negativos, o que sugere que esses modelos podem estar sobreajustando os dados e não generalizando bem para novos dados.

Portanto, a análise quantitativa do $R^2$ confirma a observação visual de que o polinômio de 2ª ordem é uma escolha adequada, apresentando um equilíbrio entre ajuste aos dados e complexidade do modelo.
