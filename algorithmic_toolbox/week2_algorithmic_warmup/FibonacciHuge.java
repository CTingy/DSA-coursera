import java.util.*;

public class FibonacciHuge {
    private static long getFibonacciHugeNaive(long n, long m) {
        if (n <= 1)
            return n;

        long previous = 0;
        long current  = 1;

        for (long i = 0; i < n - 1; ++i) {
            long tmp_previous = previous;
            previous = current;
            current = tmp_previous + current;
        }

        return current % m;
    }

    private static long getFibonacciHuge(long n, long m) {
        if (n <= 1)
            return n;

        long period = pisanoPeriodLength(m);
        n = n % period;
        if(n <= 1)
            return n;
        else {
            long prev = 0, current = 1;     
            for(int i = 0; i < n - 1; i++) {
                long temp = current;
                current = (current + prev) % m;
                prev = temp;
            }
            return current;
        }
    }

    private static long pisanoPeriodLength(long m){
        long period = 0;
        long prev = 0;
        long current = 1;
        for(int i = 0; i < m * m; i++) {    //m*m?
            long temp = current;
            current = (current + prev) % m;
            prev = temp;

            if(prev == 0 && current == 1) {
                period = i + 1; //第0個也要算進去週期，所以要i+1
                break;
            }
        }
        return period;
    }
    
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        long n = scanner.nextLong();
        long m = scanner.nextLong();
        System.out.println(getFibonacciHuge(n, m));
    }
}

