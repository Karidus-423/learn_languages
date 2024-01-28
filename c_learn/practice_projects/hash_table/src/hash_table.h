// Key value pairs.
typedef struct {
  char *key;
  char *value;
} ht_item;

// The hash table will store array of pointers to items and details about the
// size and how full it is.
typedef struct {
  int size;
  int count;
  ht_item **items;
} ht_hash_table;
