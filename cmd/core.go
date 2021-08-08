package cmd

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	namespace  string
	kubeconfig string
	clientset  *kubernetes.Clientset
	coreCmd    = &cobra.Command{
		Use:   "core",
		Short: "Run the core on kubernetes",
		Long:  `Start the roboss core on the specified kubernetes cluster.`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.TODO()
			_, err := clientset.CoreV1().Namespaces().Get(ctx, namespace, metav1.GetOptions{})
			if err != nil {
				if err, ok := err.(*errors.StatusError); ok {
					fmt.Printf("creating %s namespace\n", namespace)
					_, err := clientset.CoreV1().Namespaces().Create(ctx, &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: namespace}}, metav1.CreateOptions{})
					if err != nil {
						panic(err.Error())
					}
				} else {
					panic(err.Error())
				}
			}

			fmt.Println("deploying operator")

		},
	}
)

func init() {
	coreCmd.PersistentFlags().StringVar(&namespace, "n", "roboss", "core's namespace")
	if home := homedir.HomeDir(); home != "" {
		coreCmd.PersistentFlags().StringVar(&kubeconfig, "c", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		coreCmd.PersistentFlags().StringVar(&kubeconfig, "c", "", "absolute path to the kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
