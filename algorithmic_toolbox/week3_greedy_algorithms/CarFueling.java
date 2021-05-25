import java.util.*;
import java.io.*;

//test case
//49, 10, 6, [7, 14, 21, 28, 35, 42]
//500, 200, 4, [100, 200, 300, 400]

public class CarFueling {
    static int computeMinRefills(int tank, int[] stops) {
        if(tank >= stops[stops.length - 1])
            return 0;
        
        int refill = 0;
        int tempTank = tank;
        for(int i = 0; i < stops.length; i++) {
            if(i == stops.length - 1)
                break;
            int len = stops[i + 1] - stops[i];

            if(len > tank)
                return -1;
            else if(len > tempTank){
                refill++;
                tempTank = tank;
            } 
            tempTank -= len;
        }
        return refill;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int dist = scanner.nextInt();
        int tank = scanner.nextInt();
        int n = scanner.nextInt();
        int stops[] = new int[n + 2];
        stops[0] = 0;
        for (int i = 1; i < n + 1; i++) {
            stops[i] = scanner.nextInt();
        }
        stops[n + 1] = dist;
        System.out.println(computeMinRefills(tank, stops));
    }
}
