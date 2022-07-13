class Ascii{
    public static void main(String args[]){
      char c = 'a';
      int a = c;              // char型データ'a'を文字コードに変換
      System.out.println(a);  // 97

      String s = "aaaaa";
      c = s.charAt(0);        // sの0文字目をchar型で返す 
      System.out.println(c);  // a
      a = c;
      System.out.println(a);  // 97

      //a = s                 // Stringをそのまま文字コード変換するとコンパイルエラー
    }
}