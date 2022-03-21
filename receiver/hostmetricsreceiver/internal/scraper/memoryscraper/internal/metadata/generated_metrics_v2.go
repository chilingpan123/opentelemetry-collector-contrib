// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/memory metrics.
type MetricsSettings struct {
	SystemMemoryUsage       MetricSettings `mapstructure:"system.memory.usage"`
	SystemMemoryUtilization MetricSettings `mapstructure:"system.memory.utilization"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemMemoryUsage: MetricSettings{
			Enabled: true,
		},
		SystemMemoryUtilization: MetricSettings{
			Enabled: false,
		},
	}
}

type metricSystemMemoryUsage struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.memory.usage metric with initial data.
func (m *metricSystemMemoryUsage) init() {
	m.data.SetName("system.memory.usage")
	m.data.SetDescription("Bytes of memory in use.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemMemoryUsage) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemMemoryUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemMemoryUsage) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemMemoryUsage(settings MetricSettings) metricSystemMemoryUsage {
	m := metricSystemMemoryUsage{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemMemoryUtilization struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.memory.utilization metric with initial data.
func (m *metricSystemMemoryUtilization) init() {
	m.data.SetName("system.memory.utilization")
	m.data.SetDescription("Percentage of memory bytes in use.")
	m.data.SetUnit("1")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemMemoryUtilization) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.State, pdata.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemMemoryUtilization) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemMemoryUtilization) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemMemoryUtilization(settings MetricSettings) metricSystemMemoryUtilization {
	m := metricSystemMemoryUtilization{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                     pdata.Timestamp
	metricSystemMemoryUsage       metricSystemMemoryUsage
	metricSystemMemoryUtilization metricSystemMemoryUtilization
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                     pdata.NewTimestampFromTime(time.Now()),
		metricSystemMemoryUsage:       newMetricSystemMemoryUsage(settings.SystemMemoryUsage),
		metricSystemMemoryUtilization: newMetricSystemMemoryUtilization(settings.SystemMemoryUtilization),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	mb.metricSystemMemoryUsage.emit(metrics)
	mb.metricSystemMemoryUtilization.emit(metrics)
}

// RecordSystemMemoryUsageDataPoint adds a data point to system.memory.usage metric.
func (mb *MetricsBuilder) RecordSystemMemoryUsageDataPoint(ts pdata.Timestamp, val int64, stateAttributeValue string) {
	mb.metricSystemMemoryUsage.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// RecordSystemMemoryUtilizationDataPoint adds a data point to system.memory.utilization metric.
func (mb *MetricsBuilder) RecordSystemMemoryUtilizationDataPoint(ts pdata.Timestamp, val float64, stateAttributeValue string) {
	mb.metricSystemMemoryUtilization.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// NewMetricData creates new pdata.Metrics and sets the InstrumentationLibrary
// name on the ResourceMetrics.
func (mb *MetricsBuilder) NewMetricData() pdata.Metrics {
	md := pdata.NewMetrics()
	rm := md.ResourceMetrics().AppendEmpty()
	ilm := rm.InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/hostmetricsreceiver/memory")
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// State (Breakdown of memory usage by type.)
	State string
}{
	"state",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	Buffered          string
	Cached            string
	Inactive          string
	Free              string
	SlabReclaimable   string
	SlabUnreclaimable string
	Used              string
}{
	"buffered",
	"cached",
	"inactive",
	"free",
	"slab_reclaimable",
	"slab_unreclaimable",
	"used",
}
