import java.util.Scanner;

public class FractionalKnapsack {
    private static double getOptimalValue(int capacity, int[] values, int[] weights) {
        double value = 0;
        //sort values, weights by valuePerWeight desc
        for (int i = 0; i < values.length-1; i++){
            int max_idx = i;
            for (int j = i+1; j < values.length; j++) {
                if ((double)values[j] / weights[j] > (double)values[max_idx] / weights[max_idx])
                    max_idx = j;
            }
  
            int temp = values[max_idx];
            values[max_idx] = values[i];
            values[i] = temp;

            int temp2 = weights[max_idx];
            weights[max_idx] = weights[i];
            weights[i] = temp2;
        }

        for(int i = 0; i < values.length; i++) {
            if(capacity >= weights[i]){
                value += values[i];
                capacity -= weights[i];
            }else {
                value += ((double)values[i] / weights[i]) * capacity;
                break;
            }
        }
        return value;
    }

    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        int n = scanner.nextInt();
        int capacity = scanner.nextInt();
        int[] values = new int[n];
        int[] weights = new int[n];
        for (int i = 0; i < n; i++) {
            values[i] = scanner.nextInt();
            weights[i] = scanner.nextInt();
        }
        System.out.println(getOptimalValue(capacity, values, weights));
    }
} 
