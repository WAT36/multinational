import java.util.Arrays;
class Split{
    public static void main(String args[]){
        String s = "apple-banana-cherry-durian--";
        String[] t = s.split("-",-1);
        System.out.println(Arrays.toString(t)); //[apple, banana, cherry, durian, , ]
  
        t = s.split("-",0);
        System.out.println(Arrays.toString(t)); //[apple, banana, cherry, durian]
  
        t = s.split("-",2);
        System.out.println(Arrays.toString(t)); //[apple, banana-cherry-durian--]
    }
}