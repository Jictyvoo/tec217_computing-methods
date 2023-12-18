% Dados da tabela
xi = [0, 0.5, 1.0];
fi = [1.3, 2.5, 0.9];

% Funções dos polinômios de Lagrange
L0 = @(x) (x.^2 - 1.5*x + 0.5) / 0.5;
L1 = @(x) (x.^2 - x) / -0.25;
L2 = @(x) (x.^2 - 0.5*x) / 0.5;

% Função do polinômio P2(x)
P2 = @(x) -5.6*x.^2 + 5.2*x + 1.3;

% Intervalo de x para o gráfico
x_interval = linspace(-2, 2, 1000);

% Calcula os valores de y para cada polinômio de Lagrange
y0 = L0(x_interval);
y1 = L1(x_interval);
y2 = L2(x_interval);

% Calcula os valores de y para o polinômio P2(x)
yP2 = P2(x_interval);

% Calcula os valores de y para os pontos da tabela
fy = L0(xi) + L1(xi) + L2(xi);

% Traça os gráficos
figure;
plot(x_interval, y0, 'r', 'LineWidth', 2, 'DisplayName', 'L_0(x)');
hold on;
plot(x_interval, y1, 'g', 'LineWidth', 2, 'DisplayName', 'L_1(x)');
plot(x_interval, y2, 'b', 'LineWidth', 2, 'DisplayName', 'L_2(x)');
plot(x_interval, yP2, 'm', 'LineWidth', 2, 'DisplayName', 'P_2(x)');
scatter(xi, fi, 50, 'k', 'filled', 'DisplayName', 'Pontos da Tabela');
hold off;

% Ajusta os limites do eixo y para aproximar a tabela
ylim([min(fy) - 1.75, max(fy) + 1.75]);

% Adiciona rótulos e legenda ao gráfico
title('Curvas dos Polinômios de Lagrange e P_2(x) com Pontos da Tabela');
xlabel('x');
ylabel('y');
legend('location', 'northwest');
