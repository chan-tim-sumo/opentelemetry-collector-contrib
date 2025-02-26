// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/filter"
)

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms)
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for mongodb metrics.
type MetricsConfig struct {
	MongodbActiveReads            MetricConfig `mapstructure:"mongodb.active.reads"`
	MongodbActiveWrites           MetricConfig `mapstructure:"mongodb.active.writes"`
	MongodbCacheOperations        MetricConfig `mapstructure:"mongodb.cache.operations"`
	MongodbCollectionCount        MetricConfig `mapstructure:"mongodb.collection.count"`
	MongodbCommandsRate           MetricConfig `mapstructure:"mongodb.commands.rate"`
	MongodbConnectionCount        MetricConfig `mapstructure:"mongodb.connection.count"`
	MongodbCursorCount            MetricConfig `mapstructure:"mongodb.cursor.count"`
	MongodbCursorTimeoutCount     MetricConfig `mapstructure:"mongodb.cursor.timeout.count"`
	MongodbDataSize               MetricConfig `mapstructure:"mongodb.data.size"`
	MongodbDatabaseCount          MetricConfig `mapstructure:"mongodb.database.count"`
	MongodbDeletesRate            MetricConfig `mapstructure:"mongodb.deletes.rate"`
	MongodbDocumentOperationCount MetricConfig `mapstructure:"mongodb.document.operation.count"`
	MongodbExtentCount            MetricConfig `mapstructure:"mongodb.extent.count"`
	MongodbFlushesRate            MetricConfig `mapstructure:"mongodb.flushes.rate"`
	MongodbGetmoresRate           MetricConfig `mapstructure:"mongodb.getmores.rate"`
	MongodbGlobalLockTime         MetricConfig `mapstructure:"mongodb.global_lock.time"`
	MongodbHealth                 MetricConfig `mapstructure:"mongodb.health"`
	MongodbIndexAccessCount       MetricConfig `mapstructure:"mongodb.index.access.count"`
	MongodbIndexCount             MetricConfig `mapstructure:"mongodb.index.count"`
	MongodbIndexSize              MetricConfig `mapstructure:"mongodb.index.size"`
	MongodbInsertsRate            MetricConfig `mapstructure:"mongodb.inserts.rate"`
	MongodbLockAcquireCount       MetricConfig `mapstructure:"mongodb.lock.acquire.count"`
	MongodbLockAcquireTime        MetricConfig `mapstructure:"mongodb.lock.acquire.time"`
	MongodbLockAcquireWaitCount   MetricConfig `mapstructure:"mongodb.lock.acquire.wait_count"`
	MongodbLockDeadlockCount      MetricConfig `mapstructure:"mongodb.lock.deadlock.count"`
	MongodbMemoryUsage            MetricConfig `mapstructure:"mongodb.memory.usage"`
	MongodbNetworkIoReceive       MetricConfig `mapstructure:"mongodb.network.io.receive"`
	MongodbNetworkIoTransmit      MetricConfig `mapstructure:"mongodb.network.io.transmit"`
	MongodbNetworkRequestCount    MetricConfig `mapstructure:"mongodb.network.request.count"`
	MongodbObjectCount            MetricConfig `mapstructure:"mongodb.object.count"`
	MongodbOperationCount         MetricConfig `mapstructure:"mongodb.operation.count"`
	MongodbOperationLatencyTime   MetricConfig `mapstructure:"mongodb.operation.latency.time"`
	MongodbOperationReplCount     MetricConfig `mapstructure:"mongodb.operation.repl.count"`
	MongodbOperationTime          MetricConfig `mapstructure:"mongodb.operation.time"`
	MongodbPageFaults             MetricConfig `mapstructure:"mongodb.page_faults"`
	MongodbQueriesRate            MetricConfig `mapstructure:"mongodb.queries.rate"`
	MongodbSessionCount           MetricConfig `mapstructure:"mongodb.session.count"`
	MongodbStorageSize            MetricConfig `mapstructure:"mongodb.storage.size"`
	MongodbUpdatesRate            MetricConfig `mapstructure:"mongodb.updates.rate"`
	MongodbUptime                 MetricConfig `mapstructure:"mongodb.uptime"`
	MongodbWtcacheBytesRead       MetricConfig `mapstructure:"mongodb.wtcache.bytes.read"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		MongodbActiveReads: MetricConfig{
			Enabled: false,
		},
		MongodbActiveWrites: MetricConfig{
			Enabled: false,
		},
		MongodbCacheOperations: MetricConfig{
			Enabled: true,
		},
		MongodbCollectionCount: MetricConfig{
			Enabled: true,
		},
		MongodbCommandsRate: MetricConfig{
			Enabled: false,
		},
		MongodbConnectionCount: MetricConfig{
			Enabled: true,
		},
		MongodbCursorCount: MetricConfig{
			Enabled: true,
		},
		MongodbCursorTimeoutCount: MetricConfig{
			Enabled: true,
		},
		MongodbDataSize: MetricConfig{
			Enabled: true,
		},
		MongodbDatabaseCount: MetricConfig{
			Enabled: true,
		},
		MongodbDeletesRate: MetricConfig{
			Enabled: false,
		},
		MongodbDocumentOperationCount: MetricConfig{
			Enabled: true,
		},
		MongodbExtentCount: MetricConfig{
			Enabled: true,
		},
		MongodbFlushesRate: MetricConfig{
			Enabled: false,
		},
		MongodbGetmoresRate: MetricConfig{
			Enabled: false,
		},
		MongodbGlobalLockTime: MetricConfig{
			Enabled: true,
		},
		MongodbHealth: MetricConfig{
			Enabled: false,
		},
		MongodbIndexAccessCount: MetricConfig{
			Enabled: true,
		},
		MongodbIndexCount: MetricConfig{
			Enabled: true,
		},
		MongodbIndexSize: MetricConfig{
			Enabled: true,
		},
		MongodbInsertsRate: MetricConfig{
			Enabled: false,
		},
		MongodbLockAcquireCount: MetricConfig{
			Enabled: false,
		},
		MongodbLockAcquireTime: MetricConfig{
			Enabled: false,
		},
		MongodbLockAcquireWaitCount: MetricConfig{
			Enabled: false,
		},
		MongodbLockDeadlockCount: MetricConfig{
			Enabled: false,
		},
		MongodbMemoryUsage: MetricConfig{
			Enabled: true,
		},
		MongodbNetworkIoReceive: MetricConfig{
			Enabled: true,
		},
		MongodbNetworkIoTransmit: MetricConfig{
			Enabled: true,
		},
		MongodbNetworkRequestCount: MetricConfig{
			Enabled: true,
		},
		MongodbObjectCount: MetricConfig{
			Enabled: true,
		},
		MongodbOperationCount: MetricConfig{
			Enabled: true,
		},
		MongodbOperationLatencyTime: MetricConfig{
			Enabled: false,
		},
		MongodbOperationReplCount: MetricConfig{
			Enabled: false,
		},
		MongodbOperationTime: MetricConfig{
			Enabled: true,
		},
		MongodbPageFaults: MetricConfig{
			Enabled: false,
		},
		MongodbQueriesRate: MetricConfig{
			Enabled: false,
		},
		MongodbSessionCount: MetricConfig{
			Enabled: true,
		},
		MongodbStorageSize: MetricConfig{
			Enabled: true,
		},
		MongodbUpdatesRate: MetricConfig{
			Enabled: false,
		},
		MongodbUptime: MetricConfig{
			Enabled: false,
		},
		MongodbWtcacheBytesRead: MetricConfig{
			Enabled: false,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
	// Experimental: MetricsInclude defines a list of filters for attribute values.
	// If the list is not empty, only metrics with matching resource attribute values will be emitted.
	MetricsInclude []filter.Config `mapstructure:"metrics_include"`
	// Experimental: MetricsExclude defines a list of filters for attribute values.
	// If the list is not empty, metrics with matching resource attribute values will not be emitted.
	// MetricsInclude has higher priority than MetricsExclude.
	MetricsExclude []filter.Config `mapstructure:"metrics_exclude"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac)
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for mongodb resource attributes.
type ResourceAttributesConfig struct {
	Database      ResourceAttributeConfig `mapstructure:"database"`
	ServerAddress ResourceAttributeConfig `mapstructure:"server.address"`
	ServerPort    ResourceAttributeConfig `mapstructure:"server.port"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		Database: ResourceAttributeConfig{
			Enabled: true,
		},
		ServerAddress: ResourceAttributeConfig{
			Enabled: true,
		},
		ServerPort: ResourceAttributeConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for mongodb metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
