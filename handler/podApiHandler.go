package handler

import (
	"context"
	"encoding/json"
	errors2 "errors"
	"github.com/sirupsen/logrus"
	"github.com/yejiabin9/pod/proto/pod"
	"github.com/yejiabin9/podApi/plugin/form"
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

func (p *PodApi) AddPod(ctx context.Context, req *protoApi.Request, out *protoApi.Response) error {
	logrus.Info("accept podApi.AddPod request")
	addPodInfo := &pod.PodInfo{}
	dataSlice, ok := req.Get["pod_port"]
	if ok {
		podSlice := []*pod.PodPort{}
		for _, v := range dataSlice.Values {
			i, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				logrus.Error(err)
			}
			port := &pod.PodPort{
				ContainerPort: int32(i),
				Protocol:      "TCP",
			}
			podSlice = append(podSlice, port)
		}
		addPodInfo.PodPort = podSlice
	}
	form.FromToPodStruct(req.Post, addPodInfo)

	response, err := p.PodService.AddPod(ctx, addPodInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}

	out.StatusCode = http.StatusOK
	b, _ := json.Marshal(response)
	out.Body = string(b)
	return nil

}

func (p *PodApi) DeletePodById(ctx context.Context, req *protoApi.Request, rsp *protoApi.Response) error {
	logrus.Info("accept podApi.DeletePodById request")
	if _, ok := req.Get["pod_id"]; !ok {
		return errors2.New("参数异常")
	}

	podIdString := req.Get["pod_id"].Values[0]
	podId, err := strconv.ParseInt(podIdString, 10, 64)
	if err != nil {
		logrus.Error(err)
		return err
	}
	response, err := p.PodService.DeletePod(ctx, &pod.PodID{PodId: int32(podId)})
	if err != nil {
		logrus.Error(err)
		return err
	}

	rsp.StatusCode = http.StatusOK
	b, _ := json.Marshal(response)
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
	allPod, err := p.PodService.FindAllPod(ctx, &pod.FindAll{})
	if err != nil {
		logrus.Error(err)
		return err
	}

	rsp.StatusCode = http.StatusOK
	b, _ := json.Marshal(allPod)
	rsp.Body = string(b)
	return nil
}
