#ifndef __CONVERSION_H__
#define __CONVERSION_H__

typedef struct {
  const char *who;
  const char *text;
} dialogue_t;

typedef struct {
  const dialogue_t *dialogue; // Ending in { .who = NULL }
} conversation_t;

void run_conversation(const conversation_t *conversation);

#endif /* __CONVERSION_H__ */
