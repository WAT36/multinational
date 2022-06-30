import java.util.Scanner;

class Main{
  public static void main(String args[]){

    //Scannerのインスタンスを生成
    Scanner sc = new Scanner(System.in);

    //String型で１行読み込む
    String s = sc.nextLine();

    System.out.println("入力された値:"+s);
  }
}