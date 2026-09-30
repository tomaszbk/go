/*
Linux ELF:
gcc -gdwarf-4 -m64 -c bitfields.c -o bitfields.elf4
*/

typedef struct another_struct {
  unsigned short quix;
  int xyz[0];
  unsigned  x:1;
  long long array[40];
} t_another_struct;
t_another_struct q2;

