# 面向对象部分

> 封装 继承 多态

| 关键字    | 可以访问           | 不可访问     |
| --------- | ------------------ | ------------ |
| private   | 类自身             | 子类，类外部 |
| protected | 类自身，子类       | 类外部       |
| public    | 类自身，子类，外部 |              |

## super关键字

+ 引用父类的实例变量
+ 调用父类方法
+ super()调用父类构造函数

**方法调用值传递（传递地址和c语言类似）**

# 数据类型

> 低级可以向高级自动转换（隐式）int->long
>
> 高级向低级需要强制转换（显式）long->int

| 数据类型 | 所占空间 | 包装类型  |
| -------- | -------- | --------- |
| byte     | 8        | Byte      |
| short    | 16       | Short     |
| int      | 32       | Integer   |
| long     | 64       | Long      |
| float    | 32       | Float     |
| double   | 64       | Double    |
| char     | 16       | Character |
| boolean  | 1        | Boolean   |

自动装箱and拆箱

#### 基本类型缓冲池

+ boolean true,false
+ byte all values
+ short -128 - 127
+ int -128 - 127
+ char \u0000 - \u007F

在这些范围里的值直接使用缓冲池中对应对象



# 字符串部分

### String

```  java
public final class String
    implements java.io.Serializable, Comparable<String>, CharSequence{
    /**
     * The value is used for character storage.
     */
    @Stable
    private final byte[] value;

    /**
     * The identifier of the encoding used to encode the bytes in
     * {@code value}
     */
    private final byte coder;
}
```

final class不可被继承

字符串内容不可变

#### 不可变字符串好处

+ 可以缓存hash值

+ String Pool的需要

  String s1 = "hahaha";//以字面量形式创建，自动加入常量池

``` java
		String s1 = "123";//以字面量形式创建，自动加入String Pool
        String s2 = "123";//返回String Pool中的对象
        String s3 = new String("123");//在Java Heap中重新创建对象
        String s4 = new String("123").intern();//直接在String pool中创建对象
        System.out.println(s1 == s2);
        System.out.println(s2 == s3);
        System.out.println(s1 == s4);
```

## StringBuilder & StringBuffer

| StringBuilder | StringBuffer                 | String           |
| ------------- | ---------------------------- | ---------------- |
| 可变          | 可变                         | 不可变           |
| 线程不安全    | 有synchronized修饰，线程安全 | 不可变，线程安全 |

# 关键字

### final

**数据**

+ 基本类型，final让数值不变
+ 引用类型，final让引用不变，就是不能引用其他对象，但是引用对象本身可变

**方法**

此类方法不能被子类重写

**类**

该类不能被继承

## static

``` java
public class Base {
    public static void main(String[] args) {
        StaticTest r1 = new StaticTest("hhh");
        System.out.println(StaticTest.getCount());
        StaticTest r2 = new StaticTest("ppp");
        System.out.println(StaticTest.getCount());
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

    public static int getCount() {
        return count;
    }
}
```

+ 静态变量又称类变量，该变量属于该类，类的所有实例共享该变量，可以通过类名直接访问，该变量在内存中只存在一份
+ 实例变量在每个实例创建时都会初始化一份新的

**静态方法**

静态方法在类加载阶段就产生，不依赖于任何实例所以静态方法必须有实现，也不能是抽象方法，而且只能访问静态变量，不能访问实例变量

**静态语句块**

在类初始化的时候有且仅执行一次

//todo：静态导包



