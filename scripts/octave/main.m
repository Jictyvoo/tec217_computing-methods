% Define the function
fX = @(x)(x^3 - 9*x + 3);

% Define function parameters
a = 0;
b = 1;
epsilon = 5e-4;
maxIterations = 1000;

% Call the function
[result, totalIterations, err] = falsePositiveMethod(a, b, fX, epsilon, maxIterations);

% Check if error exists
if ~isempty(err)
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end


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
    fprintf('Failed to find root. Error: %s\n', err);
else
    fprintf('The value of the root is: %.5f. After %d iterations\n', result, totalIterations);
end
