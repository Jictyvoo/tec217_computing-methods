addpath('./methods')

%%%%%%%%%%% Start code execution %%%%%%%%%%%
% Define all constants that will be used
e0 = 8.9e-12;
F = 1.25;
r = 0.85;
q = 2e-5;
Q = q;

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;

% Define function parameters
a = 0;
b = 0.5;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = bisectionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[bisectionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[bisectionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;

% Define function parameters
a = 0;
b = 0.5;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = falsePositiveMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[falsePositionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[falsePositionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

% Define the function
fX = @(x) ((q * Q * x) / ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) - F;

% Define function parameters
a = 0;
b = 0.5;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = secantMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[secantMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[secantMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

% Define the function
gX = @(x) (F * ((4 * pi * e0) * sqrt((x^2 + r^2)^3))) / (q * Q);

% Define function parameters
x0 = 0;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[linearIterationMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[linearIterationMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%

% Define the function
f1X = @(x) (3.57652 / sqrt((x^2 + 0.7255)^3)) - ((10.7295 * x^2) / sqrt((x^2 + 0.7255)^2));

% Define function parameters
x0 = 0;
epsilon = 1e-8;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[newtonRaphsonMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[newtonRaphsonMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
