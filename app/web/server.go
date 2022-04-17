package web

import (
	"github.com/gin-gonic/gin"
	"github.com/voltento/mongo_rest/app/config"
	"github.com/voltento/mongo_rest/app/dto"
	"github.com/voltento/mongo_rest/app/service"
	"net/http"
	"strconv"
)

type Server struct {
	cfg    config.Config
	router *gin.Engine
	search *service.Search
	hc     *service.HealthChecker
}

func (s *Server) Router() *gin.Engine {
	return s.router
}

func NewServer(cfg config.Config, search *service.Search, hc *service.HealthChecker) *Server {
	router := gin.Default()

	s := &Server{cfg: cfg, router: router, search: search, hc: hc}
	return s
}

func (s *Server) bind() {

	s.router.GET("/api/v1/trivia",
		func(c *gin.Context) { basicAuth(s.cfg, c) },
		func(c *gin.Context) {
			filters, err := parseFilters(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
				return
			}
			records := s.search.FindRecords(filters)
			c.JSON(http.StatusOK, records)
		},
	)

	s.router.GET("/health", func(c *gin.Context) {
		if err := s.hc.Healthy(); err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, "healthy")
		}
	})
}

func parseFilters(c *gin.Context) (*dto.Filters, error) {
	f := dto.NewDefaultFilters()
	if val := c.Query("limit"); len(val) > 0 {
		v, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		f.Limit = int64(v)
	}

	if val := c.Query("found"); len(val) > 0 {
		v, err := strconv.ParseBool(val)
		if err != nil {
			return nil, err
		}
		f.Found = &v
	}

	if val := c.Query("number"); len(val) > 0 {
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, err
		}
		f.Number = &v
	}

	if val := c.Query("type"); len(val) > 0 {
		f.Type = &val
	}

	return f, nil
}

func (s *Server) Run() error {
	s.bind()
	return s.router.Run(s.cfg.ServiceHost)
}
