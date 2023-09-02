function [result, totalIteration, err] = newtonRaphsonMethod(x0, fX, f1X, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    relativeError = 1;
    totalIteration = 0;
    oldResult = 0;
    result = x0;

    while (abs(relativeError) >= epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        oldResult = result;
        result = result - (fX(oldResult) / f1X(oldResult));

        if result != 0 && totalIteration > 1
            relativeError = abs(oldResult - result) / result;
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
