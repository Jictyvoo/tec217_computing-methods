> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01

    Implementar e testar os seguintes métodos: (Importante: Realizar manualmente 2 iterações de cada método !)
    - RK 2a ordem Ralston;
    - RK 3a ordem clássico;
    - RK 4a ordem clássico;

    Comparação:
    	Sabendo que a expressão analítica para a corrente é: i(t) = -(1/2)eˆ(-400t) + 1/2.
    	Trace num mesmo gráfico: a solução analítica e as diversas soluções numéricas para EDOs implementadas e comente !

## Resolução

### Questão 01

Para resolver a equação diferencial $\frac{di}{dt} = 200 - 400i_{(t)}$, consideramos $f(i,t) = 200 - 400i$.

Condições iniciais: $i(0) = 0$

Intervalo de integração: $t \in [0, 0.02]$ com passo $h = 0.0001$.

A expressão analítica para a corrente é: $i(t) = -\frac{1}{2}e^{-400t} + \frac{1}{2}$.

#### Método de Runge-Kutta de 2ª ordem Ralston

1ª iteração:

$$
\begin{align}
&k_1 = f(i_n, t_n) = 0.0001 \cdot (200 - (400 \cdot 0)) = 0.02 \\
&k_2 = f(0 + \frac{3}{4} \cdot 200h, 0 + h) = 200 - 400 \cdot 0.015 = 194 \\
&y(h) = 0 + 0.0001 \cdot \left(\frac{200}{3} + \frac{2 \cdot 194}{3} \right) \\
&y(0.0001) = 0.0196
\end{align}
$$

2ª iteração:

$$
\begin{align}
&k_1 = f(0.0196, 0.0001) = 200 - 400 \cdot 0.0196 = 192.16 \\
&k_2 = f\left(0.0196 + \frac{3 \cdot 192.16 \times 10^{-4}}{4}, 2\times 10^{-4}\right) \\
&y(h+h) = 0.0196 + 10^{-4} \left(\frac{192.16}{3} + \frac{2\cdot 126.3952}{3}\right) \\
&y(2\times 10^{-4}) = 0.0384
\end{align}
$$

---

#### Método de Runge-Kutta de 3ª ordem clássico

1ª iteração:

$$
\begin{align}
&k_1 = f(0,0) = 200 \\
&k_2 = f(0+ \frac{200\times 10^{-4}}{2}, 0+\frac{10^{-4}}{2}) = 200 - 400 \cdot 0.01 = 196 \\
&k_3 = f(0 - 200 \times 10^{-4}+2.196\times 10^{-4}, 0+10^{-4}) = 200 - 400 \cdot 0.192 = 192.32 \\
\\
&y(h) = 0 + \frac{10^{-4}}{6} (200 + 4.196 + 192.32) \\
&y(0.0001) = 0.0196
\end{align}
$$

2ª iteração:

$$
\begin{align}
&k_1 = f(0.0196, 10^{-4}) = 200 - 400 \cdot 0.0196 = 192.16 \\
&k_2 = f(0.0196 - \frac{10^{-4}}{1} \cdot 192.16, 10^{-4} + \frac{10^{-4}}{2}) = 200 -400 \cdot 0.029208 = 188.3168 \\
&k_3 = f(0.0196 - 192.16 \times 10^{-4}+2\cdot 188.3168\times 10^{-4}, 2\times 10^{-4}) = 200 - 400\cdot 0.0380473 = 184.7811 \\
\\
\\
&y(2h) = 0.0196 + \frac{10^{-4}}{2} (196.16 + \cdot 188.3168 + 184.7211) \\
&y(0.0002) = 0.0385
\end{align}
$$

---

#### Método de Runge-Kutta de 4ª ordem clássico

1ª iteração:

$$
\begin{align}
&k_1 = f(0,0) = 200 - 400 \cdot 0 = 200 \\
&k_2 = f(0+ \frac{10^{-4}}{2}\cdot 200, 0 + 0.5\times 10^{-4}) = 200 - 400 \cdot 0.01 = 196 \\
&k_3 = f(0 + \frac{10^{-4}}{2} \cdot 196, 0 + 0.5 \times 10^{-4}) = 200 - 400 \cdot 0.0098 = 196.08 \\
&k_4 = f(0 + 196.08\times 10^{-4}, 0 + 10^{-4}) = 192.1568 \\
\\
\\
&y(10^{-4}) = 0 + \frac{10^{-4}}{6} (200 + 2\cdot 196 + 2\cdot 196.084 + 192.1568) \\
&y(10^{-4}) = 0.019605
\end{align}
$$

2ª iteração:

$$
\begin{align}
&k_1 = f(0.019605, 10^{-4}) = 192.158 \\
&k_2 = f(0.019605 + \frac{192.158 \times 10^{-4}}{2}, 10^{-4} + 0.5\times 10^{-4}) = 188.3148 \\
&k_3 = f(0.019605 + \frac{188.3148 \times 10^{-4}}{2}, 10^{-4}+0.5\times 10^{-4}) = 188.3917 \\
&k_4 f(0.019605 + 188.3917\times 10^{-4}, 10^{-4} + 0.5\times 10^{-4}) = 184.6223\\
\\
\\
&y(2h) = 0.019605 + \frac{10^{-4}}{6}(192.158 + 2\cdot 188.4148 + 2 \cdot 188.3917 + 184.6223) \\
&y(2\times 10^{-4}) = 0.0384
\end{align}
$$

---

Ao analisar os resultados, observamos que os métodos de Runge-Kutta de 3ª e 4ª ordem clássicos proporcionaram resultados muito próximos da solução analítica. Por outro lado, o método de 2ª ordem Ralston teve uma discrepância maior em relação à solução analítica. Isso sugere que os métodos de ordem superior são mais precisos para esta equação diferencial específica.

Essa mesma caracteristica pode ser observada de forma mais visual no gráfico abaixo:
![[quest_01_chart.png]]
