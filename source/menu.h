#ifndef __MENU_H__
#define __MENU_H__

typedef struct {
  const char *name;
  const char *description;
  void (*run)(void *context);
} menu_item_t;

typedef struct {
  const char *title;
  menu_item_t items[];
} menu_t;

void run_menu(const menu_t *menu, void *context);

#endif /* __MENU_H__ */
