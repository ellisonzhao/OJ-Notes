package AcWing.num_729;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.HashMap;
import java.util.Map;

/**
 * @author charlsonz
 * @date 2019-04-23 22:44
 */
public class Main {

    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int M = Integer.parseInt(br.readLine().trim());
        int values = 0;
        // 记录特征出现次数
        Map<String, Integer> frequencyMap = new HashMap<>();

        // 记录特征上一次出现在哪一帧
        Map<String, Integer> previousMap = new HashMap<>();
        for (int m = 1; m <= M; ++m) {
            String line = br.readLine();
            if (line == null || line.length() == 0)
                break;


            // 每行的数据存到字符串数组，方便 Map 存储
            String[] parts = line.trim().split(" ");

            // 这一帧有多少个特征
            int n = Integer.parseInt(parts[0]);

            for (int i = 1; i <= n; ++i) {

                // 特征<1, 1> 转化为字符串 "1,1"
                // String 不可变的特性,很适合用作 HashMap 的 key
                String pair = parts[2 * i - 1] + "," + parts[2 * i];

                if (previousMap.containsKey(pair)) {
                    // 这个特征出现过
                    // 上一次出现在哪一帧
                    int frameIndex = previousMap.get(pair);

                    // 连续出现该特征
                    if (frameIndex == m - 1) {
                        frequencyMap.put(pair, frequencyMap.get(pair) + 1);
                    } else {
                        frequencyMap.put(pair, 1);
                    }
                } else {
                    // 这个特征第一次出现
                    frequencyMap.put(pair, 1);
                }
                previousMap.put(pair, m);
                values = Math.max(values, frequencyMap.get(pair));
            }
        }

        System.out.println(values);

    }
}