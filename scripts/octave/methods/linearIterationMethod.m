function [result, totalIteration, err] = linearIterationMethod(x0, fX, gX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    absoluteError = 1;
    totalIteration = 0;
    oldResult = 0;
    result = x0;

    while (abs(absoluteError) >= epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        oldResult = result;
        result = gX(result);

        if result != 0 && totalIteration > 1
            absoluteError = abs(oldResult - result) / result;
        end

        rootResult = fX(result);
        if rootResult == 0
            break
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
