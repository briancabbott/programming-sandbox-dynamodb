// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// Another user submitted a request to create, update, or delete the object at the
// same time that you did. Retry the request.
type ConcurrentModification struct {
	Message *string
}

func (e *ConcurrentModification) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConcurrentModification) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConcurrentModification) ErrorCode() string             { return "ConcurrentModification" }
func (e *ConcurrentModification) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The cause of this error depends on the operation that you're performing:
//
// *
// Create a public hosted zone: Two hosted zones that have the same name or that
// have a parent/child relationship (example.com and test.example.com) can't have
// any common name servers. You tried to create a hosted zone that has the same
// name as an existing hosted zone or that's the parent or child of an existing
// hosted zone, and you specified a delegation set that shares one or more name
// servers with the existing hosted zone. For more information, see
// CreateReusableDelegationSet
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateReusableDelegationSet.html).
//
// *
// Create a private hosted zone: A hosted zone with the specified name already
// exists and is already associated with the Amazon VPC that you specified.
//
// *
// Associate VPCs with a private hosted zone: The VPC that you specified is already
// associated with another hosted zone that has the same name.
type ConflictingDomainExists struct {
	Message *string
}

func (e *ConflictingDomainExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictingDomainExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictingDomainExists) ErrorCode() string             { return "ConflictingDomainExists" }
func (e *ConflictingDomainExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You tried to update a traffic policy instance by using a traffic policy version
// that has a different DNS type than the current type for the instance. You
// specified the type in the JSON document in the CreateTrafficPolicy or
// CreateTrafficPolicyVersionrequest.
type ConflictingTypes struct {
	Message *string
}

func (e *ConflictingTypes) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictingTypes) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictingTypes) ErrorCode() string             { return "ConflictingTypes" }
func (e *ConflictingTypes) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A delegation set with the same owner and caller reference combination has
// already been created.
type DelegationSetAlreadyCreated struct {
	Message *string
}

func (e *DelegationSetAlreadyCreated) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DelegationSetAlreadyCreated) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DelegationSetAlreadyCreated) ErrorCode() string             { return "DelegationSetAlreadyCreated" }
func (e *DelegationSetAlreadyCreated) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified delegation set has already been marked as reusable.
type DelegationSetAlreadyReusable struct {
	Message *string
}

func (e *DelegationSetAlreadyReusable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DelegationSetAlreadyReusable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DelegationSetAlreadyReusable) ErrorCode() string             { return "DelegationSetAlreadyReusable" }
func (e *DelegationSetAlreadyReusable) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified delegation contains associated hosted zones which must be deleted
// before the reusable delegation set can be deleted.
type DelegationSetInUse struct {
	Message *string
}

func (e *DelegationSetInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DelegationSetInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DelegationSetInUse) ErrorCode() string             { return "DelegationSetInUse" }
func (e *DelegationSetInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can create a hosted zone that has the same name as an existing hosted zone
// (example.com is common), but there is a limit to the number of hosted zones that
// have the same name. If you get this error, Amazon Route 53 has reached that
// limit. If you own the domain name and Route 53 generates this error, contact
// Customer Support.
type DelegationSetNotAvailable struct {
	Message *string
}

func (e *DelegationSetNotAvailable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DelegationSetNotAvailable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DelegationSetNotAvailable) ErrorCode() string             { return "DelegationSetNotAvailable" }
func (e *DelegationSetNotAvailable) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A reusable delegation set with the specified ID does not exist.
type DelegationSetNotReusable struct {
	Message *string
}

func (e *DelegationSetNotReusable) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DelegationSetNotReusable) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DelegationSetNotReusable) ErrorCode() string             { return "DelegationSetNotReusable" }
func (e *DelegationSetNotReusable) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The health check you're attempting to create already exists. Amazon Route 53
// returns this error when you submit a request that has the following values:
//
// *
// The same value for CallerReference as an existing health check, and one or more
// values that differ from the existing health check that has the same caller
// reference.
//
// * The same value for CallerReference as a health check that you
// created and later deleted, regardless of the other settings in the request.
type HealthCheckAlreadyExists struct {
	Message *string
}

func (e *HealthCheckAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HealthCheckAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HealthCheckAlreadyExists) ErrorCode() string             { return "HealthCheckAlreadyExists" }
func (e *HealthCheckAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This error code is not in use.
type HealthCheckInUse struct {
	Message *string
}

func (e *HealthCheckInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HealthCheckInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HealthCheckInUse) ErrorCode() string             { return "HealthCheckInUse" }
func (e *HealthCheckInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value of HealthCheckVersion in the request doesn't match the value of
// HealthCheckVersion in the health check.
type HealthCheckVersionMismatch struct {
	Message *string
}

func (e *HealthCheckVersionMismatch) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HealthCheckVersionMismatch) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HealthCheckVersionMismatch) ErrorCode() string             { return "HealthCheckVersionMismatch" }
func (e *HealthCheckVersionMismatch) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The hosted zone you're trying to create already exists. Amazon Route 53 returns
// this error when a hosted zone has already been created with the specified
// CallerReference.
type HostedZoneAlreadyExists struct {
	Message *string
}

func (e *HostedZoneAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HostedZoneAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HostedZoneAlreadyExists) ErrorCode() string             { return "HostedZoneAlreadyExists" }
func (e *HostedZoneAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The hosted zone contains resource records that are not SOA or NS records.
type HostedZoneNotEmpty struct {
	Message *string
}

func (e *HostedZoneNotEmpty) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HostedZoneNotEmpty) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HostedZoneNotEmpty) ErrorCode() string             { return "HostedZoneNotEmpty" }
func (e *HostedZoneNotEmpty) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified HostedZone can't be found.
type HostedZoneNotFound struct {
	Message *string
}

func (e *HostedZoneNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HostedZoneNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HostedZoneNotFound) ErrorCode() string             { return "HostedZoneNotFound" }
func (e *HostedZoneNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified hosted zone is a public hosted zone, not a private hosted zone.
type HostedZoneNotPrivate struct {
	Message *string
}

func (e *HostedZoneNotPrivate) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HostedZoneNotPrivate) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HostedZoneNotPrivate) ErrorCode() string             { return "HostedZoneNotPrivate" }
func (e *HostedZoneNotPrivate) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource you're trying to access is unsupported on this Amazon Route 53
// endpoint.
type IncompatibleVersion struct {
	Message *string
}

func (e *IncompatibleVersion) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IncompatibleVersion) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IncompatibleVersion) ErrorCode() string             { return "IncompatibleVersion" }
func (e *IncompatibleVersion) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Route 53 doesn't have the permissions required to create log streams and
// send query logs to log streams. Possible causes include the following:
//
// * There
// is no resource policy that specifies the log group ARN in the value for
// Resource.
//
// * The resource policy that includes the log group ARN in the value
// for Resource doesn't have the necessary permissions.
//
// * The resource policy
// hasn't finished propagating yet.
type InsufficientCloudWatchLogsResourcePolicy struct {
	Message *string
}

func (e *InsufficientCloudWatchLogsResourcePolicy) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InsufficientCloudWatchLogsResourcePolicy) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InsufficientCloudWatchLogsResourcePolicy) ErrorCode() string {
	return "InsufficientCloudWatchLogsResourcePolicy"
}
func (e *InsufficientCloudWatchLogsResourcePolicy) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// Parameter name is invalid.
type InvalidArgument struct {
	Message *string
}

func (e *InvalidArgument) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgument) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgument) ErrorCode() string             { return "InvalidArgument" }
func (e *InvalidArgument) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This exception contains a list of messages that might contain one or more error
// messages. Each error message indicates one error in the change batch.
type InvalidChangeBatch struct {
	Message *string

	Messages []string
}

func (e *InvalidChangeBatch) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidChangeBatch) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidChangeBatch) ErrorCode() string             { return "InvalidChangeBatch" }
func (e *InvalidChangeBatch) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The specified domain name is not valid.
type InvalidDomainName struct {
	Message *string
}

func (e *InvalidDomainName) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidDomainName) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidDomainName) ErrorCode() string             { return "InvalidDomainName" }
func (e *InvalidDomainName) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The input is not valid.
type InvalidInput struct {
	Message *string
}

func (e *InvalidInput) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInput) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInput) ErrorCode() string             { return "InvalidInput" }
func (e *InvalidInput) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value that you specified to get the second or subsequent page of results is
// invalid.
type InvalidPaginationToken struct {
	Message *string
}

func (e *InvalidPaginationToken) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPaginationToken) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPaginationToken) ErrorCode() string             { return "InvalidPaginationToken" }
func (e *InvalidPaginationToken) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The format of the traffic policy document that you specified in the Document
// element is invalid.
type InvalidTrafficPolicyDocument struct {
	Message *string
}

func (e *InvalidTrafficPolicyDocument) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidTrafficPolicyDocument) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidTrafficPolicyDocument) ErrorCode() string             { return "InvalidTrafficPolicyDocument" }
func (e *InvalidTrafficPolicyDocument) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The VPC ID that you specified either isn't a valid ID or the current account is
// not authorized to access this VPC.
type InvalidVPCId struct {
	Message *string
}

func (e *InvalidVPCId) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidVPCId) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidVPCId) ErrorCode() string             { return "InvalidVPCId" }
func (e *InvalidVPCId) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The VPC that you're trying to disassociate from the private hosted zone is the
// last VPC that is associated with the hosted zone. Amazon Route 53 doesn't
// support disassociating the last VPC from a hosted zone.
type LastVPCAssociation struct {
	Message *string
}

func (e *LastVPCAssociation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LastVPCAssociation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LastVPCAssociation) ErrorCode() string             { return "LastVPCAssociation" }
func (e *LastVPCAssociation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This operation can't be completed either because the current account has reached
// the limit on reusable delegation sets that it can create or because you've
// reached the limit on the number of Amazon VPCs that you can associate with a
// private hosted zone. To get the current limit on the number of reusable
// delegation sets, see GetAccountLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
// To get the current limit on the number of Amazon VPCs that you can associate
// with a private hosted zone, see GetHostedZoneLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetHostedZoneLimit.html).
// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
// with the AWS Support Center.
type LimitsExceeded struct {
	Message *string
}

func (e *LimitsExceeded) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitsExceeded) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitsExceeded) ErrorCode() string             { return "LimitsExceeded" }
func (e *LimitsExceeded) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A change with the specified change ID does not exist.
type NoSuchChange struct {
	Message *string
}

func (e *NoSuchChange) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchChange) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchChange) ErrorCode() string             { return "NoSuchChange" }
func (e *NoSuchChange) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is no CloudWatch Logs log group with the specified ARN.
type NoSuchCloudWatchLogsLogGroup struct {
	Message *string
}

func (e *NoSuchCloudWatchLogsLogGroup) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchCloudWatchLogsLogGroup) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchCloudWatchLogsLogGroup) ErrorCode() string             { return "NoSuchCloudWatchLogsLogGroup" }
func (e *NoSuchCloudWatchLogsLogGroup) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// A reusable delegation set with the specified ID does not exist.
type NoSuchDelegationSet struct {
	Message *string
}

func (e *NoSuchDelegationSet) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchDelegationSet) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchDelegationSet) ErrorCode() string             { return "NoSuchDelegationSet" }
func (e *NoSuchDelegationSet) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Amazon Route 53 doesn't support the specified geographic location. For a list of
// supported geolocation codes, see the GeoLocation
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GeoLocation.html)
// data type.
type NoSuchGeoLocation struct {
	Message *string
}

func (e *NoSuchGeoLocation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchGeoLocation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchGeoLocation) ErrorCode() string             { return "NoSuchGeoLocation" }
func (e *NoSuchGeoLocation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No health check exists with the specified ID.
type NoSuchHealthCheck struct {
	Message *string
}

func (e *NoSuchHealthCheck) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchHealthCheck) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchHealthCheck) ErrorCode() string             { return "NoSuchHealthCheck" }
func (e *NoSuchHealthCheck) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No hosted zone exists with the ID that you specified.
type NoSuchHostedZone struct {
	Message *string
}

func (e *NoSuchHostedZone) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchHostedZone) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchHostedZone) ErrorCode() string             { return "NoSuchHostedZone" }
func (e *NoSuchHostedZone) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is no DNS query logging configuration with the specified ID.
type NoSuchQueryLoggingConfig struct {
	Message *string
}

func (e *NoSuchQueryLoggingConfig) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchQueryLoggingConfig) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchQueryLoggingConfig) ErrorCode() string             { return "NoSuchQueryLoggingConfig" }
func (e *NoSuchQueryLoggingConfig) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No traffic policy exists with the specified ID.
type NoSuchTrafficPolicy struct {
	Message *string
}

func (e *NoSuchTrafficPolicy) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchTrafficPolicy) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchTrafficPolicy) ErrorCode() string             { return "NoSuchTrafficPolicy" }
func (e *NoSuchTrafficPolicy) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// No traffic policy instance exists with the specified ID.
type NoSuchTrafficPolicyInstance struct {
	Message *string
}

func (e *NoSuchTrafficPolicyInstance) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NoSuchTrafficPolicyInstance) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NoSuchTrafficPolicyInstance) ErrorCode() string             { return "NoSuchTrafficPolicyInstance" }
func (e *NoSuchTrafficPolicyInstance) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Associating the specified VPC with the specified hosted zone has not been
// authorized.
type NotAuthorizedException struct {
	Message *string
}

func (e *NotAuthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotAuthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotAuthorizedException) ErrorCode() string             { return "NotAuthorizedException" }
func (e *NotAuthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// If Amazon Route 53 can't process a request before the next request arrives, it
// will reject subsequent requests for the same hosted zone and return an HTTP 400
// error (Bad request). If Route 53 returns this error repeatedly for the same
// request, we recommend that you wait, in intervals of increasing duration, before
// you try the request again.
type PriorRequestNotComplete struct {
	Message *string
}

func (e *PriorRequestNotComplete) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PriorRequestNotComplete) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PriorRequestNotComplete) ErrorCode() string             { return "PriorRequestNotComplete" }
func (e *PriorRequestNotComplete) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You're trying to associate a VPC with a public hosted zone. Amazon Route 53
// doesn't support associating a VPC with a public hosted zone.
type PublicZoneVPCAssociation struct {
	Message *string
}

func (e *PublicZoneVPCAssociation) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PublicZoneVPCAssociation) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PublicZoneVPCAssociation) ErrorCode() string             { return "PublicZoneVPCAssociation" }
func (e *PublicZoneVPCAssociation) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can create only one query logging configuration for a hosted zone, and a
// query logging configuration already exists for this hosted zone.
type QueryLoggingConfigAlreadyExists struct {
	Message *string
}

func (e *QueryLoggingConfigAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *QueryLoggingConfigAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *QueryLoggingConfigAlreadyExists) ErrorCode() string {
	return "QueryLoggingConfigAlreadyExists"
}
func (e *QueryLoggingConfigAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The limit on the number of requests per second was exceeded.
type ThrottlingException struct {
	Message *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This health check can't be created because the current account has reached the
// limit on the number of active health checks. For information about default
// limits, see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. For information about how to get the
// current limit for an account, see GetAccountLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
// with the AWS Support Center. You have reached the maximum number of active
// health checks for an AWS account. To request a higher limit, create a case
// (http://aws.amazon.com/route53-request) with the AWS Support Center.
type TooManyHealthChecks struct {
	Message *string
}

func (e *TooManyHealthChecks) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyHealthChecks) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyHealthChecks) ErrorCode() string             { return "TooManyHealthChecks" }
func (e *TooManyHealthChecks) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This operation can't be completed either because the current account has reached
// the limit on the number of hosted zones or because you've reached the limit on
// the number of hosted zones that can be associated with a reusable delegation
// set. For information about default limits, see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To get the current limit on hosted zones
// that can be created by an account, see GetAccountLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
// To get the current limit on hosted zones that can be associated with a reusable
// delegation set, see GetReusableDelegationSetLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetReusableDelegationSetLimit.html).
// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
// with the AWS Support Center.
type TooManyHostedZones struct {
	Message *string
}

func (e *TooManyHostedZones) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyHostedZones) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyHostedZones) ErrorCode() string             { return "TooManyHostedZones" }
func (e *TooManyHostedZones) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This traffic policy can't be created because the current account has reached the
// limit on the number of traffic policies. For information about default limits,
// see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. To get the current limit for an account,
// see GetAccountLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
// with the AWS Support Center.
type TooManyTrafficPolicies struct {
	Message *string
}

func (e *TooManyTrafficPolicies) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTrafficPolicies) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTrafficPolicies) ErrorCode() string             { return "TooManyTrafficPolicies" }
func (e *TooManyTrafficPolicies) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This traffic policy instance can't be created because the current account has
// reached the limit on the number of traffic policy instances. For information
// about default limits, see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide. For information about how to get the
// current limit for an account, see GetAccountLimit
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetAccountLimit.html).
// To request a higher limit, create a case (http://aws.amazon.com/route53-request)
// with the AWS Support Center.
type TooManyTrafficPolicyInstances struct {
	Message *string
}

func (e *TooManyTrafficPolicyInstances) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTrafficPolicyInstances) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTrafficPolicyInstances) ErrorCode() string             { return "TooManyTrafficPolicyInstances" }
func (e *TooManyTrafficPolicyInstances) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This traffic policy version can't be created because you've reached the limit of
// 1000 on the number of versions that you can create for the current traffic
// policy. To create more traffic policy versions, you can use GetTrafficPolicy
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetTrafficPolicy.html)
// to get the traffic policy document for a specified traffic policy version, and
// then use CreateTrafficPolicy
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateTrafficPolicy.html)
// to create a new traffic policy using the traffic policy document.
type TooManyTrafficPolicyVersionsForCurrentPolicy struct {
	Message *string
}

func (e *TooManyTrafficPolicyVersionsForCurrentPolicy) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTrafficPolicyVersionsForCurrentPolicy) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTrafficPolicyVersionsForCurrentPolicy) ErrorCode() string {
	return "TooManyTrafficPolicyVersionsForCurrentPolicy"
}
func (e *TooManyTrafficPolicyVersionsForCurrentPolicy) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// You've created the maximum number of authorizations that can be created for the
// specified hosted zone. To authorize another VPC to be associated with the hosted
// zone, submit a DeleteVPCAssociationAuthorization request to remove an existing
// authorization. To get a list of existing authorizations, submit a
// ListVPCAssociationAuthorizations request.
type TooManyVPCAssociationAuthorizations struct {
	Message *string
}

func (e *TooManyVPCAssociationAuthorizations) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyVPCAssociationAuthorizations) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyVPCAssociationAuthorizations) ErrorCode() string {
	return "TooManyVPCAssociationAuthorizations"
}
func (e *TooManyVPCAssociationAuthorizations) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// A traffic policy that has the same value for Name already exists.
type TrafficPolicyAlreadyExists struct {
	Message *string
}

func (e *TrafficPolicyAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrafficPolicyAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrafficPolicyAlreadyExists) ErrorCode() string             { return "TrafficPolicyAlreadyExists" }
func (e *TrafficPolicyAlreadyExists) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// There is already a traffic policy instance with the specified ID.
type TrafficPolicyInstanceAlreadyExists struct {
	Message *string
}

func (e *TrafficPolicyInstanceAlreadyExists) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrafficPolicyInstanceAlreadyExists) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrafficPolicyInstanceAlreadyExists) ErrorCode() string {
	return "TrafficPolicyInstanceAlreadyExists"
}
func (e *TrafficPolicyInstanceAlreadyExists) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// One or more traffic policy instances were created by using the specified traffic
// policy.
type TrafficPolicyInUse struct {
	Message *string
}

func (e *TrafficPolicyInUse) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TrafficPolicyInUse) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TrafficPolicyInUse) ErrorCode() string             { return "TrafficPolicyInUse" }
func (e *TrafficPolicyInUse) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The VPC that you specified is not authorized to be associated with the hosted
// zone.
type VPCAssociationAuthorizationNotFound struct {
	Message *string
}

func (e *VPCAssociationAuthorizationNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *VPCAssociationAuthorizationNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *VPCAssociationAuthorizationNotFound) ErrorCode() string {
	return "VPCAssociationAuthorizationNotFound"
}
func (e *VPCAssociationAuthorizationNotFound) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The specified VPC and hosted zone are not currently associated.
type VPCAssociationNotFound struct {
	Message *string
}

func (e *VPCAssociationNotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *VPCAssociationNotFound) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *VPCAssociationNotFound) ErrorCode() string             { return "VPCAssociationNotFound" }
func (e *VPCAssociationNotFound) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
