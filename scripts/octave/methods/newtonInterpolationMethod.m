function [yint, err] = newtonInterpolationMethod(x, y, xx)
    err = ""; % Initialize err to an empty string
    sample_size = length(x);

    % Check for consistent input
    if sample_size != length(y)
        err = "x and y must have the same length";
        yint = NaN;
        return;
    endif

    % Initialize b as a zero 2D matrix
    b = zeros(sample_size);

    % Set the first column of b as y
    b(:, 1) = y;

    % Calculate finite differences
    for ctrlIdx = 2:sample_size
        for index = 1:(sample_size - ctrlIdx + 1)
            b(index, ctrlIdx) = (b(index + 1, ctrlIdx - 1) - b(index, ctrlIdx - 1)) / (x(index + ctrlIdx - 1) - x(index));
        endfor
    endfor

    % Apply finite differences to interpolate
    xt = 1;
    yint = b(1, 1);
    for index = 1:(sample_size - 1)
        xt *= (xx - x(index));
        yint += (b(1, index + 1) * xt);
    endfor

endfunction
