#include "conversation.h"

#include <stdio.h>

void run_conversation(const conversation_t *conversation) {
  printf("#########\n");
  for (const dialogue_t *d = conversation->dialogue; d->who != NULL; d++) {
    printf("\n");
    printf("  ### %s: %s\n", d->who, d->text);
    printf("\n");
    printf("  Press enter to continue");
    getchar();
  }
}
