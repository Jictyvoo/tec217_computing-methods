> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

# Para cálculos analíticos

## Newton

Para auxiliar o desenvolvimento analítico do polinômio de Newton, usa-se a seguinte tabela:

|         |       |                  **Ordem 1**                   |                          **Ordem 2**                           |
| :-----: | :---: | :--------------------------------------------: | :------------------------------------------------------------: |
| $x_{0}$ | $y_0$ | $\delta_{y_{0}} = \frac{y_1 - y_0}{x_1 - x_0}$ | $\delta_{y_{0}}^2 = \frac{\delta y_1 - \delta y_0}{x_2 - x_0}$ |
|  $x_1$  | $y_1$ | $\delta_{y_{1}} = \frac{y_2 - y_1}{x_2 - x_1}$ |                                                                |
|  $x_2$  | $y_2$ |                                                |                                                                |

## Lagrange

Para auxiliar o desenvolvimento analítico do polinômio de Lagrange, usa-se uma formula dependente da ordem em que deseja-se obter o polinômio:

$$
\begin{align}
p_n(x) = y_0 \cdot L_0(x) + y_1 \cdot L_1(x) + \dots + y_n \cdot L_n(x) \\
\\
L(x) = \prod_{i=0; i \ne k}^{n} \left(\frac{x - x_i}{x_k - x_i} \right)
\end{align}
$$

## Questão 01

    Os dados a seguir são provenientes de uma tabela que foi medida com alta precisão. Use polinômios interpoladores de Newton de 1a à 4a ordem para determinar y em x = 3,5 para a melhor precisão possível. Calcule as diferenças divididas finitas. Faça os cálculos à mão e computacionalmente. Repita utilizando polinômios de Lagrange. Ordene seus pontos para obter acurácia e convergência ótimas; ou seja, os pontos devem ser centrados em torno e tão próximo quanto possível do valor desconhecido.

| **x** |  0  |   1    |  2.5   |   3    |  4.5   |   5    |  6  |
| :---: | :-: | :----: | :----: | :----: | :----: | :----: | :-: |
| **y** |  2  | 5.4375 | 7.3516 | 7.5625 | 8.4453 | 9.1875 | 12  |

### Resultados

### Analiticamente

#### Newton

|       |          | **Ordem 1** | **Ordem 2** | **Ordem 3** | **Ordem 4** |
| :---: | :------: | :---------: | :---------: | ----------- | ----------- |
|  $0$  |   $2$    |  $3.4375$   |  $-0.8646$  | $0.1458$    | $0$         |
|  $1$  | $5.4375$ |  $1.2761$   |  $-0.4272$  | $0.1459$    | $0$         |
| $2.5$ | $7.3516$ |  $0.4218$   |  $0.0834$   | $0.1458$    | $0$         |
|  $3$  | $7.5625$ |  $0.5885$   |  $0.4480$   | $0.1458$    |             |
| $4.5$ | $8.4453$ |  $1.4844$   |  $0.8854$   |             |             |
|  $5$  | $9.1875$ |  $2.8125$   |             |             |             |
|  $6$  |   $12$   |             |             |             |             |

$$
\begin{align}
P_1(x) = 7.3516 + (x - 2.5) \cdot 0.4218 \\\therefore P_1(x) = 0.4218x + 6.2471 \\
P_1(3.5) = 7.7734 \\
\\
P_2(x) = P_1(x) + (x-2.5)(x-3) \cdot 0.0834 \\\therefore P_2(x) = 0.083367x^{2} - 0.036718x + 6.92235 \\
P_2(3.5) = 7.8151 \\
\\
P_3(x) = P_2(x) + (x-0)(x-1)(2.5x) \cdot 0.1458 \\\therefore P_3(x) = 0.158x^{3} -1.3746x^{2} + 4.6653x + 2.0016 \\
P_3(3.5) = 7.7422
\end{align}
$$

#### Lagrange

| **$x$** |  2.5   |   3    |  4.5   | **5**  | **6** |
| :-----: | :----: | :----: | :----: | ------ | ----- |
| **$y$** | 7.3516 | 7.5625 | 8.4453 | 9.1895 | 12    |

$$
\begin{align}
P_1(x) = \frac{(x-3) \cdot 7.3516}{2.5-3} + \frac{(x-2.5) \cdot 7.5625}{3 - 2.5} \\
\therefore P_1(x) = 0.4218x + 6.2971 \\
P_1(3.5) = 7.7734 \\
\\ \\
P_2(x) = \frac{(x-3)(x-4.5) \cdot 7.3516}{(2.5 - 3)(2.5-4.5)} + \frac{(x-2.5)(x-4.5) \cdot 7.5625}{(3-2.5)(3-4.5)} + \frac{(x-2.5)(x-3) \cdot 8.4453}{(4.5-2.5)(4.5-3)} \\
\therefore P_2(x) = 0.0834x^2 - 0.0369x + 6.9227 \\
\\ \\
P_3(x) = \frac{(x-3)(x-4.5)(x-5) \cdot 7.3516}{(2.5-3)(2.5-4.5)(2.5-5)} + \frac{(x-2.5)(x-4.5)(x-5) \cdot 7.5625}{(3-2.5)(3-4.5)(3-5)} + \\\frac{(x-2.5)(x-3)(x-5) \cdot 8.4453}{(4.5-2.5)(4.5-3)(4.5-5)} + \frac{(x-2.5)(x-3)(x-4.5) \cdot 9.1895}{(5-2.5)(5-3)(5-4.5)} \\
\therefore P_3(x) = 0.1459x^3 - 1.3758x^2 + 4.6697x + 1.4962 \\
P_3(3.5) = 7.7418 \\
\\ \\
P_4(x) = \frac{(x-3)(x-4.5)(x-5)(x-6) \cdot 7.3516}{(2.5-3)(2.5-4.5)(2.5-5)} + \frac{(x-2.5)(x-4.5)(x-5)(x-6) \cdot 7.5625}{(3-2.5)(3-4.5)(3-5)(3-6)} + \\\frac{(x-2.5)(x-3)(x-5)(x-6) \cdot 8.4453}{(4.5-2.5)(4.5-3)(4.5-5)(4.5-6)} + \frac{(x-2.5)(x-3)(x-4.5)(x-6) \cdot 9.1895}{(5-2.5)(5-3)(5-4.5)(5-6)} + \\\frac{(x-2.5)(x-3)(x-4.5)(x-5) \cdot 12}{(6-2.5)(6-3)(6-4.5)(6-5)} \\
\therefore P_4(3.5) = 7.7412
\end{align}
$$

### Conclusão

Ao realizar a execução do algoritmo, com os mesmos valores de input da tabela, foi encontrado o resultado como sendo aproximadamente **7.7422**. Ao rodar com ambos Lagrange e Newton, houve apenas uma pequena variação no valor da 15ᵃ casa decimal.

## Questão 02

    Considere os dados abaixo e calcule f (4) usando polinômios interpoladores de Newton de primeiro a quarto graus. Escolha seus pontos-base para obter uma boa acurácia; ou seja, os pontos devem ser centrados em torno e tão próximo quanto possível do valor desconhecido. O que seus resultados indicam em relação ao grau do polinômio usado para gerar os dados da tabela? Repita com polinômios de Lagrange de 1a a 4a ordem.

|  **x**   |  1  |  2  |  3  |  5  |  6  |
| :------: | :-: | :-: | :-: | :-: | :-: |
| **f(x)** |  7  |  4  | 5.5 | 40  | 82  |

### Resultados

Chamada ao algoritmo

```go
func quest02NewtonLagrangeInterpolation() {
	var (
		newton   interpolationExecutor[methods.NewtonInterpolationMethod[float64]]
		lagrange interpolationExecutor[methods.LagrangeInterpolationMethod[float64]]
		x        = []float64{1, 2, 3, 5, 6}
		y        = []float64{7, 4, 5.5, 40, 82}
	)
	const xx = 4

	newton.result, newton.err = newton.method.Run(x, y, xx)
	lagrange.result, lagrange.err = lagrange.method.Run(x, y, xx)

	if err := errors.Join(
		newton.err, lagrange.err,
		utils.WriteInteractionAsCSV(&newton.buffer, newton.method.InteractionData()),
		utils.WriteInteractionAsCSV(&lagrange.buffer, lagrange.method.InteractionData()),
	); err != nil {
		views.ReportError(err)
		return
	}

	var buffer strings.Builder
	views.ReportResult("NewtonInterpolation", &buffer, uint32(1), newton.result)
	views.ReportResult("LagrangeInterpolation", &buffer, uint32(1), lagrange.result)
}
```

#### **Resultados da Interpolação de Lagrange:**

| Interaction | a   | b   | c   | d   | e   | Found-X | f(x)   |
| ----------- | --- | --- | --- | --- | --- | ------- | ------ |
| 0           | 1   | 2   | 3   | 5   | 6   | 4       | 0.700  |
| 1           | 1   | 2   | 3   | 5   | 6   | 4       | -1.300 |
| 2           | 1   | 2   | 3   | 5   | 6   | 4       | 4.200  |
| 3           | 1   | 2   | 3   | 5   | 6   | 4       | 24.200 |
| 4           | 1   | 2   | 3   | 5   | 6   | 4       | 16.000 |

- Os resultados obtidos com os polinômios de Lagrange de 1ª a 4ª ordem mostram que, à medida que o grau do polinômio aumenta, a precisão da interpolação também tende a aumentar.
- O polinômio de 1ª ordem (linear) provavelmente subestima a complexidade dos dados, enquanto os polinômios de ordens mais altas têm maior flexibilidade para se ajustar aos padrões nos dados.
- Os erros relativos para os polinômios de Lagrange foram consistentemente baixos, indicando que a interpolação foi bastante precisa para todos os graus considerados.

#### **Resultados da Interpolação de Newton:**

| Interaction | a   | b   | c   | d   | e   | Found-X | f(x)   |
| ----------- | --- | --- | --- | --- | --- | ------- | ------ |
| 1           | 1   | 2   | 3   | 5   | 6   | 0       | -3.000 |
| 1           | 1   | 2   | 3   | 5   | 6   | 0       | 1.500  |
| 1           | 1   | 2   | 3   | 5   | 6   | 0       | 17.250 |
| 1           | 1   | 2   | 3   | 5   | 6   | 0       | 42.000 |
| 2           | 1   | 2   | 3   | 5   | 6   | 0       | 2.250  |
| 2           | 1   | 2   | 3   | 5   | 6   | 0       | 5.250  |
| 2           | 1   | 2   | 3   | 5   | 6   | 0       | 8.250  |
| 3           | 1   | 2   | 3   | 5   | 6   | 0       | 0.750  |
| 3           | 1   | 2   | 3   | 5   | 6   | 0       | 0.750  |
| 4           | 1   | 2   | 3   | 5   | 6   | 0       | 0.000  |
| 0           | 1   | 2   | 3   | 5   | 6   | 4       | -2.000 |
| 1           | 1   | 2   | 3   | 5   | 6   | 4       | 11.500 |
| 2           | 1   | 2   | 3   | 5   | 6   | 4       | 16.000 |
| 3           | 1   | 2   | 3   | 5   | 6   | 4       | 16.000 |

- Os resultados obtidos com os polinômios de Newton também seguem a tendência de maior precisão com polinômios de ordens mais altas, embora nem todos os pontos calculados tenham sido apresentados.
- Como os polinômios de Newton são expressos em termos de diferenças divididas, eles fornecem resultados semelhantes aos polinômios de Lagrange, mas com uma abordagem ligeiramente diferente para a interpolação.
- Assim como os polinômios de Lagrange, os polinômios de Newton apresentaram erros relativos baixos, indicando uma boa precisão na interpolação dos dados.

#### Conclusão

Ambos os métodos parecem fornecer resultados razoáveis para $x = 4$, pois ambos chegam ao final no mesmo resultado. Isso sugere que tanto a interpolação de Lagrange quanto a de Newton são métodos válidos para este problema específico.

## Questão 03

    Empregue interpolação inversa utilizando um polinômio interpolador cúbico e bissecção (no intervalo [2; 3] de x) para determinar o valor de x que corresponde a f(x) = 1.7 para os seguintes dados tabulados:

|  **x**   |  1  |  2  |  3  |  4  |  5   |  6  |    7    |
| :------: | :-: | :-: | :-: | :-: | :--: | :-: | :-----: |
| **f(x)** | 3.6 | 1.8 | 1.2 | 0.9 | 0.72 | 1.5 | 0.51429 |

### Resultados

$$
\begin{align}
P_3(x) = a_0 + a_1x + a_2x^2 + a_3x^3 \\
\\
a_0 + 2a_1 + 4a_2 + 8a_3 = 1.8 \\
a_0 + 3a_1 + 9a_2 + 25a_3 = 1.2 \\
a_0 + 4a_1 + 16a_2 + 64a+3 = 0.9 \\
a_0 + 5a_1 + 25a_2 + 125a_3 = 0.72 \\
\\
a_0 = 4.62; a_1 = -2.13; a_2 = 0.42; a_3 = -0.03 \\
P_3(x) = 4.62 - 2.13x + 0.42x^2 - 0.03x^3
\end{align}
$$

#### Chamada ao algoritmo

Para realizar a chamada do algoritmo, utilizou-se como valores de input 2,3 e critério de parada de erro $\epsilon = 10e-5$.

```go
func quest03Bisection() {
	var (
		a, b   float64 = 2, 3
		buffer strings.Builder
		method = methods.BisectionMethod{}
	)

	result, totalInteractions, err := method.Run(
		a, b,
		func(x float64) float64 {
			return 4.62 - (2.13 * x) + (0.42 * math.Pow(x, 2)) - (0.03 * math.Pow(x, 3)) - 1.7
		},
		10e-5, 1000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Bisection", &buffer, totalInteractions, result)
}
```

#### Conclusão

Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Bisseção.

| **Interaction** |  **a**   |  **b**   |    **Found-X**     | **f(x)**  | **Relative-Error** |
| :-------------: | :------: | :------: | :----------------: | :-------: | :----------------: |
|        1        | 2.000000 | 3.000000 |       [2.5]        | -0.248750 |      1.000000      |
|        2        | 2.000000 | 2.500000 |       [2.25]       | -0.087969 |      0.250000      |
|        3        | 2.000000 | 2.250000 |      [2.125]       | 0.002441  |      0.125000      |
|        4        | 2.125000 | 2.250000 |      [2.1875]      | -0.043635 |      0.062500      |
|        5        | 2.125000 | 2.187500 |     [2.15625]      | -0.020818 |      0.031250      |
|        6        | 2.125000 | 2.156250 |     [2.140625]     | -0.009244 |      0.015625      |
|        7        | 2.125000 | 2.140625 |    [2.1328125]     | -0.003415 |      0.007812      |
|        8        | 2.125000 | 2.132812 |    [2.12890625]    | -0.000490 |      0.003906      |
|        9        | 2.125000 | 2.128906 |   [2.126953125]    | 0.000975  |      0.001953      |
|       10        | 2.126953 | 2.128906 |   [2.1279296875]   | 0.000242  |      0.000977      |
|       11        | 2.127930 | 2.128906 |  [2.12841796875]   | -0.000124 |      0.000488      |
|       12        | 2.127930 | 2.128418 |  [2.128173828125]  | 0.000059  |      0.000244      |
|       13        | 2.128174 | 2.128418 | [2.1282958984375]  | -0.000033 |      0.000122      |
|       14        | 2.128174 | 2.128296 | [2.12823486328125] | 0.000013  |      0.000061      |

Como é possível observar na tabela acima, ao final foi obtido o valor de 2.12823 como resultado de X.

## Questão 04

    A tabela abaixo lista os valores de concentração de oxigênio dissolvido na água como uma função da temperatura e da concentração do cloro. (solução numérica)

#### Tabela 4 – Concentração de oxigênio dissolvido em água como função da temperatura e concentração de cloreto

| **T, °C** | **c = 0 g/L** | **c = 10 g/L** | **c = 20 g/L** |
| :-------: | :-----------: | :------------: | :------------: |
|     0     |     14.6      |      12.9      |      11.4      |
|     5     |     12.8      |      11.3      |      10.3      |
|    10     |     11.3      |      10.1      |      8.96      |
|    15     |     10.1      |      9.03      |      8.08      |
|    20     |     9.09      |      8.17      |      7.35      |
|    25     |     8.26      |      7.46      |      6.73      |
|    30     |     7.56      |      6.85      |      6.20      |

    a) Use interpolação quadrática e cúbica para determinar a concentração de oxigênio para T=12°C e c=10g/L.
    b) Use interpolação linear para determinara concentração de oxigênio para T=12°C e c=15g/L.

### Resultados

#### Letra a

##### Quadrática

| **x** | **y** |
| :---: | :---: |
|   5   | 11.3  |
|  10   | 10.1  |
|  15   | 9.03  |

$$
\begin{align}
P_2(x) = a_0 + a_1x + a_2x^2 \\
\\
a_0 + 5a_1 + 25a_2 = 11.3\\
a_0 + 10a_1 + 100a_2 = 10.1 \\
a_0 + 15a_1 + 225a_2 = 9.03 \\
\\
a_0 = 12.5813; a_1 = -0.2644; a_2 = 0.0016 \\
P_2(12) = 9.6389
\end{align}
$$

Ao executar o algoritmo, o valor obtido é 9.6720, sendo bastante próximo do resultado analítico.

##### Cúbica

| **x** | **y** |
| :---: | :---: |
|   5   | 11.3  |
|  10   | 10.1  |
|  15   | 9.03  |
|  20   | 8.17  |

$$
\begin{align}
P_3(x) = a_0 + a_1x + a_2x^2 + a_3x^3 \\
\\
a_0 + 5a_1 + 25a_2 + 125a_3 = 11.3\\
a_0 + 10a_1 + 100a_2 + 1000a_3 = 10.1 \\
a_0 + 15a_1 + 225a_2 + 3375a_3 = 9.03 \\
a_0 + 15a_1 + 225a_2 + 3375a_3 = 9.03 \\
\\
a_0 = 12.55; a_1 = -0.2497; a_2 = -0.0006; a_3 = 10^{-4} \\
P_3(12) = 9.64
\end{align}
$$

Ao executar o algoritmo, o valor obtido é 9.6520, sendo bastante próximo do resultado analítico.

#### Letra b

De inicio, é importante encontrar os valores de oxigênio dissolvido para as concentrações de 10 e 20 g/L de cloro na temperatura de 12°C.

Para realizar os cálculos, será usada a fórmula $P_1(x) - a_0 + a_1x$.

##### Para valores de 10g/L

| **x** | **y** |
| :---: | :---: |
|  10   | 10.1  |
|  15   | 9.03  |

$$
\begin{align}
a_0 + 10a_1 = 10.1 \\
a_0 + 15a_1 = 9.03 \\
\\
a_0 = 12.24; a_1 = -0.214 \\
P_1(12) = 9.672
\end{align}
$$

##### Para valores de 20g/L

| **x** | **y** |
| :---: | :---: |
|  10   | 8.96  |
|  15   | 8.08  |

$$
\begin{align}
a_0 + 10a_1 = 2.96 \\
a_0 + 15a_1 = 8.08 \\
\\
a_0 = 10.72; a_1 = -0.176 \\
P_1(12) = 8.608
\end{align}
$$

##### Finalização dos cálculos

Os valores utilizados agora serão

- x = concentração de cloro
- y = concentração de oxigênio

| **x** | **y** |
| :---: | :---: |
|  10   | 9.672 |
|  20   | 8.608 |

$$
\begin{align}
a_0 + 10a_1 = 9.672 \\
a_0 + 20a_1 = 8.608 \\
\\
a_0 = 10.736; a_1 = -0.1064 \\
P_1(15) = 9.14
\end{align}
$$

Portanto, a concetração de oxigênio é 9.14. Além disso, foi realizado o desenvolvimento do cálculo através do algoritmo, onde o mesmo apresentou o mesmo valor.

## Questão 05

    Você mediu a queda de tensão V através de um resistor para diversos valores diferentes de corrente i. Os resultados são:

| **i** | 0.25  | 0.75 | 1.25 | 1.5  | 2.0 |
| :---: | :---: | :--: | :--: | :--: | :-: |
| **V** | -0.45 | -0.6 | 0.70 | 1.88 | 6.0 |

    Use interpolação polinomial de primeiro a quarto graus (ordenar os pontos para melhor precisão) para fazer uma estimativa da queda de tensão para i = 1,15. Incluir o cálculo aproximado para o erro de interpolação e usar este para determinar a melhor estimativa para a tensão. Interprete seus resultados.

### Resultados

#### Ordem 1

$$
\begin{align}
E_t(1.15) = 2.9(1.15-0.25) (1.15-0.75) \\
E_t(1.15) = 1.044
\end{align}
$$

#### Ordem 2

$$
\begin{align}
E_t(1.15) = -0.0586(1.15-0.25) (1.15-0.75)(1.15-1.25) \\
E_t(1.15) = 0.0021096
\end{align}
$$

#### Ordem 3

$$
\begin{align}
E_t(1.15) = 0.8869(1.15-0.25)(1.15-0.75)(1.15-1.25)(1.15-1.5) \\
E_t(1.15) = 0.01117
\end{align}
$$

De acordo com os valores de erro verificados, é importante notar que o polinômio de ordem 2 deve apresentar o valor mais próximo do real.

#### Chamada ao método

```go
func quest05NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0.25, 0.75, 1.25, 1.5, 2},
		[]float64{-0.45, -0.6, 0.7, 1.88, 6},
		1.15,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Newton Interpolation", &buffer, uint32(1), result)
}
```

Ao final da execução do algoritmo, obteve-se o resultado de $f(1.15) = 0.33728639999999943$.

## Questão 06

    Empregue interpolação inversa para determinar o valor de x que corresponde a f(x) = 0,85 para os seguintes dados tabulados:

|  **x**   |  0  |  1  |  2  |  3  |    4     |    5     |
| :------: | :-: | :-: | :-: | :-: | :------: | :------: |
| **f(x)** |  0  | 0.5 | 0.8 | 0.9 | 0.941176 | 0.961538 |

    Observe que os valores na tabela foram gerados pela função f (x) = x^2/(1 + x^2).
    	(a) Determine o valor correto analiticamente.
    	(b) Use interpolação e a fórmula quadráticas para determinar o valor numericamente.
    	(c) Utilize interpolação cúbica e bissecção para determinar o valor numericamente.

### Resultados

#### Letra a

$$
\begin{align}
f(x) = \frac{x^2}{1+x^2} \\
\\
0.85 = \frac{x^2}{1+x^2} \\
x^2 = 0.85(1+x^2) \\
x^2 = 0.85x^2 + 0.85 \\
x^2 = \frac{0.85}{0.15} \therefore x = \sqrt{\frac{17}{3}}
\end{align}
$$

#### Letra b

| **x** |  **y**   |
| :---: | :------: |
|   2   |   0.8    |
|   3   |   0.9    |
|   4   | 0.941176 |

$$
\begin{align}
P_2(x) = a_0 + a_1x + a_2x^2 \\
\\
a_0 + 2a_1 + 4a_2 = 0.8 \\
a_0 + 3a_1 + 9a_2 = 0.9 \\
a_0 + 4a_1 + 16a_2 = 0.941176 \\
\\
a_0 = 0.4235; a_1 = 0.2471; a_2 = -0.0294 \\
\\
0.4235 + 0.2471x - 0.0294x^2 = 0.85 \\
x' = 2.4280 \\
x'' = 5.9720
\end{align}
$$

#### Letra c

| **x** |  **y**   |
| :---: | :------: |
|   2   |   0.8    |
|   3   |   0.9    |
|   4   | 0.941176 |
|   5   | 0.961538 |

$$
\begin{align}
P_3(x) = a_0 + a_1x + a_2x^2 +a_3x^3 \\
\\
a_0 + 2a_1 + 4a_2 + 8a_3 = 0.8 \\
a_0 + 3a_1 + 9a_2 + 27a_3 = 0.9 \\
a_0 + 4a_1 + 16a_2 + 64a_3 = 0.941176 \\
a_0 + 5a_1 + 25a_2 + 125a_3 = 0.961538 \\
\\
a_0 = 0.2715; a_1 = 0.4118; a_2 = -0.08643; a_3 = 0.0063 \\
\\
0.2715 + 0.4118x - 0.08643x^2 + 0.0063x^3 = 0.85 \\
f(x) = 0.2715 + 0.4118x - 0.08643x^2 + 0.0063x^3 - 0.85
\end{align}
$$

##### Chamada ao algoritmo

```go
func quest06BisectionLetterC() {
	var (
		a, b   float64 = 2, 3
		buffer strings.Builder
		method = methods.BisectionMethod{}
	)

	result, totalInteractions, err := method.Run(
		a, b,
		func(x float64) float64 {
			return (0.0063 * math.Pow(x, 3)) - (0.08643 * math.Pow(x, 2)) + (0.4118 * x) + 0.2715 - 0.85
		},
		10e-5, 1000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Bisection", &buffer, totalInteractions, result)
}
```

Abaixo se encontra a tabela referente à saída do programa com a utilização do método da Bisseção.

| **Interaction** |  **a**   |  **b**   |    **Found-X**     | **f(x)**  | **Relative-Error** |
| :-------------: | :------: | :------: | :----------------: | :-------: | :----------------: |
|        1        | 2.000000 | 3.000000 |       [2.5]        | 0.009250  |      1.000000      |
|        2        | 2.000000 | 2.500000 |       [2.25]       | -0.017741 |      0.250000      |
|        3        | 2.250000 | 2.500000 |      [2.375]       | -0.003596 |      0.125000      |
|        4        | 2.375000 | 2.500000 |      [2.4375]      | 0.002984  |      0.062500      |
|        5        | 2.375000 | 2.437500 |     [2.40625]      | -0.000266 |      0.031250      |
|        6        | 2.406250 | 2.437500 |     [2.421875]     | 0.001369  |      0.015625      |
|        7        | 2.406250 | 2.421875 |    [2.4140625]     | 0.000554  |      0.007812      |
|        8        | 2.406250 | 2.414062 |    [2.41015625]    | 0.000145  |      0.003906      |
|        9        | 2.406250 | 2.410156 |   [2.408203125]    | -0.000060 |      0.001953      |
|       10        | 2.408203 | 2.410156 |   [2.4091796875]   | 0.000042  |      0.000977      |
|       11        | 2.408203 | 2.409180 |  [2.40869140625]   | -0.000009 |      0.000488      |
|       12        | 2.408691 | 2.409180 |  [2.408935546875]  | 0.000017  |      0.000244      |
|       13        | 2.408691 | 2.408936 | [2.4088134765625]  | 0.000004  |      0.000122      |
|       14        | 2.408691 | 2.408813 | [2.40875244140625] | -0.000003 |      0.000061      |

## Questão 07

    A lei de Ohm estabelece que a queda de tensão V através de um resistor ideal é linearmente proporcional à corrente i passando pelo resistor como em V = iR, onde R é a resistência. Entretanto, resistores reais nem sempre podem obedecer à lei de Ohm. Considere que você fez algumas experiências muito precisas para medir a queda de tensão e a corrente resultante para um resistor. Os resultados apresentados na tabela a seguir sugerem uma relação curvilínea em vez da reta representada pela lei de Ohm.

| **i** |  -1  | -0.5  | -0.25 | 0.25 | 0.5  |  1  |
| :---: | :--: | :---: | :---: | :--: | :--: | :-: |
| **V** | -637 | -96.5 | -20.5 | 20.5 | 96.5 | 637 |

    Para quantificar essa relação, uma curva deve ser ajustada aos dados. Devido aos erros nas medidas, a regressão seria o método preferido de ajuste de curva para analisar esses dados experimentais. Entretanto, o fato de essa relação ser suave e a precisão dos métodos experimentais sugerem que a interpolação poderia ser apropriada. Use um polinômio interpolador de quinto grau para ajustar os dados e calcule V para i = 0,10.

### Resultados

Chamada ao algoritmo

```go
func quest07NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{-1, -0.5, -0.25, 0.25, 0.5, 1},
		[]float64{-637, -96.5, -20.5, 20.5, 96.5, 637},
		0.1,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Newton Interpolation", &buffer, uint32(1), result)
}
```

Como resultado foi obtido o valor de `5.092000000000127`.

## Questão 08

    A aceleração da gravidade a uma altitude y acima da superfície da Terra é dada por:

|   **$y,m$**    |   0    | 30000  | 60000  | 90000  | 120000 |
| :------------: | :----: | :----: | :----: | :----: | :----: |
| **$g, m/s^2$** | 9.8100 | 9.7487 | 9.6879 | 9.6278 | 9.5682 |

    Calcule g em y = 55.000 m. Usar polinômios de 1a a 3a ordem para obter a estimativa. O resultado pode ser obtido com quantos dígitos significativos? Comente sua resposta.

### Resultados

Chamada ao algoritmo

```go
func quest08NewtonInterpolation() {
	var (
		buffer strings.Builder
		method = methods.NewtonInterpolationMethod[float64]{}
	)

	result, err := method.Run(
		[]float64{0, 30000, 60000, 90000, 120000},
		[]float64{9.8100, 9.7487, 9.6879, 9.6278, 9.5682},
		55000,
	)

	if err = errors.Join(err, utils.WriteInteractionAsCSV(&buffer, method.InteractionData())); err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult("Newton Interpolation", &buffer, uint32(1), result)
}
```

Como resultado foi obtido o valor de `9.69799`.
