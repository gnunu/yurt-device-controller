/*
Copyright 2022 The OpenYurt Authors.

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

package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateOptions(t *testing.T) {
	o := NewYurtDeviceControllerOptions()
	assert.Equal(t, ValidateOptions(o), nil)
}

func TestValidateEdgePlatformAddress(t *testing.T) {
	o := NewYurtDeviceControllerOptions()
	assert.Equal(t, ValidateEdgePlatformAddress(o), nil)

	o.CoreDataAddr = ""
	assert.Equal(t, ValidateEdgePlatformAddress(o), nil)

	o.CoreDataAddr = ":65536"
	assert.Equal(t, ValidateEdgePlatformAddress(o), nil)

	o.CoreDataAddr = "edgex.me:"
	assert.Equal(t, ValidateEdgePlatformAddress(o), nil)

	o.CoreDataAddr = ":"
	assert.Equal(t, ValidateEdgePlatformAddress(o), nil)

	o.CoreDataAddr = "edgex"
	assert.Contains(t, ValidateEdgePlatformAddress(o).Error(), "missing port in address")

	o.CoreDataAddr = "edgex:1:2"
	assert.Contains(t, ValidateEdgePlatformAddress(o).Error(), "too many colons in address")
}
