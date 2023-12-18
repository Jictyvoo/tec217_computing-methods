function [result, err] = lagrangeInterpolationMethod(x, y, xx)
    err = ""; % Initialize err to an empty string
    sampleSize = length(x);

    % Check for consistent input
    if sampleSize != length(y)
        err = "x and y must have the same length";
        result = NaN;
        return;
    endif

    result = 0;
    for ctrlIndex = 1:sampleSize
        term = y(ctrlIndex);

        % Start to perform the prod calculation
        for index = 1:sampleSize
            if ctrlIndex ~= index
                term *= (xx - x(index)) / (x(ctrlIndex) - x(index));
            end
        end
        result += term;
    end

endfunction
