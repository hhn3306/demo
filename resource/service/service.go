package service

//import (
//	"github.com/example/memcached-operator/api/v1alpha1"
//	corev1 "k8s.io/api/core/v1"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"k8s.io/apimachinery/pkg/runtime/schema"
//)
//
//func New(app *v1alpha1.Memcached) *corev1.Service {
//	return &corev1.Service{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "Service",
//			APIVersion: "v1",
//		},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      app.Name,
//			Namespace: app.Namespace,
//			OwnerReferences: []metav1.OwnerReference{
//				*metav1.NewControllerRef(app, schema.GroupVersionKind{
//					Group:   v1alpha1.GroupVersion.Group,
//					Version: v1alpha1.GroupVersion.Version,
//					Kind:    "App",
//				}),
//			},
//		},
//		Spec: corev1.ServiceSpec{
//			Ports: app.Spec.Ports,
//			Selector: map[string]string{
//				//"app.example.com/v1": app.Name,
//				"cache.example.com/v1alpha1": app.Name,
//			},
//		},
//	}
//}
