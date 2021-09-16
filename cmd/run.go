package cmd

import (
	"context"
	"errors"
	"github.com/kuno989/friday/backend"
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
	runCmd = &cobra.Command{
		Use: "run",
		Aliases: []string{"runserver"},
		Short: "Run web api server",
		Run: func(cmd *cobra.Command, args []string){
			s, cleanup, err := backend.InitializeServer(context.Background(), viper.GetViper())
			if err != nil {
				logrus.WithError(err).Fatal("initialize server")
			}
			defer cleanup()
			s.Logger.SetLevel(log.DEBUG)

			go func(){
				bindAddr := viper.GetString("webserver_port")
				logrus.Infof("server is running on http://localhost%s", bindAddr)
				if err := s.Start(bindAddr); err != nil {
					if !errors.Is(err, http.ErrServerClosed){
						logrus.WithError(err).Fatal("start server")
					}
				}
			}()

			sig := make(chan os.Signal, 1)
			signal.Notify(sig, os.Interrupt)
			<- sig
			signal.Reset(os.Interrupt)
			logrus.Info("shutting down server")
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			if err := s.Shutdown(ctx); err != nil {
				logrus.WithError(err).Fatal("shutdown server")
			}
		},
	}
)