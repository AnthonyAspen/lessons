package main

import (
	"fmt"
	"log"
	"context"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


func main(){
  rules := clientcmd.NewDefaultClientConfigLoadingRules()
  kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules,&clientcmd.ConfigOverrides{})
  config,err := kubeconfig.ClientConfig()
  if err != nil{
    //TODO error handling
    log.Fatal(err)
  }
  clientset := kubernetes.NewForConfigOrDie(config)

  nodeList, err := clientset.CoreV1().Nodes().List(context.Background(),metav1.ListOptions{})
  if err != nil{
    //TODO error handling
    log.Fatal(err)
  }

  for _,n := range nodeList.Items {
    fmt.Println(n.Name)
  }
}
