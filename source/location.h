#ifndef __LOCATION_H__
#define __LOCATION_H__

#include "context.h"
#include "menu.h"

typedef struct {
  int id;
  const char *name;
  int (*on_enter)(context_t *context); // Return 0 to disallow entry
  int (*on_exit)(context_t *context); // Return 0 to disallow exit
  menu_t menu;
} location_t;

#endif /* __LOCATION_H__ */
