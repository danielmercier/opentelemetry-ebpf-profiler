// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package loader loads the eBPF tracer embedded in package support.
package loader // import "go.opentelemetry.io/ebpf-profiler/support/loader"

import (
	"bytes"

	cebpf "github.com/cilium/ebpf"

	"go.opentelemetry.io/ebpf-profiler/support"
)

// LoadCollectionSpec is a wrapper around ebpf.LoadCollectionSpecFromReader and loads the eBPF
// Spec from the embedded file.
// We expect support.TracerData to hold all possible eBPF maps and programs.
func LoadCollectionSpec() (*cebpf.CollectionSpec, error) {
	return cebpf.LoadCollectionSpecFromReader(bytes.NewReader(support.TracerData()))
}
