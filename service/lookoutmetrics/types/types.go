// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A configuration that specifies the action to perform when anomalies are
// detected.
type Action struct {

	// A configuration for an AWS Lambda channel.
	LambdaConfiguration *LambdaConfiguration

	// A configuration for an Amazon SNS channel.
	SNSConfiguration *SNSConfiguration

	noSmithyDocumentSerde
}

// A configuration for Amazon SNS-integrated notifications.
type Alert struct {

	// Action that will be triggered when there is an alert.
	Action *Action

	// The ARN of the alert.
	AlertArn *string

	// A description of the alert.
	AlertDescription *string

	// The name of the alert.
	AlertName *string

	// The minimum severity for an anomaly to trigger the alert.
	AlertSensitivityThreshold int32

	// The status of the alert.
	AlertStatus AlertStatus

	// The type of the alert.
	AlertType AlertType

	// The ARN of the detector to which the alert is attached.
	AnomalyDetectorArn *string

	// The time at which the alert was created.
	CreationTime *time.Time

	// The time at which the alert was last modified.
	LastModificationTime *time.Time

	noSmithyDocumentSerde
}

// Provides a summary of an alert's configuration.
type AlertSummary struct {

	// The ARN of the alert.
	AlertArn *string

	// The name of the alert.
	AlertName *string

	// The minimum severity for an anomaly to trigger the alert.
	AlertSensitivityThreshold int32

	// The status of the alert.
	AlertStatus AlertStatus

	// The type of the alert.
	AlertType AlertType

	// The ARN of the detector to which the alert is attached.
	AnomalyDetectorArn *string

	// The time at which the alert was created.
	CreationTime *time.Time

	// The time at which the alert was last modified.
	LastModificationTime *time.Time

	// The alert's tags
	// (https://docs.aws.amazon.com/lookoutmetrics/latest/dev/detectors-tags.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

// Contains information about a detector's configuration.
type AnomalyDetectorConfig struct {

	// The frequency at which the detector analyzes its source data.
	AnomalyDetectorFrequency Frequency

	noSmithyDocumentSerde
}

// Contains information about a detector's configuration.
type AnomalyDetectorConfigSummary struct {

	// The interval at which the detector analyzes its source data.
	AnomalyDetectorFrequency Frequency

	noSmithyDocumentSerde
}

// Contains information about an an anomaly detector.
type AnomalyDetectorSummary struct {

	// The ARN of the detector.
	AnomalyDetectorArn *string

	// A description of the detector.
	AnomalyDetectorDescription *string

	// The name of the detector.
	AnomalyDetectorName *string

	// The time at which the detector was created.
	CreationTime *time.Time

	// The time at which the detector was last modified.
	LastModificationTime *time.Time

	// The status of detector.
	Status AnomalyDetectorStatus

	// The detector's tags
	// (https://docs.aws.amazon.com/lookoutmetrics/latest/dev/detectors-tags.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

// A group of anomalous metrics
type AnomalyGroup struct {

	// The ID of the anomaly group.
	AnomalyGroupId *string

	// The severity score of the group.
	AnomalyGroupScore *float64

	// The end time for the group.
	EndTime *string

	// A list of measures affected by the anomaly.
	MetricLevelImpactList []MetricLevelImpact

	// The name of the primary affected measure for the group.
	PrimaryMetricName *string

	// The start time for the group.
	StartTime *string

	noSmithyDocumentSerde
}

// Aggregated statistics for a group of anomalous metrics.
type AnomalyGroupStatistics struct {

	// The start of the time range that was searched.
	EvaluationStartDate *string

	// Statistics for individual metrics within the group.
	ItemizedMetricStatsList []ItemizedMetricStats

	// The number of groups found.
	TotalCount int32

	noSmithyDocumentSerde
}

// Details about a group of anomalous metrics.
type AnomalyGroupSummary struct {

	// The ID of the anomaly group.
	AnomalyGroupId *string

	// The severity score of the group.
	AnomalyGroupScore *float64

	// The end time for the group.
	EndTime *string

	// The name of the primary affected measure for the group.
	PrimaryMetricName *string

	// The start time for the group.
	StartTime *string

	noSmithyDocumentSerde
}

// An anomalous metric in an anomaly group.
type AnomalyGroupTimeSeries struct {

	// The ID of the anomaly group.
	//
	// This member is required.
	AnomalyGroupId *string

	// The ID of the metric.
	TimeSeriesId *string

	noSmithyDocumentSerde
}

// Feedback for an anomalous metric.
type AnomalyGroupTimeSeriesFeedback struct {

	// The ID of the anomaly group.
	//
	// This member is required.
	AnomalyGroupId *string

	// Feedback on whether the metric is a legitimate anomaly.
	//
	// This member is required.
	IsAnomaly *bool

	// The ID of the metric.
	//
	// This member is required.
	TimeSeriesId *string

	noSmithyDocumentSerde
}

// Details about an Amazon AppFlow flow datasource.
type AppFlowConfig struct {

	// name of the flow.
	FlowName *string

	// An IAM role that gives Amazon Lookout for Metrics permission to access the flow.
	RoleArn *string

	noSmithyDocumentSerde
}

// An attribute value.
type AttributeValue struct {

	// A binary value.
	B *string

	// A list of binary values.
	BS []string

	// A number.
	N *string

	// A list of numbers.
	NS []string

	// A string.
	S *string

	// A list of strings.
	SS []string

	noSmithyDocumentSerde
}

// An auto detection metric source.
type AutoDetectionMetricSource struct {

	// The source's source config.
	S3SourceConfig *AutoDetectionS3SourceConfig

	noSmithyDocumentSerde
}

// An auto detection source config.
type AutoDetectionS3SourceConfig struct {

	// The config's historical data path list.
	HistoricalDataPathList []string

	// The config's templated path list.
	TemplatedPathList []string

	noSmithyDocumentSerde
}

// Details about an Amazon CloudWatch datasource.
type CloudWatchConfig struct {

	// An IAM role that gives Amazon Lookout for Metrics permission to access data in
	// Amazon CloudWatch.
	RoleArn *string

	noSmithyDocumentSerde
}

// Details about dimensions that contributed to an anomaly.
type ContributionMatrix struct {

	// A list of contributing dimensions.
	DimensionContributionList []DimensionContribution

	noSmithyDocumentSerde
}

// Contains information about how a source CSV data file should be analyzed.
type CsvFormatDescriptor struct {

	// The character set in which the source CSV file is written.
	Charset *string

	// Whether or not the source CSV file contains a header.
	ContainsHeader *bool

	// The character used to delimit the source CSV file.
	Delimiter *string

	// The level of compression of the source CSV file.
	FileCompression CSVFileCompression

	// A list of the source CSV file's headers, if any.
	HeaderList []string

	// The character used as a quote character.
	QuoteSymbol *string

	noSmithyDocumentSerde
}

// Properties of an inferred CSV format.
type DetectedCsvFormatDescriptor struct {

	// The format's charset.
	Charset *DetectedField

	// Whether the format includes a header.
	ContainsHeader *DetectedField

	// The format's delimiter.
	Delimiter *DetectedField

	// The format's file compression.
	FileCompression *DetectedField

	// The format's header list.
	HeaderList *DetectedField

	// The format's quote symbol.
	QuoteSymbol *DetectedField

	noSmithyDocumentSerde
}

// An inferred field.
type DetectedField struct {

	// The field's confidence.
	Confidence Confidence

	// The field's message.
	Message *string

	// The field's value.
	Value *AttributeValue

	noSmithyDocumentSerde
}

// Properties of an inferred data format.
type DetectedFileFormatDescriptor struct {

	// Details about a CSV format.
	CsvFormatDescriptor *DetectedCsvFormatDescriptor

	// Details about a JSON format.
	JsonFormatDescriptor *DetectedJsonFormatDescriptor

	noSmithyDocumentSerde
}

// A detected JSON format descriptor.
type DetectedJsonFormatDescriptor struct {

	// The format's character set.
	Charset *DetectedField

	// The format's file compression.
	FileCompression *DetectedField

	noSmithyDocumentSerde
}

// An inferred dataset configuration.
type DetectedMetricSetConfig struct {

	// The dataset's interval.
	MetricSetFrequency *DetectedField

	// The dataset's data source.
	MetricSource *DetectedMetricSource

	// The dataset's offset.
	Offset *DetectedField

	noSmithyDocumentSerde
}

// An inferred data source.
type DetectedMetricSource struct {

	// The data source's source configuration.
	S3SourceConfig *DetectedS3SourceConfig

	noSmithyDocumentSerde
}

// An inferred source configuration.
type DetectedS3SourceConfig struct {

	// The source's file format descriptor.
	FileFormatDescriptor *DetectedFileFormatDescriptor

	noSmithyDocumentSerde
}

// Details about a dimension that contributed to an anomaly.
type DimensionContribution struct {

	// The name of the dimension.
	DimensionName *string

	// A list of dimension values that contributed to the anomaly.
	DimensionValueContributionList []DimensionValueContribution

	noSmithyDocumentSerde
}

// A dimension name and value.
type DimensionNameValue struct {

	// The name of the dimension.
	//
	// This member is required.
	DimensionName *string

	// The value of the dimension.
	//
	// This member is required.
	DimensionValue *string

	noSmithyDocumentSerde
}

// The severity of a value of a dimension that contributed to an anomaly.
type DimensionValueContribution struct {

	// The severity score of the value.
	ContributionScore *float64

	// The value of the dimension.
	DimensionValue *string

	noSmithyDocumentSerde
}

// The status of an anomaly detector run.
type ExecutionStatus struct {

	// The reason that the run failed, if applicable.
	FailureReason *string

	// The run's status.
	Status AnomalyDetectionTaskStatus

	// The run's timestamp.
	Timestamp *string

	noSmithyDocumentSerde
}

// Contains information about a source file's formatting.
type FileFormatDescriptor struct {

	// Contains information about how a source CSV data file should be analyzed.
	CsvFormatDescriptor *CsvFormatDescriptor

	// Contains information about how a source JSON data file should be analyzed.
	JsonFormatDescriptor *JsonFormatDescriptor

	noSmithyDocumentSerde
}

// Aggregated details about the measures contributing to the anomaly group, and the
// measures potentially impacted by the anomaly group.
type InterMetricImpactDetails struct {

	// The ID of the anomaly group.
	AnomalyGroupId *string

	// For potential causes (CAUSE_OF_INPUT_ANOMALY_GROUP), the percentage contribution
	// the measure has in causing the anomalies.
	ContributionPercentage *float64

	// The name of the measure.
	MetricName *string

	// Whether a measure is a potential cause of the anomaly group
	// (CAUSE_OF_INPUT_ANOMALY_GROUP), or whether the measure is impacted by the
	// anomaly group (EFFECT_OF_INPUT_ANOMALY_GROUP).
	RelationshipType RelationshipType

	noSmithyDocumentSerde
}

// Aggregated statistics about a measure affected by an anomaly.
type ItemizedMetricStats struct {

	// The name of the measure.
	MetricName *string

	// The number of times that the measure appears.
	OccurrenceCount int32

	noSmithyDocumentSerde
}

// Contains information about how a source JSON data file should be analyzed.
type JsonFormatDescriptor struct {

	// The character set in which the source JSON file is written.
	Charset *string

	// The level of compression of the source CSV file.
	FileCompression JsonFileCompression

	noSmithyDocumentSerde
}

// Contains information about a Lambda configuration.
type LambdaConfiguration struct {

	// The ARN of the Lambda function.
	//
	// This member is required.
	LambdaArn *string

	// The ARN of an IAM role that has permission to invoke the Lambda function.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

// A calculation made by contrasting a measure and a dimension from your source
// data.
type Metric struct {

	// The function with which the metric is calculated.
	//
	// This member is required.
	AggregationFunction AggregationFunction

	// The name of the metric.
	//
	// This member is required.
	MetricName *string

	// The namespace for the metric.
	Namespace *string

	noSmithyDocumentSerde
}

// Details about a measure affected by an anomaly.
type MetricLevelImpact struct {

	// Details about the dimensions that contributed to the anomaly.
	ContributionMatrix *ContributionMatrix

	// The name of the measure.
	MetricName *string

	// The number of anomalous metrics for the measure.
	NumTimeSeries int32

	noSmithyDocumentSerde
}

// Contains information about a dataset.
type MetricSetSummary struct {

	// The ARN of the detector to which the dataset belongs.
	AnomalyDetectorArn *string

	// The time at which the dataset was created.
	CreationTime *time.Time

	// The time at which the dataset was last modified.
	LastModificationTime *time.Time

	// The ARN of the dataset.
	MetricSetArn *string

	// The description of the dataset.
	MetricSetDescription *string

	// The name of the dataset.
	MetricSetName *string

	// The dataset's tags
	// (https://docs.aws.amazon.com/lookoutmetrics/latest/dev/detectors-tags.html).
	Tags map[string]string

	noSmithyDocumentSerde
}

// Contains information about source data used to generate a metric.
type MetricSource struct {

	// An object containing information about the AppFlow configuration.
	AppFlowConfig *AppFlowConfig

	// An object containing information about the Amazon CloudWatch monitoring
	// configuration.
	CloudWatchConfig *CloudWatchConfig

	// An object containing information about the Amazon Relational Database Service
	// (RDS) configuration.
	RDSSourceConfig *RDSSourceConfig

	// An object containing information about the Amazon Redshift database
	// configuration.
	RedshiftSourceConfig *RedshiftSourceConfig

	// Contains information about the configuration of the S3 bucket that contains
	// source files.
	S3SourceConfig *S3SourceConfig

	noSmithyDocumentSerde
}

// Contains information about the Amazon Relational Database Service (RDS)
// configuration.
type RDSSourceConfig struct {

	// A string identifying the database instance.
	DBInstanceIdentifier *string

	// The host name of the database.
	DatabaseHost *string

	// The name of the RDS database.
	DatabaseName *string

	// The port number where the database can be accessed.
	DatabasePort int32

	// The Amazon Resource Name (ARN) of the role.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	SecretManagerArn *string

	// The name of the table in the database.
	TableName *string

	// An object containing information about the Amazon Virtual Private Cloud (VPC)
	// configuration.
	VpcConfiguration *VpcConfiguration

	noSmithyDocumentSerde
}

// Provides information about the Amazon Redshift database configuration.
type RedshiftSourceConfig struct {

	// A string identifying the Redshift cluster.
	ClusterIdentifier *string

	// The name of the database host.
	DatabaseHost *string

	// The Redshift database name.
	DatabaseName *string

	// The port number where the database can be accessed.
	DatabasePort int32

	// The Amazon Resource Name (ARN) of the role providing access to the database.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the AWS Secrets Manager role.
	SecretManagerArn *string

	// The table name of the Redshift database.
	TableName *string

	// Contains information about the Amazon Virtual Private Cloud (VPC) configuration.
	VpcConfiguration *VpcConfiguration

	noSmithyDocumentSerde
}

// Contains information about the configuration of the S3 bucket that contains
// source files.
type S3SourceConfig struct {

	// Contains information about a source file's formatting.
	FileFormatDescriptor *FileFormatDescriptor

	// A list of paths to the historical data files.
	HistoricalDataPathList []string

	// The ARN of an IAM role that has read and write access permissions to the source
	// S3 bucket.
	RoleArn *string

	// A list of templated paths to the source files.
	TemplatedPathList []string

	noSmithyDocumentSerde
}

// Contains information about the source configuration in Amazon S3.
type SampleDataS3SourceConfig struct {

	// Contains information about a source file's formatting.
	//
	// This member is required.
	FileFormatDescriptor *FileFormatDescriptor

	// The Amazon Resource Name (ARN) of the role.
	//
	// This member is required.
	RoleArn *string

	// An array of strings containing the historical set of data paths.
	HistoricalDataPathList []string

	// An array of strings containing the list of templated paths.
	TemplatedPathList []string

	noSmithyDocumentSerde
}

// Contains information about the SNS topic to which you want to send your alerts
// and the IAM role that has access to that topic.
type SNSConfiguration struct {

	// The ARN of the IAM role that has access to the target SNS topic.
	//
	// This member is required.
	RoleArn *string

	// The ARN of the target SNS topic.
	//
	// This member is required.
	SnsTopicArn *string

	// The text format for alerts.
	SnsFormat SnsFormat

	noSmithyDocumentSerde
}

// Details about a metric. A metric is an aggregation of the values of a measure
// for a dimension value, such as availability in the us-east-1 Region.
type TimeSeries struct {

	// The dimensions of the metric.
	//
	// This member is required.
	DimensionList []DimensionNameValue

	// The values for the metric.
	//
	// This member is required.
	MetricValueList []float64

	// The ID of the metric.
	//
	// This member is required.
	TimeSeriesId *string

	noSmithyDocumentSerde
}

// Details about feedback submitted for an anomalous metric.
type TimeSeriesFeedback struct {

	// Feedback on whether the metric is a legitimate anomaly.
	IsAnomaly *bool

	// The ID of the metric.
	TimeSeriesId *string

	noSmithyDocumentSerde
}

// Contains information about the column used to track time in a source data file.
type TimestampColumn struct {

	// The format of the timestamp column.
	ColumnFormat *string

	// The name of the timestamp column.
	ColumnName *string

	noSmithyDocumentSerde
}

// Contains information about a a field in a validation exception.
type ValidationExceptionField struct {

	// The message with more information about the validation exception.
	//
	// This member is required.
	Message *string

	// The name of the field.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

// Contains configuration information about the Amazon Virtual Private Cloud (VPC).
type VpcConfiguration struct {

	// An array of strings containing the list of security groups.
	//
	// This member is required.
	SecurityGroupIdList []string

	// An array of strings containing the Amazon VPC subnet IDs (e.g.,
	// subnet-0bb1c79de3EXAMPLE.
	//
	// This member is required.
	SubnetIdList []string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
