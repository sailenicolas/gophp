// Copyright 2017 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

#ifndef __CONTEXT_H__
#define __CONTEXT_H__
typedef struct _engine_context {
} engine_context;

engine_context *context_new();
void context_exec(engine_context *context, char *filename);
void *context_eval(engine_context *context, char *script);
void context_bind(engine_context *context, char *name, void *value);
void context_destroy(engine_context *context);
//static void _context_bind(char *name, zval *value);
//static void _context_eval(zend_op_array *op, zval *ret);

#endif
