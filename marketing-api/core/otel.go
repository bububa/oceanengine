package core

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

const instrumentationName = "github.com/bububa/oceanengine"

type Otel struct {
	traceProvider trace.TracerProvider
	tracer        trace.Tracer
	meterProvider metric.MeterProvider
	meter         metric.Meter
	histogram     metric.Int64Histogram
	counter       metric.Int64Counter
	attrs         []attribute.KeyValue
}

func NewOtel(namespace string, appID uint64) *Otel {
	ret := &Otel{
		traceProvider: otel.GetTracerProvider(),
		meterProvider: otel.GetMeterProvider(),
		attrs: []attribute.KeyValue{
			attribute.String("sdk", "oceanengine"),
			attribute.String("appid", strconv.FormatUint(appID, 10)),
		},
	}
	if namespace == "" {
		namespace = instrumentationName
	}
	ret.tracer = ret.traceProvider.Tracer(namespace)
	ret.meter = ret.meterProvider.Meter(namespace)
	if histogram, err := ret.meter.Int64Histogram(
		"http.request_timing",
		metric.WithDescription("oeanengine api request timing histogram"),
		metric.WithUnit("milliseconds"),
	); err == nil {
		ret.histogram = histogram
	}
	if counter, err := ret.meter.Int64Counter(
		"http.request_count",
		metric.WithDescription("oeanengine api request count"),
		metric.WithUnit("1"),
	); err == nil {
		ret.counter = counter
	}
	return ret
}

func (o *Otel) WithSpan(ctx context.Context, req *http.Request, resp model.Response, payload []byte, fn func(*http.Request, model.Response) (*http.Response, error)) error {
	startTime := time.Now()
	attrs := append(o.attrs,
		semconv.HTTPURLKey.String(req.URL.String()),
		semconv.HTTPMethodKey.String(req.Method),
		semconv.HTTPTargetKey.String(req.URL.Path),
		semconv.HTTPHostKey.String(req.URL.Host),
		semconv.HTTPRequestContentLengthKey.Int64(req.ContentLength),
	)
	if payload != nil {
		attrs = append(attrs, attribute.String("payload", string(payload)))
	}

	_, span := o.tracer.Start(ctx, util.StringsJoin("http.", req.Method),
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(attrs...),
	)
	defer span.End()
	httpResp, err := fn(req, resp)
	if o.histogram != nil {
		o.histogram.Record(ctx, time.Since(startTime).Milliseconds(), metric.WithAttributes(o.attrs...))
	}
	if o.counter != nil {
		counterAttrs := append(o.attrs, semconv.HTTPTargetKey.String(req.URL.Path))
		o.counter.Add(ctx, 1, metric.WithAttributes(counterAttrs...))
	}
	if !span.IsRecording() {
		return err
	}
	if httpResp != nil {
		span.SetAttributes(semconv.HTTPStatusCodeKey.Int(httpResp.StatusCode), semconv.HTTPResponseContentLengthKey.Int64(httpResp.ContentLength))
		switch httpResp.ProtoMajor {
		case 1:
			if httpResp.ProtoMinor == 1 {
				span.SetAttributes(semconv.HTTPFlavorHTTP11)
			} else {
				span.SetAttributes(semconv.HTTPFlavorHTTP10)
			}
		case 2:
			span.SetAttributes(semconv.HTTPFlavorHTTP20)
		}
	}
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
	return err
}
