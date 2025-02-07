// Copyright 2017 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package gophp

// #include <stdlib.h>
// #include <main/php.h>
// #include "include/gophp_context.h"
import "C"

import (
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

// Context represents an individual execution context.
type Context struct {
	// Output and Log are unbuffered writers used for regular and debug output,
	// respectively. If left unset, any data written into either by the calling
	// context will be lost.
	Output io.Writer
	Log    io.Writer

	// Header represents the HTTP headers set by current PHP context.
	Header http.Header

	context *C.struct__gophp_context
	values  []*Value
}

// Bind allows for binding Go values into the current execution context under
// a certain name. Bind returns an error if attempting to bind an invalid value
// (check the documentation for NewValue for what is considered to be a "valid"
// value).
func (c *Context) Bind(name string, val interface{}) error {
	v, err := NewValue(val)
	if err != nil {
		return err
	}

	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))

	C.context_bind(c.context, n, v.Ptr())
	c.values = append(c.values, v)

	return nil
}

// Exec executes a PHP script pointed to by filename in the current execution
// context, and returns an error, if any. Output produced by the script is
// written to the context's pre-defined io.Writer instance.
func (c *Context) Exec(filename string) error {
	f := C.CString(filename)
	defer C.free(unsafe.Pointer(f))

	_, err := C.context_exec(c.context, f)
	if err != nil {
		return fmt.Errorf("Error executing script '%s' in context", filename)
	}

	return nil
}

// Eval executes the PHP expression contained in script, and returns a Value
// containing the PHP value returned by the expression, if any. Any output
// produced is written context's pre-defined io.Writer instance.
func (c *Context) Eval(script string) (*Value, error) {
	s := C.CString(script)
	defer C.free(unsafe.Pointer(s))
	result, err := C.context_eval(c.context, s)
	if err != nil {
		return nil, fmt.Errorf("error executing script '%s' in context %s", script, err.Error())
	}
	defer C.free(unsafe.Pointer(result))

	val, err := NewValueFromPtr(unsafe.Pointer(result))
	if err != nil {
		return nil, err
	}

	c.values = append(c.values, val)

	return val, nil
}

// Destroy tears down the current execution context along with any active value
// bindings for that context.
func (c *Context) Destroy() {
	if c.context == nil {
		return
	}

	for _, v := range c.values {
		v.Destroy()
	}

	c.values = nil

	C.context_destroy(c.context)
	c.context = nil
}
