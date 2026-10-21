// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package support holds the types and constants shared with the eBPF tracer, and
// the compiled tracer itself. It deliberately imports nothing beyond the standard
// library, so that packages such as libpf can depend on it without pulling the
// eBPF loader into every binary; loading the tracer lives in support/loader.
package support // import "go.opentelemetry.io/ebpf-profiler/support"

// TracerData returns the compiled eBPF tracer, which holds all possible eBPF maps
// and programs. It is empty on architectures the tracer is not built for.
func TracerData() []byte {
	return tracerData
}
