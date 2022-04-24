/*
Copyright 2022 Natun.

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

package setup

import (
	"fmt"
	"github.com/natun-ai/natun/internal/accessor"
	"github.com/natun-ai/natun/internal/engine"
	corectrl "github.com/natun-ai/natun/internal/engine/controllers"
	"github.com/natun-ai/natun/internal/historian"
	opctrl "github.com/natun-ai/natun/internal/operator"
	"github.com/natun-ai/natun/internal/stats"
	"github.com/natun-ai/natun/pkg/api"
	"github.com/natun-ai/natun/pkg/plugin"
	"github.com/spf13/viper"
	"net/http"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func setupStats(mgr manager.Manager) {
	// Setup usage reports
	stats.UID = viper.GetString("usage-reporting-uid")
	OrFail(mgr.Add(stats.Run(
		mgr.GetConfig(),
		mgr.GetClient(),
		viper.GetBool("usage-reporting"),
		ctrl.Log.WithName("stats"),
	)), "unable to add stats")
}

func historianClient(mgr manager.Manager) historian.Client {
	// Create Notifiers
	collectNotifier, err := plugin.NewCollectNotifier(viper.GetString("notifier-provider"), viper.GetViper())
	OrFail(err, "failed to create collect notifier")
	writeNotifier, err := plugin.NewWriteNotifier(viper.GetString("notifier-provider"), viper.GetViper())
	OrFail(err, "failed to create collect notifier")

	// Create a Historian Client
	hsc := historian.NewClient(historian.ClientConfig{
		CollectNotifier:            collectNotifier,
		WriteNotifier:              writeNotifier,
		Logger:                     ctrl.Log.WithName("historian"),
		CollectNotificationWorkers: 5,
		WriteNotificationWorkers:   5,
	})
	OrFail(hsc.WithManager(mgr), "failed to create historian client")

	return hsc
}

func coreControllers(mgr manager.Manager, eng api.ManagerEngine) {
	var err error

	// Setup Core Controllers
	err = (&corectrl.DataConnectorReconciler{
		Reader:        mgr.GetClient(),
		Scheme:        mgr.GetScheme(),
		EngineManager: eng,
	}).SetupWithManager(mgr)
	OrFail(err, "unable to create core controller", "controller", "DataConnector")

	err = (&corectrl.FeatureReconciler{
		Reader:         mgr.GetClient(),
		Scheme:         mgr.GetScheme(),
		UpdatesAllowed: updatesAllowed,
		EngineManager:  eng,
	}).SetupWithManager(mgr)
	OrFail(err, "unable to create core controller", "controller", "Feature")
}

func operatorControllers(mgr manager.Manager) {
	var err error

	coreAddr := viper.GetString("accessor-service")
	if coreAddr == "" {
		coreAddr, err = discoverAccessor(mgr.GetClient())
		OrFail(err, "failed to detect accessor")
	}

	err = (&opctrl.DataConnectorReconciler{
		Client:   mgr.GetClient(),
		Scheme:   mgr.GetScheme(),
		CoreAddr: coreAddr,
	}).SetupWithManager(mgr)
	OrFail(err, "unable to create controller", "operator", "DataConnector")

	err = (&opctrl.FeatureReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr)
	OrFail(err, "unable to create controller", "operator", "Feature")

	err = opctrl.SetupFeatureWebhook(mgr, updatesAllowed)
	OrFail(err, "unable to create webhook", "webhook", "Feature")
}

func Core(mgr manager.Manager) {
	// Setup usage reporting
	setupStats(mgr)

	// Create a Historian Client
	hsc := historianClient(mgr)

	// Create the state
	state, err := plugin.NewState(viper.GetString("state-provider"), viper.GetViper())
	OrFail(err, fmt.Sprintf("failed to create state for provider %s", viper.GetString("provider")))
	healthChecks = append(healthChecks, func(req *http.Request) error {
		return state.Ping(req.Context())
	})

	// Create a new Core engine
	eng := engine.New(state, hsc, ctrl.Log.WithName("engine"))

	// Create a new Accessor
	acc := accessor.New(eng, ctrl.Log.WithName("accessor"))
	OrFail(mgr.Add(acc.GRPC(viper.GetString("accessor-grpc-address"))), "unable to start gRPC accessor")
	OrFail(
		mgr.Add(acc.HTTP(viper.GetString("accessor-http-address"), viper.GetString("accessor-http-prefix"))),
		"unable to start HTTP accessor")

	coreControllers(mgr, eng)
	operatorControllers(mgr)
}
