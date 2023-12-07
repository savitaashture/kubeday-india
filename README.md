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

3. Install Tekton Dashboard
    ```bash
    kubectl apply --filename https://storage.googleapis.com/tekton-releases/dashboard/latest/release.yaml
    ```

    1. access Dashboard
       ```bash
       kubectl --namespace tekton-pipelines port-forward svc/tekton-dashboard 9097:9097
       ```
       http://localhost:9097/

4. Setup Pipelines as Code
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
       **OR**
       
       Follow : https://github.com/openshift-pipelines/pipelines-as-code/blob/main/hack/dev/kind/install.sh to create Kind, install Tekton Pipeline and setup gosmee
       
    4. Follow https://pipelinesascode.com/
       1. Create and configure the GithubApp https://pipelinesascode.com/docs/install/github_apps/
       2. Create a repository https://pipelinesascode.com/docs/guide/repositorycrd/

5. Follow Tekton Chains Tutorial https://github.com/tektoncd/chains/blob/main/docs/tutorials/signed-provenance-tutorial.md to set up Chains to sign OCI images built in Tekton
   
6. Send a pull request to https://github.com/savitaashture/kubeday-india and observe the triggering of the PipelineRun for the pull request

7. After sending a push request, check if the PipelineRun for the push request is triggered.

8. Verify that the pushed image is signed and attested using Tekton Chains

## Documentation References

Tekton Pipeline doc: https://tekton.dev/docs/

Tekton Chains doc: https://tekton.dev/docs/chains/

Pipelines as Code: https://pipelinesascode.com/

Demo Repository: https://github.com/savitaashture/kubeday-india

