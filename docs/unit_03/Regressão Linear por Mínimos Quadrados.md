> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

Encontre a reta de mínimos quadrados que aproxima estes dados. Avalie o grau de representação da reta de ajuste aos dados: calcule o desvio padrão total, o erro-padrão da estimativa e o coeficiente de determinação. Inclua a análise estatística no algoritmo para realizar tal avaliação. Trace num mesmo gráfico o conjunto de pontos e a reta do ajuste realizado.

## Fundamentos

#### Métricas de Ajuste

- **Coeficiente de Determinação**
  - Os coeficientes $a_1$ e $a_2$ representam os parâmetros da equação da reta de ajuste, onde $y = a_1 \cdot x + a_2$.
  - O coeficiente de determinação indica o quão bem a reta de ajuste representa os dados.
- **Desvio Padrão Total**
  - O desvio padrão total mede a dispersão total dos dados em relação à média.
- **Erro-Padrão da Estimativa**
  - O erro-padrão da estimativa é uma medida da variabilidade dos resíduos, que são as diferenças entre os valores observados e os valores previstos pela reta de ajuste. Um erro-padrão da estimativa menor indica uma melhor precisão do modelo.

## Algoritmo

```octave
function [a, r2, err] = minimumSquareRegressionMethod(x, y)
    err = ""; % Initialize err to an empty string

    [sampleSize] = length(x);
    % Checking if the given inputs has the same size
    if sampleSize != length(y)
        err = "x and y must have the same length";
        a = [];
        r2 = [];
        return;
    end

    sums = struct('x', 0, 'y', 0, 'xPow2', 0, 'yPow2', 0, 'xy', 0);
    for index = 1:sampleSize
        values = struct('x', x(index), 'y', y(index));
        sums.x = sums.x + values.x;
        sums.y = sums.y + values.y;
        sums.xy = sums.xy + values.x * values.y;
        sums.xPow2 = sums.xPow2 + values.x * values.x;
        sums.yPow2 = sums.yPow2 + values.y * values.y;
    end

    % Calculate coefficients
    a(1) = ((sampleSize * sums.xy) - (sums.x * sums.y)) / ((sampleSize * sums.xPow2) - (sums.x * sums.x));
    a(2) = (sums.y / sampleSize) - (a(1) * (sums.x / sampleSize));

    averageResiduum = 0;
    linearAdjustResiduum = 0;

    for index = 1:sampleSize
        averageDiff = y(index) - (sums.y / sampleSize);
        averageResiduum = averageResiduum + averageDiff * averageDiff;

        linearTempVal = y(index) - a(2) - (a(1) * x(index));
        linearAdjustResiduum = linearAdjustResiduum + linearTempVal * linearTempVal;
    end

    r2 = (averageResiduum - linearAdjustResiduum) / averageResiduum;
end
```

## Questão 01

    Considere o seguinte conjunto de dados.

| **xi** | **yi** |
| :----: | :----: |
|   1    |  1,3   |
|   2    |  3,5   |
|   3    |  4,2   |
|   4    |  5,0   |
|   5    |  7,0   |
|   6    |  8,8   |
|   7    |  10,1  |
|   8    |  12,5  |
|   9    |  13,0  |
|   10   |  15,6  |

### Chamada ao algoritmo

```octave
x = [1; 2; 3; 4; 5; 6; 7; 8; 9; 10];
y = [1.3; 3.5; 4.2; 5.0; 7.0; 8.8; 10.1; 12.5; 13.0; 15.6];

[a, r2, err] = minimumSquareRegressionMethod(x, y);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: minimumSquareRegressionMethod\n');
fprintf('Result [a]:'), disp(a);
fprintf('Result [r2]:'), disp(r2);

% Calculate the adjustment line
y_fit = a(1) * x + a(2);

% Calculate total standard deviation
mean_y = mean(y);
std_total = sqrt(sum((y - mean_y).^2) / (length(y) - 1));

% Calculate the standard error of the estimate
std_error = sqrt(sum((y - y_fit).^2) / (length(y) - 2));

% Statistical analysis
fprintf('Total Standard Deviation: %.4f\n', std_total);
fprintf('Standard error of the Estimate: %.4f\n', std_error);
fprintf('Determination coefficient (r^2): %.4f\n', r2);
```

### Resultados

#### Coeficientes da Reta de Ajuste

- $a_1 = 1.5382$ (coeficiente de inclinação)
- $a_2 = -0.3600$ (coeficiente de interseção)

#### Métricas de Ajuste

- Coeficiente de Determinação: $r^2 = 0.9881$
  - $r^2 = 0.9881$ sugere que aproximadamente 98.81% da variabilidade dos dados é explicada pela reta de ajuste. Isso indica um ajuste muito bom.
- Desvio Padrão Total: $4.6850$
- Erro-Padrão da Estimativa: $0.5414$

![[unit03_exercise_02_quest01.png]]

Os resultados sugerem que a reta de ajuste é uma representação muito eficaz dos dados fornecidos. O coeficiente de determinação $r^2$ próximo de 1 indica um ajuste robusto, e o erro-padrão da estimativa relativamente baixo sugere que a reta é capaz de prever os valores de $y$ com precisão.

## Questão 02

    Ajuste os dados da tabela a uma equação de potencia simples e depois aplique regressão linear:

$$
y = \alpha x^{\beta}%
$$

| **xi** | **yi** |
| :----: | :----: |
|   10   |  125   |
|   20   |   70   |
|   30   |  380   |
|   40   |  550   |
|   50   |  610   |
|   60   |  1220  |
|   70   |  830   |
|   80   |  1450  |

### Chamada ao algoritmo

```octave
xi = [10; 20; 30; 40; 50; 60; 70; 80];
yi = [125; 70; 380; 550; 610; 1220; 830; 1450];

% Fitting logarithms using polyfit
coefficients = polyfit(log(xi), log(yi), 1);

% Adjusted coefficients
alpha = exp(coefficients(2));
beta = coefficients(1);

% Creating a function for the fitted power equation
fitted_function = @(x) alpha * x.^beta;

new_y = fitted_function(xi)
[a, r2, err] = minimumSquareRegressionMethod(xi, new_y);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: minimumSquareRegressionMethod\n');
fprintf('Result [a]:'), disp(a);
fprintf('Result [r2]:'), disp(r2);

% Calculate the adjustment line
x_fit = linspace(min(xi), max(xi), 100);
y_fit = a(1) * xi + a(2);

% Calculate total standard deviation
mean_y = mean(new_y);
std_total = sqrt(sum((new_y - mean_y).^2) / (length(new_y) - 1));

% Calculate the standard error of the estimate
std_error = sqrt(sum((new_y - y_fit).^2) / (length(new_y) - 2));

% Statistical analysis
fprintf('Total Standard Deviation: %.4f\n', std_total);
fprintf('Standard error of the Estimate: %.4f\n', std_error);
fprintf('Determination coefficient (r^2): %.4f\n', r2);
```

### Resultados

#### Coeficientes da Reta de Ajuste

- $a_1 = 2.6529822736366486$ (coeficiente de inclinação)
- $a_2 = -11.276360390242019$ (coeficiente de interseção)

#### Métricas de Ajuste

- Coeficiente de Determinação: $r^2 = 0.9987$
  - $r^2 = 0.9920$ sugere que aproximadamente 99.87% da variabilidade dos dados é explicada pela reta de ajuste. Isso indica um ajuste muito bom.
- Desvio Padrão Total: $65.0262$
- Erro-Padrão da Estimativa: $139.2474$

![[unit03_exercise_02_quest02.png]]

Os resultados indicam que, embora o coeficiente de determinação sugira uma explicação significativa da variabilidade, o desvio padrão total e o erro-padrão da estimativa são elevados. Isso implica que há uma dispersão considerável dos dados em relação à reta de ajuste, indicando que a equação de potência simples e a reta de mínimos quadrados podem não ser modelos ideais para esses dados.
