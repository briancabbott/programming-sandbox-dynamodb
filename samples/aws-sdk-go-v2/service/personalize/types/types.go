// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Describes a custom algorithm.
type Algorithm struct {

	// The Amazon Resource Name (ARN) of the algorithm.
	AlgorithmArn *string

	// The URI of the Docker container for the algorithm image.
	AlgorithmImage *AlgorithmImage

	// The date and time (in Unix time) that the algorithm was created.
	CreationDateTime *time.Time

	// Specifies the default hyperparameters, their ranges, and whether they are
	// tunable. A tunable hyperparameter can have its value determined during
	// hyperparameter optimization (HPO).
	DefaultHyperParameterRanges *DefaultHyperParameterRanges

	// Specifies the default hyperparameters.
	DefaultHyperParameters map[string]string

	// Specifies the default maximum number of training jobs and parallel training
	// jobs.
	DefaultResourceConfig map[string]string

	// The date and time (in Unix time) that the algorithm was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the algorithm.
	Name *string

	// The Amazon Resource Name (ARN) of the role.
	RoleArn *string

	// The training input mode.
	TrainingInputMode *string
}

// Describes an algorithm image.
type AlgorithmImage struct {

	// The URI of the Docker container for the algorithm image.
	//
	// This member is required.
	DockerURI *string

	// The name of the algorithm image.
	Name *string
}

// When the solution performs AutoML (performAutoML is true in CreateSolution),
// Amazon Personalize determines which recipe, from the specified list, optimizes
// the given metric. Amazon Personalize then uses that recipe for the solution.
type AutoMLConfig struct {

	// The metric to optimize.
	MetricName *string

	// The list of candidate recipes.
	RecipeList []string
}

// When the solution performs AutoML (performAutoML is true in CreateSolution),
// specifies the recipe that best optimized the specified metric.
type AutoMLResult struct {

	// The Amazon Resource Name (ARN) of the best recipe.
	BestRecipeArn *string
}

// Contains information on a batch inference job.
type BatchInferenceJob struct {

	// The Amazon Resource Name (ARN) of the batch inference job.
	BatchInferenceJobArn *string

	// A string to string map of the configuration details of a batch inference job.
	BatchInferenceJobConfig *BatchInferenceJobConfig

	// The time at which the batch inference job was created.
	CreationDateTime *time.Time

	// If the batch inference job failed, the reason for the failure.
	FailureReason *string

	// The ARN of the filter used on the batch inference job.
	FilterArn *string

	// The Amazon S3 path that leads to the input data used to generate the batch
	// inference job.
	JobInput *BatchInferenceJobInput

	// The name of the batch inference job.
	JobName *string

	// The Amazon S3 bucket that contains the output data generated by the batch
	// inference job.
	JobOutput *BatchInferenceJobOutput

	// The time at which the batch inference job was last updated.
	LastUpdatedDateTime *time.Time

	// The number of recommendations generated by the batch inference job. This number
	// includes the error messages generated for failed input records.
	NumResults *int32

	// The ARN of the Amazon Identity and Access Management (IAM) role that requested
	// the batch inference job.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the solution version from which the batch
	// inference job was created.
	SolutionVersionArn *string

	// The status of the batch inference job. The status is one of the following
	// values:
	//
	// * PENDING
	//
	// * IN PROGRESS
	//
	// * ACTIVE
	//
	// * CREATE FAILED
	Status *string
}

// The configuration details of a batch inference job.
type BatchInferenceJobConfig struct {

	// A string to string map specifying the inference hyperparameters you wish to use
	// for hyperparameter optimization. See customizing-solution-config-hpo.
	ItemExplorationConfig map[string]string
}

// The input configuration of a batch inference job.
type BatchInferenceJobInput struct {

	// The URI of the Amazon S3 location that contains your input data. The Amazon S3
	// bucket must be in the same region as the API endpoint you are calling.
	//
	// This member is required.
	S3DataSource *S3DataConfig
}

// The output configuration parameters of a batch inference job.
type BatchInferenceJobOutput struct {

	// Information on the Amazon S3 bucket in which the batch inference job's output is
	// stored.
	//
	// This member is required.
	S3DataDestination *S3DataConfig
}

// A truncated version of the BatchInferenceJob datatype. The
// ListBatchInferenceJobs operation returns a list of batch inference job
// summaries.
type BatchInferenceJobSummary struct {

	// The Amazon Resource Name (ARN) of the batch inference job.
	BatchInferenceJobArn *string

	// The time at which the batch inference job was created.
	CreationDateTime *time.Time

	// If the batch inference job failed, the reason for the failure.
	FailureReason *string

	// The name of the batch inference job.
	JobName *string

	// The time at which the batch inference job was last updated.
	LastUpdatedDateTime *time.Time

	// The ARN of the solution version used by the batch inference job.
	SolutionVersionArn *string

	// The status of the batch inference job. The status is one of the following
	// values:
	//
	// * PENDING
	//
	// * IN PROGRESS
	//
	// * ACTIVE
	//
	// * CREATE FAILED
	Status *string
}

// Describes a deployed solution version, otherwise known as a campaign. For more
// information on campaigns, see CreateCampaign.
type Campaign struct {

	// The Amazon Resource Name (ARN) of the campaign.
	CampaignArn *string

	// The configuration details of a campaign.
	CampaignConfig *CampaignConfig

	// The date and time (in Unix format) that the campaign was created.
	CreationDateTime *time.Time

	// If a campaign fails, the reason behind the failure.
	FailureReason *string

	// The date and time (in Unix format) that the campaign was last updated.
	LastUpdatedDateTime *time.Time

	// Provides a summary of the properties of a campaign update. For a complete
	// listing, call the DescribeCampaign API.
	LatestCampaignUpdate *CampaignUpdateSummary

	// Specifies the requested minimum provisioned transactions (recommendations) per
	// second.
	MinProvisionedTPS *int32

	// The name of the campaign.
	Name *string

	// The Amazon Resource Name (ARN) of a specific version of the solution.
	SolutionVersionArn *string

	// The status of the campaign. A campaign can be in one of the following states:
	//
	// *
	// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// * DELETE
	// PENDING > DELETE IN_PROGRESS
	Status *string
}

// The configuration details of a campaign.
type CampaignConfig struct {

	// A string to string map specifying the inference hyperparameters you wish to use
	// for hyperparameter optimization. See customizing-solution-config-hpo.
	ItemExplorationConfig map[string]string
}

// Provides a summary of the properties of a campaign. For a complete listing, call
// the DescribeCampaign API.
type CampaignSummary struct {

	// The Amazon Resource Name (ARN) of the campaign.
	CampaignArn *string

	// The date and time (in Unix time) that the campaign was created.
	CreationDateTime *time.Time

	// If a campaign fails, the reason behind the failure.
	FailureReason *string

	// The date and time (in Unix time) that the campaign was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the campaign.
	Name *string

	// The status of the campaign. A campaign can be in one of the following states:
	//
	// *
	// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// * DELETE
	// PENDING > DELETE IN_PROGRESS
	Status *string
}

// Provides a summary of the properties of a campaign update. For a complete
// listing, call the DescribeCampaign API.
type CampaignUpdateSummary struct {

	// The configuration details of a campaign.
	CampaignConfig *CampaignConfig

	// The date and time (in Unix time) that the campaign update was created.
	CreationDateTime *time.Time

	// If a campaign update fails, the reason behind the failure.
	FailureReason *string

	// The date and time (in Unix time) that the campaign update was last updated.
	LastUpdatedDateTime *time.Time

	// Specifies the requested minimum provisioned transactions (recommendations) per
	// second that Amazon Personalize will support.
	MinProvisionedTPS *int32

	// The Amazon Resource Name (ARN) of the deployed solution version.
	SolutionVersionArn *string

	// The status of the campaign update. A campaign update can be in one of the
	// following states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE
	// FAILED
	//
	// * DELETE PENDING > DELETE IN_PROGRESS
	Status *string
}

// Provides the name and range of a categorical hyperparameter.
type CategoricalHyperParameterRange struct {

	// The name of the hyperparameter.
	Name *string

	// A list of the categories for the hyperparameter.
	Values []string
}

// Provides the name and range of a continuous hyperparameter.
type ContinuousHyperParameterRange struct {

	// The maximum allowable value for the hyperparameter.
	MaxValue float64

	// The minimum allowable value for the hyperparameter.
	MinValue float64

	// The name of the hyperparameter.
	Name *string
}

// Provides metadata for a dataset.
type Dataset struct {

	// The creation date and time (in Unix time) of the dataset.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset that you want metadata for.
	DatasetArn *string

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// One of the following values:
	//
	// * Interactions
	//
	// * Items
	//
	// * Users
	DatasetType *string

	// A time stamp that shows when the dataset was updated.
	LastUpdatedDateTime *time.Time

	// The name of the dataset.
	Name *string

	// The ARN of the associated schema.
	SchemaArn *string

	// The status of the dataset. A dataset can be in one of the following states:
	//
	// *
	// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// * DELETE
	// PENDING > DELETE IN_PROGRESS
	Status *string
}

// A dataset group is a collection of related datasets (Interactions, User, and
// Item). You create a dataset group by calling CreateDatasetGroup. You then create
// a dataset and add it to a dataset group by calling CreateDataset. The dataset
// group is used to create and train a solution by calling CreateSolution. A
// dataset group can contain only one of each type of dataset. You can specify an
// AWS Key Management Service (KMS) key to encrypt the datasets in the group.
type DatasetGroup struct {

	// The creation date and time (in Unix time) of the dataset group.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// If creating a dataset group fails, provides the reason why.
	FailureReason *string

	// The Amazon Resource Name (ARN) of the KMS key used to encrypt the datasets.
	KmsKeyArn *string

	// The last update date and time (in Unix time) of the dataset group.
	LastUpdatedDateTime *time.Time

	// The name of the dataset group.
	Name *string

	// The ARN of the IAM role that has permissions to create the dataset group.
	RoleArn *string

	// The current status of the dataset group. A dataset group can be in one of the
	// following states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE
	// FAILED
	//
	// * DELETE PENDING
	Status *string
}

// Provides a summary of the properties of a dataset group. For a complete listing,
// call the DescribeDatasetGroup API.
type DatasetGroupSummary struct {

	// The date and time (in Unix time) that the dataset group was created.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string

	// If creating a dataset group fails, the reason behind the failure.
	FailureReason *string

	// The date and time (in Unix time) that the dataset group was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the dataset group.
	Name *string

	// The status of the dataset group. A dataset group can be in one of the following
	// states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// *
	// DELETE PENDING
	Status *string
}

// Describes a job that imports training data from a data source (Amazon S3 bucket)
// to an Amazon Personalize dataset. For more information, see
// CreateDatasetImportJob. A dataset import job can be in one of the following
// states:
//
// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
type DatasetImportJob struct {

	// The creation date and time (in Unix time) of the dataset import job.
	CreationDateTime *time.Time

	// The Amazon S3 bucket that contains the training data to import.
	DataSource *DataSource

	// The Amazon Resource Name (ARN) of the dataset that receives the imported data.
	DatasetArn *string

	// The ARN of the dataset import job.
	DatasetImportJobArn *string

	// If a dataset import job fails, provides the reason why.
	FailureReason *string

	// The name of the import job.
	JobName *string

	// The date and time (in Unix time) the dataset was last updated.
	LastUpdatedDateTime *time.Time

	// The ARN of the AWS Identity and Access Management (IAM) role that has
	// permissions to read from the Amazon S3 data source.
	RoleArn *string

	// The status of the dataset import job. A dataset import job can be in one of the
	// following states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE
	// FAILED
	Status *string
}

// Provides a summary of the properties of a dataset import job. For a complete
// listing, call the DescribeDatasetImportJob API.
type DatasetImportJobSummary struct {

	// The date and time (in Unix time) that the dataset import job was created.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset import job.
	DatasetImportJobArn *string

	// If a dataset import job fails, the reason behind the failure.
	FailureReason *string

	// The name of the dataset import job.
	JobName *string

	// The date and time (in Unix time) that the dataset was last updated.
	LastUpdatedDateTime *time.Time

	// The status of the dataset import job. A dataset import job can be in one of the
	// following states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE
	// FAILED
	Status *string
}

// Describes the schema for a dataset. For more information on schemas, see
// CreateSchema.
type DatasetSchema struct {

	// The date and time (in Unix time) that the schema was created.
	CreationDateTime *time.Time

	// The date and time (in Unix time) that the schema was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the schema.
	Name *string

	// The schema.
	Schema *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string
}

// Provides a summary of the properties of a dataset schema. For a complete
// listing, call the DescribeSchema API.
type DatasetSchemaSummary struct {

	// The date and time (in Unix time) that the schema was created.
	CreationDateTime *time.Time

	// The date and time (in Unix time) that the schema was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the schema.
	Name *string

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string
}

// Provides a summary of the properties of a dataset. For a complete listing, call
// the DescribeDataset API.
type DatasetSummary struct {

	// The date and time (in Unix time) that the dataset was created.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string

	// The dataset type. One of the following values:
	//
	// * Interactions
	//
	// * Items
	//
	// *
	// Users
	//
	// * Event-Interactions
	DatasetType *string

	// The date and time (in Unix time) that the dataset was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the dataset.
	Name *string

	// The status of the dataset. A dataset can be in one of the following states:
	//
	// *
	// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// * DELETE
	// PENDING > DELETE IN_PROGRESS
	Status *string
}

// Describes the data source that contains the data to upload to a dataset.
type DataSource struct {

	// The path to the Amazon S3 bucket where the data that you want to upload to your
	// dataset is stored. For example: s3://bucket-name/training-data.csv
	DataLocation *string
}

// Provides the name and default range of a categorical hyperparameter and whether
// the hyperparameter is tunable. A tunable hyperparameter can have its value
// determined during hyperparameter optimization (HPO).
type DefaultCategoricalHyperParameterRange struct {

	// Whether the hyperparameter is tunable.
	IsTunable bool

	// The name of the hyperparameter.
	Name *string

	// A list of the categories for the hyperparameter.
	Values []string
}

// Provides the name and default range of a continuous hyperparameter and whether
// the hyperparameter is tunable. A tunable hyperparameter can have its value
// determined during hyperparameter optimization (HPO).
type DefaultContinuousHyperParameterRange struct {

	// Whether the hyperparameter is tunable.
	IsTunable bool

	// The maximum allowable value for the hyperparameter.
	MaxValue float64

	// The minimum allowable value for the hyperparameter.
	MinValue float64

	// The name of the hyperparameter.
	Name *string
}

// Specifies the hyperparameters and their default ranges. Hyperparameters can be
// categorical, continuous, or integer-valued.
type DefaultHyperParameterRanges struct {

	// The categorical hyperparameters and their default ranges.
	CategoricalHyperParameterRanges []DefaultCategoricalHyperParameterRange

	// The continuous hyperparameters and their default ranges.
	ContinuousHyperParameterRanges []DefaultContinuousHyperParameterRange

	// The integer-valued hyperparameters and their default ranges.
	IntegerHyperParameterRanges []DefaultIntegerHyperParameterRange
}

// Provides the name and default range of a integer-valued hyperparameter and
// whether the hyperparameter is tunable. A tunable hyperparameter can have its
// value determined during hyperparameter optimization (HPO).
type DefaultIntegerHyperParameterRange struct {

	// Indicates whether the hyperparameter is tunable.
	IsTunable bool

	// The maximum allowable value for the hyperparameter.
	MaxValue int32

	// The minimum allowable value for the hyperparameter.
	MinValue int32

	// The name of the hyperparameter.
	Name *string
}

// Provides information about an event tracker.
type EventTracker struct {

	// The Amazon AWS account that owns the event tracker.
	AccountId *string

	// The date and time (in Unix format) that the event tracker was created.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group that receives the event
	// data.
	DatasetGroupArn *string

	// The ARN of the event tracker.
	EventTrackerArn *string

	// The date and time (in Unix time) that the event tracker was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the event tracker.
	Name *string

	// The status of the event tracker. An event tracker can be in one of the following
	// states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// *
	// DELETE PENDING > DELETE IN_PROGRESS
	Status *string

	// The ID of the event tracker. Include this ID in requests to the PutEvents
	// (https://docs.aws.amazon.com/personalize/latest/dg/API_UBS_PutEvents.html) API.
	TrackingId *string
}

// Provides a summary of the properties of an event tracker. For a complete
// listing, call the DescribeEventTracker API.
type EventTrackerSummary struct {

	// The date and time (in Unix time) that the event tracker was created.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the event tracker.
	EventTrackerArn *string

	// The date and time (in Unix time) that the event tracker was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the event tracker.
	Name *string

	// The status of the event tracker. An event tracker can be in one of the following
	// states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// *
	// DELETE PENDING > DELETE IN_PROGRESS
	Status *string
}

// Provides feature transformation information. Feature transformation is the
// process of modifying raw input data into a form more suitable for model
// training.
type FeatureTransformation struct {

	// The creation date and time (in Unix time) of the feature transformation.
	CreationDateTime *time.Time

	// Provides the default parameters for feature transformation.
	DefaultParameters map[string]string

	// The Amazon Resource Name (ARN) of the FeatureTransformation object.
	FeatureTransformationArn *string

	// The last update date and time (in Unix time) of the feature transformation.
	LastUpdatedDateTime *time.Time

	// The name of the feature transformation.
	Name *string

	// The status of the feature transformation. A feature transformation can be in one
	// of the following states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or-
	// CREATE FAILED
	Status *string
}

// Contains information on a recommendation filter, including its ARN, status, and
// filter expression.
type Filter struct {

	// The time at which the filter was created.
	CreationDateTime *time.Time

	// The ARN of the dataset group to which the filter belongs.
	DatasetGroupArn *string

	// If the filter failed, the reason for its failure.
	FailureReason *string

	// The ARN of the filter.
	FilterArn *string

	// Specifies the type of item interactions to filter out of recommendation results.
	// The filter expression must follow the following format: EXCLUDE itemId WHERE
	// INTERACTIONS.event_type in ("EVENT_TYPE") Where "EVENT_TYPE" is the type of
	// event to filter out. For more information, see Using Filters with Amazon
	// Personalize (https://docs.aws.amazon.com/personalize/latest/dg/filters.html).
	FilterExpression *string

	// The time at which the filter was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the filter.
	Name *string

	// The status of the filter.
	Status *string
}

// A short summary of a filter's attributes.
type FilterSummary struct {

	// The time at which the filter was created.
	CreationDateTime *time.Time

	// The ARN of the dataset group to which the filter belongs.
	DatasetGroupArn *string

	// If the filter failed, the reason for the failure.
	FailureReason *string

	// The ARN of the filter.
	FilterArn *string

	// The time at which the filter was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the filter.
	Name *string

	// The status of the filter.
	Status *string
}

// Describes the properties for hyperparameter optimization (HPO). For use with the
// bring-your-own-recipe feature. Do not use for Amazon Personalize native recipes.
type HPOConfig struct {

	// The hyperparameters and their allowable ranges.
	AlgorithmHyperParameterRanges *HyperParameterRanges

	// The metric to optimize during HPO.
	HpoObjective *HPOObjective

	// Describes the resource configuration for HPO.
	HpoResourceConfig *HPOResourceConfig
}

// The metric to optimize during hyperparameter optimization (HPO).
type HPOObjective struct {

	// The name of the metric.
	MetricName *string

	// A regular expression for finding the metric in the training job logs.
	MetricRegex *string

	// The type of the metric. Valid values are Maximize and Minimize.
	Type *string
}

// Describes the resource configuration for hyperparameter optimization (HPO).
type HPOResourceConfig struct {

	// The maximum number of training jobs when you create a solution version. The
	// maximum value for maxNumberOfTrainingJobs is 40.
	MaxNumberOfTrainingJobs *string

	// The maximum number of parallel training jobs when you create a solution version.
	// The maximum value for maxParallelTrainingJobs is 10.
	MaxParallelTrainingJobs *string
}

// Specifies the hyperparameters and their ranges. Hyperparameters can be
// categorical, continuous, or integer-valued.
type HyperParameterRanges struct {

	// The categorical hyperparameters and their ranges.
	CategoricalHyperParameterRanges []CategoricalHyperParameterRange

	// The continuous hyperparameters and their ranges.
	ContinuousHyperParameterRanges []ContinuousHyperParameterRange

	// The integer-valued hyperparameters and their ranges.
	IntegerHyperParameterRanges []IntegerHyperParameterRange
}

// Provides the name and range of an integer-valued hyperparameter.
type IntegerHyperParameterRange struct {

	// The maximum allowable value for the hyperparameter.
	MaxValue int32

	// The minimum allowable value for the hyperparameter.
	MinValue int32

	// The name of the hyperparameter.
	Name *string
}

// Provides information about a recipe. Each recipe provides an algorithm that
// Amazon Personalize uses in model training when you use the CreateSolution
// operation.
type Recipe struct {

	// The Amazon Resource Name (ARN) of the algorithm that Amazon Personalize uses to
	// train the model.
	AlgorithmArn *string

	// The date and time (in Unix format) that the recipe was created.
	CreationDateTime *time.Time

	// The description of the recipe.
	Description *string

	// The ARN of the FeatureTransformation object.
	FeatureTransformationArn *string

	// The date and time (in Unix format) that the recipe was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the recipe.
	Name *string

	// The Amazon Resource Name (ARN) of the recipe.
	RecipeArn *string

	// One of the following values:
	//
	// * PERSONALIZED_RANKING
	//
	// * RELATED_ITEMS
	//
	// *
	// USER_PERSONALIZATION
	RecipeType *string

	// The status of the recipe.
	Status *string
}

// Provides a summary of the properties of a recipe. For a complete listing, call
// the DescribeRecipe API.
type RecipeSummary struct {

	// The date and time (in Unix time) that the recipe was created.
	CreationDateTime *time.Time

	// The date and time (in Unix time) that the recipe was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the recipe.
	Name *string

	// The Amazon Resource Name (ARN) of the recipe.
	RecipeArn *string

	// The status of the recipe.
	Status *string
}

// The configuration details of an Amazon S3 input or output bucket.
type S3DataConfig struct {

	// The file path of the Amazon S3 bucket.
	//
	// This member is required.
	Path *string

	// The Amazon Resource Name (ARN) of the Amazon Key Management Service (KMS) key
	// that Amazon Personalize uses to encrypt or decrypt the input and output files of
	// a batch inference job.
	KmsKeyArn *string
}

// An object that provides information about a solution. A solution is a trained
// model that can be deployed as a campaign.
type Solution struct {

	// When performAutoML is true, specifies the best recipe found.
	AutoMLResult *AutoMLResult

	// The creation date and time (in Unix time) of the solution.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group that provides the training
	// data.
	DatasetGroupArn *string

	// The event type (for example, 'click' or 'like') that is used for training the
	// model.
	EventType *string

	// The date and time (in Unix time) that the solution was last updated.
	LastUpdatedDateTime *time.Time

	// Describes the latest version of the solution, including the status and the ARN.
	LatestSolutionVersion *SolutionVersionSummary

	// The name of the solution.
	Name *string

	// When true, Amazon Personalize performs a search for the best
	// USER_PERSONALIZATION recipe from the list specified in the solution
	// configuration (recipeArn must not be specified). When false (the default),
	// Amazon Personalize uses recipeArn for training.
	PerformAutoML bool

	// Whether to perform hyperparameter optimization (HPO) on the chosen recipe. The
	// default is false.
	PerformHPO bool

	// The ARN of the recipe used to create the solution.
	RecipeArn *string

	// The ARN of the solution.
	SolutionArn *string

	// Describes the configuration properties for the solution.
	SolutionConfig *SolutionConfig

	// The status of the solution. A solution can be in one of the following states:
	//
	// *
	// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// * DELETE
	// PENDING > DELETE IN_PROGRESS
	Status *string
}

// Describes the configuration properties for the solution.
type SolutionConfig struct {

	// Lists the hyperparameter names and ranges.
	AlgorithmHyperParameters map[string]string

	// The AutoMLConfig object containing a list of recipes to search when AutoML is
	// performed.
	AutoMLConfig *AutoMLConfig

	// Only events with a value greater than or equal to this threshold are used for
	// training a model.
	EventValueThreshold *string

	// Lists the feature transformation parameters.
	FeatureTransformationParameters map[string]string

	// Describes the properties for hyperparameter optimization (HPO).
	HpoConfig *HPOConfig
}

// Provides a summary of the properties of a solution. For a complete listing, call
// the DescribeSolution API.
type SolutionSummary struct {

	// The date and time (in Unix time) that the solution was created.
	CreationDateTime *time.Time

	// The date and time (in Unix time) that the solution was last updated.
	LastUpdatedDateTime *time.Time

	// The name of the solution.
	Name *string

	// The Amazon Resource Name (ARN) of the solution.
	SolutionArn *string

	// The status of the solution. A solution can be in one of the following states:
	//
	// *
	// CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
	//
	// * DELETE
	// PENDING > DELETE IN_PROGRESS
	Status *string
}

// An object that provides information about a specific version of a Solution.
type SolutionVersion struct {

	// The date and time (in Unix time) that this version of the solution was created.
	CreationDateTime *time.Time

	// The Amazon Resource Name (ARN) of the dataset group providing the training data.
	DatasetGroupArn *string

	// The event type (for example, 'click' or 'like') that is used for training the
	// model.
	EventType *string

	// If training a solution version fails, the reason for the failure.
	FailureReason *string

	// The date and time (in Unix time) that the solution was last updated.
	LastUpdatedDateTime *time.Time

	// When true, Amazon Personalize searches for the most optimal recipe according to
	// the solution configuration. When false (the default), Amazon Personalize uses
	// recipeArn.
	PerformAutoML bool

	// Whether to perform hyperparameter optimization (HPO) on the chosen recipe. The
	// default is false.
	PerformHPO bool

	// The ARN of the recipe used in the solution.
	RecipeArn *string

	// The ARN of the solution.
	SolutionArn *string

	// Describes the configuration properties for the solution.
	SolutionConfig *SolutionConfig

	// The ARN of the solution version.
	SolutionVersionArn *string

	// The status of the solution version. A solution version can be in one of the
	// following states:
	//
	// * CREATE PENDING
	//
	// * CREATE IN_PROGRESS
	//
	// * ACTIVE
	//
	// * CREATE
	// FAILED
	Status *string

	// The time used to train the model. You are billed for the time it takes to train
	// a model. This field is visible only after Amazon Personalize successfully trains
	// a model.
	TrainingHours *float64

	// The scope of training used to create the solution version. The FULL option
	// trains the solution version based on the entirety of the input solution's
	// training data, while the UPDATE option processes only the training data that has
	// changed since the creation of the last solution version. Choose UPDATE when you
	// want to start recommending items added to the dataset without retraining the
	// model. The UPDATE option can only be used after you've created a solution
	// version with the FULL option and the training solution uses the
	// native-recipe-hrnn-coldstart.
	TrainingMode TrainingMode

	// If hyperparameter optimization was performed, contains the hyperparameter values
	// of the best performing model.
	TunedHPOParams *TunedHPOParams
}

// Provides a summary of the properties of a solution version. For a complete
// listing, call the DescribeSolutionVersion API.
type SolutionVersionSummary struct {

	// The date and time (in Unix time) that this version of a solution was created.
	CreationDateTime *time.Time

	// If a solution version fails, the reason behind the failure.
	FailureReason *string

	// The date and time (in Unix time) that the solution version was last updated.
	LastUpdatedDateTime *time.Time

	// The Amazon Resource Name (ARN) of the solution version.
	SolutionVersionArn *string

	// The status of the solution version. A solution version can be in one of the
	// following states:
	//
	// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE
	// FAILED
	Status *string
}

// If hyperparameter optimization (HPO) was performed, contains the hyperparameter
// values of the best performing model.
type TunedHPOParams struct {

	// A list of the hyperparameter values of the best performing model.
	AlgorithmHyperParameters map[string]string
}
