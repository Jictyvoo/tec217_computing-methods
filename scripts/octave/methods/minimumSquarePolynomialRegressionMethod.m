function [result, err] = minimumSquarePolynomialRegressionMethod(x, y, polyDegree)
    err = ""; % Initialize err to an empty string

    [sampleSize] = length(x);

    % Checking if the given inputs has the same size
    if sampleSize != length(y)
        err = "x and y must have the same length";
        result = NaN;
        return;
    end

    n = polyDegree + 1;
    a = zeros(n, n);
    b = zeros(n, 1);

    % Generating the linear system
    b(1) = sum(y);
    a(1, 1) = sampleSize;
    for index = 1:n
        k1 = index - 1;
        k2 = index;
        for subIndex = 2:n
            sumAuxiliary = sum(power(x, k2));
            a(index, subIndex) = sumAuxiliary;
            a(subIndex, index) = a(index, subIndex);
            k2 += 1;
        end

        sumAuxiliary = 0;
        if index ~= 1
            for indexIterator = 1:1:sampleSize
                sumAuxiliary += y(indexIterator) * x(indexIterator)^(k1);
            end
            b(index) = sumAuxiliary;
        end
    end

    % Start the linear system solution

    % Superior Triangular Matrix
    [a, b] = solveSuperiorTriangle(n, a, b);

    % Start the system solution calculation
    result = zeros(n, 1);
    result(n) = b(n) / a(n, n);
    for k = n - 1:-1:1
        s = 0;
        for j = k + 1:n
            s = s + a(k, j) * result(j);
        end
        result(k) = (b(k) - s) / a(k, k);
    end
end

function [a, b] = solveSuperiorTriangle(n, a, b)
    for k = 1:n - 1
        for i = k + 1:n
            m = a(i, k) / a(k, k);
            a(i, k) = 0;
            for j = k + 1:n
                a(i, j) -= m * a(k, j);
            end
            b(i) = b(i) - m * b(k);
        end
    end
    return
end

function result = power(x, y)
    result = x.^y;
end
