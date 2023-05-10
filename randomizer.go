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

// RandomInt generates a cryptographically secure random integer between min and max (inclusive)
func RandomInt(min, max int64) (int64, error) {
	// Create a new big.Int to hold the range of possible values
	rangeSize := big.NewInt(max - min + 1)

	// Generate a cryptographically secure random value within the range
	randomValue, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, err
	}

	// Convert the random value to an int64 and add the minimum value to shift the range
	return randomValue.Int64() + min, nil
}

// RandomBytes generates a slice of cryptographically secure random bytes with a given length
func RandomBytes(length int) ([]byte, error) {
	// Create a slice to hold the random bytes
	randomBytes := make([]byte, length)

	// Generate the random bytes using the crypto/rand package
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

// RandomString generates a cryptographically secure random string with a given length
func RandomString(length int) (string, error) {
	// Generate a slice of random bytes with the given length
	randomBytes, err := RandomBytes(length)
	if err != nil {
		return "", err
	}

	randomString := string(randomBytes)

	return randomString, nil
}
