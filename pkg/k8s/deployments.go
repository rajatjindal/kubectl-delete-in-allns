package k8s

import (
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//DeleteDeployments deletes the given deployment from all namespaces
func DeleteDeployments(kubeclient *kubernetes.Clientset, name string, namespaces []string) error {
	for _, ns := range namespaces {
		err := kubeclient.AppsV1().Deployments(ns).Delete(name, &metav1.DeleteOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				logrus.Infof("deployment %q not found in namespace %q", name, ns)
				continue
			}

			return err
		}

		logrus.Infof("deployment %q deleted from namespace %q", name, ns)
	}

	return nil
}
