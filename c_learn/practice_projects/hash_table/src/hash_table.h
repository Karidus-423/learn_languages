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

void ht_insert(ht_hash_table *ht, const char *key, const char *value);
char *ht_search(ht_hash_table *ht, const char *key);
void ht_delete(ht_hash_table *h, const char *key);
