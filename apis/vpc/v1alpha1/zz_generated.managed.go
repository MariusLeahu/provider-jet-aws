/*
Copyright 2021 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this Vpc.
func (mg *Vpc) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Vpc.
func (mg *Vpc) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Vpc.
func (mg *Vpc) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Vpc.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Vpc) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Vpc.
func (mg *Vpc) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Vpc.
func (mg *Vpc) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Vpc.
func (mg *Vpc) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Vpc.
func (mg *Vpc) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Vpc.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Vpc) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Vpc.
func (mg *Vpc) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcDhcpOptions.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcDhcpOptions) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcDhcpOptions.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcDhcpOptions) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcDhcpOptions.
func (mg *VpcDhcpOptions) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcDhcpOptionsAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcDhcpOptionsAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcDhcpOptionsAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcDhcpOptionsAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcDhcpOptionsAssociation.
func (mg *VpcDhcpOptionsAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcEndpoint.
func (mg *VpcEndpoint) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpoint.
func (mg *VpcEndpoint) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpoint.
func (mg *VpcEndpoint) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpoint.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpoint) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpoint.
func (mg *VpcEndpoint) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpoint.
func (mg *VpcEndpoint) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpoint.
func (mg *VpcEndpoint) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpoint.
func (mg *VpcEndpoint) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpoint.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpoint) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpoint.
func (mg *VpcEndpoint) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointConnectionNotification.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointConnectionNotification) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointConnectionNotification.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointConnectionNotification) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointConnectionNotification.
func (mg *VpcEndpointConnectionNotification) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointRouteTableAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointRouteTableAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointRouteTableAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointRouteTableAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointRouteTableAssociation.
func (mg *VpcEndpointRouteTableAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcEndpointService.
func (mg *VpcEndpointService) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointService.
func (mg *VpcEndpointService) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointService.
func (mg *VpcEndpointService) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointService.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointService) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointService.
func (mg *VpcEndpointService) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointService.
func (mg *VpcEndpointService) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointService.
func (mg *VpcEndpointService) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointService.
func (mg *VpcEndpointService) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointService.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointService) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointService.
func (mg *VpcEndpointService) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointServiceAllowedPrincipal.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointServiceAllowedPrincipal) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointServiceAllowedPrincipal.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointServiceAllowedPrincipal) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointServiceAllowedPrincipal.
func (mg *VpcEndpointServiceAllowedPrincipal) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcEndpointSubnetAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcEndpointSubnetAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcEndpointSubnetAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcEndpointSubnetAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcEndpointSubnetAssociation.
func (mg *VpcEndpointSubnetAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcIpv4CidrBlockAssociation.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcIpv4CidrBlockAssociation) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcIpv4CidrBlockAssociation.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcIpv4CidrBlockAssociation) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcIpv4CidrBlockAssociation.
func (mg *VpcIpv4CidrBlockAssociation) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcPeeringConnection.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcPeeringConnection) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcPeeringConnection.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcPeeringConnection) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VpcPeeringConnectionOptions.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VpcPeeringConnectionOptions) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VpcPeeringConnectionOptions.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VpcPeeringConnectionOptions) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this VpcPeeringConnectionOptions.
func (mg *VpcPeeringConnectionOptions) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
