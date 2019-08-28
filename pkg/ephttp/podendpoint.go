package ephttp

import (
    "k8s.io/client-go/tools/clientcmd"
    "encoding/json"
    "fmt"
    "time"

    "k8s.io/client-go/kubernetes"
)

// PodMetricsList : PodMetricsList
type PodMetricsList struct {
    Kind       string `json:"kind"`
    APIVersion string `json:"apiVersion"`
    Metadata   struct {
        Name              string    `json:"name"`
        Namespace         string    `json:"namespace"`
        SelfLink          string    `json:"selfLink"`
        UID               string    `json:"uid"`
	Resourceversion   string    `json:"resourceVersion"`
	CreationTimestamp time.Time `json:"creationTimestamp"`
	Labels struct {
		App string `json:"app"`
	} `json:"labels"`
        Annotations struct {
          Endpointk8slastchange string `json:"endpoints.kubernetes.io/last-change-trigger-time"`
	}
    } `json:"metadata"`
    Subsets []struct {
        Addresses []struct {
		IP        string `json:"ip"`
		Nodename  string `json:"nodeName"`
                Targetref struct  {
                    Kind            string `json:"kind"`
		    Namespace       string `json:"namespace"`
		    Name            string `json:"name"`
		    UID             string `json:"uid"`
		    Resourceversion string `json:"resourceversion"`
	    } `json:"targetRef"`
        } `json:"addresses"`
	Ports []struct {
		Port int `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"ports"`
    } `json:"subsets"`
}

func getMetrics(clientset *kubernetes.Clientset, pods *PodMetricsList) error {
    data, err := clientset.RESTClient().Get().AbsPath("api/v1/namespaces/credenciamento/endpoints/credenciamento-validacao-telefone").DoRaw()
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

    for _, m := range pods.Subsets {

       for _, c := range m.Addresses {

          fmt.Printf(c.IP + "\n")
          fmt.Printf(c.Nodename + "\n")

       }
    }

}
