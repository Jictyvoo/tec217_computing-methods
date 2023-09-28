function [result, totalIteration, err] = secantMethod(x0, x1, fX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    relativeError = 1;
    totalIteration = 0;

    while (abs(relativeError) >= epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        fPrevious = fX(x0);
        fCurrent = fX(x1);

        result = ((x0 * fCurrent) - (x1 * fPrevious)) / (fCurrent - fPrevious);
        if result != 0 && totalIteration > 1
            relativeError = abs(x0 - result) / result;
        end
        x1 = x0;
        x0 = result;
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
