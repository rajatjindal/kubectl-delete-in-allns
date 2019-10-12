package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//GetAllNamespaces deletes the given secret from all namespaces
func GetAllNamespaces(kubeclient *kubernetes.Clientset) ([]string, error) {
	namespacesListObject, err := kubeclient.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var namespaces []string
	for _, ns := range namespacesListObject.Items {
		namespaces = append(namespaces, ns.Name)
	}

	return namespaces, nil
}
