#ifndef __CONTEXT_H__
#define __CONTEXT_H__

#define MAX_HELD_SPIRITS (6)
#define MAX_HELD_ITEMS (256)

typedef struct {
  int opponent_spirit_id;
} battle_context_t;

typedef struct {
  int location_id;

  int held_spirit_ids[MAX_HELD_SPIRITS];
  int held_spirit_ids_count;

  int held_item_ids[MAX_HELD_ITEMS];
  int held_item_ids_count;

  battle_context_t battle_context;
} context_t;

int load_context(context_t *context, const char *file);
int store_context(context_t *context, const char *file);

#endif /* __CONTEXT_H__ */
