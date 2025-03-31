// The Computer Language Benchmarks Game
//   https:#salsa.debian.org/benchmarksgame-team/benchmarksgame/
//
//   Naive transliteration from Rex Kerr's Scala program
//   contributed by Isaac Gouy
//
// python version
// Addapted for c by sethbardsley

#include <stdio.h>
int main()
{
    int n = 10;
    int perm1[n];

    int perm1i = 0;
    while (perm1i < n)
    {
        perm1[perm1i] = perm1i;
        perm1i = perm1i + 1;
    }

    int perm[n];
    int count[n];
    int r = 0;
    int k = 0;
    int j = 0;
    int checksum = 0;
    int nperm = 0;
    int flips = 0;
    int f = 0;

    r = n;
    while (r > 0)
    {
        int i = 0;
        while (r != 1)
        {
            count[r - 1] = r;
            r = r - 1;
        }
        while (i < n)
        {
            perm[i] = perm1[i];
            i = i + 1;
        }

        // Count flips and update max and checksum
        f = 0;
        k = perm[0];
        while (k != 0)
        {
            i = 0;
            while (2 * i < k)
            {
                int t = perm[i];
                perm[i] = perm[k - i];
                perm[k - i] = t;
                i = i + 1;
            }

            k = perm[0];
            f = f + 1;
        }
        if (f > flips)
        {
            flips = f;
        }
        if ((nperm % 2) == 0)
        {
            checksum = checksum + f;
        }
        else
        {
            checksum = checksum - f;
        }

        // Use incremental change to generate another permutation
        int more = 1;
        while (more)
        {
            if (r == n)
            {
                printf("%d\n", checksum);
                r = 0;
                break;
            }

            int p0 = perm1[0];
            i = 0;
            while (i < r)
            {
                j = i + 1;
                perm1[i] = perm1[j];
                i = j;
            }

            perm1[r] = p0;

            count[r] = count[r] - 1;
            if (count[r] > 0)
            {
                more = 0;
            }
            else
            {
                r = r + 1;
            }
        }

        nperm = nperm + 1;
    }

    printf("%d\n", flips);
}