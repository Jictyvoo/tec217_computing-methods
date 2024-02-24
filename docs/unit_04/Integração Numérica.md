> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

### Algoritmo

#### Trapézio

```go
type integrationTrapezoidalKind uint8

const (
	IntegrationTrapezoidalSimple integrationTrapezoidalKind = iota
	IntegrationTrapezoidalCompound
)

type (
	funcCallback[T models.Numeric]        func(T) T
	IntegralTrapezoidal[T models.Numeric] struct {
		commonFuncZeroState[T]
		Kind              integrationTrapezoidalKind
		NumberSubdivision uint64
	}
)

func (mtd *IntegralTrapezoidal[T]) simpleCalculation(
	a, b T,
	f funcCallback[T],
) (result T, _ error) {
	h := b - a
	result = h * (f(a) + f(b)) / 2
	return
}

func (mtd *IntegralTrapezoidal[T]) compoundCalculation(
	a, b T,
	f funcCallback[T],
) (result T, _ error) {
	h := (b - a) / T(mtd.NumberSubdivision)
	previousX := a
	sum := f(previousX)
	for i := uint64(1); i < mtd.NumberSubdivision; i++ {
		previousX = previousX + h
		sum += 2 * f(previousX)
	}

	x := previousX + h
	sum += f(x)
	result = h * (sum / 2)
	return
}

func (mtd *IntegralTrapezoidal[T]) Run(a, b T, f func(T) T) (result T, err error) {
	switch mtd.Kind {
	case IntegrationTrapezoidalCompound:
		return mtd.compoundCalculation(a, b, f)
	case IntegrationTrapezoidalSimple:
		return mtd.simpleCalculation(a, b, f)
	}

	return
}
```

#### Simpson

```go
type integrationSimpsonKind uint8

const (
	IntegrationSimpson13 integrationSimpsonKind = iota
	IntegrationSimpson38
	IntegrationSimpsonCompound13
)

type (
	IntegralSimpsonMethod[T models.Numeric] struct {
		Kind              integrationSimpsonKind
		NumberSubdivision uint64
		commonFuncZeroState[T]
	}
)

func (mtd *IntegralSimpsonMethod[T]) simple13(a, b T, f funcCallback[T]) (result T, err error) {
	var (
		h  = (b - a) / 2
		f1 = f(a)
		f2 = f((a + b) / 2)
		f3 = f(b)
	)

	result = h * (f1 + 4*f2 + f3) / 3
	return
}

func (mtd *IntegralSimpsonMethod[T]) simple38(a, b T, f funcCallback[T]) (result T, err error) {
	var (
		h  = (b - a) / 3
		f1 = f(a)
		f2 = f((2*a + b) / 3)
		f3 = f((a + 2*b) / 3)
		f4 = f(b)
	)
	result = 3 * h * (f1 + 3*(f2+f3) + f4) / 8
	return
}

func (mtd *IntegralSimpsonMethod[T]) compound13(a, b T, f funcCallback[T]) (result T, err error) {
	if mtd.NumberSubdivision%2 != 0 || mtd.NumberSubdivision == 0 {
		err = errors.New("the given n should be a even number")
		return
	}
	h := (b - a) / T(mtd.NumberSubdivision)
	var sum T
	for i := uint64(1); i < mtd.NumberSubdivision; i++ {
		xi := a + T(i)*h
		switch {
		case i == 1 || i == mtd.NumberSubdivision+1:
			sum += f(xi)
		case i%2 == 0:
			sum += 4 * f(xi)
		default:
			sum += 2 * f(xi)
		}

	}
	result = h * sum / 3
	return
}

func (mtd *IntegralSimpsonMethod[T]) Run(a, b T, f func(T) T) (result T, err error) {
	switch mtd.Kind {
	case IntegrationSimpson13:
		return mtd.simple13(a, b, f)
	case IntegrationSimpson38:
		return mtd.simple38(a, b, f)
	case IntegrationSimpsonCompound13:
		return mtd.compound13(a, b, f)
	}

	return
}
```

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

    Para cada estimativa numérica, determine o erro relativo percentual com base em a). Importante: Fazer manualmente o passo a passo de cada método. Implemente os algoritmos e obtenha a solução numérica para a integral.

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

#### Letra b

$$
\begin{align}
f(x) = 1 - e^{-x} \\
\\
\int_{0}^{4} (1 - e^{-x})dx = \frac{4}{2} (f(0) + f(4)) \\
&= 2(0 + 0.9817) \\
&= 1.9634 \\
\\
\\
h = 4 - 0 \therefore h = 4 \\
Err = \frac{\left|3.0183-1.9634 \right|}{3.0183} \cdot 100 \\
Err = 34.9501\%
\end{align}
$$

#### Letra c

- Para n=2

  $$
  \begin{align}
  f(x) = 1 - e^{-x} \\
  \\
  \int_{0}^{4} (1 - e^{-x})dx = \frac{2}{2} \left(f(0) + 2f(2) + f(4) \right) \\
  &= 0 - 2 \cdot 0.8647 + 0.9817 \\
  &= 2.7111 \\
  \\
  \\
  h = (4 - 0) / 2 \therefore h = 2 \\
  Err = \frac{\left|3.0183-2.7111 \right|}{3.0183} \cdot 100 \\
  Err = 10.1779\%
  \end{align}
  $$

- Para n=4
  $$
  \begin{align}
  f(x) = 1 - e^{-x} \\
  \\
  \int_{0}^{4} (1 - e^{-x})dx = \frac{2}{4} \left(f(0) + 2f(1) + f(3) + f(4) \right) \\
  &= \frac{1}{2} \left( 0 + 2 (0.6321 + 0.8647 - 0.9502) + 0.9817 \right) \\
  &= 2.9378 \\
  \\
  \\
  h = (4 - 0) / 2 \therefore h = 2 \\
  Err = \frac{\left|3.0183-2.9378 \right|}{3.0183} \cdot 100 \\
  Err = 2.6671\%
  \end{align}
  $$

#### Letra d

$$
\begin{align}
h = (4 - 0) / 2 \therefore h = 2 \\
f(0) = 0; f(2) = 0.8647; f(4) = 0.9817 \\
\\
\int_{0}^{4} (1 - e^{-x})dx \\
&= \frac{2}{3}(0+4 \cdot 0.8647 + 0.9817) \\
&= 2.9603 \\
\\
\\
Err = \frac{\left|3.0183-2.9603 \right|}{3.0183} \cdot 100 \\
Err = 1.9216\%
\end{align}
$$

#### Letra e

$$
\begin{align}
h = (4 - 0) / 4 \therefore h = 1 \\
f(0) = 0; f(1) = 0.6321; f(2) = 0.8647; f(3) = 0.9502; f(4) = 0.9817 \\
\\
\int_{0}^{4} (1 - e^{-x})dx \\
\\
= \frac{1}{3}(0 + 4 (0.6321 + 0.9502) +2(0.8647) + 0.9817) \\
= 3.0133 \\
\\
\\
Err = \frac{\left|3.0183-3.0133 \right|}{3.0183} \cdot 100 \\
Err = 0.1657\%
\end{align}
$$

#### Letra f

$$
\begin{align}
h = (4 - 0) / 3 \therefore h = 1.3333 \\
f(0) = 0; f(1.3333) = 0.7364; f(2.6666) = 0.9305; f(4) = 0.9817 \\
\\
\int_{0}^{4} (1 - e^{-x})dx \\
\\
= \frac{4}{3}(0 + 3 (0.7364 + 0.9305) + 0.9817) \\
= 2.9912 \\
\\
\\
Err = \frac{\left|3.0183-2.9912 \right|}{3.0183} \cdot 100 \\
Err = 0.8979\%
\end{align}
$$

### Chamada ao algoritmo

#### Trapézio

```go
func quest01IntegralTrapezoidal() {
	var (
		buffer strings.Builder
		method = struct {
			simple, compound2, compound4 methods.IntegralTrapezoidal[float64]
		}{
			compound2: methods.IntegralTrapezoidal[float64]{
				Kind:              methods.IntegrationTrapezoidalCompound,
				NumberSubdivision: 2,
			},
			compound4: methods.IntegralTrapezoidal[float64]{
				Kind:              methods.IntegrationTrapezoidalCompound,
				NumberSubdivision: 4,
			},
		}
		fX = func(x float64) float64 {
			return 1 - math.Pow(math.E, -x)
		}
		results [3]struct {
			value float64
			err   error
		}
	)

	results[0].value, results[0].err = method.simple.Run(0, 4, fX)
	results[1].value, results[1].err = method.compound2.Run(0, 4, fX)
	results[2].value, results[2].err = method.compound4.Run(0, 4, fX)

	if err := errors.Join(results[0].err, results[1].err, results[2].err); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult("TrapezoidalIntegration", &buffer, uint32(0), results[0].value)
	views.ReportResult("TrapezoidalIntegration", &buffer, uint32(2), results[1].value)
	views.ReportResult("TrapezoidalIntegration", &buffer, uint32(4), results[2].value)
}
```

#### Simpson

```go
func quest01IntegralSimpson() {
	var (
		buffer strings.Builder
		method = struct {
			simple13, simple38, compound13 methods.IntegralSimpsonMethod[float64]
		}{
			simple13: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson13,
			},
			compound13: methods.IntegralSimpsonMethod[float64]{
				Kind:              methods.IntegrationSimpsonCompound13,
				NumberSubdivision: 4,
			},
			simple38: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson38,
			},
		}
		fX = func(x float64) float64 {
			return 1 - math.Pow(math.E, -x)
		}
		results [3]struct {
			value float64
			err   error
		}
	)

	results[0].value, results[0].err = method.simple13.Run(0, 4, fX)
	results[1].value, results[1].err = method.compound13.Run(0, 4, fX)
	results[2].value, results[2].err = method.simple38.Run(0, 4, fX)

	if err := errors.Join(results[0].err, results[1].err, results[2].err); err != nil {
		views.ReportError(err)
		return
	}

	views.ReportResult("SimpsonIntegration", &buffer, uint32(13), results[0].value)
	views.ReportResult("SimpsonIntegration", &buffer, uint32(113), results[1].value)
	views.ReportResult("SimpsonIntegration", &buffer, uint32(38), results[2].value)
}
```

### Resultados

Ao rodar o algoritmo de forma computacional, os valores obtidos foram

- Trapézio
  - **Simples:** 1.96337
  - **Composto**
    - **n = 2:** 2.71101
    - **n = 4:** 2.93784
- Simpson
  - **Simples 1/3:** 2.96023
  - **Composto 1/3**
    - **n = 4:** 2.99122
  - **Simples 3/8:** 1.99707

## Questão 02 - Problema aplicado: Análise de fluxo em dutos de óleo.

    A análise de um fluxo de líquido em um tubo circular aplica-se a diferentes sistemas, incluindo veias e artérias em um corpo, sistema de água de uma cidade, sistema de irrigação para uma fazenda, o sistema de tubulação que transporta fluidos em uma fábrica, as linhas hidráulicas de uma aeronave, e o jato de tinta da impressora de um computador.
    O atrito em uma tubulação circular faz com que um perfil de velocidade se desenvolva no óleo que flui. O óleo que está em contato com as paredes do tubo não está se movendo, enquanto o óleo no centro do fluxo está se movendo mais rápido. O diagrama da figura 1 mostra como a velocidade do óleo varia ao longo do diâmetro do tubo e define as variáveis utilizadas nesta análise. A equação a seguir descreve esse perfil de velocidade:

![[question02_figure.png]]

$$
v(r) = v_{max} \left( 1 - \frac{r}{r_0} \right)^{\frac{1}{n}}
$$

    A variável n é um inteiro entre 5 e 10 e define a forma do fluxo direto do óleo. A velocidade média do fluxo é a área integral do perfil de velocidade, que pode se mostrar

$$
\begin{align}
v_{media} = \frac{\int_{0}^{r_0} v(r)2\pi rdr}{\pi r_{0}^2} \\
v_{media} = \frac{2v_{max}}{r_{0}^2} \int_{0}^{r_0} r\left(1-\frac{r}{r_0}\right)^{\frac{1}{n}} dr
\end{align}
$$

    Os valores de $v_{max}$ e n podem ser medidos experimentalmente, e o valor de $r_0$ é o raio da tubulação. Escreva programas para integrar o perfil de velocidade para determinar a velocidade média do fluxo do óleo no tubo. Considere $r_0 = 0.5m$ , $n = 8$, e $v_{max} = 1.5m/s$. **Apresentar apenas a saída numérica, gráficos, etc...**

### Resultados

#### Chamada ao algoritmo

```go
func quest02IntegralSimpson() {
	const (
		n    float64 = 8
		R0           = 0.5
		vMax         = 1.5
	)

	var (
		buffer strings.Builder
		method = struct {
			simple13, simple38 methods.IntegralSimpsonMethod[float64]
		}{
			simple13: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson13,
			},
			simple38: methods.IntegralSimpsonMethod[float64]{
				Kind: methods.IntegrationSimpson38,
			},
		}
		fX = func(r float64) float64 {
			return r * math.Pow(1-(r/R0), 1/n)
		}
		results [2]struct {
			value float64
			err   error
		}
	)

	results[0].value, results[0].err = method.simple13.Run(0, R0, fX)
	results[1].value, results[1].err = method.simple38.Run(0, R0, fX)

	if err := errors.Join(results[0].err, results[1].err); err != nil {
		views.ReportError(err)
		return
	}

	var vMean [2]float64
	for index, result := range results {
		vMean[index] = (2 * vMax * result.value) / math.Pow(R0, 2)
	}

	views.ReportResult("SimpsonIntegration", &buffer, uint32(13), results[0].value)
	views.ReportSimple("Mean Velocity 1/3", vMean[0])
	views.ReportResult("SimpsonIntegration", &buffer, uint32(38), results[1].value)
	views.ReportSimple("Mean Velocity 3/8", vMean[1])
}
```

#### Conclusão

A função a ser integrada utilizada foi:

$$
f(r) = r \cdot \left(1 - \frac{r}{R_0}\right)^\frac{1}{n}
$$

De modo a obter o valor da velocidade média, o resultado obtido da integração entre os pontos $0$ e $R_0$ foi multiplicado por $\frac{2 V_{max}}{R_{0}^2}$.

Assim sendo, foram obtidos os seguintes resultados:

##### Simples 1/3

- **Resultado Integral:** 0.07642
- **Velocidade média:** 0.9170040432046711

![[quest02_chart_simpson_13.png]]

##### Simples 3/8

- **Resultado Integral:** 0.08419
- **Velocidade média:** 1.0102315915116045

![[quest02_chart_simpson_38.png]]
