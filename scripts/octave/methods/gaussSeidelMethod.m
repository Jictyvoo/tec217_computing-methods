function [foundRoots, currentIteration, err] = gaussSeidelMethod(inputMatrix, X0, maxIterations, epsilon)
    err = ""; % Initialize err to an empty string
    currentIteration = 0;

    % Checking if the given matrix is a square one
    [matrixSize, eqLength] = size(inputMatrix);
    if matrixSize != eqLength -1
        err = "the given matrix isn't a square one";
        foundRoots = NaN;
        return;
    endif

    % Checking Sassenfeld criteria
    if sassenfeld(inputMatrix) == false
        foundRoots = NaN
        err = "the given matrix does not respect the Sassenfeld criteria";
        return;
    endif

    workingMatrix = inputMatrix(:, 1:matrixSize);
    results = inputMatrix(:, matrixSize + 1);

    [C, d] = mountCAndD(workingMatrix, results);

    foundRoots = X0;
    while currentIteration < maxIterations
        currentIteration += 1;
        oldSolution = foundRoots;
        foundRoots = calculateRoots(matrixSize, workingMatrix, foundRoots, oldSolution, results);

        maximumXError = max(abs((foundRoots - oldSolution) ./ foundRoots) * 100);
        if maximumXError < epsilon
            err = "";
            return;
        endif
    endwhile

    err = printf("failed to converge in %d iterations", currentIteration);
endfunction

function ssnfd = sassenfeld(A)
    matrixSize = size(A, 1);
    betas = zeros(matrixSize, 1);

    for index = 1:matrixSize
        tempValue = 0;
        for loopCount = 1:index - 1
            tempValue += abs(A(index, loopCount)) * betas(loopCount);
        endfor
        for subIndex = index + 1:matrixSize
            tempValue += abs(A(index, subIndex));
        endfor

        betas(index) = tempValue / abs(A(index, index));

        if betas(index) >= 1
            ssnfd = false;
            return;
        endif
    endfor

    ssnfd = true;
endfunction

function [C, d] = mountCAndD(A, results)
    matrixSize = size(A, 1);
    C = zeros(matrixSize, matrixSize);
    d = zeros(matrixSize, 1);

    for eqIndex = 1:matrixSize
        for index = 1:matrixSize
            setValue = 0;
            if eqIndex != index
                setValue = -A(eqIndex, index) / A(eqIndex, eqIndex);
            endif
        endfor
        divisor = results(eqIndex);
        dividend = A(eqIndex, eqIndex);
        d(eqIndex) = divisor / dividend;
    endfor
endfunction

function foundRoots = calculateRoots(matrixSize, A, foundRoots, oldRoots, results)
    for index = 1:matrixSize
        firstSum = sum(A(index, 1:index - 1) .* foundRoots(1:index - 1));
        secondSum = sum(A(index, index + 1:matrixSize) .* oldRoots(index + 1:matrixSize));
        foundRoots(index) = (results(index) - firstSum - secondSum) / A(index, index);
    endfor
endfunction
