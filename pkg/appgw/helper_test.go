// -------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// --------------------------------------------------------------------------------------------

package appgw

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("test helpers", func() {
	Context("ensure ParseResourceID works as expected", func() {
		It("should parse appgw resourceId correctly", func() {
			subID := SubscriptionID("xxxx")
			resGp := ResourceGroup("yyyy")
			resName := ResourceName("zzzz")
			resourceID := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/publicIPAddresses/%s", subID, resGp, resName)
			outSubID, outResGp, outResName := ParseResourceID(resourceID)
			Expect(outSubID).To(Equal(subID))
			Expect(resGp).To(Equal(outResGp))
			Expect(resName).To(Equal(outResName))
		})
	})
})
