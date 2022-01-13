package k8s

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/apps/v1"
	v13 "k8s.io/client-go/kubernetes/typed/batch/v1"
	batchv1beta1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	v12 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	"valse/pkg/config"
)

type Client interface {
	ApiDiscovery() discovery.DiscoveryInterface
	ApiCoreV1() v12.CoreV1Interface
	ApiAppsV1() v1.AppsV1Interface
	ApiBatchV1beta1() batchv1beta1.BatchV1beta1Interface
	ApiBatchV1() v13.BatchV1Interface
	ApiRESTClient() rest.Interface
	clientSet() *kubernetes.Clientset
	clientConfig() *rest.Config
}

func NewClient(appConfig *config.AppConfig, logger *logrus.Logger) Client {
	return &client{
		inClusterConfig: appConfig.Client.InClusterConfig,
		logger:          logger,
	}
}

type client struct {
	inClusterConfig bool
	logger          *logrus.Logger
}

func (c *client) ApiDiscovery() discovery.DiscoveryInterface {
	return c.clientSet().Discovery()
}

func (c *client) ApiCoreV1() v12.CoreV1Interface {
	return c.clientSet().CoreV1()
}

func (c *client) ApiAppsV1() v1.AppsV1Interface {
	return c.clientSet().AppsV1()
}

func (c *client) ApiBatchV1beta1() batchv1beta1.BatchV1beta1Interface {
	return c.clientSet().BatchV1beta1()
}

func (c *client) ApiBatchV1() v13.BatchV1Interface {
	return c.clientSet().BatchV1()
}

func (c *client) ApiRESTClient() rest.Interface {
	return c.clientSet().RESTClient()
}

func (c client) clientSet() *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(c.clientConfig())
	if err != nil {
		c.logger.Fatal(err)
	}
	return clientSet
}

func (c client) clientConfig() *rest.Config {

	var (
		clientConfig *rest.Config
		err          error
	)

	if c.inClusterConfig {
		clientConfig, err = rest.InClusterConfig()
		if err != nil {
			c.logger.Fatal(err)
		}
	} else {
		kubeConfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
		clientConfig, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			c.logger.Fatal(err)
		}
	}

	return clientConfig
}
