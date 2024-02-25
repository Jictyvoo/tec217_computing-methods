> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01

    Calcule a seguinte integral:

$$
\int_{0}^{4} (1 - e^{-x})dx
$$

    a) Analiticamente.
    b) Por aplicação única da regra do trapézio.
    c) Por aplicações múltiplas da regra do trapézio, com n = 2 e 4.
    d) Por aplicação única da regra de Simpson de 1/3.
    e) Por aplicações múltiplas da regra de Simpson de 1/3, com n = 4.
    f) Por aplicação única da regra de Simpson de 3/8.
    g) Por aplicações múltiplas da regra de Simpson de 3/8, com n = 5.

    Para cada estimativa numérica de b) a g), determine o erro relativo percentual com base em a).

### Analiticamente

$$
\begin{align}
- \int_{0}^{4} e^{-x}dx + \int_{0}^{4}1dx \\
\\
\left[x\right]_{0}^{4} - \left[-e^{-x}\right]_{0}^{4} \\
\left[x + e^{-x}\right]_{0}^{4} \\
\left[(4 + e^{-4}) - (0 + e^0)\right] \\
4 + e^{-4} - 1 \\
3 + e^{-4} \\
\\
\\
\therefore - \int_{0}^{4} e^{-x}dx + \int_{0}^{4}1dx = 3.01831563889
\end{align}
$$

### Resultados

#### Letra b

Ao executar o algoritmo, o valor encontrado foi de: **1.96337**

$$
\begin{align}
Err = \frac{\left|3.0183-1.96337 \right|}{3.0183} \cdot 100 \\
Err = 34.9501\%
\end{align}
$$

#### Letra c

- Para n=2
  Ao executar o algoritmo, o valor encontrado foi de: **2.71101**

  $$
  \begin{align}
  Err = \frac{\left|3.0183-2.7111 \right|}{3.0183} \cdot 100 \\
  Err = 10.1779\%
  \end{align}
  $$

- Para n=4
  Ao executar o algoritmo, o valor encontrado foi de: **2.93784**
  $$
  \begin{align}
  Err = \frac{\left|3.0183-2.9378 \right|}{3.0183} \cdot 100 \\
  Err = 2.6671\%
  \end{align}
  $$

#### Letra d

Ao executar o algoritmo, o valor encontrado foi de: **2.96023**

$$
\begin{align}
Err = \frac{\left|3.0183-2.9603 \right|}{3.0183} \cdot 100 \\
Err = 1.9216\%
\end{align}
$$

#### Letra e

Ao executar o algoritmo, o valor encontrado foi de: **1.99707**

$$
\begin{align}
Err = \frac{\left|3.0183-1.99707 \right|}{3.0183} \cdot 100 \\
Err = 33.8346\%
\end{align}
$$

#### Letra f

Ao executar o algoritmo, o valor encontrado foi de: **2.99122**

$$
\begin{align}
Err = \frac{\left|3.0183-2.9912 \right|}{3.0183} \cdot 100 \\
Err = 0.8979\%
\end{align}
$$

## Questão 02

    Calcule a seguinte integral:

$$
\int_{0}^{\pi/2} (8 + 4\cos{x})dx
$$

    a) Analiticamente.
    b) Por aplicação única da regra do trapézio.
    c) Por aplicações múltiplas da regra do trapézio, com n = 2 e 4.
    d) Por aplicação única da regra de Simpson de 1/3.
    e) Por aplicações múltiplas da regra de Simpson de 1/3, com n = 4.
    f) Por aplicação única da regra de Simpson de 3/8.
    g) Por aplicações múltiplas da regra de Simpson de 3/8, com n = 5.

    Para cada estimativa numérica de b) a g), determine o erro relativo percentual com base em a).

### Analiticamente

$$
\begin{align}
8\int_{0}^{\pi/2} dx + 4\int_{0}^{\pi/2}\cos{x}dx \\
\\
\left[8x\right]_{0}^{\pi/2} + \left[4\sin{x} \right]_{0}^{\pi/2} \\
8\frac{\pi}{2} - 0 + 4\sin{\frac{\pi}{2}} - 4\sin{0} \\
4\pi + 4 = 16.5664\\
\\
\\
\therefore \int_{0}^{\pi/2} (8 + 4\cos{x})dx = 16.5664
\end{align}
$$

### Resultados

#### Letra b

Ao executar o algoritmo, o valor encontrado foi de: **15.70796**

$$
\begin{align}
Err = \frac{\left|16.5664-15.70796 \right|}{16.5664} \cdot 100 \\
Err = 5.1818\%
\end{align}
$$

#### Letra c

- Para n=2
  Ao executar o algoritmo, o valor encontrado foi de: **16.35861**

  $$
  \begin{align}
  Err = \frac{\left|16.5664-16.35861 \right|}{16.5664} \cdot 100 \\
  Err = 1.2543\%
  \end{align}
  $$

- Para n=4
  Ao executar o algoritmo, o valor encontrado foi de: **16.51483**
  $$
  \begin{align}
  Err = \frac{\left|16.5664-16.51483 \right|}{16.5664} \cdot 100 \\
  Err = 0.3113\%
  \end{align}
  $$

#### Letra d

Ao executar o algoritmo, o valor encontrado foi de: **16.57549**

$$
\begin{align}
Err = \frac{\left|16.5664-16.57549 \right|}{16.5664} \cdot 100 \\
Err = 0.0549\%
\end{align}
$$

#### Letra e

Ao executar o algoritmo, o valor encontrado foi de: **9.69583**

$$
\begin{align}
Err = \frac{\left|16.5664-9.69583 \right|}{16.5664} \cdot 100 \\
Err = 41.4729\%
\end{align}
$$

#### Letra f

Ao executar o algoritmo, o valor encontrado foi de: **16.5703**

$$
\begin{align}
Err = \frac{\left|16.5664-16.5703 \right|}{16.5664} \cdot 100 \\
Err = 0.0235\%
\end{align}
$$

## Questão 03

    Calcule a seguinte integral:

$$
\int_{-2}^{4} (1 -x - 4x^3 + 2x^5)dx
$$

    a) Analiticamente.
    b) Por aplicação única da regra do trapézio.
    c) Por aplicações múltiplas da regra do trapézio, com n = 2 e 4.
    d) Por aplicação única da regra de Simpson de 1/3.
    e) Por aplicações múltiplas da regra de Simpson de 1/3, com n = 4.
    f) Por aplicação única da regra de Simpson de 3/8.
    g) Por aplicações múltiplas da regra de Simpson de 3/8, com n = 5.

    Para cada estimativa numérica de b) a g), determine o erro relativo percentual com base em a).

### Analiticamente

$$
\begin{align}
\int_{-2}^{4}dx - \int_{-2}^{4}xdx - 4\int_{-2}^{4} x^3dx + 2\int_{-2}^{4} x^5dx \\
\\
\left[x - \frac{x^2}{2} - \frac{4x^4}{4} + \frac{2x^6}{6}\right]_{-2}^{4} \\
\left(4 - \frac{4^2}{2} - 4^4 + \frac{4^6}{3}\right) - \left( -2 -\frac{(-2)^2}{2} - (-2)^4 + \frac{(-2)^6}{3}\right) \\
1105.33333333 - 1.33333333333 = 1104 \\
\\
\\
\therefore \int_{-2}^{4} (1 -x - 4x^3 + 2x^5)dx = 1104
\end{align}
$$

### Resultados

#### Letra b

Ao executar o algoritmo, o valor encontrado foi de: **5280.00000**

$$
\begin{align}
Err = \frac{\left|1104-5280 \right|}{1104} \cdot 100 \\
Err = 378.2609\%
\end{align}
$$

#### Letra c

- Para n=2
  Ao executar o algoritmo, o valor encontrado foi de: **2634.00000**

  $$
  \begin{align}
  Err = \frac{\left|1104-2634 \right|}{1104} \cdot 100 \\
  Err = 138.5869\%
  \end{align}
  $$

- Para n=4
  Ao executar o algoritmo, o valor encontrado foi de: **1516.87500**
  $$
  \begin{align}
  Err = \frac{\left|1104-1516.875 \right|}{1104} \cdot 100 \\
  Err = 37.3981\%
  \end{align}
  $$

#### Letra d

Ao executar o algoritmo, o valor encontrado foi de: **1752.00000**

$$
\begin{align}
Err = \frac{\left|1104-1752 \right|}{1104} \cdot 100 \\
Err = 58.6957\%
\end{align}
$$

#### Letra e

Ao executar o algoritmo, o valor encontrado foi de: **128.28125**

$$
\begin{align}
Err = \frac{\left|1104-128.28125 \right|}{1104} \cdot 100 \\
Err = 88.3803\%
\end{align}
$$

#### Letra f

Ao executar o algoritmo, o valor encontrado foi de: **1392.00000**

$$
\begin{align}
Err = \frac{\left|1104-1392 \right|}{1104} \cdot 100 \\
Err = 26.0869\%
\end{align}
$$

## Questão 04

    A função:

$$
f(x) = e^{-x}
$$

    Pode ser usada para gerar a seguinte tabela de dados desigualmente espaçados:

|  **x**   |  0  |  0.1   |  0.3   |  0.5   |  0.7   |  0.95  |  1.2   |
| :------: | :-: | :----: | :----: | :----: | :----: | :----: | :----: |
| **f(x)** |  1  | 0.9048 | 0.7408 | 0.6065 | 0.4966 | 0.3867 | 0.3012 |

    Calcule a integral de a = 0 até b = 1.2:
    	a) Analiticamente.
    	b) Por aplicação única da regra do trapézio.
    	c) Combinação das regras do trapézio e de Simpson sempre que possível para obter uma melhor precisão.
    Para cada estimativa numérica de b) a g), calcule o erro relativo percentual com base em a).

### Resultados

#### Analiticamente

$$
\begin{align}
\int_{0}^{1.2} e^{-x} \\
\\
\left[-e^{-x}\right]_{0}^{1.2} \\
-e^{-1.2} - (-e^0) \\
-e^{-\frac{6}{5}} + 1 \\
\\
\\
\therefore \int_{0}^{1.2} e^{-x} = 0.6988
\end{align}
$$

#### Letra b

Ao executar o algoritmo com regra do trapézio simples, foi obtido o resultado `0.7807165271473213`.

#### Letra c

Para realizar a obtenção das variações dos resultados, foi executado o algoritmo pela regra dos trapézios simples, e com n igual a 2 ou 4. Ao mesmo tempo, essa mesma variação dos limites de integração foi aplicado ao método de Simpson, onde no mesmo foi usado a aplicação única 1/3 e 3/8, bem como a aplicação multipla 1/3 com n = 4.

##### Regra do Trapézio

###### Simples

- Intervalo: 0.00 ~ 0.10
  - Raiz: 0.09524
- Intervalo: 0.10 ~ 0.30
  - Raiz: 0.16457
- Intervalo: 0.30 ~ 0.50
  - Raiz: 0.13473
- Intervalo: 0.50 ~ 0.70
  - Raiz: 0.11031
- Intervalo: 0.70 ~ 0.95
  - Raiz: 0.11042
- Intervalo: 0.95 ~ 1.20
  - Raiz: 0.08599

###### Composta (n=2)

- Intervalo: 0.00 ~ 0.10
  - Raiz: 0.09518
- Intervalo: 0.10 ~ 0.30
  - Raiz: 0.16416
- Intervalo: 0.30 ~ 0.50
  - Raiz: 0.13440
- Intervalo: 0.50 ~ 0.70
  - Raiz: 0.11004
- Intervalo: 0.70 ~ 0.95
  - Raiz: 0.10999
- Intervalo: 0.95 ~ 1.20
  - Raiz: 0.08566

###### Composta (n=4)

- Intervalo: 0.00 ~ 0.10
  - Raiz: 0.09517
- Intervalo: 0.10 ~ 0.30
  - Raiz: 0.16405
- Intervalo: 0.30 ~ 0.50
  - Raiz: 0.13432
- Intervalo: 0.50 ~ 0.70
  - Raiz: 0.10997
- Intervalo: 0.70 ~ 0.95
  - Raiz: 0.10988
- Intervalo: 0.95 ~ 1.20
  - Raiz: 0.08557

##### Regra de Simpson

###### Aplicação única 1/3

- Intervalo: 0.00 ~ 0.10
  - Raiz: 0.09516
- Intervalo: 0.10 ~ 0.30
  - Raiz: 0.16402
- Intervalo: 0.30 ~ 0.50
  - Raiz: 0.13429
- Intervalo: 0.50 ~ 0.70
  - Raiz: 0.10995
- Intervalo: 0.70 ~ 0.95
  - Raiz: 0.10984
- Intervalo: 0.95 ~ 1.20
  - Raiz: 0.08555

###### Composta 1/3 (n=4)

- Intervalo: 0.00 ~ 0.10
  - Raiz: 0.05530
- Intervalo: 0.10 ~ 0.30
  - Raiz: 0.09489
- Intervalo: 0.30 ~ 0.50
  - Raiz: 0.07769
- Intervalo: 0.50 ~ 0.70
  - Raiz: 0.06360
- Intervalo: 0.70 ~ 0.95
  - Raiz: 0.06339
- Intervalo: 0.95 ~ 1.20
  - Raiz: 0.04937

###### Única 3/8

- Intervalo: 0.00 ~ 0.10
  - Raiz: 0.09516
- Intervalo: 0.10 ~ 0.30
  - Raiz: 0.16402
- Intervalo: 0.30 ~ 0.50
  - Raiz: 0.13429
- Intervalo: 0.50 ~ 0.70
  - Raiz: 0.10995
- Intervalo: 0.70 ~ 0.95
  - Raiz: 0.10984
- Intervalo: 0.95 ~ 1.20
  - Raiz: 0.08555

## Questão 05

    Determine a distância percorrida para os seguintes dados de velocidade:

| **t (min)** |  1  |  2  | 3.25 | 4.5 |  6  |  7  |  8  | 8.5 |  9  | 10  |
| :---------: | :-: | :-: | :--: | :-: | :-: | :-: | :-: | :-: | :-: | :-: |
| **v (m/s)** |  5  |  6  | 5.5  |  7  | 8.5 |  8  |  6  |  7  |  7  |  5  |

    a) Use a regra do trapézio, e também determine a velocidade média
    b) Ajuste os dados com uma equação cúbica utilizando regressão polinomial. Integre o polinômio para determinar a distância.

Utilizou-se como equação base a seguinte

$$
\begin{align}
V_m = S / \Delta t \\
\\
&V_m = \text{Velocidade média} \\
&S = \text{Distância percorrida (m)} \\
&\Delta t = \text{Variação de tempo (s)}
\end{align}
$$

#### Letra a

Ao executar os cálculos, necessitou-se adaptar os valores da variação de tempo para que os mesmos fossem dados em segundos. Assim sendo, obteve-se os seguintes valores:

$$
\Delta t = [60;120;195;270;360;420;480;510;540;600]
$$

Por fim, executou-se o cálculo usando a regra do trapézio, e obtiveram-se os valores:

- **Distância total percorrida:** 3607.5m
- **Velocidade média:** 6.5m/s

#### Letra b

Ao executar a regressão polinomial, obteve-se o seguinte polinômio

$$
\begin{align}
4.8507 + 1.00481\times10^{-3}x + 4.8699\times 10^{-5}x^2 -8.3338\times 10^{-8}x^3
\end{align}
$$

Ao executar o algoritmo para integração trapezoidal simples nesse polinomio, foi obtido o valor da distância como `2714.1474378866305` para o intervalo de 60 a 600 (os valores de tempo em segundo).

#### Chamada ao algoritmo

```go
func quest05IntegralTrapezoidalAndPolynomials() {
	t := []float64{1, 2, 3.25, 4.5, 6, 7, 8, 8.5, 9, 10}
	v := []float64{5, 6, 5.5, 7, 8.5, 8, 6, 7, 7, 5}

	// Convert time to seconds instead of minutes
	for index, inMinutes := range t {
		t[index] = inMinutes * 60
	}

	// Calculating distance traveled using the trapezoidal rule
	n := len(t) - 1 // Number of subintervals
	distance := 0.0
	for index := 0; index < n; index++ {
		h := t[index+1] - t[index]                    // Difference between time values
		distance += (h / 2) * (v[index] + v[index+1]) // Applying the trapezoidal rule
	}

	// Average speed calculation
	sum := 0.0
	for _, velocity := range v {
		sum += velocity
	}
	averageVelocity := sum / float64(len(v))

	views.ReportSimple("Distance", distance)
	views.ReportSimple("Average Velocity", averageVelocity)

	var method struct {
		polynomial          methods.MinimumSquarePolynomialRegressionMethod[float64]
		integralTrapezoidal methods.IntegralTrapezoidalMethod[float64]
	}

	result, err := method.polynomial.Run(t, v, float64(3))
	if err != nil {
		views.ReportError(err)
		return
	}
	views.ReportResult(
		"PolynomialRegression - "+strconv.FormatInt(int64(3), 10)+" degree",
		nil, uint32(1), result...,
	)

	fX := func(x float64) (funcResult float64) {
		for index, value := range result {
			var xValue float64 = 1
			if index > 0 {
				xValue = math.Pow(x, float64(index))
			}
			funcResult += value * xValue
		}

		return
	}

	/*fX = func(x float64) (funcResult float64) {
		return 4.850657832943775 + (0.0010048084709803518 * x) +
			(0.00004869867922122433 * math.Pow(x, 2)) + (-8.333763969120286e-8 * math.Pow(x, 3))
	}*/

	var integrationResult float64
	integrationResult, err = method.integralTrapezoidal.Run(60, 600, fX)
	if err != nil {
		views.ReportError(err)
	}

	views.ReportSimple("Question 05 - Simple Trapeizodal 1 ~ 10", integrationResult)
}
```

## Questão 06

    Calcule aproximações por diferenças progressivas e regressivas de O(h) e O(h^2), e aproximações por diferenças centradas de O(h^2) e O(h^4) para a primeira derivada de y = cos x em x = \pi/4 usando um valor de h = \pi/12 . Faça uma estimativa do erro relativo percentual verdadeiro \epsilon_t, para cada aproximação.

## Questão 07

    Use aproximação por diferenças centradas para obter estimativas para a primeira e segunda derivadas de y = e^x em x = 2 para h = 0.1. Empregue formulas de ordem de O(h^2) e O(h^4).
