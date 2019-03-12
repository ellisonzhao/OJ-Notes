package LeetCode.num_690;

import java.util.*;

class Employee {
    // It's the unique id of each node;
    // unique id of this employee
    public int id;
    // the importance value of this employee
    public int importance;
    // the id of direct subordinates
    public List<Integer> subordinates;
}

public class MainClass {

    public int getImportance(List<Employee> employees, int id) {
        int ans = 0;
        Queue<Employee> queue = new LinkedList<>();
        Map<Integer, Employee> employeeMap = new HashMap<>();
        for (Employee employee : employees) {
            employeeMap.put(employee.id, employee);
        }
        queue.offer(employeeMap.get(id));
        while (!queue.isEmpty()) {
            Employee temp = queue.poll();
            ans += temp.importance;
            if (temp.subordinates != null) {
                for (Integer subId : temp.subordinates) {
                    queue.offer(employeeMap.get(subId));
                }
            }
        }
        return ans;
    }
}
