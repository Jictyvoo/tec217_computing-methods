> **João Victor Oliveira Couto**
> Universidade Estadual de Feira de Santana (UEFS)
> Av. Transnordestina, S/N, Novo Horizonte
> 44036-900 - Feira de Santana - BA - Brasil

## Questão 01 - Problema predador-presa (uso do método de Euler explicito para resolver um sistema de duas EDOs de primeira ordem)

    A relação entre a população de leões (predadores), NL, e a população de gazelas (presas), NG, que residem em uma mesma área pode ser modelada por um sistema de duas EDOs. Suponha que uma comunidade seja formada por NL leões (predadores) e NG gazelas (presas), com b e d representando as taxas de natalidade e mortalidade das respectivas espécies. A taxa de variação (crescimento ou decrescimento) das populações de leões e de gazelas pode ser modelada pelas equações:

$$
\begin{align}
\frac{dN_L}{dt} = b_L N_L N_G - d_L N_L\\
&N_L(0) = 500 \\
\\
\frac{dN_G}{dt} = b_G N_G - d_G N_G N_L \\
&N_G(0) = 3000
\end{align}
$$

    Determine a população de leões e de gazelas em função do tempo, de t = 0 a t = 25 anos, se, em t = 0, N G = 3000 e N L = 500. Os coeficientes do modelo são: b G = 1,1 ano −1 , b L = 0,00025 ano −1 , d G = 0,0005 ano –1 e d L = 0,7 ano −1 .

### Resolução

$$
\begin{align}
\text{Taxa de variação da população de leões:} \\
&\frac{dN_L}{dt} = 0.00025 \cdot N_L \cdot N_G − 0.7 \cdot N_L \\
\\
\text{Taxa de variação da população de gazelas:} \\
&\frac{dN_G}{dt} = 1.1 \cdot N_G - 0.0005 \cdot N_G \cdot N_L
\end{align}
$$

Considerou-se que a população inicial de leões é o informado como 500, e o de gazelas como 3000.

#### Iterações manuais

Iteração 1:

$$
\begin{align}
N_L(0+0.01) = 500 + 0.01 (f_l(500)) \\
N_L(0.01) = 500 + 0.01 \cdot (0.00025 \cdot 500 \cdot 3000 -0.7\cdot 500) \\
N_L(0.01) = 500.25 \\
\\
\\
N_G(0+0.01) = 3000 + 0.01 \cdot f_g(3000) \\
N_G(0.01) = 3000 + 0.01\cdot (1.1 \cdot 3000 - 0.0005\cdot 500\cdot 3000) \\
N_G(0.01) = 3025.5
\end{align}
$$

Iteração 2:

$$
\begin{align}
N_L(0.02) = 500.25 + 0.01 (f_l(500.25)) \\
N_L(0.02) = 500.25 + 0.01 \cdot (0.00025 \cdot 500.25 \cdot 3000 -0.7\cdot 500.25) \\
N_L(0.02) = 500.532 \\
\\
\\
N_G(0.02) = 3025.5 + 0.01 \cdot f_g(3025.5) \\
N_G(0.02) = 3025.5 + 0.01\cdot (1.1 \cdot 3025.5 - 0.0005\cdot 500.25\cdot 3025.5) \\
N_G(0.02) = 3051.213
\end{align}
$$

#### Gráficos

##### Passo h=0.1

![[extra_step_0.01.png]]

##### Passo h=0.04

![[extra_step_0.04.png]]

##### Passo h=0.7

![[extra_step_0.07.png]]
