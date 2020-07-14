/*
Copyright 2020 Google LLC

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/googleinterns/knative-source-mongodb/pkg/apis/sources/v1alpha1"
	scheme "github.com/googleinterns/knative-source-mongodb/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MongoDbSourcesGetter has a method to return a MongoDbSourceInterface.
// A group's client should implement this interface.
type MongoDbSourcesGetter interface {
	MongoDbSources(namespace string) MongoDbSourceInterface
}

// MongoDbSourceInterface has methods to work with MongoDbSource resources.
type MongoDbSourceInterface interface {
	Create(*v1alpha1.MongoDbSource) (*v1alpha1.MongoDbSource, error)
	Update(*v1alpha1.MongoDbSource) (*v1alpha1.MongoDbSource, error)
	UpdateStatus(*v1alpha1.MongoDbSource) (*v1alpha1.MongoDbSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MongoDbSource, error)
	List(opts v1.ListOptions) (*v1alpha1.MongoDbSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MongoDbSource, err error)
	MongoDbSourceExpansion
}

// mongoDbSources implements MongoDbSourceInterface
type mongoDbSources struct {
	client rest.Interface
	ns     string
}

// newMongoDbSources returns a MongoDbSources
func newMongoDbSources(c *SourcesV1alpha1Client, namespace string) *mongoDbSources {
	return &mongoDbSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mongoDbSource, and returns the corresponding mongoDbSource object, and an error if there is any.
func (c *mongoDbSources) Get(name string, options v1.GetOptions) (result *v1alpha1.MongoDbSource, err error) {
	result = &v1alpha1.MongoDbSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mongodbsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MongoDbSources that match those selectors.
func (c *mongoDbSources) List(opts v1.ListOptions) (result *v1alpha1.MongoDbSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MongoDbSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mongodbsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mongoDbSources.
func (c *mongoDbSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mongodbsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mongoDbSource and creates it.  Returns the server's representation of the mongoDbSource, and an error, if there is any.
func (c *mongoDbSources) Create(mongoDbSource *v1alpha1.MongoDbSource) (result *v1alpha1.MongoDbSource, err error) {
	result = &v1alpha1.MongoDbSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mongodbsources").
		Body(mongoDbSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mongoDbSource and updates it. Returns the server's representation of the mongoDbSource, and an error, if there is any.
func (c *mongoDbSources) Update(mongoDbSource *v1alpha1.MongoDbSource) (result *v1alpha1.MongoDbSource, err error) {
	result = &v1alpha1.MongoDbSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mongodbsources").
		Name(mongoDbSource.Name).
		Body(mongoDbSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mongoDbSources) UpdateStatus(mongoDbSource *v1alpha1.MongoDbSource) (result *v1alpha1.MongoDbSource, err error) {
	result = &v1alpha1.MongoDbSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mongodbsources").
		Name(mongoDbSource.Name).
		SubResource("status").
		Body(mongoDbSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the mongoDbSource and deletes it. Returns an error if one occurs.
func (c *mongoDbSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mongodbsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mongoDbSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mongodbsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mongoDbSource.
func (c *mongoDbSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MongoDbSource, err error) {
	result = &v1alpha1.MongoDbSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mongodbsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}