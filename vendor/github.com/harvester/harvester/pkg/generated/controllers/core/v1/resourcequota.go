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
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	v1 "k8s.io/api/core/v1"
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

type ResourceQuotaHandler func(string, *v1.ResourceQuota) (*v1.ResourceQuota, error)

type ResourceQuotaController interface {
	generic.ControllerMeta
	ResourceQuotaClient

	OnChange(ctx context.Context, name string, sync ResourceQuotaHandler)
	OnRemove(ctx context.Context, name string, sync ResourceQuotaHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ResourceQuotaCache
}

type ResourceQuotaClient interface {
	Create(*v1.ResourceQuota) (*v1.ResourceQuota, error)
	Update(*v1.ResourceQuota) (*v1.ResourceQuota, error)
	UpdateStatus(*v1.ResourceQuota) (*v1.ResourceQuota, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.ResourceQuota, error)
	List(namespace string, opts metav1.ListOptions) (*v1.ResourceQuotaList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ResourceQuota, err error)
}

type ResourceQuotaCache interface {
	Get(namespace, name string) (*v1.ResourceQuota, error)
	List(namespace string, selector labels.Selector) ([]*v1.ResourceQuota, error)

	AddIndexer(indexName string, indexer ResourceQuotaIndexer)
	GetByIndex(indexName, key string) ([]*v1.ResourceQuota, error)
}

type ResourceQuotaIndexer func(obj *v1.ResourceQuota) ([]string, error)

type resourceQuotaController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewResourceQuotaController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ResourceQuotaController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &resourceQuotaController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromResourceQuotaHandlerToHandler(sync ResourceQuotaHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ResourceQuota
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ResourceQuota))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *resourceQuotaController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ResourceQuota))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateResourceQuotaDeepCopyOnChange(client ResourceQuotaClient, obj *v1.ResourceQuota, handler func(obj *v1.ResourceQuota) (*v1.ResourceQuota, error)) (*v1.ResourceQuota, error) {
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

func (c *resourceQuotaController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *resourceQuotaController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *resourceQuotaController) OnChange(ctx context.Context, name string, sync ResourceQuotaHandler) {
	c.AddGenericHandler(ctx, name, FromResourceQuotaHandlerToHandler(sync))
}

func (c *resourceQuotaController) OnRemove(ctx context.Context, name string, sync ResourceQuotaHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromResourceQuotaHandlerToHandler(sync)))
}

func (c *resourceQuotaController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *resourceQuotaController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *resourceQuotaController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *resourceQuotaController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *resourceQuotaController) Cache() ResourceQuotaCache {
	return &resourceQuotaCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *resourceQuotaController) Create(obj *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	result := &v1.ResourceQuota{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *resourceQuotaController) Update(obj *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	result := &v1.ResourceQuota{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *resourceQuotaController) UpdateStatus(obj *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	result := &v1.ResourceQuota{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *resourceQuotaController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *resourceQuotaController) Get(namespace, name string, options metav1.GetOptions) (*v1.ResourceQuota, error) {
	result := &v1.ResourceQuota{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *resourceQuotaController) List(namespace string, opts metav1.ListOptions) (*v1.ResourceQuotaList, error) {
	result := &v1.ResourceQuotaList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *resourceQuotaController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *resourceQuotaController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.ResourceQuota, error) {
	result := &v1.ResourceQuota{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type resourceQuotaCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *resourceQuotaCache) Get(namespace, name string) (*v1.ResourceQuota, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.ResourceQuota), nil
}

func (c *resourceQuotaCache) List(namespace string, selector labels.Selector) (ret []*v1.ResourceQuota, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ResourceQuota))
	})

	return ret, err
}

func (c *resourceQuotaCache) AddIndexer(indexName string, indexer ResourceQuotaIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ResourceQuota))
		},
	}))
}

func (c *resourceQuotaCache) GetByIndex(indexName, key string) (result []*v1.ResourceQuota, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.ResourceQuota, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.ResourceQuota))
	}
	return result, nil
}

type ResourceQuotaStatusHandler func(obj *v1.ResourceQuota, status v1.ResourceQuotaStatus) (v1.ResourceQuotaStatus, error)

type ResourceQuotaGeneratingHandler func(obj *v1.ResourceQuota, status v1.ResourceQuotaStatus) ([]runtime.Object, v1.ResourceQuotaStatus, error)

func RegisterResourceQuotaStatusHandler(ctx context.Context, controller ResourceQuotaController, condition condition.Cond, name string, handler ResourceQuotaStatusHandler) {
	statusHandler := &resourceQuotaStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromResourceQuotaHandlerToHandler(statusHandler.sync))
}

func RegisterResourceQuotaGeneratingHandler(ctx context.Context, controller ResourceQuotaController, apply apply.Apply,
	condition condition.Cond, name string, handler ResourceQuotaGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &resourceQuotaGeneratingHandler{
		ResourceQuotaGeneratingHandler: handler,
		apply:                          apply,
		name:                           name,
		gvk:                            controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterResourceQuotaStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type resourceQuotaStatusHandler struct {
	client    ResourceQuotaClient
	condition condition.Cond
	handler   ResourceQuotaStatusHandler
}

func (a *resourceQuotaStatusHandler) sync(key string, obj *v1.ResourceQuota) (*v1.ResourceQuota, error) {
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

type resourceQuotaGeneratingHandler struct {
	ResourceQuotaGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *resourceQuotaGeneratingHandler) Remove(key string, obj *v1.ResourceQuota) (*v1.ResourceQuota, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.ResourceQuota{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *resourceQuotaGeneratingHandler) Handle(obj *v1.ResourceQuota, status v1.ResourceQuotaStatus) (v1.ResourceQuotaStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.ResourceQuotaGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
