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
fprintf('Determination coefficient (r^2): %.4f\n', r2);

% Calculate the adjustment line
x_fit = linspace(min(x), max(x), 100);
y_fit = a(1) * x + a(2);

% Calculate total standard deviation
mean_y = mean(y);
std_total = sqrt(sum((y - mean_y).^2) / (length(y) - 1));

% Calculate the standard error of the estimate
std_error = sqrt(sum((y - y_fit).^2) / (length(y) - 2));

% Statistical analysis
fprintf('Total Standard Deviation: %.4f\n', std_total);
fprintf('Standard error of the Estimate: %.4f\n', std_error);
fprintf('\n');
%%%%%%%%%%% Start code execution %%%%%%%%%%%
xi = [10; 20; 30; 40; 50; 60; 70; 80];
yi = [125; 70; 380; 550; 610; 1220; 830; 1450];

% Fitting logarithms using polyfit
coefficients = polyfit(log(xi), log(yi), 1);

% Adjusted coefficients
alpha = exp(coefficients(2));
beta = coefficients(1);

% Creating a function for the fitted power equation
fitted_function = @(x) alpha * x.^beta;

new_y = fitted_function(xi)
[a, r2, err] = minimumSquareRegressionMethod(xi, new_y);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: minimumSquareRegressionMethod\n');
fprintf('Result [a]:'), disp(a);
fprintf('Determination coefficient (r^2): %.4f\n', r2);

% Calculate the adjustment line
x_fit = linspace(min(xi), max(xi), 100);
y_fit = a(1) * xi + a(2);

% Calculate total standard deviation
mean_y = mean(new_y);
std_total = sqrt(sum((new_y - mean_y).^2) / (length(new_y) - 1));

% Calculate the standard error of the estimate
std_error = sqrt(sum((new_y - y_fit).^2) / (length(new_y) - 2));

% Statistical analysis
fprintf('Total Standard Deviation: %.4f\n', std_total);
fprintf('Standard error of the Estimate: %.4f\n', std_error);
fprintf('\n');

%%%%%%%%%%% Start code execution %%%%%%%%%%%
x = [0, 1, 2, 3, 4, 5];
y = [0, 20, 60, 68, 77, 100];

for degree = 1:5
    [result, err] = minimumSquarePolynomialRegressionMethod(x, y, double(degree));
    if ~isempty(err)
        disp(err);
        return;
    end
    disp(['PolynomialRegression - ' num2str(degree) ' degree']);
    disp(result');
end
