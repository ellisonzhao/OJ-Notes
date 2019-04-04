package AcWing.num_567;


import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = sc.nextInt();
        int m = sc.nextInt();
        int result = m / n + (m % n == 0 ? 0 : 1);
        System.out.println(result);
    }
}
