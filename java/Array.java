import java.util.Arrays;

class Array{
    public static void main(String args[]){
        //int型の配列宣言
        int[] nums = new int[5];

        //配列の表示（表示方法については、後の章で述べる）
        System.out.println(Arrays.toString(nums));

        //要素の代入
        nums[0] = 1;

        //配列の表示
        System.out.println(Arrays.toString(nums));

        //要素の表示
        System.out.println(nums[0]);
    }
}