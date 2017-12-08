#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>


int main()
{
    int search_limit, desk, elf;
    const int presents = atoi(getenv("PRESENTS"));

    clock_t start, end;
    double cpu_time_used;

    start = clock();

    if (presents >= 390 && presents < 20160) {
        search_limit = presents/20;
    } else if (presents >= 20160 && presents <= 6770400) {
        search_limit = presents/30;
    } else if (presents > 6770400 ) {
        search_limit = presents/40;
    } else {
        search_limit = presents/10;
    }

    int cache[search_limit+1];
    memset(cache, 0, sizeof cache);

    for (desk = 1; desk <= search_limit; ++desk) {
        for (elf = desk; elf <= search_limit; elf += desk) {
            cache[elf] += desk;
        }
        if (cache[desk] * 10 >= presents) {
            printf("%d\n", desk);
            break;
        }
    }

    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;

    printf("%fs\n", cpu_time_used);
    return 0;
}