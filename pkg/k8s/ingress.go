package k8s

import (
	"github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//DeleteIngresses deletes the given ingress from all namespaces
func DeleteIngresses(kubeclient *kubernetes.Clientset, name string, namespaces []string) error {
	for _, ns := range namespaces {
		err := kubeclient.ExtensionsV1beta1().Ingresses(ns).Delete(name, &metav1.DeleteOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				logrus.Infof("ingress %q not found in namespace %q", name, ns)
				continue
			}
			return err
		}

		logrus.Infof("ingress %q deleted from namespace %q", name, ns)
	}

	return nil
}
