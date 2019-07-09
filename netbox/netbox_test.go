// +build integration

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netbox

import (
	"strconv"
	"testing"

	"github.com/digitalocean/go-netbox/netbox/client/dcim"
	"github.com/digitalocean/go-netbox/netbox/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

// These tests assume a netbox running locally, with the netbox test fixtures loaded.
// This is achievable with the following netbox management commands, after install:
//    python3 netbox/manage.py loaddata dcim extras ipam
//    python3 netbox/manage.py runserver 0.0.0.0:8000 --insecure

func TestRetrieveDeviceList(t *testing.T) {
	c := NewNetboxAt("localhost:8000")

	list, err := c.Dcim.DcimDevicesList(nil, nil)

	assert.NoError(t, err)
	assert.EqualValues(t, 11, *list.Payload.Count)
}

func TestSubdeviceRole(t *testing.T) {
	c := NewNetboxWithAPIKey("localhost:8000", "7b4e1ceaaf93528a41e64d048090f7fe13ed16f4")

	role := true
	manufacturerID := int64(1)
	model := "Test model"
	slug := "test-slug"
	newDeviceType := &models.WritableDeviceType{
		SubdeviceRole: role,
		Comments:      "Test device type",
		Manufacturer:  &manufacturerID,
		Model:         &model,
		Slug:          &slug,
	}
	err := newDeviceType.Validate(strfmt.Default)
	assert.NoError(t, err)

	createRequest := dcim.NewDcimDeviceTypesCreateParams().WithData(newDeviceType)
	createResponse, err := c.Dcim.DcimDeviceTypesCreate(createRequest, nil)
	assert.NoError(t, err)

	newID := strconv.Itoa(int(createResponse.Payload.ID))
	assert.NotEqual(t, 0, newID)

	retrieveResponse, err := c.Dcim.DcimDeviceTypesList(dcim.NewDcimDeviceTypesListParams().WithIDIn(&newID), nil)
	assert.NoError(t, err)
	assert.EqualValues(t, 1, *retrieveResponse.Payload.Count)
	assert.EqualValues(t, "Test device type", retrieveResponse.Payload.Results[0].Comments)

	deleteRequest := dcim.NewDcimDeviceTypesDeleteParams().WithID(createResponse.Payload.ID)
	_, err = c.Dcim.DcimDeviceTypesDelete(deleteRequest, nil)
	assert.NoError(t, err)
}
