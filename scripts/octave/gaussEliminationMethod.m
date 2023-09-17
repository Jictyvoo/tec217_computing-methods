function [det, foundRoots, err] = gaussEliminationMethod(inputMatrix, usePivot)
    err = ""; % Initialize err to an empty string

    % Checking if the given matrix is a square one
    [totalEquations, eqLength] = size(inputMatrix);
    if totalEquations != eqLength -1
        err = "the given matrix isn't a square one";
        det = NaN;
        foundRoots = NaN;
        return;
    end

    % If pivot option is enabled, do the pivot
    if usePivot
        inputMatrix = pivot(inputMatrix, 1, 1);
    end
    for eqIndex = 1:totalEquations
        if eqIndex == totalEquations
            break;
        end

        % Eliminate the next equations
        for index = eqIndex + 1:totalEquations
            divisionFactor = inputMatrix(index, eqIndex) / inputMatrix(eqIndex, eqIndex);
            inputMatrix(index, :) -= divisionFactor * inputMatrix(eqIndex, :);
        end
    end

    % Calculate determinant
    det = 1;
    for index = 1:totalEquations
        det *= inputMatrix(index, index);
    end

    % Calculate roots from bottom-up
    foundRoots = zeros(1, totalEquations);
    for eqIndex = totalEquations:-1:1
        sum = inputMatrix(eqIndex, totalEquations + 1);
        for i = eqIndex + 1:totalEquations
            sum -= inputMatrix(eqIndex, i) * foundRoots(i);
        end

        foundRoots(eqIndex) = sum / inputMatrix(eqIndex, eqIndex);
    end
end

function [matrix] = pivot(matrix, currentIndex, lookupColumn)
    % Find the index of the maxValue value in column
    maxValue = abs(matrix(currentIndex, lookupColumn));
    swapIndex = currentIndex;
    [totalEquations, eqLength] = size(matrix)

    for index = currentIndex + 1:totalEquations
        absValue = abs(matrix(index, lookupColumn));
        if absValue > maxValue
            swapIndex = index;
            maxValue = absValue;
        end
    end

    % Swap the current row and the maxValue row if necessary
    if swapIndex ~= currentIndex
        temp = matrix(swapIndex, :);
        matrix(swapIndex, :) = matrix(currentIndex, :);
        matrix(currentIndex, :) = temp;
    end
end
