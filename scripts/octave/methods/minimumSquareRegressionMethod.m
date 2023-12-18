function [a, r2, err] = minimumSquareRegressionMethod(x, y)
    err = ""; % Initialize err to an empty string

    [sampleSize] = length(x);
    % Checking if the given inputs has the same size
    if sampleSize != length(y)
        err = "x and y must have the same length";
        a = [];
        r2 = [];
        return;
    end

    sums = struct('x', 0, 'y', 0, 'xPow2', 0, 'yPow2', 0, 'xy', 0);
    for index = 1:sampleSize
        values = struct('x', x(index), 'y', y(index));
        sums.x = sums.x + values.x;
        sums.y = sums.y + values.y;
        sums.xy = sums.xy + values.x * values.y;
        sums.xPow2 = sums.xPow2 + values.x * values.x;
        sums.yPow2 = sums.yPow2 + values.y * values.y;
    end

    % Calculate coefficients
    a(1) = ((sampleSize * sums.xy) - (sums.x * sums.y)) / ((sampleSize * sums.xPow2) - (sums.x * sums.x));
    a(2) = (sums.y / sampleSize) - (a(1) * (sums.x / sampleSize));

    averageResiduum = 0;
    linearAdjustResiduum = 0;

    for index = 1:sampleSize
        averageDiff = y(index) - (sums.y / sampleSize);
        averageResiduum = averageResiduum + averageDiff * averageDiff;

        linearTempVal = y(index) - a(2) - (a(1) * x(index));
        linearAdjustResiduum = linearAdjustResiduum + linearTempVal * linearTempVal;
    end

    r2 = (averageResiduum - linearAdjustResiduum) / averageResiduum;
end
