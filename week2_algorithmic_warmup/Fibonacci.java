import java.util.Scanner;

public class Fibonacci {
  private static long calc_fib_recursive(int n) {
    if (n <= 1)
      return n;

    return calc_fib_recursive(n - 1) + calc_fib_recursive(n - 2);
  }

  private static long calc_fib_iterative(int n) {
    if (n <= 1)
      return n;
    int fn = 0, fn_1 = 1, fn_2 = 0;
    for(int i = 2; i <= n; i++) {
      fn = fn_1 + fn_2;
      fn_2 = fn_1;
      fn_1 = fn;
    }
    return fn;
  }

  public static void main(String args[]) {
    Scanner in = new Scanner(System.in);
    int n = in.nextInt();

    System.out.println(calc_fib_iterative(n));
  }
}
