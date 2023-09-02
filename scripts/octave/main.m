% Define the function
fX = @(x)(x^3 - 9 * x + 3);

% Define function parameters
a = 0;
b = 1;
epsilon = 5e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = falsePositiveMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[falsePositiveMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[falsePositiveMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%----------------------------------------------------------------------------------------------------%

% Define the function
fX = @(x)(sin(x) - x^2);

% Define function parameters
a = 0.5;
b = 1;
epsilon = 2e-2;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = bisectionMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[bisectionMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[bisectionMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%----------------------------------------------------------------------------------------------------%

% Define the function
fX = @(x)((x^2) - (3 * x) + (e^x) -2);
gX = @(x)((x^2) - 2 + (e^x)) / 3;

% Define function parameters
x0 = -1;
epsilon = 1e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[linearIterationMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[linearIterationMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%----------------------------------------------------------------------------------------------------%

% Define the function
fX = @(x)((x^2) - (3 * x) + (e^x) -2);
f1X = @(x)((2 * x) + (e^x) - 3);

% Define function parameters
x0 = -1;
epsilon = 1e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[newtonRaphsonMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[newtonRaphsonMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end

%----------------------------------------------------------------------------------------------------%

% Define the function
fX = @(x)((x^2) - (3 * x) + (e^x) -2);

% Define function parameters
x0 = 0;
x1 = -1;
epsilon = 1e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = secantMethod(x0, x1, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('[secantMethod] Failed to find root. Error: %s\n', err);
else
    fprintf('[secantMethod] The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
