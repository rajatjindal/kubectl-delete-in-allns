package k8s

import (
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//DeleteSecrets deletes the given secret from all namespaces
func DeleteSecrets(kubeclient *kubernetes.Clientset, name string, namespaces []string) error {
	for _, ns := range namespaces {
		err := kubeclient.CoreV1().Secrets(ns).Delete(name, &metav1.DeleteOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				logrus.Infof("secret %q not found in namespace %q", name, ns)
				continue
			}

			return err
		}

		logrus.Infof("secret %q deleted from namespace %q", name, ns)
	}

	return nil
}
