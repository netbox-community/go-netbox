#!/usr/bin/env bash

# Maximum and minimum values in spec that correspond to 
# `math.MaxInt64` and `math.MinInt64`, respectively, are not rendered
# properly by "go-swagger", thus causing the following errors.
#
# cannot use -9.223372036854776e+18 (untyped float constant -9.22337e+18) as int64 value in argument to validate.MinimumInt (truncated)
# cannot use 9.223372036854776e+18 (untyped float constant 9.22337e+18) as int64 value in argument to validate.MaximumInt (truncated)
#
# See: https://github.com/go-swagger/go-swagger/issues/2755

set -euo pipefail

find "${SRC_DIR}" -type f -name '*.go' -exec \
  sed -i \
    -e 's/-9.223372036854776e+18/math.MinInt64/' \
    -e 's/9.223372036854776e+18/math.MaxInt64/' \
    {} \;
