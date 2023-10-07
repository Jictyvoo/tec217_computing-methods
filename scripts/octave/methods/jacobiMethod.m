function [foundRoots, currentIteration, err] = jacobiMethod(inputMatrix, initialGuess, maxIterations, epsilon)
    err = ""; % Initialize err to an empty string
    currentIteration = 0;

    % Checking if the given matrix is a square one
    [matrixSize, eqLength] = size(inputMatrix);
    if matrixSize != eqLength -1
        err = "the given matrix isn't a square one";
        foundRoots = NaN;
        return;
    endif

    % Checking if it is diagonally dominant
    if isDiagonallyDominant(inputMatrix) == false
        foundRoots = NaN
        err = "the given matrix is not diagonally dominant";
        return
    endif

    matrixSize = size(inputMatrix, 1);

    workingMatrix = inputMatrix(:, 1:matrixSize);
    results = inputMatrix(:, matrixSize + 1);

    C = zeros(matrixSize, matrixSize);
    d = zeros(matrixSize, 1);

    for eqIndex = 1:matrixSize
        for index = 1:matrixSize
            setValue = 0;
            if eqIndex == index
                divisor = results(eqIndex);
                dividend = workingMatrix(eqIndex, eqIndex);
                d(eqIndex) = divisor / dividend;
            else
                setValue = -workingMatrix(eqIndex, index) / workingMatrix(eqIndex, eqIndex);
            endif
            C(eqIndex, index) = setValue;
        endfor
    endfor

    foundRoots = zeros(matrixSize, 1);
    foundRoots(1) = initialGuess;
    while currentIteration < maxIterations
        currentIteration += 1;
        [foundRoots, maxError, success] = calculateRoot(C, d, matrixSize, epsilon, foundRoots);
        if success
            return;
        endif
    endwhile
    err = printf("failed to converge in %d iterations", currentIteration);
endfunction

function isDiagonalDominant = isDiagonallyDominant(A)
    % Get the dimensions of the matrix
    matrixSize = size(A, 1);

    for index = 1:matrixSize
        sum = 0;
        for subIndex = 1:matrixSize
            if index != subIndex
                % Add the absolute value to the accumulator
                absoluteAij = abs(A(index, subIndex));
                sum += absoluteAij;
            endif
        endfor
        % Check if the value on the main diagonal is less than sum
        if abs(A(index, index)) < sum
            isDiagonalDominant = false;
            return;
        endif
    endfor

    % The matrix is diagonally dominant
    isDiagonalDominant = true;
endfunction

function [foundRoots, maximumXError, success] = calculateRoot(C, d, matrixSize, epsilon, xPrev)
    foundRoots = C * xPrev + d
    maximumXError = max(abs((foundRoots - xPrev) ./ foundRoots)) * 100;
    success = maximumXError < epsilon;
endfunction
