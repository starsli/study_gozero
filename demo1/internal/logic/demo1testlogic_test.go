package logic

import (
	"context"
	"demo1/demo1_pb"
	"demo1/internal/svc"
	"reflect"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestNewDemo1testLogic(t *testing.T) {
	type args struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
	}
	tests := []struct {
		name string
		args args
		want *Demo1testLogic
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDemo1testLogic(tt.args.ctx, tt.args.svcCtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDemo1testLogic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDemo1testLogic_Demo1Test(t *testing.T) {
	type fields struct {
		ctx    context.Context
		svcCtx *svc.ServiceContext
		Logger logx.Logger
	}
	type args struct {
		in *demo1_pb.Demo1Req
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *demo1_pb.Demo2Rsp
		wantErr bool
	}{
		{
			name: "Demo1Test",
			fields: fields{
				ctx:    context.Background(),
				svcCtx: &svc.ServiceContext{},
				Logger: logx.WithContext(context.Background()),
			},
			args: args{
				in: &demo1_pb.Demo1Req{
					InputParams: "123",
				},
			},
			want: &demo1_pb.Demo2Rsp{
				OutputParams: "123123",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Demo1testLogic{
				ctx:    tt.fields.ctx,
				svcCtx: tt.fields.svcCtx,
				Logger: tt.fields.Logger,
			}
			got, err := l.Demo1Test(tt.args.in)
			logx.Infof("starsli Demo1testLogic.Demo1Test() = %v, wantErr %v", got, err)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Demo1testLogic.Demo1Test() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got.OutputParams != tt.want.OutputParams {
				t.Errorf("Demo1testLogic.Demo1Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
