#include <stdio.h>

#include "battle.h"
#include "context.h"
#include "conversation.h"

int Duncee_on_enter(context_t *context) {
  // TODO: how do we generate conversation scripts...
  const dialogue_t dialogue[] = {
    {
      .who = "Djaro",
      .text = "Wake up! Grab your things and come outside.",
    },
    {
      .who = "<someone from outside>",
      .text = "What is going on?",
    },
    {
      .who = "<someone else from outside>",
      .text = "What is that thing?",
    },
    {
      .who = "You",
      .text = "*Leave house*",
    },
    {
      .who = "Djaro",
      .text = "Time to take this thing down!",
    },
    { .who = NULL },
  };
  const conversation_t convo = { .dialogue = dialogue, };
  run_conversation(&convo);

  run_battle(context);

  return 1;
}

void Duncee_menu_item_Home(context_t *context, int action) {
}
