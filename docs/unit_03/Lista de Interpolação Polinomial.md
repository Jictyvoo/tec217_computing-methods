> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01

    Considere os dados abaixo, determine: a) média; b) mediana; c) moda; d) amplitude; e) desvio-padrão; f) variância; g) coeficiente de variação; h) construa um histograma (use uma faixa de 0,8 a 2,4, com intervalos de 0,2). Desenvolva um script para calcular estatísticas descritivas para um vetor, e que gere um histograma.

|      |      |      |      |      |
|------|------|------|------|------|
| 0.90 | 1.42 | 1.30 | 1.55 | 1.63 |
| 1.32 | 1.35 | 1.47 | 1.95 | 1.66 |
| 1.96 | 1.47 | 1.92 | 1.35 | 1.05 |
| 1.85 | 1.74 | 1.65 | 1.78 | 1.71 |
| 2.29 | 1.82 | 2.06 | 2.14 | 1.27 |

### Chamada ao algoritmo

```octave
inputData = [0.90, 1.42, 1.30, 1.55, 1.63, 1.32, 1.35, 1.47, 1.95, 1.66, 1.96, 1.47, 1.92, 1.35, 1.05, 1.85, 1.74, 1.65, 1.78, 1.71, 2.29, 1.82, 2.06, 2.14, 1.27];

resultMean = mean(inputData);
resultMedian = median(inputData);
resultMode = mode(inputData);
resultRange = range(inputData);
standardDeviation = std(inputData);
variance = var(inputData);

% (g) Coefficient of variation (in percentage)
cv = standardDeviation / resultMean * 100;

% Displaying the results
printf("Mean: %f\n", resultMean);
printf("Median: %f\n", resultMedian);
printf("Mode: %f\n", resultMode);
printf("Range: %f\n", resultRange);
printf("Standard deviation: %f\n", standardDeviation);
printf("Variance: %f\n", variance);
printf("Coefficient of variation: %f%%\n", cv);

% (h) Histogram
hist(inputData, [0.8:0.2:2.4]);
title("Histogram");
xlabel("Values");
ylabel("Frequency");
```

### Resultados
a) **Média:** 1.6244000000000003
b) **Mediana:** 1.65
c) **Moda:** 1.47
d) **Amplitude:** 1.3900000000000001
e) **Desvio-padrão:** 0.33938768392503577
f) **Variância:** 0.115184
g) **Coeficiente de Variação:** 20.893110%

![[unit03_last_list_quest01.png]]

## Questão 02

	Usando a abordagem linear derive o ajuste por mínimos quadrados do modelo y = a1x + e. Ou seja, determine a inclinação para uma reta com interseção na origem. Ajuste os seguintes dados com esse modelo e mostre o resultado por meio de um gráfico, contendo os pontos e a reta ajustada.

| **x** |  2  |  4  |  6  |  7  | 10  | 11  | 14  | 17  | 20  |
| :---: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| **y** |  4  |  5  |  6  |  5  |  8  |  8  |  6  |  9  | 12  |
### Resultados
Ao executar o algoritmo para a geração do gráfico, obteu-se o seguinte resultado em formato de tabela.

![[unit03_last_list_quest02.png]]

## Questão 07

	Ajuste um polinômio cúbico aos seguintes dados. Juntamente com os coeficientes, determine r2 e Sy/x.

| **x** |  3  |  4  |  5  |  7  |  8  |  9  |  11 |  12 |
|:-----:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| **y** | 1,6 | 3,6 | 4,4 | 3,4 | 2,2 | 2,8 | 3,8 | 4,6 |

### Resultados
Os valores calculados se encontram abaixo:
$$
\begin{align}
\text{Para } P_2(x) \\
&a_0 = -11.4887 \\
&a_1 = 7.1438 \\
&a_2 = -1.0412 \\
&a_3 = 0.04667 \\
\\
\end{align}
$$
O valor do coeficiente de determinação da curva $r2 = 0.828981$.
O valor de $\frac{S_y}{x}: 0.570031$.

Ao executar o algoritmo para a geração do gráfico, obteu-se o seguinte resultado em formato de tabela.
![[unit03_last_list_quest07.png]]

## Questão 08

	Use regressão linear múltipla para ajustar os dados abaixo. Calcule os coeficientes da regressão, o erro-padrão da estimativa e o coeficiente de correlação.

| **$x_1$** |   0  |   1  |   1  |   2  |   2  |   3  |   3  |   4  |   4  |
| :-------: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| **$x_2$** |   0  |   1  |   2  |   1  |   2  |   1  |   2  |   1  |   2  |
|  **$y$**  | 15,1 | 17,9 | 12,7 | 25,6 | 20,5 | 35,1 | 29,7 | 45,4 | 40,2 |
### Resultados
Os coeficientes da regressão múltipla são dados por:

$$
\hat{y} = b_0 + b_1x_1 + b_2x_2
$$

Os coeficientes de regressão são dados por:
- $a_0 = 14.460870$
- $a_1 = 9.025217$
- $a_2 = -5.704348$

O erro-padrão da estimativa é dado por $0.888787$.
O coeficiente de correlação entre $x_1$ e $x_2$: $r_{12} = 0.995588$.
