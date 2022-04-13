// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	conventions "go.opentelemetry.io/collector/model/semconv/v1.9.0"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for hostmetricsreceiver/process metrics.
type MetricsSettings struct {
	ProcessCPUTime             MetricSettings `mapstructure:"process.cpu.time"`
	ProcessDiskIo              MetricSettings `mapstructure:"process.disk.io"`
	ProcessMemoryPhysicalUsage MetricSettings `mapstructure:"process.memory.physical_usage"`
	ProcessMemoryVirtualUsage  MetricSettings `mapstructure:"process.memory.virtual_usage"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		ProcessCPUTime: MetricSettings{
			Enabled: true,
		},
		ProcessDiskIo: MetricSettings{
			Enabled: true,
		},
		ProcessMemoryPhysicalUsage: MetricSettings{
			Enabled: true,
		},
		ProcessMemoryVirtualUsage: MetricSettings{
			Enabled: true,
		},
	}
}

type metricProcessCPUTime struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.cpu.time metric with initial data.
func (m *metricProcessCPUTime) init() {
	m.data.SetName("process.cpu.time")
	m.data.SetDescription("Total CPU seconds broken down by different states.")
	m.data.SetUnit("s")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessCPUTime) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.State, pcommon.NewValueString(stateAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessCPUTime) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessCPUTime) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessCPUTime(settings MetricSettings) metricProcessCPUTime {
	m := metricProcessCPUTime{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessDiskIo struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.disk.io metric with initial data.
func (m *metricProcessDiskIo) init() {
	m.data.SetName("process.disk.io")
	m.data.SetDescription("Disk bytes transferred.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricProcessDiskIo) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Direction, pcommon.NewValueString(directionAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessDiskIo) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessDiskIo) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessDiskIo(settings MetricSettings) metricProcessDiskIo {
	m := metricProcessDiskIo{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryPhysicalUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.physical_usage metric with initial data.
func (m *metricProcessMemoryPhysicalUsage) init() {
	m.data.SetName("process.memory.physical_usage")
	m.data.SetDescription("The amount of physical memory in use.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessMemoryPhysicalUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryPhysicalUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryPhysicalUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryPhysicalUsage(settings MetricSettings) metricProcessMemoryPhysicalUsage {
	m := metricProcessMemoryPhysicalUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricProcessMemoryVirtualUsage struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills process.memory.virtual_usage metric with initial data.
func (m *metricProcessMemoryVirtualUsage) init() {
	m.data.SetName("process.memory.virtual_usage")
	m.data.SetDescription("Virtual memory size.")
	m.data.SetUnit("By")
	m.data.SetDataType(pmetric.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.MetricAggregationTemporalityCumulative)
}

func (m *metricProcessMemoryVirtualUsage) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricProcessMemoryVirtualUsage) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricProcessMemoryVirtualUsage) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricProcessMemoryVirtualUsage(settings MetricSettings) metricProcessMemoryVirtualUsage {
	m := metricProcessMemoryVirtualUsage{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                        pcommon.Timestamp // start time that will be applied to all recorded data points.
	metricsCapacity                  int               // maximum observed number of metrics per resource.
	resourceCapacity                 int               // maximum observed number of resource attributes.
	metricsBuffer                    pmetric.Metrics   // accumulates metrics data before emitting.
	metricProcessCPUTime             metricProcessCPUTime
	metricProcessDiskIo              metricProcessDiskIo
	metricProcessMemoryPhysicalUsage metricProcessMemoryPhysicalUsage
	metricProcessMemoryVirtualUsage  metricProcessMemoryVirtualUsage
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                        pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                    pmetric.NewMetrics(),
		metricProcessCPUTime:             newMetricProcessCPUTime(settings.ProcessCPUTime),
		metricProcessDiskIo:              newMetricProcessDiskIo(settings.ProcessDiskIo),
		metricProcessMemoryPhysicalUsage: newMetricProcessMemoryPhysicalUsage(settings.ProcessMemoryPhysicalUsage),
		metricProcessMemoryVirtualUsage:  newMetricProcessMemoryVirtualUsage(settings.ProcessMemoryVirtualUsage),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceOption applies changes to provided resource.
type ResourceOption func(pcommon.Resource)

// WithProcessCommand sets provided value as "process.command" attribute for current resource.
func WithProcessCommand(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("process.command", val)
	}
}

// WithProcessCommandLine sets provided value as "process.command_line" attribute for current resource.
func WithProcessCommandLine(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("process.command_line", val)
	}
}

// WithProcessExecutableName sets provided value as "process.executable.name" attribute for current resource.
func WithProcessExecutableName(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("process.executable.name", val)
	}
}

// WithProcessExecutablePath sets provided value as "process.executable.path" attribute for current resource.
func WithProcessExecutablePath(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("process.executable.path", val)
	}
}

// WithProcessOwner sets provided value as "process.owner" attribute for current resource.
func WithProcessOwner(val string) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertString("process.owner", val)
	}
}

// WithProcessPid sets provided value as "process.pid" attribute for current resource.
func WithProcessPid(val int64) ResourceOption {
	return func(r pcommon.Resource) {
		r.Attributes().UpsertInt("process.pid", val)
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead. Resource attributes should be provided as ResourceOption arguments.
func (mb *MetricsBuilder) EmitForResource(ro ...ResourceOption) {
	rm := pmetric.NewResourceMetrics()
	rm.SetSchemaUrl(conventions.SchemaURL)
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	for _, op := range ro {
		op(rm.Resource())
	}
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/hostmetricsreceiver/process")
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricProcessCPUTime.emit(ils.Metrics())
	mb.metricProcessDiskIo.emit(ils.Metrics())
	mb.metricProcessMemoryPhysicalUsage.emit(ils.Metrics())
	mb.metricProcessMemoryVirtualUsage.emit(ils.Metrics())
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(ro ...ResourceOption) pmetric.Metrics {
	mb.EmitForResource(ro...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordProcessCPUTimeDataPoint adds a data point to process.cpu.time metric.
func (mb *MetricsBuilder) RecordProcessCPUTimeDataPoint(ts pcommon.Timestamp, val float64, stateAttributeValue string) {
	mb.metricProcessCPUTime.recordDataPoint(mb.startTime, ts, val, stateAttributeValue)
}

// RecordProcessDiskIoDataPoint adds a data point to process.disk.io metric.
func (mb *MetricsBuilder) RecordProcessDiskIoDataPoint(ts pcommon.Timestamp, val int64, directionAttributeValue string) {
	mb.metricProcessDiskIo.recordDataPoint(mb.startTime, ts, val, directionAttributeValue)
}

// RecordProcessMemoryPhysicalUsageDataPoint adds a data point to process.memory.physical_usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryPhysicalUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryPhysicalUsage.recordDataPoint(mb.startTime, ts, val)
}

// RecordProcessMemoryVirtualUsageDataPoint adds a data point to process.memory.virtual_usage metric.
func (mb *MetricsBuilder) RecordProcessMemoryVirtualUsageDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricProcessMemoryVirtualUsage.recordDataPoint(mb.startTime, ts, val)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Direction (Direction of flow of bytes (read or write).)
	Direction string
	// State (Breakdown of CPU usage by type.)
	State string
}{
	"direction",
	"state",
}

// A is an alias for Attributes.
var A = Attributes

// AttributeDirection are the possible values that the attribute "direction" can have.
var AttributeDirection = struct {
	Read  string
	Write string
}{
	"read",
	"write",
}

// AttributeState are the possible values that the attribute "state" can have.
var AttributeState = struct {
	System string
	User   string
	Wait   string
}{
	"system",
	"user",
	"wait",
}
