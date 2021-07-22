/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	"kubebuilderht/http/restcontroller"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"kubebuilderht/api/controllers"
)

func main() {
	go controllers.WatchTemplate()

	ctx := NewContextWithSigterm(context.Background())

	restful.Add(restcontroller.HTService())
	//plugin_rest.Init(restful.DefaultContainer)
	go http.ListenAndServe(":8080", nil)
	<-ctx.Done()
	time.Sleep(1 * time.Second)
}

func NewContextWithSigterm(ctx context.Context) context.Context {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	go func() {
		defer cancel()

		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

		select {
		case <-signalCh:
			fmt.Println("SIGTERM|Interrupt received")
		case <-ctx.Done():
			fmt.Println("Context was cancelled")
		}
	}()
	return ctxWithCancel
}
