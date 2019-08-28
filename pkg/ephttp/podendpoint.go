package ephttp

import (
    "k8s.io/client-go/tools/clientcmd"
    "encoding/json"
    "fmt"
    "time"

    "k8s.io/client-go/kubernetes"

    "strings"
)

// PodMetricsList : PodMetricsList
type PodMetricsList struct {
    Kind       string `json:"kind"`
    APIVersion string `json:"apiVersion"`
    Metadata   struct {
        SelfLink string `json:"selfLink"`
    } `json:"metadata"`
    Items []struct {
        Metadata struct {
            Name              string    `json:"name"`
            Namespace         string    `json:"namespace"`
            SelfLink          string    `json:"selfLink"`
            CreationTimestamp time.Time `json:"creationTimestamp"`
        } `json:"metadata"`
        Timestamp  time.Time `json:"timestamp"`
        Window     string    `json:"window"`
        Containers []struct {
            Name  string `json:"name"`
            Usage struct {
                CPU    string `json:"cpu"`
                Memory string `json:"memory"`
            } `json:"usage"`
        } `json:"containers"`
    } `json:"items"`
}

func getMetrics(clientset *kubernetes.Clientset, pods *PodMetricsList) error {
    data, err := clientset.RESTClient().Get().AbsPath("apis/metrics.k8s.io/v1beta1/pods").DoRaw()
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &pods)
    return err
}

func StartUp(kubeconfig string) {

    // use the current context in kubeconfig
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
	panic(err.Error())
    }
    // creates the clientset
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err.Error())
    }


    fmt.Printf("Performing an api http request - Current Unix Time: %v\n", time.Now().Unix())

    var pods PodMetricsList
    err = getMetrics(clientset, &pods)
    if err != nil {
       panic(err.Error())
    }

    for _, m := range pods.Items {

       for _, c := range m.Containers {

          s_mem_in_kb := strings.TrimSuffix(c.Usage.Memory, "Ki")
          s_cpu_in_n := strings.TrimSuffix(c.Usage.CPU, "n")

          fmt.Printf(s_mem_in_kb + "\n")
          fmt.Printf(s_cpu_in_n + "\n")

       }
    }

}
