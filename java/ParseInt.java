class ParseInt{
    public static void main(String args[]){
      String s = "10";
      int i = Integer.parseInt(s);
      System.out.println(i); // 10

      s = "AA";
      i = Integer.parseInt(s); // NumberFormatExceptionエラーが発生
    }
}