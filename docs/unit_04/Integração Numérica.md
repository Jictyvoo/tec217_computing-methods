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
    
    Para cada estimativa numérica, determine o erro relativo percentual com base em a). Importante: Fazer manualmente o passo a passo de cada método. Implemente os algoritmos e obtenha a solução numérica para a integral.

### Analiticamente


### Algoritmo

```octave

```

#### Chamada ao algoritmo

```octave

```

### Resultados



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
Os valores de $v_{max}$ e n podem ser medidos experimentalmente, e o valor de $r_0$ é o raio da tubulação.
Escreva programas para integrar o perfil de velocidade para determinar a velocidade média do fluxo do óleo no tubo. Considere $r_0 = 0.5m$ , $n = 8$, e $v_{max} = 1.5m/s$. **Apresentar apenas a saída numérica, gráficos, etc...**

