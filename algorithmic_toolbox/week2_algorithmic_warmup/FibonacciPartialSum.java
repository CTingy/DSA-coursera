import java.util.*;

public class FibonacciPartialSum {
    private static long getFibonacciPartialSumNaive(long from, long to) {
        long sum = 0;

        long current = 0;
        long next  = 1;

        for (long i = 0; i <= to; ++i) {
            if (i >= from) {
                sum += current;
            }

            long new_current = next;
            next = next + current;
            current = new_current;
        }

        return sum % 10;
    }

    private static long getFibonacciPartialSum(long from, long to) {
        long pisano = pisanoPeriodMod10();

        //Fm + Fm+1 + ... + Fn = Sn - Sm-1 = (Fn+2 - 1) - (Fm+1 -1) = Fn+2 - Fm+1
        long fFromPlus1 = 0, fToPlus2 = 0;
        if((from + 1) % pisano <= 1)
            fFromPlus1 = (from + 1) % pisano;
        else {
            long prev = 0, curr = 1;
            for (long i = 0; i < ((from + 1) - 1) % pisano; ++i) {
                long temp = curr;
                curr = (curr + prev) % 10;
                prev = temp;
            }
            fFromPlus1 = curr;
        }
        if((to + 2) % pisano <= 1)
            fToPlus2 = (to + 2) % pisano;
        else {
            long prev = 0, curr = 1;
            for (long i = 0; i < ((to + 2) - 1) % pisano; ++i) {
                long temp = curr;
                curr = (curr + prev) % 10;
                prev = temp;
            }
            fToPlus2 = curr;
        }
        return ((fToPlus2 + 10) - fFromPlus1) % 10; //因為fToPlus2可能小於fFromPlus1，所以要+10
    }

    private static long pisanoPeriodMod10() {
        long period = 0, prev = 0, curr = 1;
        for(int i = 0; i < 10 * 10; i++) {
            long temp = curr;
            curr = (prev + curr) % 10;
            prev = temp;
            if(prev == 0 && curr == 1) {
                period = i + 1;
                break;
            }
        }
        return period;
    }
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        long from = scanner.nextLong();
        long to = scanner.nextLong();
        System.out.println(getFibonacciPartialSum(from, to));
    }
}

