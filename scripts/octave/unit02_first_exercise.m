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
