#include <stdio.h>

#include "menu.h"

static void run1(void *context) {
  printf("run1\n");
}

static void run2(void *context) {
  printf("run2\n");
}

static const menu_t main_menu = {
  .title = "main menu",
  .items = {
    { .name = "name1", .description = "description1", .run = run1, },
    { .name = "name2", .description = "description2", .run = run2, },
    { .name = NULL, },
  },
};

int main(int argc, char *argv[]) {
  run_menu(&main_menu, NULL);
  return 0;
}
