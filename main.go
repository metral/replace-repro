package main

import (
	"fmt"

	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	storagev1beta1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/storage/v1beta1"

	//storagev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/storage/v1"
	//metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Step 1: Run update to create CSI driver in an existing k8s cluster.
		csiDriver, err := storagev1beta1.NewCSIDriver(ctx, "my-csi-driver", &storagev1beta1.CSIDriverArgs{
			ApiVersion: pulumi.String("storage.k8s.io/v1beta1"),
			Kind:       pulumi.String("CSIDriver"),
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("my-csi-driver.example.com"),
			},
			Spec: &storagev1beta1.CSIDriverSpecArgs{
				PodInfoOnMount: pulumi.Bool(true),
				VolumeLifecycleModes: pulumi.StringArray{
					pulumi.String("Ephemeral"),
				},
			},
		}, pulumi.ReplaceOnChanges([]string{"*"}), pulumi.DeleteBeforeReplace(true))
		if err != nil {
			return fmt.Errorf("CSI Driver: %q", err)
		}

		/*
			// Step 2. Uncomment this code and the storagev1 import, and comment out the code in Step 1.
			// This code simply moves from v1beta1 -> v1 of the CSI drive API resource.
			csiDriver, err := storagev1.NewCSIDriver(ctx, "my-csi-driver", &storagev1.CSIDriverArgs{
				ApiVersion: pulumi.String("storage.k8s.io/v1"),
				Kind:       pulumi.String("CSIDriver"),
				Metadata: &metav1.ObjectMetaArgs{
					Name: pulumi.String("my-csi-driver.example.com"),
				},
				Spec: &storagev1.CSIDriverSpecArgs{
					PodInfoOnMount: pulumi.Bool(true),
					VolumeLifecycleModes: pulumi.StringArray{
						pulumi.String("Ephemeral"),
					},
				},
			}, pulumi.ReplaceOnChanges([]string{"*"}), pulumi.DeleteBeforeReplace(true))
			if err != nil {
				return fmt.Errorf("CSI Driver: %q", err)
			}
		*/

		csiDriverName := csiDriver.Metadata.Name().ToStringPtrOutput().ApplyT(func(v *string) string { return *v }).(pulumi.StringOutput)
		ctx.Export("csiDriverName", csiDriverName)
		return nil
	})
}
