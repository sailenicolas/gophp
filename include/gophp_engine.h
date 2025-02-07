// Copyright 2017 Alexander Palaistras. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

#ifndef __ENGINE_H__

#define __ENGINE_H__

typedef struct _php_engine {
} php_engine;

php_engine *engine_init(void);

void engine_shutdown();

//static size_t engine_ub_write(const char *str, size_t len);

//static char *engine_read_cookies();

//static void engine_register_variables(zval *track_vars_array);

//static void engine_log_message(const char *str, int syslog_type_int);

#endif
