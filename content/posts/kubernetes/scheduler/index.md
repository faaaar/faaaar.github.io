+++
title = "KUBERNETES - 调度器的工作原理"
author = ["LIYUNFAN"]
date = 2021-01-14
tags = ["KUBERNETES"]
draft = true
+++

1.  Setup
2.  Run


## Setup {#setup}

1.  Verify input options
2.  Generate config with options
3.  Generate scheduler with the config


### Verify input options {#verify-input-options}

Just check all options.


### Generate config with options {#generate-config-with-options}

1.  Generate secure serveing config such as cert, key and so on.
2.  Copy config option from options
3.  Create **kubeClient**, **leaderElection** and **eventClient**.
4.  Fill some default values to config, such as **InsecureServing** and **InsecureMetricsServing**
5.  Init **BearerToken**
6.  Generate **EventBroadcaster**


### Generate scheduler with config {#generate-scheduler-with-config}

```go
// Config has all the context to run a Scheduler
type Config struct {
    // ComponentConfig is the scheduler server's configuration object.
    ComponentConfig kubeschedulerconfig.KubeSchedulerConfiguration

    // LoopbackClientConfig is a config for a privileged loopback connection
    LoopbackClientConfig *restclient.Config

    InsecureServing        *apiserver.DeprecatedInsecureServingInfo // nil will disable serving on an insecure port
    InsecureMetricsServing *apiserver.DeprecatedInsecureServingInfo // non-nil if metrics should be served independently
    Authentication         apiserver.AuthenticationInfo
    Authorization          apiserver.AuthorizationInfo
    SecureServing          *apiserver.SecureServingInfo

    Client          clientset.Interface
    InformerFactory informers.SharedInformerFactory

    //lint:ignore SA1019 this deprecated field still needs to be used for now. It will be removed once the migration is done.
    EventBroadcaster events.EventBroadcasterAdapter

    // LeaderElection is optional.
    LeaderElection *leaderelection.LeaderElectionConfig
}

sched, err := scheduler.New(cc.Client,
    cc.InformerFactory,
    recorderFactory,
    ctx.Done(),
    scheduler.WithProfiles(cc.ComponentConfig.Profiles...),
    scheduler.WithAlgorithmSource(cc.ComponentConfig.AlgorithmSource),
    scheduler.WithPercentageOfNodesToScore(cc.ComponentConfig.PercentageOfNodesToScore),
    scheduler.WithFrameworkOutOfTreeRegistry(outOfTreeRegistry),
    scheduler.WithPodMaxBackoffSeconds(cc.ComponentConfig.PodMaxBackoffSeconds),
    scheduler.WithPodInitialBackoffSeconds(cc.ComponentConfig.PodInitialBackoffSeconds),
    scheduler.WithExtenders(cc.ComponentConfig.Extenders...),
    scheduler.WithParallelism(cc.ComponentConfig.Parallelism),
    scheduler.WithBuildFrameworkCapturer(func(profile kubeschedulerconfig.KubeSchedulerProfile) {
        // Profiles are processed during Framework instantiation to set default plugins and configurations. Capturing them for logging
        completedProfiles = append(completedProfiles, profile)
    }),
)
```


## Run {#run}

1.  **configz["componentconfig"] = cc** . **cc** is the generated config in **Setup**.
2.  Start eventBroadcaster
3.  Setup and start healthz checks
4.  Setup and start leaderelection
5.  Setup and start metrics
6.  Start infomer
7.  Run scheduler


### Run scheduler {#run-scheduler}
