#include <stdio.h>
#include <stdlib.h>
#include <time.h>


unsigned int main()
{
    unsigned static int search_limit, desk, elf, *cache, cpu_time_used;
    static clock_t start;
    const unsigned int presents = atoi(getenv("PRESENTS"));

    start = clock();

    if (presents > 6770400) {
        search_limit = presents/40;
    } else if (presents >= 20160 && presents <= 6770400) {
        search_limit = presents / 30;
    } else if (presents >= 390 && presents < 20160) {
        search_limit = presents/20;
    } else {
        search_limit = presents/10;
    }

    cache = (int*) calloc(search_limit+1, sizeof(int));
    if (cache == NULL) {
        printf("Got messed up trying to allocate memory, bailing :'( ");
        exit(0);
    }

    for (desk = 1; desk <= search_limit; ++desk) {
        for (elf = desk; elf <= search_limit; elf += desk) {
            cache[elf] += desk;
        }
        if (cache[desk] * 10 >= presents) {
            break;
        }
    }

    cpu_time_used = (clock() - start);

    printf("%d\n", desk);
    printf("%dµs\n", cpu_time_used);
    free(cache);
    return 0;
}