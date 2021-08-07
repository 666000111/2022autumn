public class Base {
    public static void main(String[] args) {
        StaticTest r1 = new StaticTest("hhh");
        System.out.println(r1.getCount());
        StaticTest r2 = new StaticTest("ppp");
        System.out.println(r2.getCount());
        System.out.println(r1.getClass());

    }
}
class StaticTest {
    private static int count = 0;
    private String name;
    public StaticTest(String name){
        this.name = name;
        count++;
    }

    public int getCount() {
        return count;
    }
}
