/*
 * image_ref.go
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

import "strings"

// SameImageRef reports whether two container image references name the same
// image, ignoring a registry digest suffix. Admission webhooks such as Kyverno
// mutateDigest rewrite a pod image from name:tag to name:tag@sha256:..., while
// the operator's desired image stays name:tag.
func SameImageRef(left, right string) bool {
	return stripImageDigest(left) == stripImageDigest(right)
}

func stripImageDigest(image string) string {
	if i := strings.LastIndex(image, "@"); i >= 0 {
		return image[:i]
	}
	return image
}
