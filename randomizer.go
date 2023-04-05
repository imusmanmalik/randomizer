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
	"crypto/rand"
	"math/big"
)

// Generator is a struct that holds the random generator's configuration.
type Generator struct {
}

// NewGenerator returns a new instance of the random generator.
func NewGenerator() *Generator {
	return &Generator{}
}

// Intn returns a cryptographically secure random integer between 0 and n.
func (g *Generator) Intn(n int) (int, error) {
	max := big.NewInt(int64(n))
	r, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return int(r.Int64()), nil
}

// Float64 returns a cryptographically secure random float64 between 0 and 1.
func (g *Generator) Float64() (float64, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(1e17))
	if err != nil {
		return 0, err
	}
	return float64(r.Int64()) / 1e17, nil
}
