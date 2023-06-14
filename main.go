package main

import (
	consul "github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/sirupsen/logrus"
	"github.com/yejiabin9/pod/proto/pod"
	"github.com/yejiabin9/podApi/handler"
	"github.com/yejiabin9/podApi/proto/protoApi"
	"strconv"
)

var (
	consulHost       = "39.104.82.215"
	consulPort int64 = 8500

	serviceHost       = "192.168.31.50"
	servicePort int64 = 8082

	tracerHost       = ""
	tracerPort int64 = 6831

	hystrixPort int64 = 9092

	prometheusPort int64 = 9192
)

func main() {
	//1 register consul
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":" + strconv.FormatInt(consulPort, 10),
		}

	})

	service := micro.NewService(
		micro.Server(server.NewServer(func(options *server.Options) {
			options.Advertise = serviceHost + ":" + strconv.FormatInt(servicePort, 10)
		})),
		micro.Name("go.micro.api.podApi"),
		micro.Version("latest"),
		micro.Address(":"+strconv.FormatInt(servicePort, 10)),
		micro.Registry(consul),
	)
	service.Init()

	podService := pod.NewPodService("go.micro.service.pod", service.Client())

	if err := protoApi.RegisterPodApiHandler(service.Server(), &handler.PodApi{PodService: podService}); err != nil {
		logrus.Error(err.Error())
	}

	if err := service.Run(); err != nil {
		logrus.Error(err.Error())
	}
}
