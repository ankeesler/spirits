#include <assert.h>
#include <stdio.h>
#include <stdlib.h>

#include "context.h"
#include "log.h"
#include "menu.h"

static void trim_newline(char *s, int len);

void run_menu(const menu_t *menu, context_t *context) {
  while (1) {
    printf("@@@@@@@@@\n");
    printf("  %s\n", menu->title);
  
    int item_num = 0;
    for (item_num = 0; menu->items[item_num].name != NULL; item_num++) {
      const menu_item_t *mi = &menu->items[item_num];
      printf("    %d\t-\t%s\t-\t%s\n", item_num, mi->name, mi->description);
    }

    printf("\n  Enter option: ");
    char input[8] = {0,};
    fgets(input, sizeof(input) - 1, stdin);
    trim_newline(input, sizeof(input) - 1);
    log("got input '%s'", input);

    int option = atoi(input);
    if (input[0] == 0 || option < 0 || option >= item_num) {
      printf("\n Invalid option: '%s' (%d)\n", input, option);
    } else {
      (*menu->items[option].run)(context, option);
    }
  }
}

static void trim_newline(char *s, int len) {
  for (int i = 0; i < len; i++) {
    if (s[i] == '\n') {
      s[i] = '\0';
      return;
    }
  }
}
