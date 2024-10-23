/*
Copyright 2024 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"sync"
	"time"

	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type PrometheusHandler func(string, *v1.Prometheus) (*v1.Prometheus, error)

type PrometheusController interface {
	generic.ControllerMeta
	PrometheusClient

	OnChange(ctx context.Context, name string, sync PrometheusHandler)
	OnRemove(ctx context.Context, name string, sync PrometheusHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() PrometheusCache
}

type PrometheusClient interface {
	Create(*v1.Prometheus) (*v1.Prometheus, error)
	Update(*v1.Prometheus) (*v1.Prometheus, error)
	UpdateStatus(*v1.Prometheus) (*v1.Prometheus, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.Prometheus, error)
	List(namespace string, opts metav1.ListOptions) (*v1.PrometheusList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Prometheus, err error)
}

type PrometheusCache interface {
	Get(namespace, name string) (*v1.Prometheus, error)
	List(namespace string, selector labels.Selector) ([]*v1.Prometheus, error)

	AddIndexer(indexName string, indexer PrometheusIndexer)
	GetByIndex(indexName, key string) ([]*v1.Prometheus, error)
}

type PrometheusIndexer func(obj *v1.Prometheus) ([]string, error)

type prometheusController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewPrometheusController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) PrometheusController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &prometheusController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromPrometheusHandlerToHandler(sync PrometheusHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.Prometheus
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.Prometheus))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *prometheusController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.Prometheus))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdatePrometheusDeepCopyOnChange(client PrometheusClient, obj *v1.Prometheus, handler func(obj *v1.Prometheus) (*v1.Prometheus, error)) (*v1.Prometheus, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *prometheusController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *prometheusController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *prometheusController) OnChange(ctx context.Context, name string, sync PrometheusHandler) {
	c.AddGenericHandler(ctx, name, FromPrometheusHandlerToHandler(sync))
}

func (c *prometheusController) OnRemove(ctx context.Context, name string, sync PrometheusHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromPrometheusHandlerToHandler(sync)))
}

func (c *prometheusController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *prometheusController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *prometheusController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *prometheusController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *prometheusController) Cache() PrometheusCache {
	return &prometheusCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *prometheusController) Create(obj *v1.Prometheus) (*v1.Prometheus, error) {
	result := &v1.Prometheus{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *prometheusController) Update(obj *v1.Prometheus) (*v1.Prometheus, error) {
	result := &v1.Prometheus{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *prometheusController) UpdateStatus(obj *v1.Prometheus) (*v1.Prometheus, error) {
	result := &v1.Prometheus{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *prometheusController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *prometheusController) Get(namespace, name string, options metav1.GetOptions) (*v1.Prometheus, error) {
	result := &v1.Prometheus{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *prometheusController) List(namespace string, opts metav1.ListOptions) (*v1.PrometheusList, error) {
	result := &v1.PrometheusList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *prometheusController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *prometheusController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.Prometheus, error) {
	result := &v1.Prometheus{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type prometheusCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *prometheusCache) Get(namespace, name string) (*v1.Prometheus, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.Prometheus), nil
}

func (c *prometheusCache) List(namespace string, selector labels.Selector) (ret []*v1.Prometheus, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Prometheus))
	})

	return ret, err
}

func (c *prometheusCache) AddIndexer(indexName string, indexer PrometheusIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.Prometheus))
		},
	}))
}

func (c *prometheusCache) GetByIndex(indexName, key string) (result []*v1.Prometheus, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.Prometheus, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.Prometheus))
	}
	return result, nil
}

// PrometheusStatusHandler is executed for every added or modified Prometheus. Should return the new status to be updated
type PrometheusStatusHandler func(obj *v1.Prometheus, status v1.PrometheusStatus) (v1.PrometheusStatus, error)

// PrometheusGeneratingHandler is the top-level handler that is executed for every Prometheus event. It extends PrometheusStatusHandler by a returning a slice of child objects to be passed to apply.Apply
type PrometheusGeneratingHandler func(obj *v1.Prometheus, status v1.PrometheusStatus) ([]runtime.Object, v1.PrometheusStatus, error)

// RegisterPrometheusStatusHandler configures a PrometheusController to execute a PrometheusStatusHandler for every events observed.
// If a non-empty condition is provided, it will be updated in the status conditions for every handler execution
func RegisterPrometheusStatusHandler(ctx context.Context, controller PrometheusController, condition condition.Cond, name string, handler PrometheusStatusHandler) {
	statusHandler := &prometheusStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromPrometheusHandlerToHandler(statusHandler.sync))
}

// RegisterPrometheusGeneratingHandler configures a PrometheusController to execute a PrometheusGeneratingHandler for every events observed, passing the returned objects to the provided apply.Apply.
// If a non-empty condition is provided, it will be updated in the status conditions for every handler execution
func RegisterPrometheusGeneratingHandler(ctx context.Context, controller PrometheusController, apply apply.Apply,
	condition condition.Cond, name string, handler PrometheusGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &prometheusGeneratingHandler{
		PrometheusGeneratingHandler: handler,
		apply:                       apply,
		name:                        name,
		gvk:                         controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterPrometheusStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type prometheusStatusHandler struct {
	client    PrometheusClient
	condition condition.Cond
	handler   PrometheusStatusHandler
}

// sync is executed on every resource addition or modification. Executes the configured handlers and sends the updated status to the Kubernetes API
func (a *prometheusStatusHandler) sync(key string, obj *v1.Prometheus) (*v1.Prometheus, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type prometheusGeneratingHandler struct {
	PrometheusGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
	seen  sync.Map
}

// Remove handles the observed deletion of a resource, cascade deleting every associated resource previously applied
func (a *prometheusGeneratingHandler) Remove(key string, obj *v1.Prometheus) (*v1.Prometheus, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.Prometheus{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	if a.opts.UniqueApplyForResourceVersion {
		a.seen.Delete(key)
	}

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

// Handle executes the configured PrometheusGeneratingHandler and pass the resulting objects to apply.Apply, finally returning the new status of the resource
func (a *prometheusGeneratingHandler) Handle(obj *v1.Prometheus, status v1.PrometheusStatus) (v1.PrometheusStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.PrometheusGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}
	if !a.isNewResourceVersion(obj) {
		return newStatus, nil
	}

	err = generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
	if err != nil {
		return newStatus, err
	}
	a.storeResourceVersion(obj)
	return newStatus, nil
}

// isNewResourceVersion detects if a specific resource version was already successfully processed.
// Only used if UniqueApplyForResourceVersion is set in generic.GeneratingHandlerOptions
func (a *prometheusGeneratingHandler) isNewResourceVersion(obj *v1.Prometheus) bool {
	if !a.opts.UniqueApplyForResourceVersion {
		return true
	}

	// Apply once per resource version
	key := obj.Namespace + "/" + obj.Name
	previous, ok := a.seen.Load(key)
	return !ok || previous != obj.ResourceVersion
}

// storeResourceVersion keeps track of the latest resource version of an object for which Apply was executed
// Only used if UniqueApplyForResourceVersion is set in generic.GeneratingHandlerOptions
func (a *prometheusGeneratingHandler) storeResourceVersion(obj *v1.Prometheus) {
	if !a.opts.UniqueApplyForResourceVersion {
		return
	}

	key := obj.Namespace + "/" + obj.Name
	a.seen.Store(key, obj.ResourceVersion)
}
