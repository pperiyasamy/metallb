// SPDX-License-Identifier:Apache-2.0

package config

import (
	"go.universe.tf/e2etest/pkg/pointer"
	metallbv1beta1 "go.universe.tf/metallb/api/v1beta1"
)

const BGP = "bgp"
const L2 = "layer2"

func BFDProfileWithDefaults(profile metallbv1beta1.BFDProfile, multiHop bool) metallbv1beta1.BFDProfile {
	res := metallbv1beta1.BFDProfile{}
	res.Name = profile.Name
	res.Spec.ReceiveInterval = valueWithDefault(profile.Spec.ReceiveInterval, 300)
	res.Spec.TransmitInterval = valueWithDefault(profile.Spec.TransmitInterval, 300)
	res.Spec.DetectMultiplier = valueWithDefault(profile.Spec.DetectMultiplier, 3)
	res.Spec.EchoInterval = valueWithDefault(profile.Spec.EchoInterval, 50)
	res.Spec.MinimumTTL = valueWithDefault(profile.Spec.MinimumTTL, 254)
	res.Spec.EchoMode = profile.Spec.EchoMode
	res.Spec.PassiveMode = profile.Spec.PassiveMode

	if multiHop {
		res.Spec.EchoMode = pointer.BoolPtr(false)
		res.Spec.EchoInterval = pointer.Uint32Ptr(50)
	}

	return res
}

func valueWithDefault(v *uint32, def uint32) *uint32 {
	if v != nil {
		return v
	}
	return &def
}
