# Demo

## Prerequisites

1. Ensure that you have a Kubernetes cluster running version 1.25 or later.
    1.1. You can use Kind/Minikube for this purpose.
2. Have the kubectl CLI tool installed.

## Installation

1. Install Tekton Pipeline 
    ```bash
    kubectl apply -f https://storage.googleapis.com/tekton-releases/pipeline/previous/v0.54.0/release.yaml
    ```

2. Install Tekton Chains
    ```bash
    kubectl apply -f https://storage.googleapis.com/tekton-releases/chains/previous/v0.19.0/release.yaml
    ```

3. Setup Pipelines as Code
    1. Install
    ```bash
    kubectl apply -f https://github.com/openshift-pipelines/pipelines-as-code/releases/download/v0.22.4/release.k8s.yaml
    ```
    2. Port forward the pipelines-as-code controller
            a. kubectl port-forward <pipelines-as-code-controller-pod-name> 8080:8080 -n pipelines-as-code
            b. Use the gosmee client with the following command
            ```bash
            gosmee client https://hook.pipelinesascode.com/PCoifdgYPYpS http://localhost:8080
            ```
    3. Create and configure the GithubApp
    4. Create a repository

4. Send a pull request to https://github.com/savitaashture/kubeday-india and observe the triggering of the PipelineRun for the pull request

5. After sending a push request, check if the PipelineRun for the push request is triggered.

6. Verify that the pushed image is signed and attested using Tekton Chains

## Documentation References

Tekton Pipeline doc: https://tekton.dev/docs/
Tekton Chains doc: https://tekton.dev/docs/chains/
Pipelines as Code: https://pipelinesascode.com/

Demo Repository: https://github.com/savitaashture/kubeday-india

