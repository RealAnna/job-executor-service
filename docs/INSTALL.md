## Installation

The *job-executor-service* can be installed as a part of [Keptn's uniform](https://keptn.sh) using `helm`.
It is recommended to install the *job-executor-service* on the remote execution-plane, which can be on the same cluster, or on a completely
separate Kubernetes environment (see [Keptn docs: Multi-cluster setup](https://keptn.sh/docs/0.11.x/operate/multi_cluster/) for details).

During the installation process various parameters of the *job-executor-service* can be customized, for a complete list
of these helm values have a look at the [documentation](../chart/README.md).
In order to install the *job-executor-service* on the remote execution plane, some values of the helm chart need to be configured:
* `remoteControlPlane.topicSubscription` - list of Keptn CloudEvent types that this instance should listen to, e.g., `sh.keptn.event.remote-task.triggered`
* `remoteControlPlane.api.protocol` - protocol (`http` or `https`) used to connect to the remote control plane
* `remoteControlPlane.api.hostname` - Keptn API Hostname (e.g., `1.2.3.4.nip.io`). If Keptn is installed on the same cluster the API is usually reachable under `api-gateway-nginx.keptn`.
* `remoteControlPlane.api.token` - Keptn API Token (can be obtained from Bridge)

**Example**

```bash
KEPTN_API_PROTOCOL=http # or https
KEPTN_API_HOST=<INSERT-YOUR-HOSTNAME-HERE> # e.g., 1.2.3.4.nip.io
 KEPTN_API_TOKEN=<INSERT-YOUR-KEPTN-API-TOKEN-HERE>

TASK_SUBSCRIPTION=sh.keptn.event.remote-task.triggered

helm upgrade --install --create-namespace -n <NAMESPACE> \
  job-executor-service https://github.com/keptn-contrib/job-executor-service/releases/download/<VERSION>/job-executor-service-<VERSION>.tgz \
 --set remoteControlPlane.topicSubscription=${TASK_SUBSCRIPTION},remoteControlPlane.api.protocol=${KEPTN_API_PROTOCOL},remoteControlPlane.api.hostname=${KEPTN_API_HOST},remoteControlPlane.api.token=${KEPTN_API_TOKEN}
```

Please replace `<VERSION>` with the actual version you want to install from the compatibility matrix above or the
[GitHub releases page](https://github.com/keptn-contrib/job-executor-service/releases).

To verify that everything works you can visit Bridge, select a project, go to Uniform, and verify that `job-executor-service`  is registered as "remote execution plane" with the correct version and event type.

**Example (with auto-detection)**

For easier installation of the *job-executor-service* a Keptn installation on the same kubernetes cluster can be discovered automatically. This only
works if no `remoteControlPlane.api.token`, `remoteControlPlane.api.protocol` or `remoteControlPlane.api.hostname` is provided.
If multiple Keptn installations are present, the `remoteControlPlane.autoDetect.namespace` must be set to the desired Keptn instance.
The auto-detection feature can be enabled by setting `remoteControlPlane.autoDetect.enabled` to `true`.

```bash
TASK_SUBSCRIPTION=sh.keptn.event.remote-task.triggered

helm upgrade --install --create-namespace -n <NAMESPACE> \
  job-executor-service https://github.com/keptn-contrib/job-executor-service/releases/download/<VERSION>/job-executor-service-<VERSION>.tgz \
 --set remoteControlPlane.autoDetect.enabled=true,remoteControlPlane.topicSubscription=${TASK_SUBSCRIPTION},remoteControlPlane.api.token="",remoteControlPlane.api.hostname="",remoteControlPlane.api.protocol=""
```


### Update API Token on Remote Execution-Plane

To update your `KEPTN_API_TOKEN` on an existing installation, please execute the following command (make sure to use the same `<VERSION>` is currently installed):

```bash
 KEPTN_API_TOKEN=<INSERT-YOUR-NEW-KEPTN-API-TOKEN-HERE>

helm upgrade -n <NAMESPACE> \
  job-executor-service https://github.com/keptn-contrib/job-executor-service/releases/download/<VERSION>/job-executor-service-<VERSION>.tgz \
  --reuse-values \
 --set remoteControlPlane.api.token=${KEPTN_API_TOKEN}
```

### Update Topic Subscriptions

To update your `TASK_SUBSCRIPTION` (as in the Cloud Event types that job-executor-service is listening to), please execute the following command (make sure to use the same `<VERSION>` is currently installed):

```bash
TASK_SUBSCRIPTION=sh.keptn.event.remote-task.triggered,sh.keptn.event.some-other-task.triggered

helm upgrade -n <NAMESPACE> \
  job-executor-service https://github.com/keptn-contrib/job-executor-service/releases/download/<VERSION>/job-executor-service-<VERSION>.tgz \
  --reuse-values \
 --set remoteControlPlane.topicSubscription=${TASK_SUBSCRIPTION}
```

## Upgrade

To upgrade to a newer version of *job-executor-service*, first save the existing installation values to a helm override
file
```shell
helm -n <NAMESPACE> get values job-executor-service > values.yaml
```
then upgrade using the previous installation value file as an override
```bash
helm upgrade -n <NAMESPACE> \
  job-executor-service https://github.com/keptn-contrib/job-executor-service/releases/download/<VERSION>/job-executor-service-<VERSION>.tgz \
  -f values.yaml
```

To upgrade to a newer version of *job-executor-service* and automatically use the auto-detection to configure the Keptn 
API token, the `helm upgrade` command should be:
```bash
helm upgrade -n <NAMESPACE> \
  job-executor-service https://github.com/keptn-contrib/job-executor-service/releases/download/<VERSION>/job-executor-service-<VERSION>.tgz \
  -f values.yaml \
  --set remoteControlPlane.api.token="",remoteControlPlane.api.hostname="",remoteControlPlane.api.protocol=""
```

**Note:** during upgrade we dump the existing installation values to a file and use it as an override file to work around
some of the `helm upgrade --reuse-values` [limitations](https://github.com/helm/helm/issues/8085) and make sure that new
default values are applied as expected when upgrading job-executor-service to a newer version.

## Uninstall

To uninstall *job-executor-service*, run

```bash
helm uninstall -n <NAMESPACE> job-executor-service
```