function [LMatrix, UMatrix, det, roots, err] = luDecompositionMethod(inputMatrix)
    err = ""; % Initialize err to an empty string

    % Checking if the given matrix is a square one
    [matrixSize, eqLength] = size(inputMatrix);
    if matrixSize != eqLength -1
        err = "the given matrix isn't a square one";
        det = NaN;
        roots = NaN;
        LMatrix = NaN;
        UMatrix = NaN;
        return;
    end

    det = 1;
    % Initialize the L and U matrices
    LMatrix = eye(matrixSize);

    % Split the input matrix to obtain uMatrix and results separatedly
    UMatrix = inputMatrix(:, 1:matrixSize);
    results = inputMatrix(:, matrixSize + 1);

    for eqIndex = 1:matrixSize
        pivot = UMatrix(eqIndex, eqIndex);

        for rowIndex = eqIndex + 1:matrixSize
            factor = UMatrix(rowIndex, eqIndex) / pivot;
            UMatrix(rowIndex, :) -= factor * UMatrix(eqIndex, :);
            LMatrix(rowIndex, eqIndex) = factor; % store the multiplier in L matrix
        end

        det *= UMatrix(eqIndex, eqIndex);
    end

    % Calculate roots by forward and backward substitution
    y = forwardSubstitution(LMatrix, matrixSize, results);
    roots = backwardSubstitution(UMatrix, matrixSize, y);

end

% Solve Inferior triangle
function [x] = forwardSubstitution(L, matrixSize, b)
    x = zeros(matrixSize, 1);
    x(1) = b(1) / L(1, 1);

    for index = 2:matrixSize
        x(index) = (b(index) - L(index, 1:index - 1) * x(1:index - 1)) / L(index, index);
    end
end

% Solve superior triangle
function [x] = backwardSubstitution(U, matrixSize, b)
    x = zeros(matrixSize, 1);
    x(matrixSize) = b(matrixSize) / U(matrixSize, matrixSize);

    for index = matrixSize - 1:-1:1
        x(index) = (b(index) - U(index, index + 1:matrixSize) * x(index + 1:matrixSize)) / U(index, index);
    end
end
