import java.util.*;

public class FibonacciSumSquares {
    private static long getFibonacciSumSquaresNaive(long n) {
        if (n <= 1)
            return n;

        long previous = 0;
        long current  = 1;
        long sum      = 1;

        for (long i = 0; i < n - 1; ++i) {
            long tmp_previous = previous;
            previous = current;
            current = tmp_previous + current;
            sum += current * current;
        }

        return sum % 10;
    }

    private static long getFibonacciSumSquares(long n) {
        if (n <= 1)
            return n;

        //F1^2 + F2^2 + ... + Fn^2 = Fn * (Fn + Fn-1)
        long fn = 0, fn_1 = 0;

        long pisano = pisanoPeriodMod10();
        long n_1 = (n - 1) % pisano;
        n = n % pisano;
        if(n <= 1)
            fn = n;
        else {
            long previous = 0, current  = 1;
            for (long i = 0; i < n - 1; ++i) {  
                long tmp_previous = previous;
                previous = current;
                current = (tmp_previous + current) % 10;   
            }
            fn = current;
        }
        if(n_1 <= 1)
            fn_1 = n_1;
        else {
            long previous = 0, current  = 1;
            for (long i = 0; i < n_1 - 1; ++i) {  
                long tmp_previous = previous;
                previous = current;
                current = (tmp_previous + current) % 10;   
            }
            fn_1 = current;
        }
        return (fn * (fn + fn_1)) % 10;
    }

    private static long pisanoPeriodMod10(){
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
        long n = scanner.nextLong();
        long s = getFibonacciSumSquares(n);
        System.out.println(s);
    }
}

