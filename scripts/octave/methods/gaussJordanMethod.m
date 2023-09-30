function [det, invMatrix, roots, err] = gaussJordanMethod(inputMatrix)
    err = ""; % Initialize err to an empty string

    % Checking if the given matrix is a square one
    [totalEquations, eqLength] = size(inputMatrix);
    if totalEquations != eqLength -1
        err = "the given matrix isn't a square one";
        det = NaN;
        roots = NaN;
        invMatrix = NaN;
        return;
    end

    % Generate the augmentedMatrix by insert the identity on it
    augmentedMatrix = [inputMatrix, eye(totalEquations)];

    det = 1
    for eqIndex = 1:totalEquations
        % Swap equation by the one in pivot element and use the value to normalize it next
        [augmentedMatrix, pivotElement] = pivot(augmentedMatrix, eqIndex, eqIndex);
        det *= pivotElement

        % Normalization
        augmentedMatrix(eqIndex, :) = augmentedMatrix(eqIndex, :) / augmentedMatrix(eqIndex, eqIndex);

        % Elimination
        for index = setdiff(1:totalEquations, eqIndex)
            augmentedMatrix(index, :) = augmentedMatrix(index, :) - augmentedMatrix(eqIndex, :) * augmentedMatrix(index, eqIndex);
        end
    end

    invMatrix = augmentedMatrix(:, totalEquations + 2:(totalEquations * 2) + 1);
    roots = augmentedMatrix(:, totalEquations + 1);
end

function [matrix, maxValue] = pivot(matrix, currentIndex, lookupColumn)
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
