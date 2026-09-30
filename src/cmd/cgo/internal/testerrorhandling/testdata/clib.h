// C helpers for the paired error-handling examples. The functions live in
// clib.c (not as static functions here) so that every Go file in the package
// observes the same counters.

typedef struct {
	int *p; // a pointer member makes cgo check &buf arguments
	int n;
} c_buf;

// Traced calls: each of these increments the call counter.
int c_div(int num, int den);   // errno = EDOM and result -1 when den == 0
void c_touch(int x);           // void; errno = EINVAL when x < 0
int c_add3(int a, int b, int c);
int c_strlen(const char *s);   // errno = EINVAL and result -1 for ""
int c_bufsum(c_buf *b, int k); // b->n + k; errno = EDOM when k == 0

// Observation helpers: these do not count as traced calls.
void c_note_handled(void);
int c_nhandled(void);
int c_ncalls(void);
void c_reset(void);
