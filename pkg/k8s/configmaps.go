package k8s

import (
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//DeleteConfigMaps deletes the given configmap from all namespaces
func DeleteConfigMaps(kubeclient *kubernetes.Clientset, name string, namespaces []string) error {
	for _, ns := range namespaces {
		err := kubeclient.CoreV1().ConfigMaps(ns).Delete(name, &metav1.DeleteOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				logrus.Infof("configmap %q not found in namespace %q", name, ns)
				continue
			}

			return err
		}

		logrus.Infof("configmap %q deleted from namespace %q", name, ns)
	}

	return nil
}
