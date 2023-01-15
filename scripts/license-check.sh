#!/bin/bash

# Verify that the correct license block is present in all Go source
# files.
IFS=$'\n' read -r -d '' -a EXPECTED <<EndOfLicense
// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
EndOfLicense
AUTHOR_REGEX='^// Copyright 20[0-9][0-9] The go-netbox Authors\.$'

# Scan each Go source file for license.
EXIT=0
GOFILES=$(find . -name "*.go")

for FILE in $GOFILES; do
	IFS=$'\n' read -r -d '' -a BLOCK < <(tail -n +3 $FILE | head -n 14)
	IFS=$'\n' read -r -d '' -a BLOCK2 < <(head -n 14 $FILE)

	tmp_block=${BLOCK[@]:1}
	tmp_block2=${BLOCK2[@]:1}
	tmp_expected=${EXPECTED[@]:1}
	if [[ $tmp_block != $tmp_expected && $tmp_block2 != $tmp_expected ]]; then
		echo "file missing license: $FILE"
		EXIT=1
	fi
	if ! [[ "${BLOCK[0]}" =~ $AUTHOR_REGEX || "${BLOCK2[0]}" =~ $AUTHOR_REGEX ]]; then
		echo "file missing author line: $FILE"
		EXIT=1
	fi
done

exit $EXIT
