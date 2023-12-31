addpath('./methods')

%%%%%%%%%%% Start code execution %%%%%%%%%%%
inputMatrix = [
            10, 2, -1, 27;
            -3, -5, 2, -61.5;
            1, 1, 6, -21.5
            ];

[det, foundRoots, err] = gaussEliminationMethod(inputMatrix, false);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussElimination\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:'), disp(foundRoots);

fprintf('\n');
%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
inputMatrix = [
            2, -6, -1, -38;
            -3, -1, 7, -34;
            -8, 1, -2, -20
            ];

[det, foundRoots, err] = gaussEliminationMethod(inputMatrix, true);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussElimination with Pivot\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:'), disp(foundRoots);

fprintf('\n');

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
inputMatrix = [
            -0.04, 0.04, 0.12, 3;
            0.56, -1.56, 0.32, 1;
            -0.24, 1.24, -0.28, 0
            ];

[det, invMatrix, foundRoots, err] = gaussJordanMethod(inputMatrix);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussJordan\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:\n'), disp(foundRoots);
fprintf('Inverse:\n'), disp(invMatrix);

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
inputMatrix = [
            2, 1, -1, 3;
            -1, 3, 2, 1;
            3, 1, -3, 2;
            ];

[LMatrix, UMatrix, det, foundRoots, err] = luDecompositionMethod(inputMatrix);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: LU Decomposition\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:\n'), disp(foundRoots);

% Report L and U Matrices
disp("LMatrix: "); disp(LMatrix);
disp("UMatrix: "); disp(UMatrix);

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
inputMatrix = [
            5, 1, 1, 5;
            3, 4, 1, 6;
            3, 3, 6, 0
            ];
initialGuess = 0;
maxIterations = 1000;
epsilon = 0.05;

% Run the Jacobi method.
[foundRoots, totalIterations, err] = jacobiMethod(inputMatrix, initialGuess, maxIterations, epsilon);

% Report error if there is one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: Jacobi\n');
fprintf('Total Iterations: %d\n', totalIterations);
fprintf('Roots:\n'), disp(foundRoots);

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
inputMatrix = [
            1, 0.5, -0.1, 0.2;
            0.2, 1, -0.2, -2;
            -0.1, -0.2, 1, 1
            ];
X0 = [0, 0, 0];
maxIterations = 1000;
epsilon = 1e-5;

% Run the Gauss-Seidel method.
[foundRoots, totalIterations, err] = gaussSeidelMethod(inputMatrix, X0, maxIterations, epsilon);

% Report error if there is one.
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: Gauss-Seidel\n');
fprintf('Total Iterations: %d\n', totalIterations);
fprintf('Roots:\n'), disp(foundRoots);
