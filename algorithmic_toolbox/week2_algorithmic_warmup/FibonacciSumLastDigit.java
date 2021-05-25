import java.util.*;

public class FibonacciSumLastDigit {
    private static long getFibonacciSumNaive(long n) {
        if (n <= 1)
            return n;

        long previous = 0;
        long current  = 1;
        long sum      = 1;

        for (long i = 0; i < n - 1; ++i) {
            long tmp_previous = previous;
            previous = current;
            current = tmp_previous + current;
            sum += current;
        }

        return sum % 10;
    }

    private static long getFibonacciSum(long n) {
        if (n <= 1)
            return n;
        long pisano = pisanoPeriodMod10();
        //sum(n) = f(n + 2) - 1
        long fn2 = 0;
        if((n + 2) % pisano <= 1)
            fn2 = (n + 2) % pisano;
        else {
            long prev = 0, curr = 1;
            for(int i = 0; i < (n + 2 - 1) % pisano; i++) {
                long temp = curr;
                curr = (curr + prev) % 10;
                prev = temp;
            }
            fn2 = curr;
        }
        return (fn2 + 10 - 1) % 10; //+10以防fn2-1<0
    }

    private static long pisanoPeriodMod10() {
        long prev = 0;
        long curr = 1;
        long period = 0;
        for(int i = 0; i < 10 * 10; i++) {
            long temp = curr;
            curr = (curr + prev) % 10;
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
        long n = scanner.nextLong();
        long s = getFibonacciSum(n);
        System.out.println(s);
    }
}

