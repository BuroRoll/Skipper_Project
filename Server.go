package Skipper

//
//import (
//	"context"
//	"net/http"
//	"time"
//)
//
//type Server struct {
//	httpServer *http.Server
//}
//
//// Запуск сервера
//func (s *Server) Run(port string, handler http.Handler) error {
//	s.httpServer = &http.Server{
//		Addr:           ":" + port,
//		Handler:        handler,
//		MaxHeaderBytes: 5 << 20, // 5 MB
//		ReadTimeout:    10 * time.Second,
//		WriteTimeout:   10 * time.Second,
//	}
//
//	return s.httpServer.ListenAndServe()
//}
//
//func (s *Server) Shutdown(ctx context.Context) error {
//	return s.httpServer.Shutdown(ctx)
//}
