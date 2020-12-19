// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AdminStatus string

// Enum values for AdminStatus
const (
	AdminStatusEnabled             AdminStatus = "ENABLED"
	AdminStatusDisablingInProgress AdminStatus = "DISABLING_IN_PROGRESS"
)

// Values returns all known values for AdminStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AdminStatus) Values() []AdminStatus {
	return []AdminStatus{
		"ENABLED",
		"DISABLING_IN_PROGRESS",
	}
}

type Currency string

// Enum values for Currency
const (
	CurrencyUsd Currency = "USD"
)

// Values returns all known values for Currency. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Currency) Values() []Currency {
	return []Currency{
		"USD",
	}
}

type DayOfWeek string

// Enum values for DayOfWeek
const (
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
)

// Values returns all known values for DayOfWeek. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (DayOfWeek) Values() []DayOfWeek {
	return []DayOfWeek{
		"SUNDAY",
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
	}
}

type EffectivePermission string

// Enum values for EffectivePermission
const (
	EffectivePermissionPublic    EffectivePermission = "PUBLIC"
	EffectivePermissionNotPublic EffectivePermission = "NOT_PUBLIC"
	EffectivePermissionUnknown   EffectivePermission = "UNKNOWN"
)

// Values returns all known values for EffectivePermission. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EffectivePermission) Values() []EffectivePermission {
	return []EffectivePermission{
		"PUBLIC",
		"NOT_PUBLIC",
		"UNKNOWN",
	}
}

type EncryptionType string

// Enum values for EncryptionType
const (
	EncryptionTypeNone    EncryptionType = "NONE"
	EncryptionTypeAes256  EncryptionType = "AES256"
	EncryptionTypeAwsKms  EncryptionType = "aws:kms"
	EncryptionTypeUnknown EncryptionType = "UNKNOWN"
)

// Values returns all known values for EncryptionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncryptionType) Values() []EncryptionType {
	return []EncryptionType{
		"NONE",
		"AES256",
		"aws:kms",
		"UNKNOWN",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeClienterror   ErrorCode = "ClientError"
	ErrorCodeInternalerror ErrorCode = "InternalError"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"ClientError",
		"InternalError",
	}
}

type FindingActionType string

// Enum values for FindingActionType
const (
	FindingActionTypeAwsApiCall FindingActionType = "AWS_API_CALL"
)

// Values returns all known values for FindingActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingActionType) Values() []FindingActionType {
	return []FindingActionType{
		"AWS_API_CALL",
	}
}

type FindingCategory string

// Enum values for FindingCategory
const (
	FindingCategoryClassification FindingCategory = "CLASSIFICATION"
	FindingCategoryPolicy         FindingCategory = "POLICY"
)

// Values returns all known values for FindingCategory. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingCategory) Values() []FindingCategory {
	return []FindingCategory{
		"CLASSIFICATION",
		"POLICY",
	}
}

type FindingPublishingFrequency string

// Enum values for FindingPublishingFrequency
const (
	FindingPublishingFrequencyFifteenMinutes FindingPublishingFrequency = "FIFTEEN_MINUTES"
	FindingPublishingFrequencyOneHour        FindingPublishingFrequency = "ONE_HOUR"
	FindingPublishingFrequencySixHours       FindingPublishingFrequency = "SIX_HOURS"
)

// Values returns all known values for FindingPublishingFrequency. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (FindingPublishingFrequency) Values() []FindingPublishingFrequency {
	return []FindingPublishingFrequency{
		"FIFTEEN_MINUTES",
		"ONE_HOUR",
		"SIX_HOURS",
	}
}

type FindingsFilterAction string

// Enum values for FindingsFilterAction
const (
	FindingsFilterActionArchive FindingsFilterAction = "ARCHIVE"
	FindingsFilterActionNoop    FindingsFilterAction = "NOOP"
)

// Values returns all known values for FindingsFilterAction. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FindingsFilterAction) Values() []FindingsFilterAction {
	return []FindingsFilterAction{
		"ARCHIVE",
		"NOOP",
	}
}

type FindingStatisticsSortAttributeName string

// Enum values for FindingStatisticsSortAttributeName
const (
	FindingStatisticsSortAttributeNameGroupkey FindingStatisticsSortAttributeName = "groupKey"
	FindingStatisticsSortAttributeNameCount    FindingStatisticsSortAttributeName = "count"
)

// Values returns all known values for FindingStatisticsSortAttributeName. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (FindingStatisticsSortAttributeName) Values() []FindingStatisticsSortAttributeName {
	return []FindingStatisticsSortAttributeName{
		"groupKey",
		"count",
	}
}

type FindingType string

// Enum values for FindingType
const (
	FindingTypeSensitivedataS3objectMultiple             FindingType = "SensitiveData:S3Object/Multiple"
	FindingTypeSensitivedataS3objectFinancial            FindingType = "SensitiveData:S3Object/Financial"
	FindingTypeSensitivedataS3objectPersonal             FindingType = "SensitiveData:S3Object/Personal"
	FindingTypeSensitivedataS3objectCredentials          FindingType = "SensitiveData:S3Object/Credentials"
	FindingTypeSensitivedataS3objectCustomidentifier     FindingType = "SensitiveData:S3Object/CustomIdentifier"
	FindingTypePolicyIamuserS3bucketpublic               FindingType = "Policy:IAMUser/S3BucketPublic"
	FindingTypePolicyIamuserS3bucketsharedexternally     FindingType = "Policy:IAMUser/S3BucketSharedExternally"
	FindingTypePolicyIamuserS3bucketreplicatedexternally FindingType = "Policy:IAMUser/S3BucketReplicatedExternally"
	FindingTypePolicyIamuserS3bucketencryptiondisabled   FindingType = "Policy:IAMUser/S3BucketEncryptionDisabled"
	FindingTypePolicyIamuserS3blockpublicaccessdisabled  FindingType = "Policy:IAMUser/S3BlockPublicAccessDisabled"
)

// Values returns all known values for FindingType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FindingType) Values() []FindingType {
	return []FindingType{
		"SensitiveData:S3Object/Multiple",
		"SensitiveData:S3Object/Financial",
		"SensitiveData:S3Object/Personal",
		"SensitiveData:S3Object/Credentials",
		"SensitiveData:S3Object/CustomIdentifier",
		"Policy:IAMUser/S3BucketPublic",
		"Policy:IAMUser/S3BucketSharedExternally",
		"Policy:IAMUser/S3BucketReplicatedExternally",
		"Policy:IAMUser/S3BucketEncryptionDisabled",
		"Policy:IAMUser/S3BlockPublicAccessDisabled",
	}
}

type GroupBy string

// Enum values for GroupBy
const (
	GroupByResourcesaffectedS3bucketName GroupBy = "resourcesAffected.s3Bucket.name"
	GroupByType                          GroupBy = "type"
	GroupByClassificationdetailsJobid    GroupBy = "classificationDetails.jobId"
	GroupBySeverityDescription           GroupBy = "severity.description"
)

// Values returns all known values for GroupBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (GroupBy) Values() []GroupBy {
	return []GroupBy{
		"resourcesAffected.s3Bucket.name",
		"type",
		"classificationDetails.jobId",
		"severity.description",
	}
}

type IsDefinedInJob string

// Enum values for IsDefinedInJob
const (
	IsDefinedInJobTrue    IsDefinedInJob = "TRUE"
	IsDefinedInJobFalse   IsDefinedInJob = "FALSE"
	IsDefinedInJobUnknown IsDefinedInJob = "UNKNOWN"
)

// Values returns all known values for IsDefinedInJob. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IsDefinedInJob) Values() []IsDefinedInJob {
	return []IsDefinedInJob{
		"TRUE",
		"FALSE",
		"UNKNOWN",
	}
}

type IsMonitoredByJob string

// Enum values for IsMonitoredByJob
const (
	IsMonitoredByJobTrue    IsMonitoredByJob = "TRUE"
	IsMonitoredByJobFalse   IsMonitoredByJob = "FALSE"
	IsMonitoredByJobUnknown IsMonitoredByJob = "UNKNOWN"
)

// Values returns all known values for IsMonitoredByJob. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IsMonitoredByJob) Values() []IsMonitoredByJob {
	return []IsMonitoredByJob{
		"TRUE",
		"FALSE",
		"UNKNOWN",
	}
}

type JobComparator string

// Enum values for JobComparator
const (
	JobComparatorEq       JobComparator = "EQ"
	JobComparatorGt       JobComparator = "GT"
	JobComparatorGte      JobComparator = "GTE"
	JobComparatorLt       JobComparator = "LT"
	JobComparatorLte      JobComparator = "LTE"
	JobComparatorNe       JobComparator = "NE"
	JobComparatorContains JobComparator = "CONTAINS"
)

// Values returns all known values for JobComparator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (JobComparator) Values() []JobComparator {
	return []JobComparator{
		"EQ",
		"GT",
		"GTE",
		"LT",
		"LTE",
		"NE",
		"CONTAINS",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusRunning    JobStatus = "RUNNING"
	JobStatusPaused     JobStatus = "PAUSED"
	JobStatusCancelled  JobStatus = "CANCELLED"
	JobStatusComplete   JobStatus = "COMPLETE"
	JobStatusIdle       JobStatus = "IDLE"
	JobStatusUserPaused JobStatus = "USER_PAUSED"
)

// Values returns all known values for JobStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"RUNNING",
		"PAUSED",
		"CANCELLED",
		"COMPLETE",
		"IDLE",
		"USER_PAUSED",
	}
}

type JobType string

// Enum values for JobType
const (
	JobTypeOneTime   JobType = "ONE_TIME"
	JobTypeScheduled JobType = "SCHEDULED"
)

// Values returns all known values for JobType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (JobType) Values() []JobType {
	return []JobType{
		"ONE_TIME",
		"SCHEDULED",
	}
}

type LastRunErrorStatusCode string

// Enum values for LastRunErrorStatusCode
const (
	LastRunErrorStatusCodeNone  LastRunErrorStatusCode = "NONE"
	LastRunErrorStatusCodeError LastRunErrorStatusCode = "ERROR"
)

// Values returns all known values for LastRunErrorStatusCode. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (LastRunErrorStatusCode) Values() []LastRunErrorStatusCode {
	return []LastRunErrorStatusCode{
		"NONE",
		"ERROR",
	}
}

type ListJobsFilterKey string

// Enum values for ListJobsFilterKey
const (
	ListJobsFilterKeyJobtype   ListJobsFilterKey = "jobType"
	ListJobsFilterKeyJobstatus ListJobsFilterKey = "jobStatus"
	ListJobsFilterKeyCreatedat ListJobsFilterKey = "createdAt"
	ListJobsFilterKeyName      ListJobsFilterKey = "name"
)

// Values returns all known values for ListJobsFilterKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ListJobsFilterKey) Values() []ListJobsFilterKey {
	return []ListJobsFilterKey{
		"jobType",
		"jobStatus",
		"createdAt",
		"name",
	}
}

type ListJobsSortAttributeName string

// Enum values for ListJobsSortAttributeName
const (
	ListJobsSortAttributeNameCreatedat ListJobsSortAttributeName = "createdAt"
	ListJobsSortAttributeNameJobstatus ListJobsSortAttributeName = "jobStatus"
	ListJobsSortAttributeNameName      ListJobsSortAttributeName = "name"
	ListJobsSortAttributeNameJobtype   ListJobsSortAttributeName = "jobType"
)

// Values returns all known values for ListJobsSortAttributeName. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ListJobsSortAttributeName) Values() []ListJobsSortAttributeName {
	return []ListJobsSortAttributeName{
		"createdAt",
		"jobStatus",
		"name",
		"jobType",
	}
}

type MacieStatus string

// Enum values for MacieStatus
const (
	MacieStatusPaused  MacieStatus = "PAUSED"
	MacieStatusEnabled MacieStatus = "ENABLED"
)

// Values returns all known values for MacieStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (MacieStatus) Values() []MacieStatus {
	return []MacieStatus{
		"PAUSED",
		"ENABLED",
	}
}

type OrderBy string

// Enum values for OrderBy
const (
	OrderByAsc  OrderBy = "ASC"
	OrderByDesc OrderBy = "DESC"
)

// Values returns all known values for OrderBy. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (OrderBy) Values() []OrderBy {
	return []OrderBy{
		"ASC",
		"DESC",
	}
}

type RelationshipStatus string

// Enum values for RelationshipStatus
const (
	RelationshipStatusEnabled                     RelationshipStatus = "Enabled"
	RelationshipStatusPaused                      RelationshipStatus = "Paused"
	RelationshipStatusInvited                     RelationshipStatus = "Invited"
	RelationshipStatusCreated                     RelationshipStatus = "Created"
	RelationshipStatusRemoved                     RelationshipStatus = "Removed"
	RelationshipStatusResigned                    RelationshipStatus = "Resigned"
	RelationshipStatusEmailverificationinprogress RelationshipStatus = "EmailVerificationInProgress"
	RelationshipStatusEmailverificationfailed     RelationshipStatus = "EmailVerificationFailed"
	RelationshipStatusRegiondisabled              RelationshipStatus = "RegionDisabled"
	RelationshipStatusAccountsuspended            RelationshipStatus = "AccountSuspended"
)

// Values returns all known values for RelationshipStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RelationshipStatus) Values() []RelationshipStatus {
	return []RelationshipStatus{
		"Enabled",
		"Paused",
		"Invited",
		"Created",
		"Removed",
		"Resigned",
		"EmailVerificationInProgress",
		"EmailVerificationFailed",
		"RegionDisabled",
		"AccountSuspended",
	}
}

type ScopeFilterKey string

// Enum values for ScopeFilterKey
const (
	ScopeFilterKeyBucketCreationDate     ScopeFilterKey = "BUCKET_CREATION_DATE"
	ScopeFilterKeyObjectExtension        ScopeFilterKey = "OBJECT_EXTENSION"
	ScopeFilterKeyObjectLastModifiedDate ScopeFilterKey = "OBJECT_LAST_MODIFIED_DATE"
	ScopeFilterKeyObjectSize             ScopeFilterKey = "OBJECT_SIZE"
	ScopeFilterKeyTag                    ScopeFilterKey = "TAG"
)

// Values returns all known values for ScopeFilterKey. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ScopeFilterKey) Values() []ScopeFilterKey {
	return []ScopeFilterKey{
		"BUCKET_CREATION_DATE",
		"OBJECT_EXTENSION",
		"OBJECT_LAST_MODIFIED_DATE",
		"OBJECT_SIZE",
		"TAG",
	}
}

type SensitiveDataItemCategory string

// Enum values for SensitiveDataItemCategory
const (
	SensitiveDataItemCategoryFinancialInformation SensitiveDataItemCategory = "FINANCIAL_INFORMATION"
	SensitiveDataItemCategoryPersonalInformation  SensitiveDataItemCategory = "PERSONAL_INFORMATION"
	SensitiveDataItemCategoryCredentials          SensitiveDataItemCategory = "CREDENTIALS"
	SensitiveDataItemCategoryCustomIdentifier     SensitiveDataItemCategory = "CUSTOM_IDENTIFIER"
)

// Values returns all known values for SensitiveDataItemCategory. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (SensitiveDataItemCategory) Values() []SensitiveDataItemCategory {
	return []SensitiveDataItemCategory{
		"FINANCIAL_INFORMATION",
		"PERSONAL_INFORMATION",
		"CREDENTIALS",
		"CUSTOM_IDENTIFIER",
	}
}

type SeverityDescription string

// Enum values for SeverityDescription
const (
	SeverityDescriptionLow    SeverityDescription = "Low"
	SeverityDescriptionMedium SeverityDescription = "Medium"
	SeverityDescriptionHigh   SeverityDescription = "High"
)

// Values returns all known values for SeverityDescription. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SeverityDescription) Values() []SeverityDescription {
	return []SeverityDescription{
		"Low",
		"Medium",
		"High",
	}
}

type SharedAccess string

// Enum values for SharedAccess
const (
	SharedAccessExternal  SharedAccess = "EXTERNAL"
	SharedAccessInternal  SharedAccess = "INTERNAL"
	SharedAccessNotShared SharedAccess = "NOT_SHARED"
	SharedAccessUnknown   SharedAccess = "UNKNOWN"
)

// Values returns all known values for SharedAccess. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SharedAccess) Values() []SharedAccess {
	return []SharedAccess{
		"EXTERNAL",
		"INTERNAL",
		"NOT_SHARED",
		"UNKNOWN",
	}
}

type StorageClass string

// Enum values for StorageClass
const (
	StorageClassStandard           StorageClass = "STANDARD"
	StorageClassReducedRedundancy  StorageClass = "REDUCED_REDUNDANCY"
	StorageClassStandardIa         StorageClass = "STANDARD_IA"
	StorageClassIntelligentTiering StorageClass = "INTELLIGENT_TIERING"
	StorageClassDeepArchive        StorageClass = "DEEP_ARCHIVE"
	StorageClassOnezoneIa          StorageClass = "ONEZONE_IA"
	StorageClassGlacier            StorageClass = "GLACIER"
)

// Values returns all known values for StorageClass. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (StorageClass) Values() []StorageClass {
	return []StorageClass{
		"STANDARD",
		"REDUCED_REDUNDANCY",
		"STANDARD_IA",
		"INTELLIGENT_TIERING",
		"DEEP_ARCHIVE",
		"ONEZONE_IA",
		"GLACIER",
	}
}

type TagTarget string

// Enum values for TagTarget
const (
	TagTargetS3Object TagTarget = "S3_OBJECT"
)

// Values returns all known values for TagTarget. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (TagTarget) Values() []TagTarget {
	return []TagTarget{
		"S3_OBJECT",
	}
}

type Unit string

// Enum values for Unit
const (
	UnitTerabytes Unit = "TERABYTES"
)

// Values returns all known values for Unit. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Unit) Values() []Unit {
	return []Unit{
		"TERABYTES",
	}
}

type UsageStatisticsFilterComparator string

// Enum values for UsageStatisticsFilterComparator
const (
	UsageStatisticsFilterComparatorGt       UsageStatisticsFilterComparator = "GT"
	UsageStatisticsFilterComparatorGte      UsageStatisticsFilterComparator = "GTE"
	UsageStatisticsFilterComparatorLt       UsageStatisticsFilterComparator = "LT"
	UsageStatisticsFilterComparatorLte      UsageStatisticsFilterComparator = "LTE"
	UsageStatisticsFilterComparatorEq       UsageStatisticsFilterComparator = "EQ"
	UsageStatisticsFilterComparatorNe       UsageStatisticsFilterComparator = "NE"
	UsageStatisticsFilterComparatorContains UsageStatisticsFilterComparator = "CONTAINS"
)

// Values returns all known values for UsageStatisticsFilterComparator. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (UsageStatisticsFilterComparator) Values() []UsageStatisticsFilterComparator {
	return []UsageStatisticsFilterComparator{
		"GT",
		"GTE",
		"LT",
		"LTE",
		"EQ",
		"NE",
		"CONTAINS",
	}
}

type UsageStatisticsFilterKey string

// Enum values for UsageStatisticsFilterKey
const (
	UsageStatisticsFilterKeyAccountid          UsageStatisticsFilterKey = "accountId"
	UsageStatisticsFilterKeyServicelimit       UsageStatisticsFilterKey = "serviceLimit"
	UsageStatisticsFilterKeyFreetrialstartdate UsageStatisticsFilterKey = "freeTrialStartDate"
	UsageStatisticsFilterKeyTotal              UsageStatisticsFilterKey = "total"
)

// Values returns all known values for UsageStatisticsFilterKey. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageStatisticsFilterKey) Values() []UsageStatisticsFilterKey {
	return []UsageStatisticsFilterKey{
		"accountId",
		"serviceLimit",
		"freeTrialStartDate",
		"total",
	}
}

type UsageStatisticsSortKey string

// Enum values for UsageStatisticsSortKey
const (
	UsageStatisticsSortKeyAccountid          UsageStatisticsSortKey = "accountId"
	UsageStatisticsSortKeyTotal              UsageStatisticsSortKey = "total"
	UsageStatisticsSortKeyServicelimitvalue  UsageStatisticsSortKey = "serviceLimitValue"
	UsageStatisticsSortKeyFreetrialstartdate UsageStatisticsSortKey = "freeTrialStartDate"
)

// Values returns all known values for UsageStatisticsSortKey. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UsageStatisticsSortKey) Values() []UsageStatisticsSortKey {
	return []UsageStatisticsSortKey{
		"accountId",
		"total",
		"serviceLimitValue",
		"freeTrialStartDate",
	}
}

type UsageType string

// Enum values for UsageType
const (
	UsageTypeDataInventoryEvaluation UsageType = "DATA_INVENTORY_EVALUATION"
	UsageTypeSensitiveDataDiscovery  UsageType = "SENSITIVE_DATA_DISCOVERY"
)

// Values returns all known values for UsageType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (UsageType) Values() []UsageType {
	return []UsageType{
		"DATA_INVENTORY_EVALUATION",
		"SENSITIVE_DATA_DISCOVERY",
	}
}

type UserIdentityType string

// Enum values for UserIdentityType
const (
	UserIdentityTypeAssumedrole   UserIdentityType = "AssumedRole"
	UserIdentityTypeIamuser       UserIdentityType = "IAMUser"
	UserIdentityTypeFederateduser UserIdentityType = "FederatedUser"
	UserIdentityTypeRoot          UserIdentityType = "Root"
	UserIdentityTypeAwsaccount    UserIdentityType = "AWSAccount"
	UserIdentityTypeAwsservice    UserIdentityType = "AWSService"
)

// Values returns all known values for UserIdentityType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (UserIdentityType) Values() []UserIdentityType {
	return []UserIdentityType{
		"AssumedRole",
		"IAMUser",
		"FederatedUser",
		"Root",
		"AWSAccount",
		"AWSService",
	}
}
