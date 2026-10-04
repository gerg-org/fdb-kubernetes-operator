/*
 * image_ref_test.go
 *
 * This source file is part of the FoundationDB open source project
 *
 * Copyright 2018-2026 Apple Inc. and the FoundationDB project authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package internal

import "testing"

func TestSameImageRef(t *testing.T) {
	const base = "registry.example/gergorg/foundationdb-kubernetes-sidecar:7.1.61-1"
	const digested = base + "@sha256:630e5388047a6211c5a1f222b3dbd827230c7c6e2644a484c59223b7d43bc93f"

	if !SameImageRef(base, digested) {
		t.Fatal("digest suffix should not count as a different image")
	}
	if !SameImageRef(digested, digested) {
		t.Fatal("identical digested refs should match")
	}
	if SameImageRef(base, base+"-other") {
		t.Fatal("different tags should not match")
	}
}
