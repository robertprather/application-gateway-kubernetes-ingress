// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

package appgw

import (
	"fmt"
	"strings"
	"github.com/golang/glog"
)


// SubscriptionID in the resourceID
type SubscriptionID string

// ResourceGroup in the resourceID
type ResourceGroup string

// ResourceName in the resourceID
type ResourceName string

// ParseResourceID gets subscriptionId, resource group, resource name from resourceID
func ParseResourceID(ID string) (SubscriptionID, ResourceGroup, ResourceName) {
	split := strings.Split(ID, "/")
	if len(split) < 9 {
		glog.Errorf("resourceID %s is invalid. There should be atleast 9 segments in resourceID", ID)
		return "", "", ""
	}

	return SubscriptionID(split[2]), ResourceGroup(split[4]), ResourceName(split[8])
}

// ConvertToClusterResourceGroup converts infra resource group to aks cluster ID
func ConvertToClusterResourceGroup(nodeResourceID string) string {
	subscription, resourceGroup, _ := ParseResourceID(nodeResourceID)
	split := strings.Split(string(resourceGroup), "_")
	return fmt.Sprintf("/subscriptions/%s/resourcegroups/%s/providers/Microsoft.ContainerService/managedClusters/%s", subscription, split[1], split[2])
}