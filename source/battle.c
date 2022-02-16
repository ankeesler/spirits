#include "battle.h"

#include <stdio.h>

#include "context.h"
#include "menu.h"

static void battle_action_attack(context_t *context, int action);
static void battle_action_spirits(context_t *context, int action);
static void battle_action_item(context_t *context, int action);
static void battle_action_flee(context_t *context, int action);

static void view_spirit(context_t *context, int action);

static const menu_item_t battle_menu_items[] = {
  {
    .name = "Attack",
    .description = "Perform attack with current spirit",
    .run = battle_action_attack,
  },
  {
    .name = "Spirits",
    .description = "Examine spirits",
    .run = battle_action_spirits,
  },
  {
    .name = "Item",
    .description = "Use item",
    .run = battle_action_item,
  },
  {
    .name = "Flee",
    .description = "Run away",
    .run = battle_action_flee,
  },
};

static const menu_t battle_menu = {
  .title = "Battle!",
  .items = battle_menu_items,
};

void run_battle(context_t *context) {
  run_menu(&battle_menu, context);
}

static void battle_action_attack(context_t *context, int action) {
}

static void battle_action_spirits(context_t *context, int action) {
  menu_item_t spirits_menu_items[MAX_HELD_SPIRITS] = { 0, };
  for (int i = 0; i < context->held_spirit_ids_count; i++) {
    /* const spirit_t *spirit = &spirits[context->held_spirit_ids[i]]; */
    /* spirits_menu_items[i].name = spirit->name; */
    spirits_menu_items[i].name = "xxx";
    spirits_menu_items[i].description = "View xxx";
    spirits_menu_items[i].run = view_spirit;
  }

  const menu_t spirits_menu = {
    .title = "Spirits",
    .items = spirits_menu_items,
  };
  run_menu(&spirits_menu, context);
}

static void battle_action_item(context_t *context, int action) {
  if (context->held_item_ids_count == 0) {
    printf("  You don't have any items!\n");
    return;
  }

  /* const menu_item_t item_menu_items[MAX_HELD_ITEMS] = { 0, }; */
  /* for (int i = 0; i < context->held_item_ids_count; i++) { */
  /*   const item_t *item = &items[context->held_item_ids[i]]; */
  /*   item_menu_items[i].name = item->name; */
  /*   item_menu_items[i].description = item->description; */
  /*   item_menu_items[i].run = NULL; */
  /* } */
}

static void battle_action_flee(context_t *context, int action) {
}

static void view_spirit(context_t *context, int action) {
  /* int spirit_id = context->held_spirit_ids[action]; */
  printf("  Name: xxx\n");
  printf("  Description: xxx\n");
  printf("  Type: xxx\n");
  printf("  HP: xxx\n");
  printf("  Attack: xxx\n");
  printf("  Defense: xxx\n");
  printf("  Speed: xxx\n");
}
