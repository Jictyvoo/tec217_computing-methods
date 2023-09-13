> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## 1. Considere o sistema de ponto flutuante normalizado SPF(2,4,-1,2) de base 2, 4 dígitos na mantissa, menor expoente -1, e maior expoente 2. Para este sistema:
	a) Qual é o menor positivo exatamente representável?
	b) Qual é o maior positivo exatamente representável?
	c) Quantos são os exatamente representáveis positivos?
	d) Qual é o número total de reais exatamente representáveis?
	e) Represente na reta todos os positivos exatamente representáveis.
	f) Defina a região de overflow e underflow.

SPF(2,4,-1,2) de base 2, 4 dígitos na mantissa, menor expoente -1 e maior expoente 2.

a) O menor positivo exatamente representável ocorre quando a mantissa é a menor possível (0001) e o expoente é o menor possível (-1). Portanto, o menor positivo exatamente representável é 

$$
\begin{align}
2^{-1} * 0.1000 \\
2^{-1} * 2^{-4} = 0.03125
\end{align}
$$

b) O maior positivo exatamente representável ocorre quando a mantissa é a maior possível (1111) e o expoente é o maior possível (2). Portanto, o maior positivo exatamente representável é
$$
\begin{align}
2^2 * 0.1111 \\
(2^{-1}+2^{-2}+2^{-3}+2^{-4}) * 2^2 \\
2^{1}+2^{0}+2^{-1}+2^{-2} = 3.75

\end{align}
$$

c) Para determinar quantos números positivos exatamente representáveis existem, basta contar as combinações únicas de mantissa e expoente possíveis. Há 16 combinações possíveis de mantissa (2^4) e 4 combinações possíveis de expoente (de -1 a 2), resultando em um total de 16 * 4 = 64 números positivos exatamente representáveis.

d) Para encontrar o número total de reais exatamente representáveis, precisamos considerar os números negativos e zero, além dos positivos. Como há 64 positivos exatamente representáveis, haverá o mesmo número de negativos exatamente representáveis. Portanto, o total de reais exatamente representáveis é 64 (positivos) + 64 (negativos) + 1 (zero) = 129.

e) Representar os positivos exatamente representáveis na reta:

```
   0.125         0.375         0.625         0.875          1.125         1.375         1.625         1.875
    |-------------|-------------|-------------|--------------|-------------|-------------|-------------|
```

f) A região de overflow ocorre quando um número excede o maior expoente possível. No caso deste sistema, o maior expoente é 2. Portanto, qualquer número com expoente maior do que 2 resultará em overflow.

A região de underflow ocorre quando um número é muito pequeno para ser representado com precisão no sistema. No caso deste sistema, o menor expoente é -1, então qualquer número com expoente menor do que -1 resultará em underflow.

## 2. Converta os seguintes números na base 2 para a base 10:
	a) 1011001
	b) 0,01011
	c) 110,01001

Vou converter os números da base 2 para a base 10 conforme solicitado:

a) 89
$$
\begin{align}
1011001_{2} = 1 \cdot 2^{6} + 0 \cdot 2^{5} + 1 \cdot 2^{4} + 1 \cdot 2^{3} + 0 \cdot 2^{2} + 0 \cdot 2^{1} + 1 \cdot 2^{0} \\
\\
1011001_{2} = 89
\end{align}
$$

b) 0.34375

$$
\begin{align}
(0.01011)_{2} = (0 \cdot 2^{0}) + (0 \cdot 2^{-1}) + (1 \cdot 2^{-2}) + (0 \cdot 2^{-3}) + (1 \cdot 2^{-4}) + (1 \cdot 2^{-5}) \\
\\
(0.01011)₂ = (0.34375)₁₀
\end{align}
$$

c) 6.28125

$$
\begin{align}
(110.01001)_{2} = (1 \cdot 2^{2}) + (1 \cdot 2^{1}) + (0 \cdot 2^{0}) + (0 \cdot 2^{-1}) + (1 \cdot 2^{-2}) + (0 \cdot 2^{-3}) + (0 \cdot 2^{-4}) + (1 \cdot 2^{-5}) \\
\\
(110.01001)_{2} = (6.28125)_{10}
\end{align}
$$

## 3. Calcular (0,1011 x 2^-1) - (0,1010 x 2^-1) em SPF(2, 4, -3, 3).

$$
\begin{align}
0.1011 \cdot 2^{-1} = (2^{-1}+2^{-3}+2^{-4}) \cdot2^{-1} \\
0.1010 \cdot 2^{-1} = (2^{-1}+2^{-3}) \cdot2^{-1}\\
\\
\text{agora subtrai um numero pelo outro} \\
(0.1011 - 0.1010)\cdot 2^{-1} = 0.0001\cdot 2^{-1} \\
0.0001\cdot 2^{-1} = \frac{1}{32} = 0.03125 \\

\end{align}
$$
O resultado dessa subtração dará *underflow*.

## 4. A derivada de f(x) = 1/(1-3x^2)^2 é dada por 6x/(1-3x^2)^2. Espera-se dificuldades calculando esta função em x = 0,577? Tente usando aritmética com 3 e 4 algarismos significativos com truncamento.

$$
\begin{align}
\frac{6x}{(1-3x^2)^2} = 2352910.7926,
\text{para x = 0.577}
\end{align}
$$

## 5. Uma máquina de calcular armazena 4 dígitos na mantissa utilizando arredondamento. Para os valores (x=17534) e (y=21178) calcule:
	a) o armazenamento dos valores no processo de utilização da máquina digital para cada número x e y?
	b) o erro absoluto e relativos das variáveis x e y.
	c) o erro relativo ao realizar as operações x+y e x×y.

a) Lembrando que a máquina armazena 4 dígitos na mantissa utilizando arredondamento, o que significa que apenas os 4 dígitos mais significativos serão armazenados.

Para x=17534x=17534: A representação de xx utilizando arredondamento é 17501750 (4 dígitos mais significativos).

Para y=21178y=21178: A representação de yy utilizando arredondamento é 21182118 (4 dígitos mais significativos).

b) O erro absoluto E é dado pela diferença entre o valor real e o valor aproximado:

$$E = \text{Valor real} - \text{Valor aproximado}$$

 $$
\begin{align}
\text{O erro relativo } e \text{ é dado pelo quociente entre o erro absoluto e o valor real:} \\
e = \frac{E}{\text{Valor real}}
\end{align}
$$
$$
\begin{align}
\text{Para } x = 17534: \\
\text{Erro absoluto de }x = 17534 - 1750 = 15784 \\
\text{Erro relativo de }x = \frac{15784}{17534} \\

\text{Para } y = 21178: \\
\text{Erro absoluto de } y = 21178 - 2118 = 19060 \\
\text{Erro relativo de } y = \frac{19060}{21178} \\
\end{align}
$$

c) Erro relativo ao realizar as operações x+y e x * y:

Ao realizar operações com números aproximados, os erros se propagam. O erro relativo em uma operação é dado pela soma dos erros relativos dos operandos.

Para a operação \(x + y\):
Erro relativo = Erro relativo de \(x\) + Erro relativo de \(y\)

Para a operação \(x \times y\):
Erro relativo = Erro relativo de \(x\) + Erro relativo de \(y\)

Lembrando que estamos trabalhando com aproximações, então esses erros são estimativas das incertezas associadas aos cálculos na máquina de calcular.
