/*
#
# Copyright Usman Malik - https://github.com/imusmanmalik
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
*/

package randomizer

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	g := NewGenerator()

	// Test Intn function.
	n, err := g.Intn(100)
	if err != nil {
		t.Fatalf("Intn returned error: %s", err)
	}
	if n < 0 || n >= 100 {
		t.Fatalf("Intn returned invalid value: %d", n)
	}

	// Test Float64 function.
	f, err := g.Float64()
	if err != nil {
		t.Fatalf("Float64 returned error: %s", err)
	}
	if f < 0.0 || f >= 1.0 {
		t.Fatalf("Float64 returned invalid value: %f", f)
	}
}
