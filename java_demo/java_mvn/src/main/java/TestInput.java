

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class TestInput {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int times = sc.nextInt();
        sc.nextLine();
        List<String> list = new ArrayList<>();
        for(int i = 0;i < times;i++){
            String s = sc.nextLine();
            list.add(s);
        }
        for(String s: list){
            System.out.println("output"+s);
        }
    }
}
