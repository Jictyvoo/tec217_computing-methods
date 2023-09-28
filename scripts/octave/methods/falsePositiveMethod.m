function [result, totalIteration, err] = falsePositiveMethod(a, b, fX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    if fX(a) * fX(b) >= 0
        err = "there is no root on given interval";
        result = NaN;
        totalIteration = 0;
        return
    end

    absoluteError = 1;
    totalIteration = 0;
    result = 0;

    while (abs(absoluteError) > epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        fResultA = fX(a);
        fResultB = fX(b);
        oldResult = result;

        if fResultA == fResultB
            err = "division by zero error";
            return;
        end

        result = b - (fResultB * (a - b)) / (fResultA - fResultB);
        if totalIteration > 1
            absoluteError = abs(oldResult - result);
        end

        rootResult = fX(result);

        if rootResult == 0
            return;
        end

        if fResultA * rootResult < 0
            b = result;
        else
            a = result;
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
