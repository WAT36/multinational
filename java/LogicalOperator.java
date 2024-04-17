class LogicalOperator{
    public static void main(String args[]){
        boolean a = true;
        boolean b = false;

        int c = 3; //0011
        int d = 5; //0101

        System.out.println(a && b);
        System.out.println(a || b);
        System.out.println(a ^ b);
        System.out.println(!a);

        System.out.println(c & d); // 0001
        System.out.println(c | d); // 0111
        System.out.println(c ^ d); // 0110
        System.out.println(~c);    // 0110
    }
}