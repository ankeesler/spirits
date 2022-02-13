#ifndef __LOG_H__
#define __LOG_H__

#include <stdio.h>

#define log(...)                                                  \
  do {                                                            \
    fprintf(stderr, "spirits: %s:%d: ", __FILE__, __LINE__);      \
    fprintf(stderr, __VA_ARGS__);                                 \
    fprintf(stderr, "\n");                                        \
    fflush(stderr);                                               \
  } while(0);

#endif /* __LOG_H__ */
