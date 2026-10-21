#!/usr/bin/env bash

set -eu

# generate <dir> <package path> <definitions file> <generated file> regenerates the
# Go definitions of <definitions file> and fails if <generated file> is out of date.
generate() {
    local dir=$1 pkgpath=$2 defs=$3 out=$4
    local pkg=${pkgpath##*/}

    (
        cd "$dir"

        # Put license header at the top of the file
        cat <<EOF >types_gen.go
// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

EOF

        # Generate Go definitions from C
        go tool cgo -godefs -- -DEBPF_TRACE_HEADER_ONLY "$defs" >> types_gen.go

        # Properly format the generated code
        go fmt .

        # Set correct package path
        sed -i "s|^package ${pkg}\$|package ${pkg} // import \"${pkgpath}\"|" types_gen.go

        if ! diff types_gen.go "$out"; then
            echo "Auto generated and existing code differ, please review and update ${dir}/${out}"
            exit 1
        fi

        # Clean up temporary files
        rm -rf _obj/ types_gen.go
    )
}

generate . go.opentelemetry.io/ebpf-profiler/support types_def.go types.go
generate loader go.opentelemetry.io/ebpf-profiler/support/loader metrics_def.go metrics.go
