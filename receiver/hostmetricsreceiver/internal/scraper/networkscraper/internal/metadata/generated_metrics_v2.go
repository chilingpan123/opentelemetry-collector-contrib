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

// MetricsSettings provides settings for hostmetricsreceiver/network metrics.
type MetricsSettings struct {
	SystemNetworkConnections MetricSettings `mapstructure:"system.network.connections"`
	SystemNetworkDropped     MetricSettings `mapstructure:"system.network.dropped"`
	SystemNetworkErrors      MetricSettings `mapstructure:"system.network.errors"`
	SystemNetworkIo          MetricSettings `mapstructure:"system.network.io"`
	SystemNetworkPackets     MetricSettings `mapstructure:"system.network.packets"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		SystemNetworkConnections: MetricSettings{
			Enabled: true,
		},
		SystemNetworkDropped: MetricSettings{
			Enabled: true,
		},
		SystemNetworkErrors: MetricSettings{
			Enabled: true,
		},
		SystemNetworkIo: MetricSettings{
			Enabled: true,
		},
		SystemNetworkPackets: MetricSettings{
			Enabled: true,
		},
	}
}

type metricSystemNetworkConnections struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.connections metric with initial data.
func (m *metricSystemNetworkConnections) init() {
	m.data.SetName("system.network.connections")
	m.data.SetDescription("The number of connections.")
	m.data.SetUnit("{connections}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkConnections) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, protocolAttributeValue string, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Protocol, pdata.NewAttributeValueString(protocolAttributeValue))
	dp.Attributes().Insert(A.State, pdata.NewAttributeValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkConnections) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkConnections) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkConnections(settings MetricSettings) metricSystemNetworkConnections {
	m := metricSystemNetworkConnections{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkDropped struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.dropped metric with initial data.
func (m *metricSystemNetworkDropped) init() {
	m.data.SetName("system.network.dropped")
	m.data.SetDescription("The number of packets dropped.")
	m.data.SetUnit("{packets}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkDropped) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewAttributeValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewAttributeValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkDropped) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkDropped) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkDropped(settings MetricSettings) metricSystemNetworkDropped {
	m := metricSystemNetworkDropped{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkErrors struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.errors metric with initial data.
func (m *metricSystemNetworkErrors) init() {
	m.data.SetName("system.network.errors")
	m.data.SetDescription("The number of errors encountered.")
	m.data.SetUnit("{errors}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkErrors) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewAttributeValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewAttributeValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkErrors) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkErrors) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkErrors(settings MetricSettings) metricSystemNetworkErrors {
	m := metricSystemNetworkErrors{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkIo struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.io metric with initial data.
func (m *metricSystemNetworkIo) init() {
	m.data.SetName("system.network.io")
	m.data.SetDescription("The number of bytes transmitted and received.")
	m.data.SetUnit("By")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkIo) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewAttributeValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewAttributeValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkIo) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkIo(settings MetricSettings) metricSystemNetworkIo {
	m := metricSystemNetworkIo{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricSystemNetworkPackets struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills system.network.packets metric with initial data.
func (m *metricSystemNetworkPackets) init() {
	m.data.SetName("system.network.packets")
	m.data.SetDescription("The number of packets transferred.")
	m.data.SetUnit("{packets}")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricSystemNetworkPackets) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Device, pdata.NewAttributeValueString(deviceAttributeValue))
	dp.Attributes().Insert(A.Direction, pdata.NewAttributeValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricSystemNetworkPackets) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricSystemNetworkPackets) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricSystemNetworkPackets(settings MetricSettings) metricSystemNetworkPackets {
	m := metricSystemNetworkPackets{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                      pdata.Timestamp
	metricSystemNetworkConnections metricSystemNetworkConnections
	metricSystemNetworkDropped     metricSystemNetworkDropped
	metricSystemNetworkErrors      metricSystemNetworkErrors
	metricSystemNetworkIo          metricSystemNetworkIo
	metricSystemNetworkPackets     metricSystemNetworkPackets
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
		startTime:                      pdata.NewTimestampFromTime(time.Now()),
		metricSystemNetworkConnections: newMetricSystemNetworkConnections(settings.SystemNetworkConnections),
		metricSystemNetworkDropped:     newMetricSystemNetworkDropped(settings.SystemNetworkDropped),
		metricSystemNetworkErrors:      newMetricSystemNetworkErrors(settings.SystemNetworkErrors),
		metricSystemNetworkIo:          newMetricSystemNetworkIo(settings.SystemNetworkIo),
		metricSystemNetworkPackets:     newMetricSystemNetworkPackets(settings.SystemNetworkPackets),
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
	mb.metricSystemNetworkConnections.emit(metrics)
	mb.metricSystemNetworkDropped.emit(metrics)
	mb.metricSystemNetworkErrors.emit(metrics)
	mb.metricSystemNetworkIo.emit(metrics)
	mb.metricSystemNetworkPackets.emit(metrics)
}

// RecordSystemNetworkConnectionsDataPoint adds a data point to system.network.connections metric.
func (mb *MetricsBuilder) RecordSystemNetworkConnectionsDataPoint(ts pdata.Timestamp, val int64, protocolAttributeValue string, stateAttributeValue string) {
	mb.metricSystemNetworkConnections.recordDataPoint(mb.startTime, ts, val, protocolAttributeValue, stateAttributeValue)
}

// RecordSystemNetworkDroppedDataPoint adds a data point to system.network.dropped metric.
func (mb *MetricsBuilder) RecordSystemNetworkDroppedDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemNetworkDropped.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemNetworkErrorsDataPoint adds a data point to system.network.errors metric.
func (mb *MetricsBuilder) RecordSystemNetworkErrorsDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemNetworkErrors.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemNetworkIoDataPoint adds a data point to system.network.io metric.
func (mb *MetricsBuilder) RecordSystemNetworkIoDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemNetworkIo.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
}

// RecordSystemNetworkPacketsDataPoint adds a data point to system.network.packets metric.
func (mb *MetricsBuilder) RecordSystemNetworkPacketsDataPoint(ts pdata.Timestamp, val int64, deviceAttributeValue string, directionAttributeValue string) {
	mb.metricSystemNetworkPackets.recordDataPoint(mb.startTime, ts, val, deviceAttributeValue, directionAttributeValue)
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
	ilm.InstrumentationLibrary().SetName("otelcol/hostmetricsreceiver/network")
	return md
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Device (Name of the network interface.)
	Device string
	// Direction (Direction of flow of bytes/operations (receive or transmit).)
	Direction string
	// Protocol (Network protocol, e.g. TCP or UDP.)
	Protocol string
	// State (State of the network connection.)
	State string
}{
	"device",
	"direction",
	"protocol",
	"state",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	Receive  string
	Transmit string
}{
	"receive",
	"transmit",
}

// AttributeProtocol are the possible values that the attribute "protocol" can have.
var AttributeProtocol = struct {
	Tcp string
}{
	"tcp",
}
