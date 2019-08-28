package ephttp

import (
    "k8s.io/client-go/tools/clientcmd"
    "encoding/json"
    "log"
    "time"

    "k8s.io/client-go/kubernetes"
)

type PodEndpointList struct {
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

func getEndpoints(clientset *kubernetes.Clientset, pods *PodEndpointList, svc string, ns string) error {
    data, err := clientset.RESTClient().Get().AbsPath("api/v1/namespaces/" + ns + "/endpoints/" + svc).DoRaw()
    if err != nil {
        return err
    }
    err = json.Unmarshal(data, &pods)
    return err
}

func StartUp(kubeconfig string, svc string, ns string) []string {

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

    var pods PodEndpointList
    err = getEndpoints(clientset, &pods, svc, ns)
    if err != nil {
       panic(err.Error())
    }

    var ipsep []string
    for _, m := range pods.Subsets {

       for _, c := range m.Addresses {

          log.Printf(c.IP)
          log.Printf(c.Nodename)

          ipsep = append(ipsep, c.IP)

       }
    }

    return ipsep
}
