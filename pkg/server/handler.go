package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron/v2"

	"ainokiseki/binance_rush/api"
	"ainokiseki/binance_rush/pkg/trade"
	"ainokiseki/binance_rush/util"
)

type handler struct {
	api.UnimplementedBinanceServer
	scheduler gocron.Scheduler
	client    *trade.BinanceClient
}

func newHandler(c *trade.BinanceClient) (*handler, error) {
	scheduler, err := gocron.NewScheduler()
	scheduler.Start()
	return &handler{
		scheduler: scheduler,
		client:    c,
	}, err
}

func (h *handler) close() {
	h.scheduler.Shutdown()
}

func (h *handler) Ping(ctx context.Context, request *api.HelloRequest) (*api.HelloReply, error) {
	//TODO implement me
	panic("implement me")
}

func (h *handler) CreateCoinRushTask(ctx context.Context, request *api.CreateCoinRushTaskRequest) (*api.CreateCoinRushTaskReply, error) {

	subctx := context.Background()
	ch := make(chan struct{})
	for i := 0; i < int(request.GetExecuteTimes()); i++ {
		go func() {
			<-ch
			err := h.client.LimitTakerBuyOrder(subctx, request.GetSymbol(), "", request.GetPrice(), request.GetBidQuantity())
			log.Print(err)
		}()
	}

	milliTimeStamp := request.GetStartTimestampMilli()
	mean, sd := util.GetDelay(ctx, h.client.Client)

	milliTimeStamp -= int64(mean) - int64(sd)

	tm := getTimeFromMilli(request.GetStartTimestampMilli())

	j, err := h.scheduler.NewJob(
		gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(tm)),
		gocron.NewTask(func() {
			close(ch)
			log.Println("close channel in:", time.Now().Format("2006-01-02 15:04:05.000"))

		}),
		gocron.WithName(request.GetSymbol()+":"+request.GetPrice()+":"+request.GetBidQuantity()),
	)
	log.Println("create task:", tm.String())

	if err != nil {
		return nil, err
	}
	return &api.CreateCoinRushTaskReply{
		Success: false,
		Id:      j.ID().String(),
	}, err
}

func (h *handler) ListTask(ctx context.Context, req *api.ListTaskRequest) (*api.ListTaskReply, error) {
	fmt.Println("handle task")
	jobs := h.scheduler.Jobs()
	res := &api.ListTaskReply{Tasks: make([]*api.Task, len(jobs))}

	loc, _ := time.LoadLocation("Asia/Shanghai") // 对于UTC+8可以使用 'Asia/Shanghai' 时区

	for index, i := range jobs {
		lastTime, err := i.LastRun()
		if err != nil {
			log.Println("last time error", err)
			continue
		}
		nextTime, err := i.LastRun()
		if err != nil {
			log.Println("next time error", err)
			continue
		}
		fmt.Println(lastTime.String(), nextTime.String())
		res.Tasks[index] = &api.Task{
			Id:       i.ID().String(),
			Name:     i.Name(),
			LastTime: lastTime.In(loc).Format("2006-01-02 15:04:05.000"),
			NextTime: nextTime.In(loc).Format("2006-01-02 15:04:05.000"),
		}
	}
	return res, nil
}

func getTimeFromMilli(timestamp int64) time.Time {
	return time.Unix(timestamp/1e3, (timestamp%1e3)*1e6)
}
