package AcWing.num_57;

public class MainClass {
    public int digitAtIndex(int n) {
        if (n <= 9)
            return n;
        // 确定所在的数字有几位
        int digit = 1, count = 10;
        while (n >= count) {
            n -= count;
            digit += 1;
            // 比如两位数总共占180位
            count = (int) (Math.pow(10, digit) - Math.pow(10, digit - 1));
        }
        // n所在数字有digit位
        // 确定所在的数字 = digit位数的第一个数字 + more
        int num = (int) Math.pow(10, digit - 1) + n / digit;

        //确定是所在数字（num）的第几位(从右往左)
        // 注意是从第0位开始
        int index = digit - n % digit;
        for (int i = 1; i < index; ++i) {
            num /= 10;
        }
        return num % 10;
    }

    public static void main(String[] args) {
        System.out.println(new MainClass().digitAtIndex(10000));
    }
}
