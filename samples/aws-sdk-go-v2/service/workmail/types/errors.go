// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
)

// The directory is already in use by another WorkMail organization in the same
// account and Region.
type DirectoryInUseException struct {
	Message *string
}

func (e *DirectoryInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryInUseException) ErrorCode() string             { return "DirectoryInUseException" }
func (e *DirectoryInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The directory service doesn't recognize the credentials supplied by WorkMail.
type DirectoryServiceAuthenticationFailedException struct {
	Message *string
}

func (e *DirectoryServiceAuthenticationFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryServiceAuthenticationFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryServiceAuthenticationFailedException) ErrorCode() string {
	return "DirectoryServiceAuthenticationFailedException"
}
func (e *DirectoryServiceAuthenticationFailedException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}

// The directory is unavailable. It might be located in another Region or deleted.
type DirectoryUnavailableException struct {
	Message *string
}

func (e *DirectoryUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DirectoryUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DirectoryUnavailableException) ErrorCode() string             { return "DirectoryUnavailableException" }
func (e *DirectoryUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The email address that you're trying to assign is already created for a
// different user, group, or resource.
type EmailAddressInUseException struct {
	Message *string
}

func (e *EmailAddressInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EmailAddressInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EmailAddressInUseException) ErrorCode() string             { return "EmailAddressInUseException" }
func (e *EmailAddressInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user, group, or resource that you're trying to register is already
// registered.
type EntityAlreadyRegisteredException struct {
	Message *string
}

func (e *EntityAlreadyRegisteredException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityAlreadyRegisteredException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityAlreadyRegisteredException) ErrorCode() string {
	return "EntityAlreadyRegisteredException"
}
func (e *EntityAlreadyRegisteredException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The identifier supplied for the user, group, or resource does not exist in your
// organization.
type EntityNotFoundException struct {
	Message *string
}

func (e *EntityNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityNotFoundException) ErrorCode() string             { return "EntityNotFoundException" }
func (e *EntityNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You are performing an operation on a user, group, or resource that isn't in the
// expected state, such as trying to delete an active user.
type EntityStateException struct {
	Message *string
}

func (e *EntityStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *EntityStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *EntityStateException) ErrorCode() string             { return "EntityStateException" }
func (e *EntityStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The configuration for a resource isn't valid. A resource must either be able to
// auto-respond to requests or have at least one delegate associated that can do so
// on its behalf.
type InvalidConfigurationException struct {
	Message *string
}

func (e *InvalidConfigurationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidConfigurationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidConfigurationException) ErrorCode() string             { return "InvalidConfigurationException" }
func (e *InvalidConfigurationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// One or more of the input parameters don't match the service's restrictions.
type InvalidParameterException struct {
	Message *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The supplied password doesn't match the minimum security constraints, such as
// length or use of special characters.
type InvalidPasswordException struct {
	Message *string
}

func (e *InvalidPasswordException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidPasswordException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidPasswordException) ErrorCode() string             { return "InvalidPasswordException" }
func (e *InvalidPasswordException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The request exceeds the limit of the resource.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// For an email or alias to be created in Amazon WorkMail, the included domain must
// be defined in the organization.
type MailDomainNotFoundException struct {
	Message *string
}

func (e *MailDomainNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MailDomainNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MailDomainNotFoundException) ErrorCode() string             { return "MailDomainNotFoundException" }
func (e *MailDomainNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// After a domain has been added to the organization, it must be verified. The
// domain is not yet verified.
type MailDomainStateException struct {
	Message *string
}

func (e *MailDomainStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *MailDomainStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *MailDomainStateException) ErrorCode() string             { return "MailDomainStateException" }
func (e *MailDomainStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The user, group, or resource name isn't unique in Amazon WorkMail.
type NameAvailabilityException struct {
	Message *string
}

func (e *NameAvailabilityException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NameAvailabilityException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NameAvailabilityException) ErrorCode() string             { return "NameAvailabilityException" }
func (e *NameAvailabilityException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// An operation received a valid organization identifier that either doesn't belong
// or exist in the system.
type OrganizationNotFoundException struct {
	Message *string
}

func (e *OrganizationNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OrganizationNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OrganizationNotFoundException) ErrorCode() string             { return "OrganizationNotFoundException" }
func (e *OrganizationNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The organization must have a valid state to perform certain operations on the
// organization or its members.
type OrganizationStateException struct {
	Message *string
}

func (e *OrganizationStateException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OrganizationStateException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OrganizationStateException) ErrorCode() string             { return "OrganizationStateException" }
func (e *OrganizationStateException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// This user, group, or resource name is not allowed in Amazon WorkMail.
type ReservedNameException struct {
	Message *string
}

func (e *ReservedNameException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ReservedNameException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ReservedNameException) ErrorCode() string             { return "ReservedNameException" }
func (e *ReservedNameException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource cannot be found.
type ResourceNotFoundException struct {
	Message *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource can have up to 50 user-applied tags.
type TooManyTagsException struct {
	Message *string
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You can't perform a write operation against a read-only directory.
type UnsupportedOperationException struct {
	Message *string
}

func (e *UnsupportedOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedOperationException) ErrorCode() string             { return "UnsupportedOperationException" }
func (e *UnsupportedOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
