// Code generated by Thrift Compiler (0.14.1). DO NOT EDIT.

package addsvc

import (
	"bytes"
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"time"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = time.Now
var _ = bytes.Equal

// Attributes:
//  - Value
//  - Err
type SumReply struct {
	Value int64  `thrift:"value,1" db:"value" json:"value"`
	Err   string `thrift:"err,2" db:"err" json:"err"`
}

func NewSumReply() *SumReply {
	return &SumReply{}
}

func (p *SumReply) GetValue() int64 {
	return p.Value
}

func (p *SumReply) GetErr() string {
	return p.Err
}
func (p *SumReply) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SumReply) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *SumReply) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Err = v
	}
	return nil
}

func (p *SumReply) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "SumReply"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SumReply) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "value", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:value: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.Value)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:value: ", p), err)
	}
	return err
}

func (p *SumReply) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "err", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:err: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Err)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.err (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:err: ", p), err)
	}
	return err
}

func (p *SumReply) Equals(other *SumReply) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Value != other.Value {
		return false
	}
	if p.Err != other.Err {
		return false
	}
	return true
}

func (p *SumReply) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SumReply(%+v)", *p)
}

// Attributes:
//  - Value
//  - Err
type ConcatReply struct {
	Value string `thrift:"value,1" db:"value" json:"value"`
	Err   string `thrift:"err,2" db:"err" json:"err"`
}

func NewConcatReply() *ConcatReply {
	return &ConcatReply{}
}

func (p *ConcatReply) GetValue() string {
	return p.Value
}

func (p *ConcatReply) GetErr() string {
	return p.Err
}
func (p *ConcatReply) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *ConcatReply) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *ConcatReply) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Err = v
	}
	return nil
}

func (p *ConcatReply) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "ConcatReply"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *ConcatReply) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "value", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:value: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Value)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.value (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:value: ", p), err)
	}
	return err
}

func (p *ConcatReply) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "err", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:err: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.Err)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.err (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:err: ", p), err)
	}
	return err
}

func (p *ConcatReply) Equals(other *ConcatReply) bool {
	if p == other {
		return true
	} else if p == nil || other == nil {
		return false
	}
	if p.Value != other.Value {
		return false
	}
	if p.Err != other.Err {
		return false
	}
	return true
}

func (p *ConcatReply) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ConcatReply(%+v)", *p)
}

type AddService interface {
	// Parameters:
	//  - A
	//  - B
	Sum(ctx context.Context, a int64, b int64) (_r *SumReply, _err error)
	// Parameters:
	//  - A
	//  - B
	Concat(ctx context.Context, a string, b string) (_r *ConcatReply, _err error)
}

type AddServiceClient struct {
	c    thrift.TClient
	meta thrift.ResponseMeta
}

func NewAddServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *AddServiceClient {
	return &AddServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewAddServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *AddServiceClient {
	return &AddServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewAddServiceClient(c thrift.TClient) *AddServiceClient {
	return &AddServiceClient{
		c: c,
	}
}

func (p *AddServiceClient) Client_() thrift.TClient {
	return p.c
}

func (p *AddServiceClient) LastResponseMeta_() thrift.ResponseMeta {
	return p.meta
}

func (p *AddServiceClient) SetLastResponseMeta_(meta thrift.ResponseMeta) {
	p.meta = meta
}

// Parameters:
//  - A
//  - B
func (p *AddServiceClient) Sum(ctx context.Context, a int64, b int64) (_r *SumReply, _err error) {
	var _args0 AddServiceSumArgs
	_args0.A = a
	_args0.B = b
	var _result2 AddServiceSumResult
	var _meta1 thrift.ResponseMeta
	_meta1, _err = p.Client_().Call(ctx, "Sum", &_args0, &_result2)
	p.SetLastResponseMeta_(_meta1)
	if _err != nil {
		return
	}
	return _result2.GetSuccess(), nil
}

// Parameters:
//  - A
//  - B
func (p *AddServiceClient) Concat(ctx context.Context, a string, b string) (_r *ConcatReply, _err error) {
	var _args3 AddServiceConcatArgs
	_args3.A = a
	_args3.B = b
	var _result5 AddServiceConcatResult
	var _meta4 thrift.ResponseMeta
	_meta4, _err = p.Client_().Call(ctx, "Concat", &_args3, &_result5)
	p.SetLastResponseMeta_(_meta4)
	if _err != nil {
		return
	}
	return _result5.GetSuccess(), nil
}

type AddServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      AddService
}

func (p *AddServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *AddServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *AddServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewAddServiceProcessor(handler AddService) *AddServiceProcessor {

	self6 := &AddServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self6.processorMap["Sum"] = &addServiceProcessorSum{handler: handler}
	self6.processorMap["Concat"] = &addServiceProcessorConcat{handler: handler}
	return self6
}

func (p *AddServiceProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err2 := iprot.ReadMessageBegin(ctx)
	if err2 != nil {
		return false, thrift.WrapTException(err2)
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(ctx, thrift.STRUCT)
	iprot.ReadMessageEnd(ctx)
	x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(ctx, name, thrift.EXCEPTION, seqId)
	x7.Write(ctx, oprot)
	oprot.WriteMessageEnd(ctx)
	oprot.Flush(ctx)
	return false, x7

}

type addServiceProcessorSum struct {
	handler AddService
}

func (p *addServiceProcessorSum) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AddServiceSumArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "Sum", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := AddServiceSumResult{}
	var retval *SumReply
	if retval, err2 = p.handler.Sum(ctx, args.A, args.B); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Sum: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "Sum", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "Sum", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

type addServiceProcessorConcat struct {
	handler AddService
}

func (p *addServiceProcessorConcat) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := AddServiceConcatArgs{}
	var err2 error
	if err2 = args.Read(ctx, iprot); err2 != nil {
		iprot.ReadMessageEnd(ctx)
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err2.Error())
		oprot.WriteMessageBegin(ctx, "Concat", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return false, thrift.WrapTException(err2)
	}
	iprot.ReadMessageEnd(ctx)

	tickerCancel := func() {}
	// Start a goroutine to do server side connectivity check.
	if thrift.ServerConnectivityCheckInterval > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithCancel(ctx)
		defer cancel()
		var tickerCtx context.Context
		tickerCtx, tickerCancel = context.WithCancel(context.Background())
		defer tickerCancel()
		go func(ctx context.Context, cancel context.CancelFunc) {
			ticker := time.NewTicker(thrift.ServerConnectivityCheckInterval)
			defer ticker.Stop()
			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					if !iprot.Transport().IsOpen() {
						cancel()
						return
					}
				}
			}
		}(tickerCtx, cancel)
	}

	result := AddServiceConcatResult{}
	var retval *ConcatReply
	if retval, err2 = p.handler.Concat(ctx, args.A, args.B); err2 != nil {
		tickerCancel()
		if err2 == thrift.ErrAbandonRequest {
			return false, thrift.WrapTException(err2)
		}
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Concat: "+err2.Error())
		oprot.WriteMessageBegin(ctx, "Concat", thrift.EXCEPTION, seqId)
		x.Write(ctx, oprot)
		oprot.WriteMessageEnd(ctx)
		oprot.Flush(ctx)
		return true, thrift.WrapTException(err2)
	} else {
		result.Success = retval
	}
	tickerCancel()
	if err2 = oprot.WriteMessageBegin(ctx, "Concat", thrift.REPLY, seqId); err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = result.Write(ctx, oprot); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.WriteMessageEnd(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = thrift.WrapTException(err2)
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - A
//  - B
type AddServiceSumArgs struct {
	A int64 `thrift:"a,1" db:"a" json:"a"`
	B int64 `thrift:"b,2" db:"b" json:"b"`
}

func NewAddServiceSumArgs() *AddServiceSumArgs {
	return &AddServiceSumArgs{}
}

func (p *AddServiceSumArgs) GetA() int64 {
	return p.A
}

func (p *AddServiceSumArgs) GetB() int64 {
	return p.B
}
func (p *AddServiceSumArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.I64 {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AddServiceSumArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.A = v
	}
	return nil
}

func (p *AddServiceSumArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.B = v
	}
	return nil
}

func (p *AddServiceSumArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Sum_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AddServiceSumArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "a", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:a: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.A)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.a (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:a: ", p), err)
	}
	return err
}

func (p *AddServiceSumArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "b", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:b: ", p), err)
	}
	if err := oprot.WriteI64(ctx, int64(p.B)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.b (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:b: ", p), err)
	}
	return err
}

func (p *AddServiceSumArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddServiceSumArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AddServiceSumResult struct {
	Success *SumReply `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAddServiceSumResult() *AddServiceSumResult {
	return &AddServiceSumResult{}
}

var AddServiceSumResult_Success_DEFAULT *SumReply

func (p *AddServiceSumResult) GetSuccess() *SumReply {
	if !p.IsSetSuccess() {
		return AddServiceSumResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AddServiceSumResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddServiceSumResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AddServiceSumResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &SumReply{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AddServiceSumResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Sum_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AddServiceSumResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AddServiceSumResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddServiceSumResult(%+v)", *p)
}

// Attributes:
//  - A
//  - B
type AddServiceConcatArgs struct {
	A string `thrift:"a,1" db:"a" json:"a"`
	B string `thrift:"b,2" db:"b" json:"b"`
}

func NewAddServiceConcatArgs() *AddServiceConcatArgs {
	return &AddServiceConcatArgs{}
}

func (p *AddServiceConcatArgs) GetA() string {
	return p.A
}

func (p *AddServiceConcatArgs) GetB() string {
	return p.B
}
func (p *AddServiceConcatArgs) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField2(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AddServiceConcatArgs) ReadField1(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.A = v
	}
	return nil
}

func (p *AddServiceConcatArgs) ReadField2(ctx context.Context, iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(ctx); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.B = v
	}
	return nil
}

func (p *AddServiceConcatArgs) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Concat_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(ctx, oprot); err != nil {
			return err
		}
		if err := p.writeField2(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AddServiceConcatArgs) writeField1(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "a", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:a: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.A)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.a (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:a: ", p), err)
	}
	return err
}

func (p *AddServiceConcatArgs) writeField2(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin(ctx, "b", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:b: ", p), err)
	}
	if err := oprot.WriteString(ctx, string(p.B)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.b (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:b: ", p), err)
	}
	return err
}

func (p *AddServiceConcatArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddServiceConcatArgs(%+v)", *p)
}

// Attributes:
//  - Success
type AddServiceConcatResult struct {
	Success *ConcatReply `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewAddServiceConcatResult() *AddServiceConcatResult {
	return &AddServiceConcatResult{}
}

var AddServiceConcatResult_Success_DEFAULT *ConcatReply

func (p *AddServiceConcatResult) GetSuccess() *ConcatReply {
	if !p.IsSetSuccess() {
		return AddServiceConcatResult_Success_DEFAULT
	}
	return p.Success
}
func (p *AddServiceConcatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddServiceConcatResult) Read(ctx context.Context, iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin(ctx)
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField0(ctx, iprot); err != nil {
					return err
				}
			} else {
				if err := iprot.Skip(ctx, fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(ctx, fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(ctx); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(ctx); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *AddServiceConcatResult) ReadField0(ctx context.Context, iprot thrift.TProtocol) error {
	p.Success = &ConcatReply{}
	if err := p.Success.Read(ctx, iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *AddServiceConcatResult) Write(ctx context.Context, oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(ctx, "Concat_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(ctx, oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(ctx); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(ctx); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *AddServiceConcatResult) writeField0(ctx context.Context, oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin(ctx, "success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(ctx, oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(ctx); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *AddServiceConcatResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddServiceConcatResult(%+v)", *p)
}
