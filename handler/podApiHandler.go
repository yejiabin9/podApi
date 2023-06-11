package handler

import (
	"context"
	"encoding/json"
	errors2 "errors"
	"github.com/sirupsen/logrus"
	"github.com/yejiabin9/pod/proto/pod"
	"github.com/yejiabin9/podApi/proto/protoApi"
	"net/http"
	"strconv"
)

type PodApi struct {
	PodService pod.PodService
}

func (p *PodApi) FindPodById(ctx context.Context, in *protoApi.Request, out *protoApi.Response) error {
	logrus.Info("accept podApiFindPodById request")
	if _, ok := in.Get["pod_id"]; !ok {
		out.StatusCode = 500
		return errors2.New("request params error")
	}
	podIDString := in.Get["pod_id"].Values[0]
	podID, err := strconv.ParseInt(podIDString, 10, 64)
	if err != nil {
		return err
	}
	podInfo, err := p.PodService.FindPodByID(ctx, &pod.PodID{
		PodId: int32(podID),
	})
	if err != nil {
		return err
	}
	out.StatusCode = 200
	b, _ := json.Marshal(podInfo)
	out.Body = string(b)
	return nil
}

func (p *PodApi) AddPod(ctx context.Context, in *protoApi.Request, out *protoApi.Response) error {
	logrus.Info("accept podApi.AddPod request")
	out.StatusCode = http.StatusOK
	b, _ := json.Marshal("{success :'request /podApi/AddPod'}")
	out.Body = string(b)
	return nil

}

func (p *PodApi) DeletePodById(ctx context.Context, req *protoApi.Request, rsp *protoApi.Response) error {
	logrus.Info("accept podApi.DeletePodById request")
	rsp.StatusCode = http.StatusOK
	b, _ := json.Marshal("{success :'request /podApi/DeletePodById'}")
	rsp.Body = string(b)
	return nil
}

func (p *PodApi) UpdatePod(ctx context.Context, req *protoApi.Request, rsp *protoApi.Response) error {
	logrus.Info("accept podApi.UpdatePod request")
	rsp.StatusCode = http.StatusOK
	b, _ := json.Marshal("{success :'request /podApi/UpdatePod'}")
	rsp.Body = string(b)
	return nil
}

// 一个默认方法
func (p *PodApi) Call(ctx context.Context, req *protoApi.Request, rsp *protoApi.Response) error {
	logrus.Info("accept podApi.Call request")
	rsp.StatusCode = http.StatusOK
	b, _ := json.Marshal("{success :'request /podApi/Call'}")
	rsp.Body = string(b)
	return nil
}
