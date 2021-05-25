import java.util.Scanner;

public class Change {
    private static int getChange(int m) {
        int min = 0;
        min += m / 10;
        int mod10 = m % 10;
        min += mod10 / 5;
        int mod5 = mod10 % 5;
        min += mod5;
        return min;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int m = scanner.nextInt();
        System.out.println(getChange(m));

    }
}

