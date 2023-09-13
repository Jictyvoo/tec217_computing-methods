% Define all constants that will be used
Pumax = 80000;
ku = 0.05;
Pumin = 110000;
Psmax = 320000;
P0 = 10000;
ks = 0.09;

% Define the function
fX = @(x) (Psmax / (1 + (((Psmax / P0) - 1) * e^(-ks * x)))) - (1.2 * ((Pumax * e^(-ku * x)) + Pumin));

% Define function parameters
a = 0;
b = 100;
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
fX = @(x) (Psmax / (1 + (((Psmax / P0) - 1) * e^(-ks * x)))) - (1.2 * ((Pumax * e^(-ku * x)) + Pumin));

% Define function parameters
a = 0;
b = 100;
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
fX = @(x) (Psmax / (1 + (((Psmax / P0) - 1) * e^(-ks * x)))) - (1.2 * ((Pumax * e^(-ku * x)) + Pumin));

% Define function parameters
a = 0;
b = 100;
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
fX = @(x) (32 / (1 + (31 * exp(-ks * x)))) - 13.2 - (9.6 * exp(-ku * x));
gX = @(x) log(((32 / ((9.6 * exp(-0.05 * x)) + 13.2)) - 1) / 31) / -0.09;

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
fX = @(x) (32 / (1 + (31 * exp(-ks * x)))) - 13.2 - (9.6 * exp(-ku * x));
f1X = @(x) (0.48 * exp(-0.05 * x)) + ((89.28 * exp(0.09 * x)) / (exp(0.09 * x) + 31)^2);

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
