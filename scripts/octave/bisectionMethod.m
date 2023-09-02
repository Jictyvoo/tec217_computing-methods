function [result, totalIteration, err] = bisectionMethod(a, b, fX, epsilon, maxIterations)
    err = ""; % Initialize err to an empty string
    if fX(a) * fX(b) >= 0
        err = "you have not assumed right a and b";
        result = NaN;
        totalIteration = 0;
        return
    end

    absoluteError = 1;
    totalIteration = 0;
    result = 0;

    while (absoluteError > epsilon && totalIteration < maxIterations)
        totalIteration += 1;
        oldResult = result;
        result = (a + b) / 2;
        if totalIteration > 1
            absoluteError = abs(oldResult - result);
        end

        fR = fX(result);
        if fR == 0
            return
        end

        if fX(a) * fR < 0
            b = result;
        else
            a = result;
        end
    end

    if totalIteration >= maxIterations
        err = ["failed to find root after ", num2str(totalIteration), " iterations"];
    end

endfunction
