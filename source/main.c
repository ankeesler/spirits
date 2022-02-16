#include <stdio.h>

#include "context.h"
#include "location.h"
#include "data/locations.h"

static void run_from_context(context_t *context);

int main(int argc, char *argv[]) {
  const char *context_file = argv[1];

  // Load context
  // If no context, set to default.
  context_t context;
  if (!load_context(&context, context_file)) {
    // TODO: extract this logic to storyline
    context.location_id = 0;
    context.held_spirit_ids[0] = 0;
    context.held_spirit_ids_count = 1;
  }

  while (1) {
    run_from_context(&context);
    store_context(&context, context_file);
  }

  return 0;
}

static void run_from_context(context_t *context) {
  const location_t *location = &locations[context->location_id];

  if (!(*location->on_enter)(context)) {
    // TODO: not allowed in here...
  }

  run_menu(&location->menu, context);
}
