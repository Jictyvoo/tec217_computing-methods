addpath('./methods')

%%%%%%%%%%% Start code execution %%%%%%%%%%%
x = [0; 0.5; 1];
y = [1; 2.12; 3.55];

[foundY, err] = newtonInterpolationMethod(x, y, 0.7);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: NewtonInterpolationMethod\n');
fprintf('Result:'), disp(foundY);

fprintf('\n');

%%%%%%%%%%% Start code execution %%%%%%%%%%%
x = [0; 0.5; 1];
y = [1.3; 2.5; 0.9];

[foundY, err] = lagrangeInterpolationMethod(x, y, 0.8);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: lagrangeInterpolationMethod\n');
fprintf('Result:'), disp(foundY);

fprintf('\n');
%%%%%%%%%%% Start code execution %%%%%%%%%%%
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

fprintf('\n');

%%%%%%%%%%% Start code execution %%%%%%%%%%%
x = [0; 0.25; 0.5; 0.75; 1];
y = [1; 1.284; 1.6487; 2.1170; 2.7183];
%p = minimumSquarePolinomialRegressionMethod(x, y, 1);
%p = minimumSquarePolinomialRegressionMethod(x, y, 2);
%p = minimumSquarePolinomialRegressionMethod(x, y, 3);

f1 = @(x) (0.8997 * x) + 1.7078;
y1 = f1(x);

f2 = @(x) (1.0051 * (x.^2)) + (0.8642 * x) + 0.8437;
y2 = f2(x);

f3 = @(x) (0.9999 * (x.^3)) + (1.0141 * (x.^2)) + (0.4253 * x) + 0.2789;
y3 = f3(x);

x = [0.1; 0.4; 0.5; 0.7; 0.7; 0.9];
y = [0.61; 0.92; 0.99; 1.52; 1.47; 2.03];
p = minimumSquarePolynomialRegressionMethod(x, y, 2)
scatter(x, y)
