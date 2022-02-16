#ifndef __MENU_H__
#define __MENU_H__

#include "context.h"

typedef struct {
  const char *name;
  const char *description;
  void (*run)(context_t *context, int option);
} menu_item_t;

typedef struct {
  const char *title;
  const menu_item_t *items; // Ending in { .name = NULL }
} menu_t;

void run_menu(const menu_t *menu, context_t *context);

#endif /* __MENU_H__ */
