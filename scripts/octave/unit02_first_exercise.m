inputMatrix = [
            10, 2, -1, 27;
            -3, -5, 2, -61.5;
            1, 1, 6, -21.5
            ];

[det, foundRoots, err] = gaussEliminationMethod(inputMatrix);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussElimination\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:'), disp(foundRoots);

%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%
inputMatrix = [
            2, -6, -1, -38;
            -3, -1, 7, -34;
            -8, 1, -2, -20
            ];

[det, foundRoots, err] = gaussEliminationMethod(inputMatrix);

% Report error if the matrix isn't a square one:
if ~isempty(err)
    disp(['Error: ', err]);
    return;
end

% Report Result:
fprintf('Method: GaussElimination\n');
fprintf('Determinant: %d\n', det);
fprintf('Roots:'), disp(foundRoots);
