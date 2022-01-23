package service

import (
	restful "github.com/emicklei/go-restful"
	"gorpc/common/blog"
	"gorpc/common/rpc"
	"gorpc/scene_server/tmserver/core"
	"net/http"
)

const (
	CommandRDBOperation              = "RDB"
	CommandMigrateOperation          = "DBMigrate"
	CommandWatchTransactionOperation = "WatchTransaction"
)

// Service service methods
type Service interface {
	WebService() *restful.WebService
	SetConfig() error
}

type CoreService struct {
	rpc        *rpc.Server
	listenIP   string
	listenPort uint
	core       core.Core
}

// New create a new service instance
func New(ip string, port uint) Service {
	return &CoreService{
		listenIP:   ip,
		listenPort: port,
	}
}

func (s *CoreService) SetConfig() error {
	s.rpc = rpc.NewServer()

	// init all handlers
	s.rpc.Handle(CommandRDBOperation, s.DBOperation)
	s.rpc.HandleStream(CommandWatchTransactionOperation, s.WatchTransaction)
	s.core = core.New()
	return nil
}

func (s *CoreService) WebService() *restful.WebService {
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.SetLogger(&blog.GlogWriter{})
	restful.TraceLogger(&blog.GlogWriter{})
	ws := &restful.WebService{}
	// ws.Path("/txn/v3").Filter(s.engine.Metric().RestfulMiddleWare)
	ws.Route(ws.Method(http.MethodConnect).Path("rpc").To(func(req *restful.Request, resp *restful.Response) {
		if sub, ok := resp.ResponseWriter.(*restful.Response); ok {
			s.rpc.ServeHTTP(sub.ResponseWriter, req.Request)
			return
		}
		s.rpc.ServeHTTP(resp.ResponseWriter, req.Request)
	}))

	return ws
}
