// THIS FILE IS GENERATED BY script/locations.awk

#include "locations.h"

#include <stdlib.h>

#include "location.h"

#include "locations_handlers.c"

static menu_item_t Duncee_menu_items[] = {
  {
    .name = "Home",
    .description = "Visit your home",
    .run = Duncee_menu_item_Home,
  },
  { .name = NULL, },
};

const location_t locations[] = {
  {
    .id = 0,
    .name = "Duncee",
    .on_enter = Duncee_on_enter,
    .menu = {
      .title = "Duncee",
      .items = Duncee_menu_items,
    },
  },
};
