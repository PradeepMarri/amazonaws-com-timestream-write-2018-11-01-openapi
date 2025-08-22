package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DescribeBatchLoadTaskRequest represents the DescribeBatchLoadTaskRequest schema from the OpenAPI specification
type DescribeBatchLoadTaskRequest struct {
	Taskid interface{} `json:"TaskId"`
}

// PartitionKey represents the PartitionKey schema from the OpenAPI specification
type PartitionKey struct {
	Enforcementinrecord interface{} `json:"EnforcementInRecord,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	TypeField interface{} `json:"Type"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"Tags,omitempty"`
}

// ReportS3Configuration represents the ReportS3Configuration schema from the OpenAPI specification
type ReportS3Configuration struct {
	Bucketname interface{} `json:"BucketName"`
	Encryptionoption interface{} `json:"EncryptionOption,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Objectkeyprefix interface{} `json:"ObjectKeyPrefix,omitempty"`
}

// MagneticStoreRejectedDataLocation represents the MagneticStoreRejectedDataLocation schema from the OpenAPI specification
type MagneticStoreRejectedDataLocation struct {
	S3configuration interface{} `json:"S3Configuration,omitempty"`
}

// MeasureValue represents the MeasureValue schema from the OpenAPI specification
type MeasureValue struct {
	Name interface{} `json:"Name"`
	TypeField interface{} `json:"Type"`
	Value interface{} `json:"Value"`
}

// Schema represents the Schema schema from the OpenAPI specification
type Schema struct {
	Compositepartitionkey interface{} `json:"CompositePartitionKey,omitempty"`
}

// ListTablesRequest represents the ListTablesRequest schema from the OpenAPI specification
type ListTablesRequest struct {
	Databasename interface{} `json:"DatabaseName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// ResumeBatchLoadTaskRequest represents the ResumeBatchLoadTaskRequest schema from the OpenAPI specification
type ResumeBatchLoadTaskRequest struct {
	Taskid interface{} `json:"TaskId"`
}

// DescribeDatabaseResponse represents the DescribeDatabaseResponse schema from the OpenAPI specification
type DescribeDatabaseResponse struct {
	Database interface{} `json:"Database,omitempty"`
}

// DescribeDatabaseRequest represents the DescribeDatabaseRequest schema from the OpenAPI specification
type DescribeDatabaseRequest struct {
	Databasename interface{} `json:"DatabaseName"`
}

// BatchLoadProgressReport represents the BatchLoadProgressReport schema from the OpenAPI specification
type BatchLoadProgressReport struct {
	Bytesmetered interface{} `json:"BytesMetered,omitempty"`
	Filefailures interface{} `json:"FileFailures,omitempty"`
	Parsefailures interface{} `json:"ParseFailures,omitempty"`
	Recordingestionfailures interface{} `json:"RecordIngestionFailures,omitempty"`
	Recordsingested interface{} `json:"RecordsIngested,omitempty"`
	Recordsprocessed interface{} `json:"RecordsProcessed,omitempty"`
}

// UpdateTableResponse represents the UpdateTableResponse schema from the OpenAPI specification
type UpdateTableResponse struct {
	Table interface{} `json:"Table,omitempty"`
}

// CreateTableResponse represents the CreateTableResponse schema from the OpenAPI specification
type CreateTableResponse struct {
	Table interface{} `json:"Table,omitempty"`
}

// CreateBatchLoadTaskRequest represents the CreateBatchLoadTaskRequest schema from the OpenAPI specification
type CreateBatchLoadTaskRequest struct {
	Targettablename interface{} `json:"TargetTableName"`
	Clienttoken interface{} `json:"ClientToken,omitempty"`
	Datamodelconfiguration DataModelConfiguration `json:"DataModelConfiguration,omitempty"` // <p/>
	Datasourceconfiguration interface{} `json:"DataSourceConfiguration"`
	Recordversion interface{} `json:"RecordVersion,omitempty"`
	Reportconfiguration ReportConfiguration `json:"ReportConfiguration"` // Report configuration for a batch load task. This contains details about where error reports are stored.
	Targetdatabasename interface{} `json:"TargetDatabaseName"`
}

// Dimension represents the Dimension schema from the OpenAPI specification
type Dimension struct {
	Dimensionvaluetype interface{} `json:"DimensionValueType,omitempty"`
	Name interface{} `json:"Name"`
	Value interface{} `json:"Value"`
}

// Table represents the Table schema from the OpenAPI specification
type Table struct {
	Arn interface{} `json:"Arn,omitempty"`
	Schema interface{} `json:"Schema,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Tablestatus interface{} `json:"TableStatus,omitempty"`
	Magneticstorewriteproperties interface{} `json:"MagneticStoreWriteProperties,omitempty"`
	Tablename interface{} `json:"TableName,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Databasename interface{} `json:"DatabaseName,omitempty"`
	Retentionproperties interface{} `json:"RetentionProperties,omitempty"`
}

// BatchLoadTask represents the BatchLoadTask schema from the OpenAPI specification
type BatchLoadTask struct {
	Databasename interface{} `json:"DatabaseName,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Resumableuntil interface{} `json:"ResumableUntil,omitempty"`
	Tablename interface{} `json:"TableName,omitempty"`
	Taskid interface{} `json:"TaskId,omitempty"`
	Taskstatus interface{} `json:"TaskStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tagkeys interface{} `json:"TagKeys"`
}

// ResumeBatchLoadTaskResponse represents the ResumeBatchLoadTaskResponse schema from the OpenAPI specification
type ResumeBatchLoadTaskResponse struct {
}

// CreateTableRequest represents the CreateTableRequest schema from the OpenAPI specification
type CreateTableRequest struct {
	Tags interface{} `json:"Tags,omitempty"`
	Databasename interface{} `json:"DatabaseName"`
	Magneticstorewriteproperties interface{} `json:"MagneticStoreWriteProperties,omitempty"`
	Retentionproperties interface{} `json:"RetentionProperties,omitempty"`
	Schema interface{} `json:"Schema,omitempty"`
	Tablename interface{} `json:"TableName"`
}

// RecordsIngested represents the RecordsIngested schema from the OpenAPI specification
type RecordsIngested struct {
	Total interface{} `json:"Total,omitempty"`
	Magneticstore interface{} `json:"MagneticStore,omitempty"`
	Memorystore interface{} `json:"MemoryStore,omitempty"`
}

// Endpoint represents the Endpoint schema from the OpenAPI specification
type Endpoint struct {
	Address interface{} `json:"Address"`
	Cacheperiodinminutes interface{} `json:"CachePeriodInMinutes"`
}

// DataSourceS3Configuration represents the DataSourceS3Configuration schema from the OpenAPI specification
type DataSourceS3Configuration struct {
	Objectkeyprefix interface{} `json:"ObjectKeyPrefix,omitempty"`
	Bucketname interface{} `json:"BucketName"`
}

// UpdateDatabaseRequest represents the UpdateDatabaseRequest schema from the OpenAPI specification
type UpdateDatabaseRequest struct {
	Databasename interface{} `json:"DatabaseName"`
	Kmskeyid interface{} `json:"KmsKeyId"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
}

// WriteRecordsResponse represents the WriteRecordsResponse schema from the OpenAPI specification
type WriteRecordsResponse struct {
	Recordsingested interface{} `json:"RecordsIngested,omitempty"`
}

// ReportConfiguration represents the ReportConfiguration schema from the OpenAPI specification
type ReportConfiguration struct {
	Reports3configuration interface{} `json:"ReportS3Configuration,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// DescribeTableRequest represents the DescribeTableRequest schema from the OpenAPI specification
type DescribeTableRequest struct {
	Databasename interface{} `json:"DatabaseName"`
	Tablename interface{} `json:"TableName"`
}

// RetentionProperties represents the RetentionProperties schema from the OpenAPI specification
type RetentionProperties struct {
	Magneticstoreretentionperiodindays interface{} `json:"MagneticStoreRetentionPeriodInDays"`
	Memorystoreretentionperiodinhours interface{} `json:"MemoryStoreRetentionPeriodInHours"`
}

// MixedMeasureMapping represents the MixedMeasureMapping schema from the OpenAPI specification
type MixedMeasureMapping struct {
	Sourcecolumn interface{} `json:"SourceColumn,omitempty"`
	Targetmeasurename interface{} `json:"TargetMeasureName,omitempty"`
	Measurename interface{} `json:"MeasureName,omitempty"`
	Measurevaluetype interface{} `json:"MeasureValueType"`
	Multimeasureattributemappings interface{} `json:"MultiMeasureAttributeMappings,omitempty"`
}

// CreateDatabaseResponse represents the CreateDatabaseResponse schema from the OpenAPI specification
type CreateDatabaseResponse struct {
	Database interface{} `json:"Database,omitempty"`
}

// UpdateTableRequest represents the UpdateTableRequest schema from the OpenAPI specification
type UpdateTableRequest struct {
	Databasename interface{} `json:"DatabaseName"`
	Magneticstorewriteproperties interface{} `json:"MagneticStoreWriteProperties,omitempty"`
	Retentionproperties interface{} `json:"RetentionProperties,omitempty"`
	Schema interface{} `json:"Schema,omitempty"`
	Tablename interface{} `json:"TableName"`
}

// S3Configuration represents the S3Configuration schema from the OpenAPI specification
type S3Configuration struct {
	Objectkeyprefix interface{} `json:"ObjectKeyPrefix,omitempty"`
	Bucketname interface{} `json:"BucketName,omitempty"`
	Encryptionoption interface{} `json:"EncryptionOption,omitempty"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
}

// DataModel represents the DataModel schema from the OpenAPI specification
type DataModel struct {
	Timecolumn interface{} `json:"TimeColumn,omitempty"`
	Timeunit interface{} `json:"TimeUnit,omitempty"`
	Dimensionmappings interface{} `json:"DimensionMappings"`
	Measurenamecolumn interface{} `json:"MeasureNameColumn,omitempty"`
	Mixedmeasuremappings interface{} `json:"MixedMeasureMappings,omitempty"`
	Multimeasuremappings interface{} `json:"MultiMeasureMappings,omitempty"`
}

// ListDatabasesResponse represents the ListDatabasesResponse schema from the OpenAPI specification
type ListDatabasesResponse struct {
	Databases interface{} `json:"Databases,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// MultiMeasureMappings represents the MultiMeasureMappings schema from the OpenAPI specification
type MultiMeasureMappings struct {
	Multimeasureattributemappings interface{} `json:"MultiMeasureAttributeMappings"`
	Targetmultimeasurename interface{} `json:"TargetMultiMeasureName,omitempty"`
}

// DescribeEndpointsRequest represents the DescribeEndpointsRequest schema from the OpenAPI specification
type DescribeEndpointsRequest struct {
}

// DeleteTableRequest represents the DeleteTableRequest schema from the OpenAPI specification
type DeleteTableRequest struct {
	Databasename interface{} `json:"DatabaseName"`
	Tablename interface{} `json:"TableName"`
}

// DataSourceConfiguration represents the DataSourceConfiguration schema from the OpenAPI specification
type DataSourceConfiguration struct {
	Csvconfiguration CsvConfiguration `json:"CsvConfiguration,omitempty"` // A delimited data format where the column separator can be a comma and the record separator is a newline character.
	Dataformat interface{} `json:"DataFormat"`
	Datasources3configuration interface{} `json:"DataSourceS3Configuration"`
}

// ListDatabasesRequest represents the ListDatabasesRequest schema from the OpenAPI specification
type ListDatabasesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
}

// CreateBatchLoadTaskResponse represents the CreateBatchLoadTaskResponse schema from the OpenAPI specification
type CreateBatchLoadTaskResponse struct {
	Taskid interface{} `json:"TaskId"`
}

// DescribeTableResponse represents the DescribeTableResponse schema from the OpenAPI specification
type DescribeTableResponse struct {
	Table interface{} `json:"Table,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"Value"`
	Key interface{} `json:"Key"`
}

// MagneticStoreWriteProperties represents the MagneticStoreWriteProperties schema from the OpenAPI specification
type MagneticStoreWriteProperties struct {
	Enablemagneticstorewrites interface{} `json:"EnableMagneticStoreWrites"`
	Magneticstorerejecteddatalocation interface{} `json:"MagneticStoreRejectedDataLocation,omitempty"`
}

// CreateDatabaseRequest represents the CreateDatabaseRequest schema from the OpenAPI specification
type CreateDatabaseRequest struct {
	Databasename interface{} `json:"DatabaseName"`
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// ListBatchLoadTasksResponse represents the ListBatchLoadTasksResponse schema from the OpenAPI specification
type ListBatchLoadTasksResponse struct {
	Batchloadtasks interface{} `json:"BatchLoadTasks,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// DataModelS3Configuration represents the DataModelS3Configuration schema from the OpenAPI specification
type DataModelS3Configuration struct {
	Bucketname interface{} `json:"BucketName,omitempty"`
	Objectkey interface{} `json:"ObjectKey,omitempty"`
}

// DescribeEndpointsResponse represents the DescribeEndpointsResponse schema from the OpenAPI specification
type DescribeEndpointsResponse struct {
	Endpoints interface{} `json:"Endpoints"`
}

// ListBatchLoadTasksRequest represents the ListBatchLoadTasksRequest schema from the OpenAPI specification
type ListBatchLoadTasksRequest struct {
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Taskstatus interface{} `json:"TaskStatus,omitempty"`
}

// DeleteDatabaseRequest represents the DeleteDatabaseRequest schema from the OpenAPI specification
type DeleteDatabaseRequest struct {
	Databasename interface{} `json:"DatabaseName"`
}

// MultiMeasureAttributeMapping represents the MultiMeasureAttributeMapping schema from the OpenAPI specification
type MultiMeasureAttributeMapping struct {
	Measurevaluetype interface{} `json:"MeasureValueType,omitempty"`
	Sourcecolumn interface{} `json:"SourceColumn"`
	Targetmultimeasureattributename interface{} `json:"TargetMultiMeasureAttributeName,omitempty"`
}

// Record represents the Record schema from the OpenAPI specification
type Record struct {
	Measurevalues interface{} `json:"MeasureValues,omitempty"`
	Time interface{} `json:"Time,omitempty"`
	Timeunit interface{} `json:"TimeUnit,omitempty"`
	Version interface{} `json:"Version,omitempty"`
	Dimensions interface{} `json:"Dimensions,omitempty"`
	Measurename interface{} `json:"MeasureName,omitempty"`
	Measurevalue interface{} `json:"MeasureValue,omitempty"`
	Measurevaluetype interface{} `json:"MeasureValueType,omitempty"`
}

// DescribeBatchLoadTaskResponse represents the DescribeBatchLoadTaskResponse schema from the OpenAPI specification
type DescribeBatchLoadTaskResponse struct {
	Batchloadtaskdescription interface{} `json:"BatchLoadTaskDescription"`
}

// BatchLoadTaskDescription represents the BatchLoadTaskDescription schema from the OpenAPI specification
type BatchLoadTaskDescription struct {
	Progressreport interface{} `json:"ProgressReport,omitempty"`
	Resumableuntil interface{} `json:"ResumableUntil,omitempty"`
	Targetdatabasename interface{} `json:"TargetDatabaseName,omitempty"`
	Errormessage interface{} `json:"ErrorMessage,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Recordversion interface{} `json:"RecordVersion,omitempty"`
	Reportconfiguration interface{} `json:"ReportConfiguration,omitempty"`
	Datamodelconfiguration interface{} `json:"DataModelConfiguration,omitempty"`
	Targettablename interface{} `json:"TargetTableName,omitempty"`
	Taskid interface{} `json:"TaskId,omitempty"`
	Taskstatus interface{} `json:"TaskStatus,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Datasourceconfiguration interface{} `json:"DataSourceConfiguration,omitempty"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"Tags"`
	Resourcearn interface{} `json:"ResourceARN"`
}

// UpdateDatabaseResponse represents the UpdateDatabaseResponse schema from the OpenAPI specification
type UpdateDatabaseResponse struct {
	Database Database `json:"Database,omitempty"` // A top-level container for a table. Databases and tables are the fundamental management concepts in Amazon Timestream. All tables in a database are encrypted with the same KMS key.
}

// DimensionMapping represents the DimensionMapping schema from the OpenAPI specification
type DimensionMapping struct {
	Destinationcolumn interface{} `json:"DestinationColumn,omitempty"`
	Sourcecolumn interface{} `json:"SourceColumn,omitempty"`
}

// Database represents the Database schema from the OpenAPI specification
type Database struct {
	Kmskeyid interface{} `json:"KmsKeyId,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Tablecount interface{} `json:"TableCount,omitempty"`
	Arn interface{} `json:"Arn,omitempty"`
	Creationtime interface{} `json:"CreationTime,omitempty"`
	Databasename interface{} `json:"DatabaseName,omitempty"`
}

// CsvConfiguration represents the CsvConfiguration schema from the OpenAPI specification
type CsvConfiguration struct {
	Columnseparator interface{} `json:"ColumnSeparator,omitempty"`
	Escapechar interface{} `json:"EscapeChar,omitempty"`
	Nullvalue interface{} `json:"NullValue,omitempty"`
	Quotechar interface{} `json:"QuoteChar,omitempty"`
	Trimwhitespace interface{} `json:"TrimWhiteSpace,omitempty"`
}

// DataModelConfiguration represents the DataModelConfiguration schema from the OpenAPI specification
type DataModelConfiguration struct {
	Datamodel interface{} `json:"DataModel,omitempty"`
	Datamodels3configuration interface{} `json:"DataModelS3Configuration,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// WriteRecordsRequest represents the WriteRecordsRequest schema from the OpenAPI specification
type WriteRecordsRequest struct {
	Databasename interface{} `json:"DatabaseName"`
	Records interface{} `json:"Records"`
	Tablename interface{} `json:"TableName"`
	Commonattributes interface{} `json:"CommonAttributes,omitempty"`
}

// ListTablesResponse represents the ListTablesResponse schema from the OpenAPI specification
type ListTablesResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tables interface{} `json:"Tables,omitempty"`
}
