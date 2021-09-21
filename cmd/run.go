package cmd

import (
	"context"
	"errors"
	"github.com/kuno989/friday/backend"
	"github.com/kuno989/friday/fridayEngine"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	runBackend = &cobra.Command{
		Use:     "server",
		Aliases: []string{"runserver"},
		Short:   "Run friday web api server",
		Run: func(cmd *cobra.Command, args []string) {
			s, cleanup, err := backend.InitializeServer(context.Background(), viper.GetViper())
			if err != nil {
				logrus.WithError(err).Fatal("initialize server")
			}
			defer cleanup()
			s.Logger.SetLevel(log.DEBUG)

			go func() {
				bindAddr := viper.GetString("webserver_port")
				logrus.Infof("friday api Server Running on http://localhost%s", bindAddr)
				if err := s.Start(bindAddr); err != nil {
					if !errors.Is(err, http.ErrServerClosed) {
						logrus.WithError(err).Fatal("start server")
					}
				}
			}()

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<-sig
			signal.Reset(os.Interrupt)
			logrus.Info("shutting down server")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := s.Shutdown(ctx); err != nil {
				logrus.WithError(err).Fatal("shutdown server")
			}
		},
	}
	runFridayEngine = &cobra.Command{
		Use:     "engine",
		Aliases: []string{"runserver"},
		Short:   "Run friday engine server",
		Run: func(cmd *cobra.Command, args []string) {
			s, cleanup, err := fridayEngine.InitializeServer(context.Background(), viper.GetViper())
			if err != nil {
				logrus.WithError(err).Fatal("initialize server")
			}
			defer cleanup()
			ch, err := s.Rb.Channel()
			if err != nil {
				logrus.WithError(err).Fatal("initialize rabbitmq")
			}
			logrus.Info("friday Engine Server Running")
			queue, err := ch.QueueDeclare(s.Rb.Config.FileScanQueue, false, false, false, false, nil)
			if err != nil {
				logrus.WithError(err).Fatal("initialize queue")
			}
			m, err := ch.Consume(queue.Name, "", true, false, false, false, nil)
			if err != nil {
				logrus.WithError(err).Fatal("initialize consume")
			}
			go func() {
				for delivery := range m {
					err := s.AmqpHandler(delivery)
					if err != nil {
						return
					}
				}
			}()
			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<-sig
			signal.Reset(os.Interrupt)
			logrus.Info("shutting down server")
			if err := ch.Close(); err != nil {
				logrus.WithError(err).Fatal("shutdown server")
			}
		},
	}
)
