/*
Copyright (c) 2022 RaptorML authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"context"
	"fmt"
	manifests "github.com/raptor-ml/raptor/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
	"time"
)

const FeatureSetBuilder = "featureset"
const ExpressionBuilder = "expression"

// DataSource is a parsed abstracted representation of a manifests.DataSource
type DataSource struct {
	FQN    string                 `json:"fqn"`
	Kind   string                 `json:"kind"`
	Config manifests.ParsedConfig `json:"config"`
	//Todo Schema
}

// DataSourceFromManifest returns a DataSource from a manifests.DataSource
func DataSourceFromManifest(ctx context.Context, src *manifests.DataSource, r client.Reader) (DataSource, error) {
	pc, err := src.ParseConfig(ctx, r)
	if err != nil {
		return DataSource{}, fmt.Errorf("failed to parse config: %w", err)
	}

	return DataSource{
		FQN:    src.FQN(),
		Kind:   src.Spec.Kind,
		Config: pc,
	}, nil
}

// FeatureDescriptor is describing a feature definition for an internal use of the Core.
type FeatureDescriptor struct {
	FQN        string        `json:"FQN"`
	Primitive  PrimitiveType `json:"primitive"`
	Aggr       []AggrFn      `json:"aggr"`
	Freshness  time.Duration `json:"freshness"`
	Staleness  time.Duration `json:"staleness"`
	Timeout    time.Duration `json:"timeout"`
	Builder    string        `json:"builder"`
	DataSource string        `json:"data_source"`
}

// ValidWindow checks if the feature have aggregation enabled, and if it is valid
func (fd FeatureDescriptor) ValidWindow() bool {
	if fd.Freshness < 1 {
		return false
	}
	if fd.Staleness < fd.Freshness {
		return false
	}
	if len(fd.Aggr) == 0 {
		return false
	}
	if !(fd.Primitive == PrimitiveTypeInteger || fd.Primitive == PrimitiveTypeFloat) {
		return false
	}
	return true
}
func aggrsToStrings(a []manifests.AggrFn) []string {
	var res []string
	for _, v := range a {
		res = append(res, string(v))
	}
	return res
}

// FeatureDescriptorFromManifest returns a FeatureDescriptor from a manifests.Feature
func FeatureDescriptorFromManifest(in *manifests.Feature) (*FeatureDescriptor, error) {
	primitive := StringToPrimitiveType(string(in.Spec.Primitive))
	if primitive == PrimitiveTypeUnknown {
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedPrimitiveError, in.Spec.Primitive)
	}
	aggr, err := StringsToAggrFns(aggrsToStrings(in.Spec.Builder.Aggr))
	if err != nil {
		return nil, fmt.Errorf("failed to parse aggregation functions: %w", err)
	}
	if len(aggr) > 0 && primitive != PrimitiveTypeInteger && primitive != PrimitiveTypeFloat {
		return nil, fmt.Errorf("%w with Aggregation: %s", ErrUnsupportedPrimitiveError, in.Spec.Primitive)
	}
	if in.Spec.Builder.AggrGranularity.Milliseconds() > 0 && len(aggr) > 0 {
		in.Spec.Freshness = in.Spec.Builder.AggrGranularity
	}

	fd := &FeatureDescriptor{
		FQN:       in.FQN(),
		Primitive: primitive,
		Aggr:      aggr,
		Freshness: in.Spec.Freshness.Duration,
		Staleness: in.Spec.Staleness.Duration,
		Timeout:   in.Spec.Timeout.Duration,
		Builder:   strings.ToLower(in.Spec.Builder.Kind),
	}
	if in.Spec.DataSource != nil {
		fd.DataSource = in.Spec.DataSource.FQN()
	}
	if fd.Builder == "" {
		fd.Builder = ExpressionBuilder
	}

	if len(fd.Aggr) > 0 && !fd.ValidWindow() {
		return nil, fmt.Errorf("invalid feature specification for windowed feature")
	}
	return fd, nil
}

func NormalizeFQN(fqn, defaultNamespace string) string {
	ns := strings.Index(fqn, ".")
	if ns != -1 {
		return fqn
	}

	fn := strings.Index(fqn, "[")
	if fn != -1 {
		return fmt.Sprintf("%s.%s%s", fqn[:fn], defaultNamespace, fqn[fn:])
	}
	return fmt.Sprintf("%s.%s", fqn, defaultNamespace)
}
