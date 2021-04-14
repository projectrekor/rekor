/*
Copyright The Rekor Authors.

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

package signer

import (
	"context"
	"fmt"
	"strings"

	"github.com/sigstore/cosign/pkg/cosign/kms/gcp"
	"github.com/sigstore/sigstore/pkg/signature"
)

func New(ctx context.Context, signer string) (signature.Signer, error) {
	switch {
	case strings.HasPrefix(signer, gcp.ReferenceScheme):
		return gcp.NewGCP(ctx, signer)
	case signer == MemoryScheme:
		return NewMemory()
	default:
		return nil, fmt.Errorf("please provide a valid signer, %v is not valid", signer)
	}
}
