import java.util.ArrayList;
import java.util.List;
class Main{
  public static void main(String args[]){

    List<String> l = new ArrayList<>();
    l.add("a");
    l.add("bb");
    l.add("ccc");
    l.add("dddd");

    System.out.println(l + " -> " + String.join("-",l));
  }
}